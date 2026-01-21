// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/tr64desc/ethifconfig"
	"log"
	"net/http"
	"testing"
)

var ethifconfigMock = &mock.ServiceMock{
	Path:       "/upnp/control/lanethernetifcfg",
	HandleFunc: ethifconfigHandler,
}

func TestLANEthernetInterfaceConfig(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", ethifconfigMock)
	defer tr064Mock.Stop(t.Context())
	// Actual test
	client := tr064.NewClient(tr064Mock.Server())
	client.Debug = true
	serviceClient := &ethifconfig.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceSpec:       tr064.ServiceSpec("tr64desc"),
			ServiceType:       "urn:dslforum-org:service:LANEthernetInterfaceConfig:1",
			ServiceId:         "urn:LANEthernetIfCfg-com:serviceId:LANEthernetInterfaceConfig1",
			ServiceControlUrl: "/upnp/control/lanethernetifcfg",
		},
	}
	{
		in := &ethifconfig.SetEnableRequest{}
		require.NoError(t, serviceClient.SetEnable(in))
	}
	{
		out := &ethifconfig.GetInfoResponse{}
		require.NoError(t, serviceClient.GetInfo(out))
	}
	{
		out := &ethifconfig.GetStatisticsResponse{}
		require.NoError(t, serviceClient.GetStatistics(out))
	}
}

func ethifconfigHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "SetEnable":
		ethifconfig_SetEnable(w)
	case "GetInfo":
		ethifconfig_GetInfo(w)
	case "GetStatistics":
		ethifconfig_GetStatistics(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func ethifconfig_SetEnable(w http.ResponseWriter) {
	out := ethifconfig.SetEnableResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func ethifconfig_GetInfo(w http.ResponseWriter) {
	out := ethifconfig.GetInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func ethifconfig_GetStatistics(w http.ResponseWriter) {
	out := ethifconfig.GetStatisticsResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
