// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/tr64desc/x_webdav"
	"log"
	"net/http"
	"testing"
)

var x_webdavMock = &mock.ServiceMock{
	Path:       "/upnp/control/x_webdav",
	HandleFunc: x_webdavHandler,
}

func TestX_AVM_DE_WebDAVClient(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", x_webdavMock)
	defer tr064Mock.Stop(t.Context())
	// Actual test
	client := tr064.NewClient(tr064Mock.Server())
	client.Debug = true
	serviceClient := &x_webdav.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceSpec:       tr064.ServiceSpec("tr64desc"),
			ServiceType:       "urn:dslforum-org:service:X_AVM-DE_WebDAVClient:1",
			ServiceId:         "urn:X_AVM-DE_WebDAV-com:serviceId:X_AVM-DE_WebDAVClient1",
			ServiceControlUrl: "/upnp/control/x_webdav",
		},
	}
	{
		out := &x_webdav.GetInfoResponse{}
		require.NoError(t, serviceClient.GetInfo(out))
	}
	{
		in := &x_webdav.SetConfigRequest{}
		require.NoError(t, serviceClient.SetConfig(in))
	}
}

func x_webdavHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetInfo":
		x_webdav_GetInfo(w)
	case "SetConfig":
		x_webdav_SetConfig(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func x_webdav_GetInfo(w http.ResponseWriter) {
	out := x_webdav.GetInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_webdav_SetConfig(w http.ResponseWriter) {
	out := x_webdav.SetConfigResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
