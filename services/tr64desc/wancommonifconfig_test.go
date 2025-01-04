// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/tr64desc/wancommonifconfig"
	"log"
	"net/http"
	"testing"
)

var wancommonifconfigMock = &mock.ServiceMock{
	Path:       "/upnp/control/wancommonifconfig1",
	HandleFunc: wancommonifconfigHandler,
}

func TestWANCommonInterfaceConfig(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", wancommonifconfigMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server(), tr064.TR064Spec("tr64desc"))
	client.Debug = true
	serviceClient := &wancommonifconfig.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceName: "WANCommonInterfaceConfig",
			ServiceType: "urn:dslforum-org:service:WANCommonInterfaceConfig:1",
			ServiceId:   "urn:WANCIfConfig-com:serviceId:WANCommonInterfaceConfig1",
			ServiceUrl:  "/upnp/control/wancommonifconfig1",
		},
	}
	{
		out := &wancommonifconfig.GetCommonLinkPropertiesResponse{}
		require.NoError(t, serviceClient.GetCommonLinkProperties(out))
	}
	{
		out := &wancommonifconfig.GetTotalBytesSentResponse{}
		require.NoError(t, serviceClient.GetTotalBytesSent(out))
	}
	{
		out := &wancommonifconfig.GetTotalBytesReceivedResponse{}
		require.NoError(t, serviceClient.GetTotalBytesReceived(out))
	}
	{
		out := &wancommonifconfig.GetTotalPacketsSentResponse{}
		require.NoError(t, serviceClient.GetTotalPacketsSent(out))
	}
	{
		out := &wancommonifconfig.GetTotalPacketsReceivedResponse{}
		require.NoError(t, serviceClient.GetTotalPacketsReceived(out))
	}
	{
		in := &wancommonifconfig.X_AVM_DE_SetWANAccessTypeRequest{}
		require.NoError(t, serviceClient.X_AVM_DE_SetWANAccessType(in))
	}
	{
		out := &wancommonifconfig.X_AVM_DE_GetActiveProviderResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetActiveProvider(out))
	}
	{
		in := &wancommonifconfig.X_AVM_DE_GetOnlineMonitorRequest{}
		out := &wancommonifconfig.X_AVM_DE_GetOnlineMonitorResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetOnlineMonitor(in, out))
	}
}

func wancommonifconfigHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetCommonLinkProperties":
		wancommonifconfig_GetCommonLinkProperties(w)
	case "GetTotalBytesSent":
		wancommonifconfig_GetTotalBytesSent(w)
	case "GetTotalBytesReceived":
		wancommonifconfig_GetTotalBytesReceived(w)
	case "GetTotalPacketsSent":
		wancommonifconfig_GetTotalPacketsSent(w)
	case "GetTotalPacketsReceived":
		wancommonifconfig_GetTotalPacketsReceived(w)
	case "X_AVM-DE_SetWANAccessType":
		wancommonifconfig_X_AVM_DE_SetWANAccessType(w)
	case "X_AVM-DE_GetActiveProvider":
		wancommonifconfig_X_AVM_DE_GetActiveProvider(w)
	case "X_AVM-DE_GetOnlineMonitor":
		wancommonifconfig_X_AVM_DE_GetOnlineMonitor(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func wancommonifconfig_GetCommonLinkProperties(w http.ResponseWriter) {
	out := wancommonifconfig.GetCommonLinkPropertiesResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wancommonifconfig_GetTotalBytesSent(w http.ResponseWriter) {
	out := wancommonifconfig.GetTotalBytesSentResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wancommonifconfig_GetTotalBytesReceived(w http.ResponseWriter) {
	out := wancommonifconfig.GetTotalBytesReceivedResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wancommonifconfig_GetTotalPacketsSent(w http.ResponseWriter) {
	out := wancommonifconfig.GetTotalPacketsSentResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wancommonifconfig_GetTotalPacketsReceived(w http.ResponseWriter) {
	out := wancommonifconfig.GetTotalPacketsReceivedResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wancommonifconfig_X_AVM_DE_SetWANAccessType(w http.ResponseWriter) {
	out := wancommonifconfig.X_AVM_DE_SetWANAccessTypeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wancommonifconfig_X_AVM_DE_GetActiveProvider(w http.ResponseWriter) {
	out := wancommonifconfig.X_AVM_DE_GetActiveProviderResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wancommonifconfig_X_AVM_DE_GetOnlineMonitor(w http.ResponseWriter) {
	out := wancommonifconfig.X_AVM_DE_GetOnlineMonitorResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
