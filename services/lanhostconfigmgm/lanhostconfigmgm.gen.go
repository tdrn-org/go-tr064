// generated from spec version: 1.0
package lanhostconfigmgm

import (
	"encoding/xml"
	"github.com/tdrn-org/go-tr064"
)

type ServiceClient struct {
	TR064Client *tr064.Client
	Service     tr064.ServiceDescriptor
}

type GetInfoRequest struct {
	XMLName      xml.Name `xml:"u:GetInfoRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetInfoResponse struct {
	XMLName                   xml.Name `xml:"GetInfoResponse"`
	NewDHCPServerConfigurable bool     `xml:"NewDHCPServerConfigurable"`
	NewDHCPRelay              bool     `xml:"NewDHCPRelay"`
	NewMinAddress             string   `xml:"NewMinAddress"`
	NewMaxAddress             string   `xml:"NewMaxAddress"`
	NewReservedAddresses      string   `xml:"NewReservedAddresses"`
	NewDHCPServerEnable       bool     `xml:"NewDHCPServerEnable"`
	NewDNSServers             string   `xml:"NewDNSServers"`
	NewDomainName             string   `xml:"NewDomainName"`
	NewIPRouters              string   `xml:"NewIPRouters"`
	NewSubnetMask             string   `xml:"NewSubnetMask"`
}

func (client *ServiceClient) GetInfo(out *GetInfoResponse) error {
	in := &GetInfoRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetInfoRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetInfoRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetInfoResponse]{
		Body: &tr064.SOAPResponseBody[GetInfoResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetInfo", soapRequest, soapResponse)
}

type SetDHCPServerEnableRequest struct {
	XMLName             xml.Name `xml:"u:SetDHCPServerEnableRequest"`
	XMLNameSpace        string   `xml:"xmlns:u,attr"`
	NewDHCPServerEnable bool     `xml:"NewDHCPServerEnable"`
}

type SetDHCPServerEnableResponse struct {
	XMLName xml.Name `xml:"SetDHCPServerEnableResponse"`
}

func (client *ServiceClient) SetDHCPServerEnable(in *SetDHCPServerEnableRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetDHCPServerEnableRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[SetDHCPServerEnableRequest]{
			In: in,
		},
	}
	out := &SetDHCPServerEnableResponse{}
	soapResponse := &tr064.SOAPResponse[SetDHCPServerEnableResponse]{
		Body: &tr064.SOAPResponseBody[SetDHCPServerEnableResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetDHCPServerEnable", soapRequest, soapResponse)
}

type SetIPInterfaceRequest struct {
	XMLName             xml.Name `xml:"u:SetIPInterfaceRequest"`
	XMLNameSpace        string   `xml:"xmlns:u,attr"`
	NewEnable           bool     `xml:"NewEnable"`
	NewIPAddress        string   `xml:"NewIPAddress"`
	NewSubnetMask       string   `xml:"NewSubnetMask"`
	NewIPAddressingType string   `xml:"NewIPAddressingType"`
}

type SetIPInterfaceResponse struct {
	XMLName xml.Name `xml:"SetIPInterfaceResponse"`
}

func (client *ServiceClient) SetIPInterface(in *SetIPInterfaceRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetIPInterfaceRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[SetIPInterfaceRequest]{
			In: in,
		},
	}
	out := &SetIPInterfaceResponse{}
	soapResponse := &tr064.SOAPResponse[SetIPInterfaceResponse]{
		Body: &tr064.SOAPResponseBody[SetIPInterfaceResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetIPInterface", soapRequest, soapResponse)
}

type GetAddressRangeRequest struct {
	XMLName      xml.Name `xml:"u:GetAddressRangeRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetAddressRangeResponse struct {
	XMLName       xml.Name `xml:"GetAddressRangeResponse"`
	NewMinAddress string   `xml:"NewMinAddress"`
	NewMaxAddress string   `xml:"NewMaxAddress"`
}

