package service

import (
	"context"
	"errors"
	"fmt"

	pbAS "github.com/go-ocf/cloud/authorization/pb"
	"github.com/go-ocf/cloud/coap-gateway/coapconv"
	pbCQRS "github.com/go-ocf/cloud/resource-aggregate/pb"
	"github.com/go-ocf/go-coap/v2/message"
	coapCodes "github.com/go-ocf/go-coap/v2/message/codes"
	"github.com/go-ocf/go-coap/v2/mux"
	"github.com/go-ocf/kit/codec/cbor"
	"github.com/go-ocf/kit/log"
	"github.com/go-ocf/kit/net/coap"
	"google.golang.org/grpc/status"
)

type CoapSignInReq struct {
	DeviceID    string `json:"di"`
	UserID      string `json:"uid"`
	AccessToken string `json:"accesstoken"`
	Login       bool   `json:"login"`
}

type CoapSignInResp struct {
	ExpiresIn int64 `json:"expiresin"`
}

func validateSignIn(req CoapSignInReq) error {
	if req.DeviceID == "" {
		return errors.New("cannot sign in to auth server: invalid deviceID")
	}
	if req.AccessToken == "" {
		return errors.New("cannot sign in to auth server: invalid accessToken")
	}
	if req.UserID == "" {
		return errors.New("cannot sign in to auth server: invalid userId")
	}
	return nil
}

// https://github.com/openconnectivityfoundation/security/blob/master/swagger2.0/oic.sec.session.swagger.json
func signInPostHandler(s mux.ResponseWriter, req *mux.Message, client *Client, signIn CoapSignInReq) {
	err := validateSignIn(signIn)
	if err != nil {
		if err != nil {
			client.logAndWriteErrorResponse(fmt.Errorf("cannot handle sign in: %v", err), coapCodes.BadRequest, req.Token)
			return
		}
	}

	resp, err := client.server.asClient.SignIn(req.Context, &pbAS.SignInRequest{
		DeviceId: signIn.DeviceID,
		UserId:   signIn.UserID,
	})
	if err != nil {
		client.logAndWriteErrorResponse(fmt.Errorf("cannot handle sign in: %v", err), coapconv.GrpcCode2CoapCode(status.Convert(err).Code(), coapCodes.POST), req.Token)
		return
	}

	coapResp := CoapSignInResp{
		ExpiresIn: resp.ExpiresIn,
	}

	accept := coap.GetAccept(req.Options)
	encode, err := coap.GetEncoder(accept)
	if err != nil {
		client.logAndWriteErrorResponse(fmt.Errorf("cannot handle sign in: %v", err), coapCodes.InternalServerError, req.Token)
		return
	}
	out, err := encode(coapResp)
	if err != nil {
		client.logAndWriteErrorResponse(fmt.Errorf("cannot handle sign in: %v", err), coapCodes.InternalServerError, req.Token)
		return
	}

	authCtx := authCtx{
		AuthorizationContext: pbCQRS.AuthorizationContext{
			DeviceId: signIn.DeviceID,
			UserId:   signIn.UserID,
		},
		AccessToken: signIn.AccessToken,
	}

	err = client.UpdateCloudDeviceStatus(req.Context, signIn.DeviceID, authCtx.AuthorizationContext, true)
	if err != nil {
		// Events from resources of device will be comes but device is offline. To recover cloud state, client need to reconnect to cloud.
		client.logAndWriteErrorResponse(fmt.Errorf("cannot handle sign in: cannot update cloud device status: %v", err), coapCodes.InternalServerError, req.Token)
		client.Close()
		return
	}

	oldDeviceID := client.storeAuthorizationContext(authCtx)
	newDevice := false

	switch {
	case oldDeviceID == "":
		newDevice = true
	case oldDeviceID != signIn.DeviceID:
		err := client.server.projection.Unregister(oldDeviceID)
		client.server.clientContainerByDeviceID.Remove(oldDeviceID)
		if err != nil {
			client.logAndWriteErrorResponse(fmt.Errorf("cannot handle sign in: %v", err), coapCodes.InternalServerError, req.Token)
			client.Close()
			return
		}
		newDevice = true
	}

	if newDevice {
		client.server.clientContainerByDeviceID.Add(signIn.DeviceID, client)
		loaded, err := client.server.projection.Register(context.Background(), signIn.DeviceID)
		if err != nil {
			client.server.projection.Unregister(signIn.DeviceID)
			client.server.clientContainerByDeviceID.Remove(signIn.DeviceID)

			client.logAndWriteErrorResponse(fmt.Errorf("cannot handle sign in: %v", err), coapCodes.InternalServerError, req.Token)
			client.Close()
			return
		}
		if !loaded {
			models := client.server.projection.Models(signIn.DeviceID, "")
			if len(models) == 0 {
				log.Errorf("cannot load models for deviceID %v", signIn.DeviceID)
			} else {
				for _, r := range models {
					r.(*resourceCtx).TriggerSignIn()
				}
			}
		}
	}
	client.sendResponse(coapCodes.Changed, req.Token, accept, out)
}

func signOutPostHandler(s mux.ResponseWriter, req *mux.Message, client *Client) {
	authCtxOld := client.loadAuthorizationContext()

	if authCtxOld.DeviceId != "" {
		err := client.UpdateCloudDeviceStatus(req.Context, authCtxOld.DeviceId, authCtxOld.AuthorizationContext, false)
		if err != nil {
			// Device will be still reported as online and it can fix his state by next calls online, offline commands.
			log.Errorf("DeviceId %v: cannot handle sign out: cannot update cloud device status: %v", authCtxOld.DeviceId, err)
			return
		}

		client.storeAuthorizationContext(authCtx{})
	}

	client.sendResponse(coapCodes.Changed, req.Token, message.AppOcfCbor, []byte{0xA0}) // empty object
}

// Sign-in
// https://github.com/openconnectivityfoundation/security/blob/master/swagger2.0/oic.sec.session.swagger.json
func signInHandler(s mux.ResponseWriter, req *mux.Message, client *Client) {
	switch req.Code {
	case coapCodes.POST:
		var signIn CoapSignInReq
		err := cbor.ReadFrom(req.Body, &signIn)
		if err != nil {
			client.logAndWriteErrorResponse(fmt.Errorf("cannot handle sign in: %v", err), coapCodes.BadRequest, req.Token)
			return
		}
		switch signIn.Login {
		case true:
			signInPostHandler(s, req, client, signIn)
		default:
			signOutPostHandler(s, req, client)
		}
	default:
		client.logAndWriteErrorResponse(fmt.Errorf("Forbidden request from %v", client.remoteAddrString()), coapCodes.Forbidden, req.Token)
	}
}
