package service

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pbCA "github.com/plgd-dev/cloud/certificate-authority/pb"
	"github.com/plgd-dev/cloud/grpc-gateway/client"
	"github.com/plgd-dev/cloud/grpc-gateway/pb"
	"github.com/plgd-dev/cloud/http-gateway/uri"
	"github.com/plgd-dev/cloud/pkg/log"
	kitNetGrpc "github.com/plgd-dev/cloud/pkg/net/grpc"
	kitHttp "github.com/plgd-dev/cloud/pkg/net/http"
	"github.com/tmc/grpc-websocket-proxy/wsproxy"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/google/uuid"
	router "github.com/gorilla/mux"
)

//RequestHandler for handling incoming request
type RequestHandler struct {
	client   *client.Client
	config   *Config
	caClient pbCA.CertificateAuthorityClient
	mux      *runtime.ServeMux
}

type jsonPrettyMarshaller struct {
	*runtime.JSONPb
}

func newJsonPrettyMarshaller() *jsonPrettyMarshaller {
	return &jsonPrettyMarshaller{
		JSONPb: &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				EmitUnpopulated: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		},
	}
}

// ContentType always returns "application/json".
func (*jsonPrettyMarshaller) ContentType(_ interface{}) string {
	return "application/json-pretty"
}

func isContent(val map[string]interface{}) bool {
	_, ok := val["contentType"]
	if !ok {
		return false
	}
	_, ok = val["data"]
	if !ok {
		return false
	}
	return true
}

func modify(v interface{}) {
	valMap, ok := v.(map[string]interface{})
	if ok {
		if isContent(valMap) {

			return
		}
		for _, v := range valMap {
			modify(v)
		}
		return
	}
	valArr, ok := v.([]interface{})
	if ok {
		for _, v := range valArr {
			modify(v)
		}
	}
}

// Marshal marshals "v" into JSON.
func (j *jsonPrettyMarshaller) Marshal(orig interface{}) ([]byte, error) {
	val, ok := orig.(map[string]interface{})
	if !ok {
		return j.JSONPb.Marshal(v)
	}
	var string contentType
	var data interface{}
	for key, v := range val {
		if key == "contentType" {

		}

		if key == "data" {

		}
	}
}

//NewRequestHandler factory for new RequestHandler
func NewRequestHandler(config *Config, client *client.Client, caClient pbCA.CertificateAuthorityClient) *RequestHandler {
	return &RequestHandler{
		client:   client,
		caClient: caClient,
		config:   config,
		mux: runtime.NewServeMux(
			runtime.WithIncomingHeaderMatcher(func(key string) (string, bool) {
				switch strings.ToLower(key) {
				/*
					case "accept-content":
						return grpc.AcceptContentHeaderKey, true
				*/
				}
				return key, false
			}),
			runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					EmitUnpopulated: true,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{
					DiscardUnknown: true,
				},
			}),
		),
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := httputil.DumpRequest(r, false)
		if err != nil {
			log.Infof("Request: %v %v", r.Method, r.RequestURI)
		} else {
			log.Infof("Request: %v", string(data))
		}

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

type logger struct{}

func (logger) Warnln(v ...interface{})  { log.Warnf("%v", v) }
func (logger) Debugln(v ...interface{}) { log.Debugf("%v\n", v) }

func splitDevicePath(requestURI, prefix string) []string {
	p := kitHttp.CanonicalHref(requestURI)
	p = strings.TrimPrefix(p, prefix) // remove core prefix
	p = strings.TrimLeft(p, "/")
	return strings.Split(p, "/")
}

func resourcePendingCommandsMatcher(r *http.Request, rm *router.RouteMatch) bool {
	paths := splitDevicePath(r.RequestURI, uri.Devices)
	if len(paths) > 2 && paths[1] == uri.ResourcesPathKey && strings.Contains(paths[len(paths)-1], uri.PendingCommandsPathKey) {
		if rm.Vars == nil {
			rm.Vars = make(map[string]string)
		}
		rm.Vars[uri.DeviceIDKey] = paths[0]
		rm.Vars[uri.ResourceHrefKey] = strings.Split("/"+strings.Join(paths[2:len(paths)-1], "/"), "?")[0]
		return true
	}
	return false
}

