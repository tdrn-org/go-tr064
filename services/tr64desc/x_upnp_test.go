// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/tr64desc/x_upnp"
	"log"
	"net/http"
	"testing"
)

var x_upnpMock = &mock.ServiceMock{
	Path:       "/upnp/control/x_upnp",
	HandleFunc: x_upnpHandler,
}

func TestX_AVM_DE_UPnP(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", x_upnpMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server(), tr064.ServiceSpec("tr64desc"))
	client.Debug = true
	serviceClient := &x_upnp.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceSpec:       tr064.ServiceSpec("tr64desc"),
			ServiceType:       "urn:dslforum-org:service:X_AVM-DE_UPnP:1",
			ServiceId:         "urn:X_AVM-DE_UPnP-com:serviceId:X_AVM-DE_UPnP1",
			ServiceControlUrl: "/upnp/control/x_upnp",
		},
	}
	{
		out := &x_upnp.GetInfoResponse{}
		require.NoError(t, serviceClient.GetInfo(out))
	}
	{
		in := &x_upnp.SetConfigRequest{}
		require.NoError(t, serviceClient.SetConfig(in))
	}
}

func x_upnpHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetInfo":
		x_upnp_GetInfo(w)
	case "SetConfig":
		x_upnp_SetConfig(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func x_upnp_GetInfo(w http.ResponseWriter) {
	out := x_upnp.GetInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_upnp_SetConfig(w http.ResponseWriter) {
	out := x_upnp.SetConfigResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
