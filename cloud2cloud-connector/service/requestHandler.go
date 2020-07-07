package service

import (
	"net/http"
	"net/url"
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/go-ocf/cloud/cloud2cloud-connector/events"
	"github.com/go-ocf/cloud/cloud2cloud-connector/store"
	"github.com/go-ocf/cloud/cloud2cloud-connector/uri"
	"github.com/go-ocf/kit/log"

	router "github.com/gorilla/mux"

	pbAS "github.com/go-ocf/cloud/authorization/pb"
	pbRA "github.com/go-ocf/cloud/resource-aggregate/pb"
)

const cloudIDKey = "CloudId"
const accountIDKey = "AccountId"

type provisionCacheData struct {
	linkedAccount store.LinkedAccount
	linkedCloud   store.LinkedCloud
}

//RequestHandler for handling incoming request
type RequestHandler struct {
	oauthCallback string
	store         *Store

	asClient pbAS.AuthorizationServiceClient
	raClient pbRA.ResourceAggregateClient

	provisionCache *cache.Cache
	subManager     *SubscriptionManager
	triggerTask    func(Task)
}

func logAndWriteErrorResponse(err error, statusCode int, w http.ResponseWriter) {
	log.Errorf("%v", err)
	w.Header().Set(events.ContentTypeKey, "text/plain")
	w.WriteHeader(statusCode)
	w.Write([]byte(err.Error()))
}

//NewRequestHandler factory for new RequestHandler
func NewRequestHandler(
	oauthCallback string,
	subManager *SubscriptionManager,
	asClient pbAS.AuthorizationServiceClient,
	raClient pbRA.ResourceAggregateClient,
	store *Store,
	triggerTask func(Task),
) *RequestHandler {
	return &RequestHandler{
		oauthCallback:  oauthCallback,
		subManager:     subManager,
		asClient:       asClient,
		raClient:       raClient,
		store:          store,
		provisionCache: cache.New(5*time.Minute, 10*time.Minute),
		triggerTask:    triggerTask,
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("%v %v %+v", r.Method, r.RequestURI, r.Header)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// NewHTTP returns HTTP server
func NewHTTP(requestHandler *RequestHandler) *http.Server {
	r := router.NewRouter()
	r.StrictSlash(true)
	r.Use(loggingMiddleware)

	// health check
	r.HandleFunc("/", healthCheck).Methods("GET")

	// retrieve all linked clouds
	r.HandleFunc(uri.LinkedClouds, requestHandler.RetrieveLinkedClouds).Methods("GET")
	// add linked cloud
	r.HandleFunc(uri.LinkedClouds, requestHandler.AddLinkedCloud).Methods("POST")
	// delete linked cloud
	r.HandleFunc(uri.LinkedCloud, requestHandler.DeleteLinkedCloud).Methods("DELETE")
	// add linked account
	r.HandleFunc(uri.LinkedAccounts, requestHandler.AddLinkedAccount).Methods("GET")
	// delete linked cloud
	r.HandleFunc(uri.LinkedAccount, requestHandler.DeleteLinkedAccount).Methods("DELETE")
	// notify linked cloud
	r.HandleFunc(uri.Events, requestHandler.ProcessEvent).Methods("POST")

	// OAuthCallback
	oauthURL, _ := url.Parse(requestHandler.oauthCallback)
	r.HandleFunc(oauthURL.Path, requestHandler.OAuthCallback).Methods("GET")

	return &http.Server{Handler: r}
}