func resourceMatcher(r *http.Request, rm *router.RouteMatch) bool {
	paths := splitDevicePath(r.RequestURI, uri.Devices)
	if len(paths) > 2 && paths[1] == uri.ResourcesPathKey {
		if rm.Vars == nil {
			rm.Vars = make(map[string]string)
		}
		rm.Vars[uri.DeviceIDKey] = paths[0]
		rm.Vars[uri.ResourceHrefKey] = strings.Split("/"+strings.Join(paths[2:], "/"), "?")[0]
		return true
	}
	return false
}

// NewHTTP returns HTTP server
func NewHTTP(requestHandler *RequestHandler, authInterceptor kitHttp.Interceptor) *http.Server {
	r := router.NewRouter()
	r.Use(loggingMiddleware)
	r.Use(kitHttp.CreateAuthMiddleware(authInterceptor, func(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
		writeError(w, fmt.Errorf("cannot access to %v: %w", r.RequestURI, err))
	}))
	r.StrictSlash(true)

	// certifica authority sign
	r.HandleFunc(uri.CertificaAuthoritySign, requestHandler.signCertificate).Methods(http.MethodPost)

	// Aliases
	r.HandleFunc(uri.AliasDevice, requestHandler.getDevice).Methods(http.MethodGet)
	r.HandleFunc(uri.AliasDeviceResourceLinks, requestHandler.getDeviceResourceLinks).Methods(http.MethodGet)
	r.HandleFunc(uri.AliasDeviceResources, requestHandler.getDeviceResources).Methods(http.MethodGet)
	r.HandleFunc(uri.AliasDevicePendingCommands, requestHandler.getDevicePendingCommands).Methods(http.MethodGet)

	r.PathPrefix(uri.Devices).Methods(http.MethodGet).MatcherFunc(resourcePendingCommandsMatcher).HandlerFunc(requestHandler.getResourcePendingCommands)
	r.PathPrefix(uri.Devices).Methods(http.MethodGet).MatcherFunc(resourceMatcher).HandlerFunc(requestHandler.getResource)

	// register grpc-proxy handler
	pb.RegisterGrpcGatewayHandlerClient(context.Background(), requestHandler.mux, requestHandler.client.GrpcGatewayClient())

	// ws grpc-proxy
	ws := wsproxy.WebsocketProxy(requestHandler.mux, wsproxy.WithRequestMutator(func(incoming, outgoing *http.Request) *http.Request {
		outgoing.Method = http.MethodPost
		return outgoing
	}), wsproxy.WithLogger(logger{}))
	r.PathPrefix(uri.APIWS + "/").HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ws.ServeHTTP(rw, r)
	})

	// api grpc-proxy
	r.Handle(uri.ClientConfiguration, requestHandler.mux)
	r.PathPrefix(uri.API + "/").HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		requestHandler.mux.ServeHTTP(rw, r)
	})

	// serve www directory
	if requestHandler.config.UI.Enabled {
		r.HandleFunc(uri.OAuthConfiguration, requestHandler.getOAuthConfiguration).Methods(http.MethodGet)
		fs := http.FileServer(http.Dir(requestHandler.config.UI.Directory))
		r.PathPrefix("/").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c := httptest.NewRecorder()
			fs.ServeHTTP(c, r)
			if c.Code == http.StatusNotFound {
				c = httptest.NewRecorder()
				r.URL.Path = "/"
				fs.ServeHTTP(c, r)
			}
			for k, v := range c.HeaderMap {
				w.Header().Set(k, strings.Join(v, ""))
			}
			w.WriteHeader(c.Code)
			c.Body.WriteTo(w)
		}))
	}

	return &http.Server{Handler: r}
}

func (requestHandler *RequestHandler) makeCtx(r *http.Request) context.Context {
	token := getAccessToken(r.Header)
	return kitNetGrpc.CtxWithToken(r.Context(), token)
}

func getAccessToken(h http.Header) string {
	accessToken := h.Get("Authorization")
	if len(accessToken) < 7 {
		return ""
	}
	return accessToken[7:]
}

func getCorrelationID(h http.Header) string {
	correlationID := h.Get("Correlation-ID")
	if correlationID == "" {
		return uuid.New().String()
	}
	return correlationID
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
