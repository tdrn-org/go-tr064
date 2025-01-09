// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/tr64desc/x_speedtest"
	"log"
	"net/http"
	"testing"
)

var x_speedtestMock = &mock.ServiceMock{
	Path:       "/upnp/control/x_speedtest",
	HandleFunc: x_speedtestHandler,
}

func TestX_AVM_DE_Speedtest(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", x_speedtestMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server())
	client.Debug = true
	serviceClient := &x_speedtest.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceSpec:       tr064.ServiceSpec("tr64desc"),
			ServiceType:       "urn:dslforum-org:service:X_AVM-DE_Speedtest:1",
			ServiceId:         "urn:X_AVM-DE_Speedtest-com:serviceId:X_AVM-DE_Speedtest1",
			ServiceControlUrl: "/upnp/control/x_speedtest",
		},
	}
	{
		out := &x_speedtest.GetInfoResponse{}
		require.NoError(t, serviceClient.GetInfo(out))
	}
	{
		in := &x_speedtest.SetConfigRequest{}
		require.NoError(t, serviceClient.SetConfig(in))
	}
	{
		out := &x_speedtest.GetStatisticsResponse{}
		require.NoError(t, serviceClient.GetStatistics(out))
	}
	{
		require.NoError(t, serviceClient.ResetStatistics())
	}
}

func x_speedtestHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetInfo":
		x_speedtest_GetInfo(w)
	case "SetConfig":
		x_speedtest_SetConfig(w)
	case "GetStatistics":
		x_speedtest_GetStatistics(w)
	case "ResetStatistics":
		x_speedtest_ResetStatistics(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func x_speedtest_GetInfo(w http.ResponseWriter) {
	out := x_speedtest.GetInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_speedtest_SetConfig(w http.ResponseWriter) {
	out := x_speedtest.SetConfigResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_speedtest_GetStatistics(w http.ResponseWriter) {
	out := x_speedtest.GetStatisticsResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_speedtest_ResetStatistics(w http.ResponseWriter) {
	out := x_speedtest.ResetStatisticsResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
