// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/igddesc/igddsl"
	"log"
	"net/http"
	"testing"
)

var igddslMock = &mock.ServiceMock{
	Path:       "/igdupnp/control/WANDSLLinkC1",
	HandleFunc: igddslHandler,
}

func TestWANDSLLinkConfig(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", igddslMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server(), tr064.ServiceSpec("igddesc"))
	client.Debug = true
	serviceClient := &igddsl.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceSpec: tr064.ServiceSpec("igddesc"),
			ServiceId:   "urn:upnp-org:serviceId:WANDSLLinkC1",
			ServiceUrl:  "/igdupnp/control/WANDSLLinkC1",
		},
	}
	{
		in := &igddsl.SetDSLLinkTypeRequest{}
		require.NoError(t, serviceClient.SetDSLLinkType(in))
	}
	{
		out := &igddsl.GetDSLLinkInfoResponse{}
		require.NoError(t, serviceClient.GetDSLLinkInfo(out))
	}
	{
		out := &igddsl.GetAutoConfigResponse{}
		require.NoError(t, serviceClient.GetAutoConfig(out))
	}
	{
		out := &igddsl.GetModulationTypeResponse{}
		require.NoError(t, serviceClient.GetModulationType(out))
	}
	{
		in := &igddsl.SetDestinationAddressRequest{}
		require.NoError(t, serviceClient.SetDestinationAddress(in))
	}
	{
		out := &igddsl.GetDestinationAddressResponse{}
		require.NoError(t, serviceClient.GetDestinationAddress(out))
	}
	{
		in := &igddsl.SetATMEncapsulationRequest{}
		require.NoError(t, serviceClient.SetATMEncapsulation(in))
	}
	{
		out := &igddsl.GetATMEncapsulationResponse{}
		require.NoError(t, serviceClient.GetATMEncapsulation(out))
	}
	{
		in := &igddsl.SetFCSPreservedRequest{}
		require.NoError(t, serviceClient.SetFCSPreserved(in))
	}
	{
		out := &igddsl.GetFCSPreservedResponse{}
		require.NoError(t, serviceClient.GetFCSPreserved(out))
	}
}

func igddslHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "SetDSLLinkType":
		igddsl_SetDSLLinkType(w)
	case "GetDSLLinkInfo":
		igddsl_GetDSLLinkInfo(w)
	case "GetAutoConfig":
		igddsl_GetAutoConfig(w)
	case "GetModulationType":
		igddsl_GetModulationType(w)
	case "SetDestinationAddress":
		igddsl_SetDestinationAddress(w)
	case "GetDestinationAddress":
		igddsl_GetDestinationAddress(w)
	case "SetATMEncapsulation":
		igddsl_SetATMEncapsulation(w)
	case "GetATMEncapsulation":
		igddsl_GetATMEncapsulation(w)
	case "SetFCSPreserved":
		igddsl_SetFCSPreserved(w)
	case "GetFCSPreserved":
		igddsl_GetFCSPreserved(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func igddsl_SetDSLLinkType(w http.ResponseWriter) {
	out := igddsl.SetDSLLinkTypeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igddsl_GetDSLLinkInfo(w http.ResponseWriter) {
	out := igddsl.GetDSLLinkInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igddsl_GetAutoConfig(w http.ResponseWriter) {
	out := igddsl.GetAutoConfigResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igddsl_GetModulationType(w http.ResponseWriter) {
	out := igddsl.GetModulationTypeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igddsl_SetDestinationAddress(w http.ResponseWriter) {
	out := igddsl.SetDestinationAddressResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igddsl_GetDestinationAddress(w http.ResponseWriter) {
	out := igddsl.GetDestinationAddressResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igddsl_SetATMEncapsulation(w http.ResponseWriter) {
	out := igddsl.SetATMEncapsulationResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igddsl_GetATMEncapsulation(w http.ResponseWriter) {
	out := igddsl.GetATMEncapsulationResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igddsl_SetFCSPreserved(w http.ResponseWriter) {
	out := igddsl.SetFCSPreservedResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func igddsl_GetFCSPreserved(w http.ResponseWriter) {
	out := igddsl.GetFCSPreservedResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
