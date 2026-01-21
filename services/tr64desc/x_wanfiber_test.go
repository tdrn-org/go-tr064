// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/tr64desc/x_wanfiber"
	"log"
	"net/http"
	"testing"
)

var x_wanfiberMock = &mock.ServiceMock{
	Path:       "/upnp/control/x_wanfiber",
	HandleFunc: x_wanfiberHandler,
}

func TestX_AVM_DE_WANFiber(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", x_wanfiberMock)
	defer tr064Mock.Stop(t.Context())
	// Actual test
	client := tr064.NewClient(tr064Mock.Server())
	client.Debug = true
	serviceClient := &x_wanfiber.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceSpec:       tr064.ServiceSpec("tr64desc"),
			ServiceType:       "urn:dslforum-org:service:X_AVM-DE_WANFiber:1",
			ServiceId:         "urn:X_AVM-DE_WANFiber-com:serviceId:X_AVM-DE_WANFiber1",
			ServiceControlUrl: "/upnp/control/x_wanfiber",
		},
	}
	{
		out := &x_wanfiber.GetInfoResponse{}
		require.NoError(t, serviceClient.GetInfo(out))
	}
	{
		out := &x_wanfiber.GetInfoGPONResponse{}
		require.NoError(t, serviceClient.GetInfoGPON(out))
	}
	{
		out := &x_wanfiber.GetStatisticsResponse{}
		require.NoError(t, serviceClient.GetStatistics(out))
	}
}

func x_wanfiberHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetInfo":
		x_wanfiber_GetInfo(w)
	case "GetInfoGPON":
		x_wanfiber_GetInfoGPON(w)
	case "GetStatistics":
		x_wanfiber_GetStatistics(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func x_wanfiber_GetInfo(w http.ResponseWriter) {
	out := x_wanfiber.GetInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_wanfiber_GetInfoGPON(w http.ResponseWriter) {
	out := x_wanfiber.GetInfoGPONResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_wanfiber_GetStatistics(w http.ResponseWriter) {
	out := x_wanfiber.GetStatisticsResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