func (client *ServiceClient) GetAddressRange(out *GetAddressRangeResponse) error {
	in := &GetAddressRangeRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetAddressRangeRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetAddressRangeRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetAddressRangeResponse]{
		Body: &tr064.SOAPResponseBody[GetAddressRangeResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetAddressRange", soapRequest, soapResponse)
}

type SetAddressRangeRequest struct {
	XMLName       xml.Name `xml:"u:SetAddressRangeRequest"`
	XMLNameSpace  string   `xml:"xmlns:u,attr"`
	NewMinAddress string   `xml:"NewMinAddress"`
	NewMaxAddress string   `xml:"NewMaxAddress"`
}

type SetAddressRangeResponse struct {
	XMLName xml.Name `xml:"SetAddressRangeResponse"`
}

func (client *ServiceClient) SetAddressRange(in *SetAddressRangeRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetAddressRangeRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[SetAddressRangeRequest]{
			In: in,
		},
	}
	out := &SetAddressRangeResponse{}
	soapResponse := &tr064.SOAPResponse[SetAddressRangeResponse]{
		Body: &tr064.SOAPResponseBody[SetAddressRangeResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetAddressRange", soapRequest, soapResponse)
}

type GetIPRoutersListRequest struct {
	XMLName      xml.Name `xml:"u:GetIPRoutersListRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetIPRoutersListResponse struct {
	XMLName      xml.Name `xml:"GetIPRoutersListResponse"`
	NewIPRouters string   `xml:"NewIPRouters"`
}

func (client *ServiceClient) GetIPRoutersList(out *GetIPRoutersListResponse) error {
	in := &GetIPRoutersListRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetIPRoutersListRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetIPRoutersListRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetIPRoutersListResponse]{
		Body: &tr064.SOAPResponseBody[GetIPRoutersListResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetIPRoutersList", soapRequest, soapResponse)
}

type SetIPRouterRequest struct {
	XMLName      xml.Name `xml:"u:SetIPRouterRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewIPRouters string   `xml:"NewIPRouters"`
}

type SetIPRouterResponse struct {
	XMLName xml.Name `xml:"SetIPRouterResponse"`
}

func (client *ServiceClient) SetIPRouter(in *SetIPRouterRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetIPRouterRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[SetIPRouterRequest]{
			In: in,
		},
	}
	out := &SetIPRouterResponse{}
	soapResponse := &tr064.SOAPResponse[SetIPRouterResponse]{
		Body: &tr064.SOAPResponseBody[SetIPRouterResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetIPRouter", soapRequest, soapResponse)
}

type GetSubnetMaskRequest struct {
	XMLName      xml.Name `xml:"u:GetSubnetMaskRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetSubnetMaskResponse struct {
	XMLName       xml.Name `xml:"GetSubnetMaskResponse"`
	NewSubnetMask string   `xml:"NewSubnetMask"`
}

func (client *ServiceClient) GetSubnetMask(out *GetSubnetMaskResponse) error {
	in := &GetSubnetMaskRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetSubnetMaskRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetSubnetMaskRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetSubnetMaskResponse]{
		Body: &tr064.SOAPResponseBody[GetSubnetMaskResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetSubnetMask", soapRequest, soapResponse)
}

type SetSubnetMaskRequest struct {
	XMLName       xml.Name `xml:"u:SetSubnetMaskRequest"`
	XMLNameSpace  string   `xml:"xmlns:u,attr"`
	NewSubnetMask string   `xml:"NewSubnetMask"`
}

type SetSubnetMaskResponse struct {
	XMLName xml.Name `xml:"SetSubnetMaskResponse"`
}

func (client *ServiceClient) SetSubnetMask(in *SetSubnetMaskRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetSubnetMaskRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[SetSubnetMaskRequest]{
			In: in,
		},
	}
	out := &SetSubnetMaskResponse{}
	soapResponse := &tr064.SOAPResponse[SetSubnetMaskResponse]{
		Body: &tr064.SOAPResponseBody[SetSubnetMaskResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetSubnetMask", soapRequest, soapResponse)
}

type GetDNSServersRequest struct {
	XMLName      xml.Name `xml:"u:GetDNSServersRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetDNSServersResponse struct {
	XMLName       xml.Name `xml:"GetDNSServersResponse"`
	NewDNSServers string   `xml:"NewDNSServers"`
}

func (client *ServiceClient) GetDNSServers(out *GetDNSServersResponse) error {
	in := &GetDNSServersRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetDNSServersRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetDNSServersRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetDNSServersResponse]{
		Body: &tr064.SOAPResponseBody[GetDNSServersResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetDNSServers", soapRequest, soapResponse)
}

type GetIPInterfaceNumberOfEntriesRequest struct {
	XMLName      xml.Name `xml:"u:GetIPInterfaceNumberOfEntriesRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetIPInterfaceNumberOfEntriesResponse struct {
	XMLName                       xml.Name `xml:"GetIPInterfaceNumberOfEntriesResponse"`
	NewIPInterfaceNumberOfEntries uint16   `xml:"NewIPInterfaceNumberOfEntries"`
}

func (client *ServiceClient) GetIPInterfaceNumberOfEntries(out *GetIPInterfaceNumberOfEntriesResponse) error {
	in := &GetIPInterfaceNumberOfEntriesRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetIPInterfaceNumberOfEntriesRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetIPInterfaceNumberOfEntriesRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetIPInterfaceNumberOfEntriesResponse]{
		Body: &tr064.SOAPResponseBody[GetIPInterfaceNumberOfEntriesResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetIPInterfaceNumberOfEntries", soapRequest, soapResponse)
}
