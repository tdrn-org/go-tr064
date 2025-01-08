// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/igddesc/igdicfg"
	"log"
	"net/http"
	"testing"
)

var igdicfgMock = &mock.ServiceMock{
	Path:       "/igdupnp/control/WANCommonIFC1",
	HandleFunc: igdicfgHandler,
}

func TestWANCommonInterfaceConfig(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", igdicfgMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server(), tr064.ServiceSpec("igddesc"))
	client.Debug = true
	serviceClient := &igdicfg.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceSpec:       tr064.ServiceSpec("igddesc"),
			ServiceType:       "urn:schemas-upnp-org:service:WANCommonInterfaceConfig:1",
			ServiceId:         "urn:upnp-org:serviceId:WANCommonIFC1",
			ServiceControlUrl: "/igdupnp/control/WANCommonIFC1",
		},
	}
	{
		out := &igdicfg.GetCommonLinkPropertiesResponse{}
		require.NoError(t, serviceClient.GetCommonLinkProperties(out))
	}
	{
		out := &igdicfg.GetTotalBytesSentResponse{}
		require.NoError(t, serviceClient.GetTotalBytesSent(out))
	}
	{
		out := &igdicfg.GetTotalBytesReceivedResponse{}
		require.NoError(t, serviceClient.GetTotalBytesReceived(out))
	}
	{
		out := &igdicfg.GetTotalPacketsSentResponse{}
		require.NoError(t, serviceClient.GetTotalPacketsSent(out))
	}
	{
		out := &igdicfg.GetTotalPacketsReceivedResponse{}
		require.NoError(t, serviceClient.GetTotalPacketsReceived(out))
	}
	{
		out := &igdicfg.GetAddonInfosResponse{}
		require.NoError(t, serviceClient.GetAddonInfos(out))
	}
	{
		out := &igdicfg.X_AVM_DE_GetDsliteStatusResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetDsliteStatus(out))
	}
	{
		out := &igdicfg.X_AVM_DE_GetIPTVInfosResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetIPTVInfos(out))
	}
}

func igdicfgHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetCommonLinkProperties":
		igdicfg_GetCommonLinkProperties(w)
	case "GetTotalBytesSent":
		igdicfg_GetTotalBytesSent(w)
	case "GetTotalBytesReceived":
		igdicfg_GetTotalBytesReceived(w)
	case "GetTotalPacketsSent":
		igdicfg_GetTotalPacketsSent(w)
	case "GetTotalPacketsReceived":
		igdicfg_GetTotalPacketsReceived(w)
	case "GetAddonInfos":
		igdicfg_GetAddonInfos(w)
	case "X_AVM_DE_GetDsliteStatus":
		igdicfg_X_AVM_DE_GetDsliteStatus(w)
	case "X_AVM_DE_GetIPTVInfos":
		igdicfg_X_AVM_DE_GetIPTVInfos(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func igdicfg_GetCommonLinkProperties(w http.ResponseWriter) {
	out := igdicfg.GetCommonLinkPropertiesResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igdicfg_GetTotalBytesSent(w http.ResponseWriter) {
	out := igdicfg.GetTotalBytesSentResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igdicfg_GetTotalBytesReceived(w http.ResponseWriter) {
	out := igdicfg.GetTotalBytesReceivedResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igdicfg_GetTotalPacketsSent(w http.ResponseWriter) {
	out := igdicfg.GetTotalPacketsSentResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igdicfg_GetTotalPacketsReceived(w http.ResponseWriter) {
	out := igdicfg.GetTotalPacketsReceivedResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igdicfg_GetAddonInfos(w http.ResponseWriter) {
	out := igdicfg.GetAddonInfosResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igdicfg_X_AVM_DE_GetDsliteStatus(w http.ResponseWriter) {
	out := igdicfg.X_AVM_DE_GetDsliteStatusResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igdicfg_X_AVM_DE_GetIPTVInfos(w http.ResponseWriter) {
	out := igdicfg.X_AVM_DE_GetIPTVInfosResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
