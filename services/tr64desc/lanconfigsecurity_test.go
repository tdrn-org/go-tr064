// generated from spec version: 1.0
package services_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/tr64desc/lanconfigsecurity"
)

var lanconfigsecurityMock = &mock.ServiceMock{
	Path:       "/upnp/control/lanconfigsecurity",
	HandleFunc: lanconfigsecurityHandler,
}

func TestLANConfigSecurity(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", lanconfigsecurityMock)
	defer tr064Mock.Stop(t.Context())
	// Actual test
	client := tr064.NewClient(tr064Mock.Server())
	client.Debug = true
	serviceClient := &lanconfigsecurity.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceSpec:       tr064.ServiceSpec("tr64desc"),
			ServiceType:       "urn:dslforum-org:service:LANConfigSecurity:1",
			ServiceId:         "urn:LANConfigSecurity-com:serviceId:LANConfigSecurity1",
			ServiceControlUrl: "/upnp/control/lanconfigsecurity",
		},
	}
	{
		out := &lanconfigsecurity.GetInfoResponse{}
		require.NoError(t, serviceClient.GetInfo(out))
	}
	{
		out := &lanconfigsecurity.X_AVM_DE_GetCurrentUserResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetCurrentUser(out))
	}
	{
		out := &lanconfigsecurity.X_AVM_DE_GetAnonymousLoginResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetAnonymousLogin(out))
	}
	{
		in := &lanconfigsecurity.SetConfigPasswordRequest{}
		require.NoError(t, serviceClient.SetConfigPassword(in))
	}
	{
		out := &lanconfigsecurity.X_AVM_DE_GetUserListResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetUserList(out))
	}
}

func lanconfigsecurityHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetInfo":
		lanconfigsecurity_GetInfo(w)
	case "X_AVM-DE_GetCurrentUser":
		lanconfigsecurity_X_AVM_DE_GetCurrentUser(w)
	case "X_AVM-DE_GetAnonymousLogin":
		lanconfigsecurity_X_AVM_DE_GetAnonymousLogin(w)
	case "SetConfigPassword":
		lanconfigsecurity_SetConfigPassword(w)
	case "X_AVM-DE_GetUserList":
		lanconfigsecurity_X_AVM_DE_GetUserList(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func lanconfigsecurity_GetInfo(w http.ResponseWriter) {
	out := lanconfigsecurity.GetInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func lanconfigsecurity_X_AVM_DE_GetCurrentUser(w http.ResponseWriter) {
	out := lanconfigsecurity.X_AVM_DE_GetCurrentUserResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func lanconfigsecurity_X_AVM_DE_GetAnonymousLogin(w http.ResponseWriter) {
	out := lanconfigsecurity.X_AVM_DE_GetAnonymousLoginResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func lanconfigsecurity_SetConfigPassword(w http.ResponseWriter) {
	out := lanconfigsecurity.SetConfigPasswordResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func lanconfigsecurity_X_AVM_DE_GetUserList(w http.ResponseWriter) {
	out := lanconfigsecurity.X_AVM_DE_GetUserListResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
