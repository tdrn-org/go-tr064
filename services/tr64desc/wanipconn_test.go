// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/tr64desc/wanipconn"
	"log"
	"net/http"
	"testing"
)

var wanipconnMock = &mock.ServiceMock{
	Path:       "/upnp/control/wanipconnection1",
	HandleFunc: wanipconnHandler,
}

func TestWANIPConnection(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", wanipconnMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server(), tr064.ServiceSpec("tr64desc"))
	client.Debug = true
	serviceClient := &wanipconn.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceSpec: tr064.ServiceSpec("tr64desc"),
			ServiceId:   "urn:WANIPConnection-com:serviceId:WANIPConnection1",
			ServiceUrl:  "/upnp/control/wanipconnection1",
		},
	}
	{
		out := &wanipconn.GetInfoResponse{}
		require.NoError(t, serviceClient.GetInfo(out))
	}
	{
		out := &wanipconn.GetConnectionTypeInfoResponse{}
		require.NoError(t, serviceClient.GetConnectionTypeInfo(out))
	}
	{
		in := &wanipconn.SetConnectionTypeRequest{}
		require.NoError(t, serviceClient.SetConnectionType(in))
	}
	{
		out := &wanipconn.GetStatusInfoResponse{}
		require.NoError(t, serviceClient.GetStatusInfo(out))
	}
	{
		out := &wanipconn.GetNATRSIPStatusResponse{}
		require.NoError(t, serviceClient.GetNATRSIPStatus(out))
	}
	{
		in := &wanipconn.SetConnectionTriggerRequest{}
		require.NoError(t, serviceClient.SetConnectionTrigger(in))
	}
	{
		require.NoError(t, serviceClient.ForceTermination())
	}
	{
		require.NoError(t, serviceClient.RequestConnection())
	}
	{
		in := &wanipconn.GetGenericPortMappingEntryRequest{}
		out := &wanipconn.GetGenericPortMappingEntryResponse{}
		require.NoError(t, serviceClient.GetGenericPortMappingEntry(in, out))
	}
	{
		in := &wanipconn.GetSpecificPortMappingEntryRequest{}
		out := &wanipconn.GetSpecificPortMappingEntryResponse{}
		require.NoError(t, serviceClient.GetSpecificPortMappingEntry(in, out))
	}
	{
		in := &wanipconn.AddPortMappingRequest{}
		require.NoError(t, serviceClient.AddPortMapping(in))
	}
	{
		in := &wanipconn.DeletePortMappingRequest{}
		require.NoError(t, serviceClient.DeletePortMapping(in))
	}
	{
		out := &wanipconn.GetExternalIPAddressResponse{}
		require.NoError(t, serviceClient.GetExternalIPAddress(out))
	}
	{
		out := &wanipconn.X_GetDNSServersResponse{}
		require.NoError(t, serviceClient.X_GetDNSServers(out))
	}
	{
		out := &wanipconn.GetPortMappingNumberOfEntriesResponse{}
		require.NoError(t, serviceClient.GetPortMappingNumberOfEntries(out))
	}
	{
		in := &wanipconn.SetRouteProtocolRxRequest{}
		require.NoError(t, serviceClient.SetRouteProtocolRx(in))
	}
	{
		in := &wanipconn.SetIdleDisconnectTimeRequest{}
		require.NoError(t, serviceClient.SetIdleDisconnectTime(in))
	}
}

func wanipconnHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetInfo":
		wanipconn_GetInfo(w)
	case "GetConnectionTypeInfo":
		wanipconn_GetConnectionTypeInfo(w)
	case "SetConnectionType":
		wanipconn_SetConnectionType(w)
	case "GetStatusInfo":
		wanipconn_GetStatusInfo(w)
	case "GetNATRSIPStatus":
		wanipconn_GetNATRSIPStatus(w)
	case "SetConnectionTrigger":
		wanipconn_SetConnectionTrigger(w)
	case "ForceTermination":
		wanipconn_ForceTermination(w)
	case "RequestConnection":
		wanipconn_RequestConnection(w)
	case "GetGenericPortMappingEntry":
		wanipconn_GetGenericPortMappingEntry(w)
	case "GetSpecificPortMappingEntry":
		wanipconn_GetSpecificPortMappingEntry(w)
	case "AddPortMapping":
		wanipconn_AddPortMapping(w)
	case "DeletePortMapping":
		wanipconn_DeletePortMapping(w)
	case "GetExternalIPAddress":
		wanipconn_GetExternalIPAddress(w)
	case "X_GetDNSServers":
		wanipconn_X_GetDNSServers(w)
	case "GetPortMappingNumberOfEntries":
		wanipconn_GetPortMappingNumberOfEntries(w)
	case "SetRouteProtocolRx":
		wanipconn_SetRouteProtocolRx(w)
	case "SetIdleDisconnectTime":
		wanipconn_SetIdleDisconnectTime(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func wanipconn_GetInfo(w http.ResponseWriter) {
	out := wanipconn.GetInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanipconn_GetConnectionTypeInfo(w http.ResponseWriter) {
	out := wanipconn.GetConnectionTypeInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanipconn_SetConnectionType(w http.ResponseWriter) {
	out := wanipconn.SetConnectionTypeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanipconn_GetStatusInfo(w http.ResponseWriter) {
	out := wanipconn.GetStatusInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanipconn_GetNATRSIPStatus(w http.ResponseWriter) {
	out := wanipconn.GetNATRSIPStatusResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanipconn_SetConnectionTrigger(w http.ResponseWriter) {
	out := wanipconn.SetConnectionTriggerResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanipconn_ForceTermination(w http.ResponseWriter) {
	out := wanipconn.ForceTerminationResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanipconn_RequestConnection(w http.ResponseWriter) {
	out := wanipconn.RequestConnectionResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanipconn_GetGenericPortMappingEntry(w http.ResponseWriter) {
	out := wanipconn.GetGenericPortMappingEntryResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanipconn_GetSpecificPortMappingEntry(w http.ResponseWriter) {
	out := wanipconn.GetSpecificPortMappingEntryResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanipconn_AddPortMapping(w http.ResponseWriter) {
	out := wanipconn.AddPortMappingResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanipconn_DeletePortMapping(w http.ResponseWriter) {
	out := wanipconn.DeletePortMappingResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanipconn_GetExternalIPAddress(w http.ResponseWriter) {
	out := wanipconn.GetExternalIPAddressResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanipconn_X_GetDNSServers(w http.ResponseWriter) {
	out := wanipconn.X_GetDNSServersResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanipconn_GetPortMappingNumberOfEntries(w http.ResponseWriter) {
	out := wanipconn.GetPortMappingNumberOfEntriesResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanipconn_SetRouteProtocolRx(w http.ResponseWriter) {
	out := wanipconn.SetRouteProtocolRxResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wanipconn_SetIdleDisconnectTime(w http.ResponseWriter) {
	out := wanipconn.SetIdleDisconnectTimeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
