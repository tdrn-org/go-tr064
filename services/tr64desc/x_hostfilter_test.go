// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/tr64desc/x_hostfilter"
	"log"
	"net/http"
	"testing"
)

var x_hostfilterMock = &mock.ServiceMock{
	Path:       "/upnp/control/x_hostfilter",
	HandleFunc: x_hostfilterHandler,
}

func TestX_AVM_DE_HostFilter(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", x_hostfilterMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server())
	client.Debug = true
	serviceClient := &x_hostfilter.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceSpec:       tr064.ServiceSpec("tr64desc"),
			ServiceType:       "urn:dslforum-org:service:X_AVM-DE_HostFilter:1",
			ServiceId:         "urn:X_AVM-DE_HostFilter-com:serviceId:X_AVM-DE_HostFilter1",
			ServiceControlUrl: "/upnp/control/x_hostfilter",
		},
	}
	{
		out := &x_hostfilter.MarkTicketResponse{}
		require.NoError(t, serviceClient.MarkTicket(out))
	}
	{
		in := &x_hostfilter.GetTicketIDStatusRequest{}
		out := &x_hostfilter.GetTicketIDStatusResponse{}
		require.NoError(t, serviceClient.GetTicketIDStatus(in, out))
	}
	{
		require.NoError(t, serviceClient.DiscardAllTickets())
	}
	{
		in := &x_hostfilter.DisallowWANAccessByIPRequest{}
		require.NoError(t, serviceClient.DisallowWANAccessByIP(in))
	}
	{
		in := &x_hostfilter.GetWANAccessByIPRequest{}
		out := &x_hostfilter.GetWANAccessByIPResponse{}
		require.NoError(t, serviceClient.GetWANAccessByIP(in, out))
	}
	{
		in := &x_hostfilter.GetHostEntryByIPRequest{}
		out := &x_hostfilter.GetHostEntryByIPResponse{}
		require.NoError(t, serviceClient.GetHostEntryByIP(in, out))
	}
	{
		in := &x_hostfilter.GetFilterProfileByIDRequest{}
		out := &x_hostfilter.GetFilterProfileByIDResponse{}
		require.NoError(t, serviceClient.GetFilterProfileByID(in, out))
	}
	{
		in := &x_hostfilter.AddTicketTimeToHostEntryByIPRequest{}
		out := &x_hostfilter.AddTicketTimeToHostEntryByIPResponse{}
		require.NoError(t, serviceClient.AddTicketTimeToHostEntryByIP(in, out))
	}
	{
		in := &x_hostfilter.AddHostEntryToFilterProfileRequest{}
		require.NoError(t, serviceClient.AddHostEntryToFilterProfile(in))
	}
	{
		out := &x_hostfilter.GetFilterProfilesResponse{}
		require.NoError(t, serviceClient.GetFilterProfiles(out))
	}
}

func x_hostfilterHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "MarkTicket":
		x_hostfilter_MarkTicket(w)
	case "GetTicketIDStatus":
		x_hostfilter_GetTicketIDStatus(w)
	case "DiscardAllTickets":
		x_hostfilter_DiscardAllTickets(w)
	case "DisallowWANAccessByIP":
		x_hostfilter_DisallowWANAccessByIP(w)
	case "GetWANAccessByIP":
		x_hostfilter_GetWANAccessByIP(w)
	case "GetHostEntryByIP":
		x_hostfilter_GetHostEntryByIP(w)
	case "GetFilterProfileByID":
		x_hostfilter_GetFilterProfileByID(w)
	case "AddTicketTimeToHostEntryByIP":
		x_hostfilter_AddTicketTimeToHostEntryByIP(w)
	case "AddHostEntryToFilterProfile":
		x_hostfilter_AddHostEntryToFilterProfile(w)
	case "GetFilterProfiles":
		x_hostfilter_GetFilterProfiles(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func x_hostfilter_MarkTicket(w http.ResponseWriter) {
	out := x_hostfilter.MarkTicketResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_hostfilter_GetTicketIDStatus(w http.ResponseWriter) {
	out := x_hostfilter.GetTicketIDStatusResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_hostfilter_DiscardAllTickets(w http.ResponseWriter) {
	out := x_hostfilter.DiscardAllTicketsResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_hostfilter_DisallowWANAccessByIP(w http.ResponseWriter) {
	out := x_hostfilter.DisallowWANAccessByIPResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_hostfilter_GetWANAccessByIP(w http.ResponseWriter) {
	out := x_hostfilter.GetWANAccessByIPResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_hostfilter_GetHostEntryByIP(w http.ResponseWriter) {
	out := x_hostfilter.GetHostEntryByIPResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_hostfilter_GetFilterProfileByID(w http.ResponseWriter) {
	out := x_hostfilter.GetFilterProfileByIDResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_hostfilter_AddTicketTimeToHostEntryByIP(w http.ResponseWriter) {
	out := x_hostfilter.AddTicketTimeToHostEntryByIPResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_hostfilter_AddHostEntryToFilterProfile(w http.ResponseWriter) {
	out := x_hostfilter.AddHostEntryToFilterProfileResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_hostfilter_GetFilterProfiles(w http.ResponseWriter) {
	out := x_hostfilter.GetFilterProfilesResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
