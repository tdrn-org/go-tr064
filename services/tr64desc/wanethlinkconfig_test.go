// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/tr64desc/wanethlinkconfig"
	"log"
	"net/http"
	"testing"
)

var wanethlinkconfigMock = &mock.ServiceMock{
	Path:       "/upnp/control/wanethlinkconfig1",
	HandleFunc: wanethlinkconfigHandler,
}

func TestWANEthernetLinkConfig(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", wanethlinkconfigMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server(), tr064.TR064Spec("tr64desc"))
	client.Debug = true
	serviceClient := &wanethlinkconfig.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceName: "WANEthernetLinkConfig",
			ServiceType: "urn:dslforum-org:service:WANEthernetLinkConfig:1",
			ServiceId:   "urn:WANEthernetLinkConfig-com:serviceId:WANEthernetLinkConfig1",
			ServiceUrl:  "/upnp/control/wanethlinkconfig1",
		},
	}
	{
		out := &wanethlinkconfig.GetEthernetLinkStatusResponse{}
		require.NoError(t, serviceClient.GetEthernetLinkStatus(out))
	}
}

func wanethlinkconfigHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetEthernetLinkStatus":
		wanethlinkconfig_GetEthernetLinkStatus(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func wanethlinkconfig_GetEthernetLinkStatus(w http.ResponseWriter) {
	out := wanethlinkconfig.GetEthernetLinkStatusResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
