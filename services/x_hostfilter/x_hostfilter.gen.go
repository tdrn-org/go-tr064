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
	soapRequest := &tr064.SOAPRequest[MarkTicketRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[MarkTicketRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[MarkTicketResponse]{
		Body: &tr064.SOAPResponseBody[MarkTicketResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "MarkTicket", soapRequest, soapResponse)
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
	soapRequest := &tr064.SOAPRequest[GetTicketIDStatusRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetTicketIDStatusRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetTicketIDStatusResponse]{
		Body: &tr064.SOAPResponseBody[GetTicketIDStatusResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetTicketIDStatus", soapRequest, soapResponse)
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
	soapRequest := &tr064.SOAPRequest[DiscardAllTicketsRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[DiscardAllTicketsRequest]{
			In: in,
		},
	}
	out := &DiscardAllTicketsResponse{}
	soapResponse := &tr064.SOAPResponse[DiscardAllTicketsResponse]{
		Body: &tr064.SOAPResponseBody[DiscardAllTicketsResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "DiscardAllTickets", soapRequest, soapResponse)
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
	soapRequest := &tr064.SOAPRequest[DisallowWANAccessByIPRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[DisallowWANAccessByIPRequest]{
			In: in,
		},
	}
	out := &DisallowWANAccessByIPResponse{}
	soapResponse := &tr064.SOAPResponse[DisallowWANAccessByIPResponse]{
		Body: &tr064.SOAPResponseBody[DisallowWANAccessByIPResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "DisallowWANAccessByIP", soapRequest, soapResponse)
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
	soapRequest := &tr064.SOAPRequest[GetWANAccessByIPRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetWANAccessByIPRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetWANAccessByIPResponse]{
		Body: &tr064.SOAPResponseBody[GetWANAccessByIPResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetWANAccessByIP", soapRequest, soapResponse)
}
