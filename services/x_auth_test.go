// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/x_auth"
	"log"
	"net/http"
	"testing"
)

var x_authMock = &mock.ServiceMock{
	Path:       "/upnp/control/x_auth",
	HandleFunc: x_authHandler,
}

func TestX_AVM_DE_Auth(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", x_authMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server(), mock.User, mock.Password)
	client.Debug = true
	serviceClient := &x_auth.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceName: "X_AVM_DE_Auth",
			ServiceType: "urn:dslforum-org:service:X_AVM-DE_Auth:1",
			ServiceId:   "urn:X_AVM-DE_Auth-com:serviceId:X_AVM-DE_Auth1",
			ServiceUrl:  "/upnp/control/x_auth",
		},
	}
	{
		out := &x_auth.GetInfoResponse{}
		require.NoError(t, serviceClient.GetInfo(out))
	}
	{
		out := &x_auth.GetStateResponse{}
		require.NoError(t, serviceClient.GetState(out))
	}
	{
		in := &x_auth.SetConfigRequest{}
		out := &x_auth.SetConfigResponse{}
		require.NoError(t, serviceClient.SetConfig(in, out))
	}
}

func x_authHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetInfo":
		x_auth_GetInfo(w)
	case "GetState":
		x_auth_GetState(w)
	case "SetConfig":
		x_auth_SetConfig(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func x_auth_GetInfo(w http.ResponseWriter) {
	out := x_auth.GetInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_auth_GetState(w http.ResponseWriter) {
	out := x_auth.GetStateResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_auth_SetConfig(w http.ResponseWriter) {
	out := x_auth.SetConfigResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
