// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/tr64desc/wanpppconn"
	"log"
	"net/http"
	"testing"
)

var wanpppconnMock = &mock.ServiceMock{
	Path:       "/upnp/control/wanpppconn1",
	HandleFunc: wanpppconnHandler,
}

func TestWANPPPConnection(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", wanpppconnMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server(), tr064.TR064Spec("tr64desc"))
	client.Debug = true
	serviceClient := &wanpppconn.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceName: "WANPPPConnection",
			ServiceType: "urn:dslforum-org:service:WANPPPConnection:1",
			ServiceId:   "urn:WANPPPConnection-com:serviceId:WANPPPConnection1",
			ServiceUrl:  "/upnp/control/wanpppconn1",
		},
	}
	{
		out := &wanpppconn.GetInfoResponse{}
		require.NoError(t, serviceClient.GetInfo(out))
	}
	{
		out := &wanpppconn.GetConnectionTypeInfoResponse{}
		require.NoError(t, serviceClient.GetConnectionTypeInfo(out))
	}
	{
		in := &wanpppconn.SetConnectionTypeRequest{}
		require.NoError(t, serviceClient.SetConnectionType(in))
	}
	{
		out := &wanpppconn.GetStatusInfoResponse{}
		require.NoError(t, serviceClient.GetStatusInfo(out))
	}
	{
		out := &wanpppconn.GetUserNameResponse{}
		require.NoError(t, serviceClient.GetUserName(out))
	}
	{
		in := &wanpppconn.SetUserNameRequest{}
		require.NoError(t, serviceClient.SetUserName(in))
	}
	{
		in := &wanpppconn.SetPasswordRequest{}
		require.NoError(t, serviceClient.SetPassword(in))
	}
	{
		out := &wanpppconn.GetNATRSIPStatusResponse{}
		require.NoError(t, serviceClient.GetNATRSIPStatus(out))
	}
	{
		in := &wanpppconn.SetConnectionTriggerRequest{}
		require.NoError(t, serviceClient.SetConnectionTrigger(in))
	}
	{
		require.NoError(t, serviceClient.ForceTermination())
	}
	{
		require.NoError(t, serviceClient.RequestConnection())
	}
	{
		in := &wanpppconn.GetGenericPortMappingEntryRequest{}
		out := &wanpppconn.GetGenericPortMappingEntryResponse{}
		require.NoError(t, serviceClient.GetGenericPortMappingEntry(in, out))
	}
	{
		in := &wanpppconn.GetSpecificPortMappingEntryRequest{}
		out := &wanpppconn.GetSpecificPortMappingEntryResponse{}
		require.NoError(t, serviceClient.GetSpecificPortMappingEntry(in, out))
	}
	{
		in := &wanpppconn.AddPortMappingRequest{}
		require.NoError(t, serviceClient.AddPortMapping(in))
	}
	{
		in := &wanpppconn.DeletePortMappingRequest{}
		require.NoError(t, serviceClient.DeletePortMapping(in))
	}
	{
		out := &wanpppconn.GetExternalIPAddressResponse{}
		require.NoError(t, serviceClient.GetExternalIPAddress(out))
	}
	{
		out := &wanpppconn.X_GetDNSServersResponse{}
		require.NoError(t, serviceClient.X_GetDNSServers(out))
	}
	{
		out := &wanpppconn.GetLinkLayerMaxBitRatesResponse{}
		require.NoError(t, serviceClient.GetLinkLayerMaxBitRates(out))
	}
	{
		out := &wanpppconn.GetPortMappingNumberOfEntriesResponse{}
		require.NoError(t, serviceClient.GetPortMappingNumberOfEntries(out))
	}
	{
		in := &wanpppconn.SetRouteProtocolRxRequest{}
		require.NoError(t, serviceClient.SetRouteProtocolRx(in))
	}
	{
		in := &wanpppconn.SetIdleDisconnectTimeRequest{}
		require.NoError(t, serviceClient.SetIdleDisconnectTime(in))
	}
	{
		out := &wanpppconn.X_AVM_DE_GetAutoDisconnectTimeSpanResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetAutoDisconnectTimeSpan(out))
	}
	{
		in := &wanpppconn.X_AVM_DE_SetAutoDisconnectTimeSpanRequest{}
		require.NoError(t, serviceClient.X_AVM_DE_SetAutoDisconnectTimeSpan(in))
	}
}

func wanpppconnHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetInfo":
		wanpppconn_GetInfo(w)
	case "GetConnectionTypeInfo":
		wanpppconn_GetConnectionTypeInfo(w)
	case "SetConnectionType":
		wanpppconn_SetConnectionType(w)
	case "GetStatusInfo":
		wanpppconn_GetStatusInfo(w)
	case "GetUserName":
		wanpppconn_GetUserName(w)
	case "SetUserName":
		wanpppconn_SetUserName(w)
	case "SetPassword":
		wanpppconn_SetPassword(w)
	case "GetNATRSIPStatus":
		wanpppconn_GetNATRSIPStatus(w)
	case "SetConnectionTrigger":
		wanpppconn_SetConnectionTrigger(w)
	case "ForceTermination":
		wanpppconn_ForceTermination(w)
	case "RequestConnection":
		wanpppconn_RequestConnection(w)
	case "GetGenericPortMappingEntry":
		wanpppconn_GetGenericPortMappingEntry(w)
	case "GetSpecificPortMappingEntry":
		wanpppconn_GetSpecificPortMappingEntry(w)
	case "AddPortMapping":
		wanpppconn_AddPortMapping(w)
	case "DeletePortMapping":
		wanpppconn_DeletePortMapping(w)
	case "GetExternalIPAddress":
		wanpppconn_GetExternalIPAddress(w)
	case "X_GetDNSServers":
		wanpppconn_X_GetDNSServers(w)
	case "GetLinkLayerMaxBitRates":
		wanpppconn_GetLinkLayerMaxBitRates(w)
	case "GetPortMappingNumberOfEntries":
		wanpppconn_GetPortMappingNumberOfEntries(w)
	case "SetRouteProtocolRx":
		wanpppconn_SetRouteProtocolRx(w)
	case "SetIdleDisconnectTime":
		wanpppconn_SetIdleDisconnectTime(w)
	case "X_AVM-DE_GetAutoDisconnectTimeSpan":
		wanpppconn_X_AVM_DE_GetAutoDisconnectTimeSpan(w)
	case "X_AVM-DE_SetAutoDisconnectTimeSpan":
		wanpppconn_X_AVM_DE_SetAutoDisconnectTimeSpan(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func wanpppconn_GetInfo(w http.ResponseWriter) {
	out := wanpppconn.GetInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanpppconn_GetConnectionTypeInfo(w http.ResponseWriter) {
	out := wanpppconn.GetConnectionTypeInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanpppconn_SetConnectionType(w http.ResponseWriter) {
	out := wanpppconn.SetConnectionTypeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanpppconn_GetStatusInfo(w http.ResponseWriter) {
	out := wanpppconn.GetStatusInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanpppconn_GetUserName(w http.ResponseWriter) {
	out := wanpppconn.GetUserNameResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanpppconn_SetUserName(w http.ResponseWriter) {
	out := wanpppconn.SetUserNameResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanpppconn_SetPassword(w http.ResponseWriter) {
	out := wanpppconn.SetPasswordResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanpppconn_GetNATRSIPStatus(w http.ResponseWriter) {
	out := wanpppconn.GetNATRSIPStatusResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanpppconn_SetConnectionTrigger(w http.ResponseWriter) {
	out := wanpppconn.SetConnectionTriggerResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanpppconn_ForceTermination(w http.ResponseWriter) {
	out := wanpppconn.ForceTerminationResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanpppconn_RequestConnection(w http.ResponseWriter) {
	out := wanpppconn.RequestConnectionResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanpppconn_GetGenericPortMappingEntry(w http.ResponseWriter) {
	out := wanpppconn.GetGenericPortMappingEntryResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanpppconn_GetSpecificPortMappingEntry(w http.ResponseWriter) {
	out := wanpppconn.GetSpecificPortMappingEntryResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanpppconn_AddPortMapping(w http.ResponseWriter) {
	out := wanpppconn.AddPortMappingResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanpppconn_DeletePortMapping(w http.ResponseWriter) {
	out := wanpppconn.DeletePortMappingResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanpppconn_GetExternalIPAddress(w http.ResponseWriter) {
	out := wanpppconn.GetExternalIPAddressResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanpppconn_X_GetDNSServers(w http.ResponseWriter) {
	out := wanpppconn.X_GetDNSServersResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanpppconn_GetLinkLayerMaxBitRates(w http.ResponseWriter) {
	out := wanpppconn.GetLinkLayerMaxBitRatesResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanpppconn_GetPortMappingNumberOfEntries(w http.ResponseWriter) {
	out := wanpppconn.GetPortMappingNumberOfEntriesResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanpppconn_SetRouteProtocolRx(w http.ResponseWriter) {
	out := wanpppconn.SetRouteProtocolRxResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanpppconn_SetIdleDisconnectTime(w http.ResponseWriter) {
	out := wanpppconn.SetIdleDisconnectTimeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanpppconn_X_AVM_DE_GetAutoDisconnectTimeSpan(w http.ResponseWriter) {
	out := wanpppconn.X_AVM_DE_GetAutoDisconnectTimeSpanResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanpppconn_X_AVM_DE_SetAutoDisconnectTimeSpan(w http.ResponseWriter) {
	out := wanpppconn.X_AVM_DE_SetAutoDisconnectTimeSpanResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
