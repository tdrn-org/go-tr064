// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/tr64desc/hosts"
	"log"
	"net/http"
	"testing"
)

var hostsMock = &mock.ServiceMock{
	Path:       "/upnp/control/hosts",
	HandleFunc: hostsHandler,
}

func TestHosts(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", hostsMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server(), tr064.ServiceSpec("tr64desc"))
	client.Debug = true
	serviceClient := &hosts.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceSpec:       tr064.ServiceSpec("tr64desc"),
			ServiceType:       "urn:dslforum-org:service:Hosts:1",
			ServiceId:         "urn:LanDeviceHosts-com:serviceId:Hosts1",
			ServiceControlUrl: "/upnp/control/hosts",
		},
	}
	{
		out := &hosts.GetHostNumberOfEntriesResponse{}
		require.NoError(t, serviceClient.GetHostNumberOfEntries(out))
	}
	{
		in := &hosts.GetSpecificHostEntryRequest{}
		out := &hosts.GetSpecificHostEntryResponse{}
		require.NoError(t, serviceClient.GetSpecificHostEntry(in, out))
	}
	{
		in := &hosts.GetGenericHostEntryRequest{}
		out := &hosts.GetGenericHostEntryResponse{}
		require.NoError(t, serviceClient.GetGenericHostEntry(in, out))
	}
	{
		out := &hosts.X_AVM_DE_GetInfoResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetInfo(out))
	}
	{
		out := &hosts.X_AVM_DE_GetChangeCounterResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetChangeCounter(out))
	}
	{
		in := &hosts.X_AVM_DE_SetHostNameByMACAddressRequest{}
		require.NoError(t, serviceClient.X_AVM_DE_SetHostNameByMACAddress(in))
	}
	{
		in := &hosts.X_AVM_DE_GetAutoWakeOnLANByMACAddressRequest{}
		out := &hosts.X_AVM_DE_GetAutoWakeOnLANByMACAddressResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetAutoWakeOnLANByMACAddress(in, out))
	}
	{
		in := &hosts.X_AVM_DE_SetAutoWakeOnLANByMACAddressRequest{}
		require.NoError(t, serviceClient.X_AVM_DE_SetAutoWakeOnLANByMACAddress(in))
	}
	{
		in := &hosts.X_AVM_DE_WakeOnLANByMACAddressRequest{}
		require.NoError(t, serviceClient.X_AVM_DE_WakeOnLANByMACAddress(in))
	}
	{
		in := &hosts.X_AVM_DE_GetSpecificHostEntryByIPRequest{}
		out := &hosts.X_AVM_DE_GetSpecificHostEntryByIPResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetSpecificHostEntryByIP(in, out))
	}
	{
		require.NoError(t, serviceClient.X_AVM_DE_HostsCheckUpdate())
	}
	{
		in := &hosts.X_AVM_DE_HostDoUpdateRequest{}
		require.NoError(t, serviceClient.X_AVM_DE_HostDoUpdate(in))
	}
	{
		in := &hosts.X_AVM_DE_SetPrioritizationByIPRequest{}
		require.NoError(t, serviceClient.X_AVM_DE_SetPrioritizationByIP(in))
	}
	{
		out := &hosts.X_AVM_DE_GetHostListPathResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetHostListPath(out))
	}
	{
		out := &hosts.X_AVM_DE_GetMeshListPathResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetMeshListPath(out))
	}
	{
		out := &hosts.X_AVM_DE_GetFriendlyNameResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetFriendlyName(out))
	}
	{
		in := &hosts.X_AVM_DE_SetFriendlyNameRequest{}
		require.NoError(t, serviceClient.X_AVM_DE_SetFriendlyName(in))
	}
	{
		in := &hosts.X_AVM_DE_SetFriendlyNameByIPRequest{}
		require.NoError(t, serviceClient.X_AVM_DE_SetFriendlyNameByIP(in))
	}
	{
		in := &hosts.X_AVM_DE_SetFriendlyNameByMACRequest{}
		require.NoError(t, serviceClient.X_AVM_DE_SetFriendlyNameByMAC(in))
	}
}

func hostsHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetHostNumberOfEntries":
		hosts_GetHostNumberOfEntries(w)
	case "GetSpecificHostEntry":
		hosts_GetSpecificHostEntry(w)
	case "GetGenericHostEntry":
		hosts_GetGenericHostEntry(w)
	case "X_AVM-DE_GetInfo":
		hosts_X_AVM_DE_GetInfo(w)
	case "X_AVM-DE_GetChangeCounter":
		hosts_X_AVM_DE_GetChangeCounter(w)
	case "X_AVM-DE_SetHostNameByMACAddress":
		hosts_X_AVM_DE_SetHostNameByMACAddress(w)
	case "X_AVM-DE_GetAutoWakeOnLANByMACAddress":
		hosts_X_AVM_DE_GetAutoWakeOnLANByMACAddress(w)
	case "X_AVM-DE_SetAutoWakeOnLANByMACAddress":
		hosts_X_AVM_DE_SetAutoWakeOnLANByMACAddress(w)
	case "X_AVM-DE_WakeOnLANByMACAddress":
		hosts_X_AVM_DE_WakeOnLANByMACAddress(w)
	case "X_AVM-DE_GetSpecificHostEntryByIP":
		hosts_X_AVM_DE_GetSpecificHostEntryByIP(w)
	case "X_AVM-DE_HostsCheckUpdate":
		hosts_X_AVM_DE_HostsCheckUpdate(w)
	case "X_AVM-DE_HostDoUpdate":
		hosts_X_AVM_DE_HostDoUpdate(w)
	case "X_AVM-DE_SetPrioritizationByIP":
		hosts_X_AVM_DE_SetPrioritizationByIP(w)
	case "X_AVM-DE_GetHostListPath":
		hosts_X_AVM_DE_GetHostListPath(w)
	case "X_AVM-DE_GetMeshListPath":
		hosts_X_AVM_DE_GetMeshListPath(w)
	case "X_AVM-DE_GetFriendlyName":
		hosts_X_AVM_DE_GetFriendlyName(w)
	case "X_AVM-DE_SetFriendlyName":
		hosts_X_AVM_DE_SetFriendlyName(w)
	case "X_AVM-DE_SetFriendlyNameByIP":
		hosts_X_AVM_DE_SetFriendlyNameByIP(w)
	case "X_AVM-DE_SetFriendlyNameByMAC":
		hosts_X_AVM_DE_SetFriendlyNameByMAC(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func hosts_GetHostNumberOfEntries(w http.ResponseWriter) {
	out := hosts.GetHostNumberOfEntriesResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func hosts_GetSpecificHostEntry(w http.ResponseWriter) {
	out := hosts.GetSpecificHostEntryResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func hosts_GetGenericHostEntry(w http.ResponseWriter) {
	out := hosts.GetGenericHostEntryResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func hosts_X_AVM_DE_GetInfo(w http.ResponseWriter) {
	out := hosts.X_AVM_DE_GetInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func hosts_X_AVM_DE_GetChangeCounter(w http.ResponseWriter) {
	out := hosts.X_AVM_DE_GetChangeCounterResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func hosts_X_AVM_DE_SetHostNameByMACAddress(w http.ResponseWriter) {
	out := hosts.X_AVM_DE_SetHostNameByMACAddressResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func hosts_X_AVM_DE_GetAutoWakeOnLANByMACAddress(w http.ResponseWriter) {
	out := hosts.X_AVM_DE_GetAutoWakeOnLANByMACAddressResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func hosts_X_AVM_DE_SetAutoWakeOnLANByMACAddress(w http.ResponseWriter) {
	out := hosts.X_AVM_DE_SetAutoWakeOnLANByMACAddressResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func hosts_X_AVM_DE_WakeOnLANByMACAddress(w http.ResponseWriter) {
	out := hosts.X_AVM_DE_WakeOnLANByMACAddressResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func hosts_X_AVM_DE_GetSpecificHostEntryByIP(w http.ResponseWriter) {
	out := hosts.X_AVM_DE_GetSpecificHostEntryByIPResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func hosts_X_AVM_DE_HostsCheckUpdate(w http.ResponseWriter) {
	out := hosts.X_AVM_DE_HostsCheckUpdateResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func hosts_X_AVM_DE_HostDoUpdate(w http.ResponseWriter) {
	out := hosts.X_AVM_DE_HostDoUpdateResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func hosts_X_AVM_DE_SetPrioritizationByIP(w http.ResponseWriter) {
	out := hosts.X_AVM_DE_SetPrioritizationByIPResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func hosts_X_AVM_DE_GetHostListPath(w http.ResponseWriter) {
	out := hosts.X_AVM_DE_GetHostListPathResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func hosts_X_AVM_DE_GetMeshListPath(w http.ResponseWriter) {
	out := hosts.X_AVM_DE_GetMeshListPathResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func hosts_X_AVM_DE_GetFriendlyName(w http.ResponseWriter) {
	out := hosts.X_AVM_DE_GetFriendlyNameResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func hosts_X_AVM_DE_SetFriendlyName(w http.ResponseWriter) {
	out := hosts.X_AVM_DE_SetFriendlyNameResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func hosts_X_AVM_DE_SetFriendlyNameByIP(w http.ResponseWriter) {
	out := hosts.X_AVM_DE_SetFriendlyNameByIPResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func hosts_X_AVM_DE_SetFriendlyNameByMAC(w http.ResponseWriter) {
	out := hosts.X_AVM_DE_SetFriendlyNameByMACResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
