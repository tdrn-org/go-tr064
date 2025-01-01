// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/x_hostfilter"
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
			ServiceName: "X_AVM_DE_HostFilter",
			ServiceType: "urn:dslforum-org:service:X_AVM-DE_HostFilter:1",
			ServiceId:   "urn:X_AVM-DE_HostFilter-com:serviceId:X_AVM-DE_HostFilter1",
			ServiceUrl:  "/upnp/control/x_hostfilter",
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
