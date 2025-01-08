// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/tr64desc/x_wanmobileconn"
	"log"
	"net/http"
	"testing"
)

var x_wanmobileconnMock = &mock.ServiceMock{
	Path:       "/upnp/control/x_wanmobileconn",
	HandleFunc: x_wanmobileconnHandler,
}

func TestX_AVM_DE_WANMobileConnection(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", x_wanmobileconnMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server(), tr064.ServiceSpec("tr64desc"))
	client.Debug = true
	serviceClient := &x_wanmobileconn.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceSpec:       tr064.ServiceSpec("tr64desc"),
			ServiceType:       "urn:dslforum-org:service:X_AVM-DE_WANMobileConnection:1",
			ServiceId:         "urn:X_AVM-DE_WANMobileConnection-com:serviceId:X_AVM-DE_WANMobileConnection1",
			ServiceControlUrl: "/upnp/control/x_wanmobileconn",
		},
	}
	{
		out := &x_wanmobileconn.GetInfoResponse{}
		require.NoError(t, serviceClient.GetInfo(out))
	}
	{
		out := &x_wanmobileconn.GetInfoExResponse{}
		require.NoError(t, serviceClient.GetInfoEx(out))
	}
	{
		in := &x_wanmobileconn.SetPINRequest{}
		require.NoError(t, serviceClient.SetPIN(in))
	}
	{
		in := &x_wanmobileconn.SetPUKRequest{}
		require.NoError(t, serviceClient.SetPUK(in))
	}
	{
		in := &x_wanmobileconn.SetAccessTechnologyRequest{}
		require.NoError(t, serviceClient.SetAccessTechnology(in))
	}
	{
		out := &x_wanmobileconn.GetAccessTechnologyResponse{}
		require.NoError(t, serviceClient.GetAccessTechnology(out))
	}
	{
		in := &x_wanmobileconn.SetPreferredAccessTechnologyRequest{}
		require.NoError(t, serviceClient.SetPreferredAccessTechnology(in))
	}
	{
		out := &x_wanmobileconn.GetPreferredAccessTechnologyResponse{}
		require.NoError(t, serviceClient.GetPreferredAccessTechnology(out))
	}
	{
		in := &x_wanmobileconn.SetEnabledBandCapabilitiesRequest{}
		require.NoError(t, serviceClient.SetEnabledBandCapabilities(in))
	}
	{
		out := &x_wanmobileconn.GetEnabledBandCapabilitiesResponse{}
		require.NoError(t, serviceClient.GetEnabledBandCapabilities(out))
	}
	{
		out := &x_wanmobileconn.GetBandCapabilitiesResponse{}
		require.NoError(t, serviceClient.GetBandCapabilities(out))
	}
}

func x_wanmobileconnHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetInfo":
		x_wanmobileconn_GetInfo(w)
	case "GetInfoEx":
		x_wanmobileconn_GetInfoEx(w)
	case "SetPIN":
		x_wanmobileconn_SetPIN(w)
	case "SetPUK":
		x_wanmobileconn_SetPUK(w)
	case "SetAccessTechnology":
		x_wanmobileconn_SetAccessTechnology(w)
	case "GetAccessTechnology":
		x_wanmobileconn_GetAccessTechnology(w)
	case "SetPreferredAccessTechnology":
		x_wanmobileconn_SetPreferredAccessTechnology(w)
	case "GetPreferredAccessTechnology":
		x_wanmobileconn_GetPreferredAccessTechnology(w)
	case "SetEnabledBandCapabilities":
		x_wanmobileconn_SetEnabledBandCapabilities(w)
	case "GetEnabledBandCapabilities":
		x_wanmobileconn_GetEnabledBandCapabilities(w)
	case "GetBandCapabilities":
		x_wanmobileconn_GetBandCapabilities(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func x_wanmobileconn_GetInfo(w http.ResponseWriter) {
	out := x_wanmobileconn.GetInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_wanmobileconn_GetInfoEx(w http.ResponseWriter) {
	out := x_wanmobileconn.GetInfoExResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_wanmobileconn_SetPIN(w http.ResponseWriter) {
	out := x_wanmobileconn.SetPINResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_wanmobileconn_SetPUK(w http.ResponseWriter) {
	out := x_wanmobileconn.SetPUKResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_wanmobileconn_SetAccessTechnology(w http.ResponseWriter) {
	out := x_wanmobileconn.SetAccessTechnologyResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_wanmobileconn_GetAccessTechnology(w http.ResponseWriter) {
	out := x_wanmobileconn.GetAccessTechnologyResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_wanmobileconn_SetPreferredAccessTechnology(w http.ResponseWriter) {
	out := x_wanmobileconn.SetPreferredAccessTechnologyResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_wanmobileconn_GetPreferredAccessTechnology(w http.ResponseWriter) {
	out := x_wanmobileconn.GetPreferredAccessTechnologyResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_wanmobileconn_SetEnabledBandCapabilities(w http.ResponseWriter) {
	out := x_wanmobileconn.SetEnabledBandCapabilitiesResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_wanmobileconn_GetEnabledBandCapabilities(w http.ResponseWriter) {
	out := x_wanmobileconn.GetEnabledBandCapabilitiesResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_wanmobileconn_GetBandCapabilities(w http.ResponseWriter) {
	out := x_wanmobileconn.GetBandCapabilitiesResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
