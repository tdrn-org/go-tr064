// generated from spec version: 1.0
package x_hostfilter

import (
	"encoding/xml"
	"github.com/tdrn-org/go-tr064"
)

type ServiceClient struct {
	TR064Client *tr064.Client
	Service     tr064.ServiceDescriptor
}

type MarkTicketRequest struct {
	XMLName      xml.Name `xml:"u:MarkTicketRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type MarkTicketResponse struct {
	XMLName     xml.Name `xml:"MarkTicketResponse"`
	NewTicketID string   `xml:"NewTicketID"`
}

func (client *ServiceClient) MarkTicket(out *MarkTicketResponse) error {
	in := &MarkTicketRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "MarkTicket", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetTicketIDStatusRequest struct {
	XMLName      xml.Name `xml:"u:GetTicketIDStatusRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewTicketID  string   `xml:"NewTicketID"`
}

type GetTicketIDStatusResponse struct {
	XMLName           xml.Name `xml:"GetTicketIDStatusResponse"`
	NewTicketIDStatus string   `xml:"NewTicketIDStatus"`
}

func (client *ServiceClient) GetTicketIDStatus(in *GetTicketIDStatusRequest, out *GetTicketIDStatusResponse) error {
	in.XMLNameSpace = client.Service.Type()
	return client.TR064Client.InvokeService(client.Service, "GetTicketIDStatus", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type DiscardAllTicketsRequest struct {
	XMLName      xml.Name `xml:"u:DiscardAllTicketsRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type DiscardAllTicketsResponse struct {
	XMLName xml.Name `xml:"DiscardAllTicketsResponse"`
}

func (client *ServiceClient) DiscardAllTickets() error {
	in := &DiscardAllTicketsRequest{XMLNameSpace: client.Service.Type()}
	out := &DiscardAllTicketsResponse{}
	return client.TR064Client.InvokeService(client.Service, "DiscardAllTickets", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type DisallowWANAccessByIPRequest struct {
	XMLName        xml.Name `xml:"u:DisallowWANAccessByIPRequest"`
	XMLNameSpace   string   `xml:"xmlns:u,attr"`
	NewIPv4Address string   `xml:"NewIPv4Address"`
	NewDisallow    bool     `xml:"NewDisallow"`
}

type DisallowWANAccessByIPResponse struct {
	XMLName xml.Name `xml:"DisallowWANAccessByIPResponse"`
}

func (client *ServiceClient) DisallowWANAccessByIP(in *DisallowWANAccessByIPRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &DisallowWANAccessByIPResponse{}
	return client.TR064Client.InvokeService(client.Service, "DisallowWANAccessByIP", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetWANAccessByIPRequest struct {
	XMLName        xml.Name `xml:"u:GetWANAccessByIPRequest"`
	XMLNameSpace   string   `xml:"xmlns:u,attr"`
	NewIPv4Address string   `xml:"NewIPv4Address"`
}

type GetWANAccessByIPResponse struct {
	XMLName      xml.Name `xml:"GetWANAccessByIPResponse"`
	NewDisallow  bool     `xml:"NewDisallow"`
	NewWANAccess string   `xml:"NewWANAccess"`
}

func (client *ServiceClient) GetWANAccessByIP(in *GetWANAccessByIPRequest, out *GetWANAccessByIPResponse) error {
	in.XMLNameSpace = client.Service.Type()
	return client.TR064Client.InvokeService(client.Service, "GetWANAccessByIP", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}
