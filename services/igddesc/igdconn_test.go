// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/igddesc/igdconn"
	"log"
	"net/http"
	"testing"
)

var igdconnMock = &mock.ServiceMock{
	Path:       "/igdupnp/control/WANIPConn1",
	HandleFunc: igdconnHandler,
}

func TestWANIPConnection(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", igdconnMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server(), tr064.TR064Spec("igddesc"))
	client.Debug = true
	serviceClient := &igdconn.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceName: "WANIPConnection",
			ServiceType: "urn:schemas-upnp-org:service:WANIPConnection:1",
			ServiceId:   "urn:upnp-org:serviceId:WANIPConn1",
			ServiceUrl:  "/igdupnp/control/WANIPConn1",
		},
	}
	{
		in := &igdconn.SetConnectionTypeRequest{}
		require.NoError(t, serviceClient.SetConnectionType(in))
	}
	{
		out := &igdconn.GetConnectionTypeInfoResponse{}
		require.NoError(t, serviceClient.GetConnectionTypeInfo(out))
	}
	{
		out := &igdconn.GetAutoDisconnectTimeResponse{}
		require.NoError(t, serviceClient.GetAutoDisconnectTime(out))
	}
	{
		out := &igdconn.GetIdleDisconnectTimeResponse{}
		require.NoError(t, serviceClient.GetIdleDisconnectTime(out))
	}
	{
		require.NoError(t, serviceClient.RequestConnection())
	}
	{
		require.NoError(t, serviceClient.RequestTermination())
	}
	{
		require.NoError(t, serviceClient.ForceTermination())
	}
	{
		out := &igdconn.GetStatusInfoResponse{}
		require.NoError(t, serviceClient.GetStatusInfo(out))
	}
	{
		out := &igdconn.GetNATRSIPStatusResponse{}
		require.NoError(t, serviceClient.GetNATRSIPStatus(out))
	}
	{
		in := &igdconn.GetGenericPortMappingEntryRequest{}
		out := &igdconn.GetGenericPortMappingEntryResponse{}
		require.NoError(t, serviceClient.GetGenericPortMappingEntry(in, out))
	}
	{
		in := &igdconn.GetSpecificPortMappingEntryRequest{}
		out := &igdconn.GetSpecificPortMappingEntryResponse{}
		require.NoError(t, serviceClient.GetSpecificPortMappingEntry(in, out))
	}
	{
		in := &igdconn.AddPortMappingRequest{}
		require.NoError(t, serviceClient.AddPortMapping(in))
	}
	{
		in := &igdconn.DeletePortMappingRequest{}
		require.NoError(t, serviceClient.DeletePortMapping(in))
	}
	{
		out := &igdconn.GetExternalIPAddressResponse{}
		require.NoError(t, serviceClient.GetExternalIPAddress(out))
	}
	{
		out := &igdconn.X_AVM_DE_GetExternalIPv6AddressResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetExternalIPv6Address(out))
	}
	{
		out := &igdconn.X_AVM_DE_GetIPv6PrefixResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetIPv6Prefix(out))
	}
	{
		out := &igdconn.X_AVM_DE_GetDNSServerResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetDNSServer(out))
	}
	{
		out := &igdconn.X_AVM_DE_GetIPv6DNSServerResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetIPv6DNSServer(out))
	}
}

func igdconnHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "SetConnectionType":
		igdconn_SetConnectionType(w)
	case "GetConnectionTypeInfo":
		igdconn_GetConnectionTypeInfo(w)
	case "GetAutoDisconnectTime":
		igdconn_GetAutoDisconnectTime(w)
	case "GetIdleDisconnectTime":
		igdconn_GetIdleDisconnectTime(w)
	case "RequestConnection":
		igdconn_RequestConnection(w)
	case "RequestTermination":
		igdconn_RequestTermination(w)
	case "ForceTermination":
		igdconn_ForceTermination(w)
	case "GetStatusInfo":
		igdconn_GetStatusInfo(w)
	case "GetNATRSIPStatus":
		igdconn_GetNATRSIPStatus(w)
	case "GetGenericPortMappingEntry":
		igdconn_GetGenericPortMappingEntry(w)
	case "GetSpecificPortMappingEntry":
		igdconn_GetSpecificPortMappingEntry(w)
	case "AddPortMapping":
		igdconn_AddPortMapping(w)
	case "DeletePortMapping":
		igdconn_DeletePortMapping(w)
	case "GetExternalIPAddress":
		igdconn_GetExternalIPAddress(w)
	case "X_AVM_DE_GetExternalIPv6Address":
		igdconn_X_AVM_DE_GetExternalIPv6Address(w)
	case "X_AVM_DE_GetIPv6Prefix":
		igdconn_X_AVM_DE_GetIPv6Prefix(w)
	case "X_AVM_DE_GetDNSServer":
		igdconn_X_AVM_DE_GetDNSServer(w)
	case "X_AVM_DE_GetIPv6DNSServer":
		igdconn_X_AVM_DE_GetIPv6DNSServer(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func igdconn_SetConnectionType(w http.ResponseWriter) {
	out := igdconn.SetConnectionTypeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igdconn_GetConnectionTypeInfo(w http.ResponseWriter) {
	out := igdconn.GetConnectionTypeInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igdconn_GetAutoDisconnectTime(w http.ResponseWriter) {
	out := igdconn.GetAutoDisconnectTimeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igdconn_GetIdleDisconnectTime(w http.ResponseWriter) {
	out := igdconn.GetIdleDisconnectTimeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igdconn_RequestConnection(w http.ResponseWriter) {
	out := igdconn.RequestConnectionResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igdconn_RequestTermination(w http.ResponseWriter) {
	out := igdconn.RequestTerminationResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igdconn_ForceTermination(w http.ResponseWriter) {
	out := igdconn.ForceTerminationResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igdconn_GetStatusInfo(w http.ResponseWriter) {
	out := igdconn.GetStatusInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igdconn_GetNATRSIPStatus(w http.ResponseWriter) {
	out := igdconn.GetNATRSIPStatusResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igdconn_GetGenericPortMappingEntry(w http.ResponseWriter) {
	out := igdconn.GetGenericPortMappingEntryResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igdconn_GetSpecificPortMappingEntry(w http.ResponseWriter) {
	out := igdconn.GetSpecificPortMappingEntryResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igdconn_AddPortMapping(w http.ResponseWriter) {
	out := igdconn.AddPortMappingResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igdconn_DeletePortMapping(w http.ResponseWriter) {
	out := igdconn.DeletePortMappingResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igdconn_GetExternalIPAddress(w http.ResponseWriter) {
	out := igdconn.GetExternalIPAddressResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igdconn_X_AVM_DE_GetExternalIPv6Address(w http.ResponseWriter) {
	out := igdconn.X_AVM_DE_GetExternalIPv6AddressResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igdconn_X_AVM_DE_GetIPv6Prefix(w http.ResponseWriter) {
	out := igdconn.X_AVM_DE_GetIPv6PrefixResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igdconn_X_AVM_DE_GetDNSServer(w http.ResponseWriter) {
	out := igdconn.X_AVM_DE_GetDNSServerResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igdconn_X_AVM_DE_GetIPv6DNSServer(w http.ResponseWriter) {
	out := igdconn.X_AVM_DE_GetIPv6DNSServerResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
