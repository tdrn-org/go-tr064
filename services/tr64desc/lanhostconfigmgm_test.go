// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/tr64desc/lanhostconfigmgm"
	"log"
	"net/http"
	"testing"
)

var lanhostconfigmgmMock = &mock.ServiceMock{
	Path:       "/upnp/control/lanhostconfigmgm",
	HandleFunc: lanhostconfigmgmHandler,
}

func TestLANHostConfigManagement(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", lanhostconfigmgmMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server(), tr064.TR064Spec("tr64desc"))
	client.Debug = true
	serviceClient := &lanhostconfigmgm.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceName: "LANHostConfigManagement",
			ServiceType: "urn:dslforum-org:service:LANHostConfigManagement:1",
			ServiceId:   "urn:LANHCfgMgm-com:serviceId:LANHostConfigManagement1",
			ServiceUrl:  "/upnp/control/lanhostconfigmgm",
		},
	}
	{
		out := &lanhostconfigmgm.GetInfoResponse{}
		require.NoError(t, serviceClient.GetInfo(out))
	}
	{
		in := &lanhostconfigmgm.SetDHCPServerEnableRequest{}
		require.NoError(t, serviceClient.SetDHCPServerEnable(in))
	}
	{
		in := &lanhostconfigmgm.SetIPInterfaceRequest{}
		require.NoError(t, serviceClient.SetIPInterface(in))
	}
	{
		out := &lanhostconfigmgm.GetAddressRangeResponse{}
		require.NoError(t, serviceClient.GetAddressRange(out))
	}
	{
		in := &lanhostconfigmgm.SetAddressRangeRequest{}
		require.NoError(t, serviceClient.SetAddressRange(in))
	}
	{
		out := &lanhostconfigmgm.GetIPRoutersListResponse{}
		require.NoError(t, serviceClient.GetIPRoutersList(out))
	}
	{
		in := &lanhostconfigmgm.SetIPRouterRequest{}
		require.NoError(t, serviceClient.SetIPRouter(in))
	}
	{
		out := &lanhostconfigmgm.GetSubnetMaskResponse{}
		require.NoError(t, serviceClient.GetSubnetMask(out))
	}
	{
		in := &lanhostconfigmgm.SetSubnetMaskRequest{}
		require.NoError(t, serviceClient.SetSubnetMask(in))
	}
	{
		out := &lanhostconfigmgm.GetDNSServersResponse{}
		require.NoError(t, serviceClient.GetDNSServers(out))
	}
	{
		out := &lanhostconfigmgm.GetIPInterfaceNumberOfEntriesResponse{}
		require.NoError(t, serviceClient.GetIPInterfaceNumberOfEntries(out))
	}
}

func lanhostconfigmgmHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetInfo":
		lanhostconfigmgm_GetInfo(w)
	case "SetDHCPServerEnable":
		lanhostconfigmgm_SetDHCPServerEnable(w)
	case "SetIPInterface":
		lanhostconfigmgm_SetIPInterface(w)
	case "GetAddressRange":
		lanhostconfigmgm_GetAddressRange(w)
	case "SetAddressRange":
		lanhostconfigmgm_SetAddressRange(w)
	case "GetIPRoutersList":
		lanhostconfigmgm_GetIPRoutersList(w)
	case "SetIPRouter":
		lanhostconfigmgm_SetIPRouter(w)
	case "GetSubnetMask":
		lanhostconfigmgm_GetSubnetMask(w)
	case "SetSubnetMask":
		lanhostconfigmgm_SetSubnetMask(w)
	case "GetDNSServers":
		lanhostconfigmgm_GetDNSServers(w)
	case "GetIPInterfaceNumberOfEntries":
		lanhostconfigmgm_GetIPInterfaceNumberOfEntries(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func lanhostconfigmgm_GetInfo(w http.ResponseWriter) {
	out := lanhostconfigmgm.GetInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func lanhostconfigmgm_SetDHCPServerEnable(w http.ResponseWriter) {
	out := lanhostconfigmgm.SetDHCPServerEnableResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func lanhostconfigmgm_SetIPInterface(w http.ResponseWriter) {
	out := lanhostconfigmgm.SetIPInterfaceResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func lanhostconfigmgm_GetAddressRange(w http.ResponseWriter) {
	out := lanhostconfigmgm.GetAddressRangeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func lanhostconfigmgm_SetAddressRange(w http.ResponseWriter) {
	out := lanhostconfigmgm.SetAddressRangeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func lanhostconfigmgm_GetIPRoutersList(w http.ResponseWriter) {
	out := lanhostconfigmgm.GetIPRoutersListResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func lanhostconfigmgm_SetIPRouter(w http.ResponseWriter) {
	out := lanhostconfigmgm.SetIPRouterResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func lanhostconfigmgm_GetSubnetMask(w http.ResponseWriter) {
	out := lanhostconfigmgm.GetSubnetMaskResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func lanhostconfigmgm_SetSubnetMask(w http.ResponseWriter) {
	out := lanhostconfigmgm.SetSubnetMaskResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func lanhostconfigmgm_GetDNSServers(w http.ResponseWriter) {
	out := lanhostconfigmgm.GetDNSServersResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func lanhostconfigmgm_GetIPInterfaceNumberOfEntries(w http.ResponseWriter) {
	out := lanhostconfigmgm.GetIPInterfaceNumberOfEntriesResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
