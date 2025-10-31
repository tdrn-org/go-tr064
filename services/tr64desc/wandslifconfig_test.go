// generated from spec version: 1.0
package services_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/tr64desc/wandslifconfig"
)

var wandslifconfigMock = &mock.ServiceMock{
	Path:       "/upnp/control/wandslifconfig1",
	HandleFunc: wandslifconfigHandler,
}

func TestWANDSLInterfaceConfig(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", wandslifconfigMock)
	defer tr064Mock.Stop(t.Context())
	// Actual test
	client := tr064.NewClient(tr064Mock.Server())
	client.Debug = true
	serviceClient := &wandslifconfig.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceSpec:       tr064.ServiceSpec("tr64desc"),
			ServiceType:       "urn:dslforum-org:service:WANDSLInterfaceConfig:1",
			ServiceId:         "urn:WANDSLIfConfig-com:serviceId:WANDSLInterfaceConfig1",
			ServiceControlUrl: "/upnp/control/wandslifconfig1",
		},
	}
	{
		out := &wandslifconfig.GetInfoResponse{}
		require.NoError(t, serviceClient.GetInfo(out))
	}
	{
		out := &wandslifconfig.GetStatisticsTotalResponse{}
		require.NoError(t, serviceClient.GetStatisticsTotal(out))
	}
	{
		out := &wandslifconfig.X_AVM_DE_GetDSLDiagnoseInfoResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetDSLDiagnoseInfo(out))
	}
	{
		out := &wandslifconfig.X_AVM_DE_GetDSLInfoResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetDSLInfo(out))
	}
}

func wandslifconfigHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetInfo":
		wandslifconfig_GetInfo(w)
	case "GetStatisticsTotal":
		wandslifconfig_GetStatisticsTotal(w)
	case "X_AVM-DE_GetDSLDiagnoseInfo":
		wandslifconfig_X_AVM_DE_GetDSLDiagnoseInfo(w)
	case "X_AVM-DE_GetDSLInfo":
		wandslifconfig_X_AVM_DE_GetDSLInfo(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func wandslifconfig_GetInfo(w http.ResponseWriter) {
	out := wandslifconfig.GetInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wandslifconfig_GetStatisticsTotal(w http.ResponseWriter) {
	out := wandslifconfig.GetStatisticsTotalResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wandslifconfig_X_AVM_DE_GetDSLDiagnoseInfo(w http.ResponseWriter) {
	out := wandslifconfig.X_AVM_DE_GetDSLDiagnoseInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wandslifconfig_X_AVM_DE_GetDSLInfo(w http.ResponseWriter) {
	out := wandslifconfig.X_AVM_DE_GetDSLInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
