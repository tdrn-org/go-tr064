// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/igddesc/igd2ipv6fwc"
	"log"
	"net/http"
	"testing"
)

var igd2ipv6fwcMock = &mock.ServiceMock{
	Path:       "/igd2upnp/control/WANIPv6Firewall1",
	HandleFunc: igd2ipv6fwcHandler,
}

func TestWANIPv6FirewallControl(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", igd2ipv6fwcMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server())
	client.Debug = true
	serviceClient := &igd2ipv6fwc.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceSpec:       tr064.ServiceSpec("igddesc"),
			ServiceType:       "urn:schemas-upnp-org:service:WANIPv6FirewallControl:1",
			ServiceId:         "urn:upnp-org:serviceId:WANIPv6Firewall1",
			ServiceControlUrl: "/igd2upnp/control/WANIPv6Firewall1",
		},
	}
	{
		out := &igd2ipv6fwc.GetFirewallStatusResponse{}
		require.NoError(t, serviceClient.GetFirewallStatus(out))
	}
	{
		in := &igd2ipv6fwc.GetOutboundPinholeTimeoutRequest{}
		out := &igd2ipv6fwc.GetOutboundPinholeTimeoutResponse{}
		require.NoError(t, serviceClient.GetOutboundPinholeTimeout(in, out))
	}
	{
		in := &igd2ipv6fwc.AddPinholeRequest{}
		out := &igd2ipv6fwc.AddPinholeResponse{}
		require.NoError(t, serviceClient.AddPinhole(in, out))
	}
	{
		in := &igd2ipv6fwc.UpdatePinholeRequest{}
		require.NoError(t, serviceClient.UpdatePinhole(in))
	}
	{
		in := &igd2ipv6fwc.DeletePinholeRequest{}
		require.NoError(t, serviceClient.DeletePinhole(in))
	}
	{
		in := &igd2ipv6fwc.GetPinholePacketsRequest{}
		out := &igd2ipv6fwc.GetPinholePacketsResponse{}
		require.NoError(t, serviceClient.GetPinholePackets(in, out))
	}
	{
		in := &igd2ipv6fwc.CheckPinholeWorkingRequest{}
		out := &igd2ipv6fwc.CheckPinholeWorkingResponse{}
		require.NoError(t, serviceClient.CheckPinholeWorking(in, out))
	}
}

func igd2ipv6fwcHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetFirewallStatus":
		igd2ipv6fwc_GetFirewallStatus(w)
	case "GetOutboundPinholeTimeout":
		igd2ipv6fwc_GetOutboundPinholeTimeout(w)
	case "AddPinhole":
		igd2ipv6fwc_AddPinhole(w)
	case "UpdatePinhole":
		igd2ipv6fwc_UpdatePinhole(w)
	case "DeletePinhole":
		igd2ipv6fwc_DeletePinhole(w)
	case "GetPinholePackets":
		igd2ipv6fwc_GetPinholePackets(w)
	case "CheckPinholeWorking":
		igd2ipv6fwc_CheckPinholeWorking(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func igd2ipv6fwc_GetFirewallStatus(w http.ResponseWriter) {
	out := igd2ipv6fwc.GetFirewallStatusResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igd2ipv6fwc_GetOutboundPinholeTimeout(w http.ResponseWriter) {
	out := igd2ipv6fwc.GetOutboundPinholeTimeoutResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igd2ipv6fwc_AddPinhole(w http.ResponseWriter) {
	out := igd2ipv6fwc.AddPinholeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igd2ipv6fwc_UpdatePinhole(w http.ResponseWriter) {
	out := igd2ipv6fwc.UpdatePinholeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igd2ipv6fwc_DeletePinhole(w http.ResponseWriter) {
	out := igd2ipv6fwc.DeletePinholeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igd2ipv6fwc_GetPinholePackets(w http.ResponseWriter) {
	out := igd2ipv6fwc.GetPinholePacketsResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igd2ipv6fwc_CheckPinholeWorking(w http.ResponseWriter) {
	out := igd2ipv6fwc.CheckPinholeWorkingResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
