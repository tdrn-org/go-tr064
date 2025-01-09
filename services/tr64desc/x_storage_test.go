// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/tr64desc/x_storage"
	"log"
	"net/http"
	"testing"
)

var x_storageMock = &mock.ServiceMock{
	Path:       "/upnp/control/x_storage",
	HandleFunc: x_storageHandler,
}

func TestX_AVM_DE_Storage(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", x_storageMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server())
	client.Debug = true
	serviceClient := &x_storage.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceSpec:       tr064.ServiceSpec("tr64desc"),
			ServiceType:       "urn:dslforum-org:service:X_AVM-DE_Storage:1",
			ServiceId:         "urn:X_AVM-DE_Storage-com:serviceId:X_AVM-DE_Storage1",
			ServiceControlUrl: "/upnp/control/x_storage",
		},
	}
	{
		out := &x_storage.GetInfoResponse{}
		require.NoError(t, serviceClient.GetInfo(out))
	}
	{
		out := &x_storage.RequestFTPServerWANResponse{}
		require.NoError(t, serviceClient.RequestFTPServerWAN(out))
	}
	{
		in := &x_storage.SetFTPServerRequest{}
		require.NoError(t, serviceClient.SetFTPServer(in))
	}
	{
		in := &x_storage.SetFTPServerWANRequest{}
		require.NoError(t, serviceClient.SetFTPServerWAN(in))
	}
	{
		in := &x_storage.SetSMBServerRequest{}
		require.NoError(t, serviceClient.SetSMBServer(in))
	}
	{
		out := &x_storage.GetUserInfoResponse{}
		require.NoError(t, serviceClient.GetUserInfo(out))
	}
	{
		in := &x_storage.SetUserConfigRequest{}
		require.NoError(t, serviceClient.SetUserConfig(in))
	}
}

func x_storageHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetInfo":
		x_storage_GetInfo(w)
	case "RequestFTPServerWAN":
		x_storage_RequestFTPServerWAN(w)
	case "SetFTPServer":
		x_storage_SetFTPServer(w)
	case "SetFTPServerWAN":
		x_storage_SetFTPServerWAN(w)
	case "SetSMBServer":
		x_storage_SetSMBServer(w)
	case "GetUserInfo":
		x_storage_GetUserInfo(w)
	case "SetUserConfig":
		x_storage_SetUserConfig(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func x_storage_GetInfo(w http.ResponseWriter) {
	out := x_storage.GetInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_storage_RequestFTPServerWAN(w http.ResponseWriter) {
	out := x_storage.RequestFTPServerWANResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_storage_SetFTPServer(w http.ResponseWriter) {
	out := x_storage.SetFTPServerResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_storage_SetFTPServerWAN(w http.ResponseWriter) {
	out := x_storage.SetFTPServerWANResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_storage_SetSMBServer(w http.ResponseWriter) {
	out := x_storage.SetSMBServerResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_storage_GetUserInfo(w http.ResponseWriter) {
	out := x_storage.GetUserInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_storage_SetUserConfig(w http.ResponseWriter) {
	out := x_storage.SetUserConfigResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
