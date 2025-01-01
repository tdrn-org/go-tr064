// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/x_remote"
	"log"
	"net/http"
	"testing"
)

var x_remoteMock = &mock.ServiceMock{
	Path:       "/upnp/control/x_remote",
	HandleFunc: x_remoteHandler,
}

func TestX_AVM_DE_RemoteAccess(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", x_remoteMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server(), mock.User, mock.Password)
	client.Debug = true
	serviceClient := &x_remote.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceName: "X_AVM_DE_RemoteAccess",
			ServiceType: "urn:dslforum-org:service:X_AVM-DE_RemoteAccess:1",
			ServiceId:   "urn:X_AVM-DE_RemoteAccess-com:serviceId:X_AVM-DE_RemoteAccess1",
			ServiceUrl:  "/upnp/control/x_remote",
		},
	}
	{
		out := &x_remote.GetInfoResponse{}
		require.NoError(t, serviceClient.GetInfo(out))
	}
	{
		in := &x_remote.SetConfigRequest{}
		require.NoError(t, serviceClient.SetConfig(in))
	}
	{
		in := &x_remote.SetEnableRequest{}
		out := &x_remote.SetEnableResponse{}
		require.NoError(t, serviceClient.SetEnable(in, out))
	}
	{
		in := &x_remote.SetLetsEncryptEnableRequest{}
		require.NoError(t, serviceClient.SetLetsEncryptEnable(in))
	}
	{
		out := &x_remote.GetDDNSInfoResponse{}
		require.NoError(t, serviceClient.GetDDNSInfo(out))
	}
	{
		out := &x_remote.GetDDNSProvidersResponse{}
		require.NoError(t, serviceClient.GetDDNSProviders(out))
	}
	{
		in := &x_remote.SetDDNSConfigRequest{}
		require.NoError(t, serviceClient.SetDDNSConfig(in))
	}
}

func x_remoteHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetInfo":
		x_remote_GetInfo(w)
	case "SetConfig":
		x_remote_SetConfig(w)
	case "SetEnable":
		x_remote_SetEnable(w)
	case "SetLetsEncryptEnable":
		x_remote_SetLetsEncryptEnable(w)
	case "GetDDNSInfo":
		x_remote_GetDDNSInfo(w)
	case "GetDDNSProviders":
		x_remote_GetDDNSProviders(w)
	case "SetDDNSConfig":
		x_remote_SetDDNSConfig(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func x_remote_GetInfo(w http.ResponseWriter) {
	out := x_remote.GetInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_remote_SetConfig(w http.ResponseWriter) {
	out := x_remote.SetConfigResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_remote_SetEnable(w http.ResponseWriter) {
	out := x_remote.SetEnableResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_remote_SetLetsEncryptEnable(w http.ResponseWriter) {
	out := x_remote.SetLetsEncryptEnableResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_remote_GetDDNSInfo(w http.ResponseWriter) {
	out := x_remote.GetDDNSInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_remote_GetDDNSProviders(w http.ResponseWriter) {
	out := x_remote.GetDDNSProvidersResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_remote_SetDDNSConfig(w http.ResponseWriter) {
	out := x_remote.SetDDNSConfigResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
