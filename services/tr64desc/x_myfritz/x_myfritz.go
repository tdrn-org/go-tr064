// generated from spec version: 1.0
package x_myfritz

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
	XMLName             xml.Name `xml:"GetInfoResponse"`
	NewEnabled          bool     `xml:"NewEnabled"`
	NewDeviceRegistered bool     `xml:"NewDeviceRegistered"`
	NewDynDNSName       string   `xml:"NewDynDNSName"`
	NewPort             uint32   `xml:"NewPort"`
	NewState            string   `xml:"NewState"`
	NewEmail            string   `xml:"NewEmail"`
}

func (client *ServiceClient) GetInfo(out *GetInfoResponse) error {
	in := &GetInfoRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetInfoRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
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

type SetMyFRITZRequest struct {
	XMLName      xml.Name `xml:"u:SetMyFRITZRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewEnabled   bool     `xml:"NewEnabled"`
	NewEmail     string   `xml:"NewEmail"`
}

type SetMyFRITZResponse struct {
	XMLName xml.Name `xml:"SetMyFRITZResponse"`
}

func (client *ServiceClient) SetMyFRITZ(in *SetMyFRITZRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetMyFRITZRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[SetMyFRITZRequest]{
			In: in,
		},
	}
	out := &SetMyFRITZResponse{}
	soapResponse := &tr064.SOAPResponse[SetMyFRITZResponse]{
		Body: &tr064.SOAPResponseBody[SetMyFRITZResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetMyFRITZ", soapRequest, soapResponse)
}

type GetNumberOfServicesRequest struct {
	XMLName      xml.Name `xml:"u:GetNumberOfServicesRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetNumberOfServicesResponse struct {
	XMLName             xml.Name `xml:"GetNumberOfServicesResponse"`
	NewNumberOfServices uint32   `xml:"NewNumberOfServices"`
}

func (client *ServiceClient) GetNumberOfServices(out *GetNumberOfServicesResponse) error {
	in := &GetNumberOfServicesRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetNumberOfServicesRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetNumberOfServicesRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetNumberOfServicesResponse]{
		Body: &tr064.SOAPResponseBody[GetNumberOfServicesResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetNumberOfServices", soapRequest, soapResponse)
}

type GetServiceByIndexRequest struct {
	XMLName      xml.Name `xml:"u:GetServiceByIndexRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewIndex     uint32   `xml:"NewIndex"`
}

type GetServiceByIndexResponse struct {
	XMLName                  xml.Name `xml:"GetServiceByIndexResponse"`
	NewEnabled               bool     `xml:"NewEnabled"`
	NewName                  string   `xml:"NewName"`
	NewScheme                string   `xml:"NewScheme"`
	NewPort                  uint32   `xml:"NewPort"`
	NewURLPath               string   `xml:"NewURLPath"`
	NewType                  string   `xml:"NewType"`
	NewIPv4ForwardingWarning uint8    `xml:"NewIPv4ForwardingWarning"`
	NewIPv4Addresses         string   `xml:"NewIPv4Addresses"`
	NewIPv6Addresses         string   `xml:"NewIPv6Addresses"`
	NewIPv6InterfaceIDs      string   `xml:"NewIPv6InterfaceIDs"`
	NewMACAddress            string   `xml:"NewMACAddress"`
	NewHostName              string   `xml:"NewHostName"`
	NewDynDnsLabel           string   `xml:"NewDynDnsLabel"`
	NewStatus                uint32   `xml:"NewStatus"`
}

func (client *ServiceClient) GetServiceByIndex(in *GetServiceByIndexRequest, out *GetServiceByIndexResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[GetServiceByIndexRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetServiceByIndexRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetServiceByIndexResponse]{
		Body: &tr064.SOAPResponseBody[GetServiceByIndexResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetServiceByIndex", soapRequest, soapResponse)
}

type SetServiceByIndexRequest struct {
	XMLName            xml.Name `xml:"u:SetServiceByIndexRequest"`
	XMLNameSpace       string   `xml:"xmlns:u,attr"`
	NewIndex           uint32   `xml:"NewIndex"`
	NewEnabled         bool     `xml:"NewEnabled"`
	NewName            string   `xml:"NewName"`
	NewScheme          string   `xml:"NewScheme"`
	NewPort            uint32   `xml:"NewPort"`
	NewURLPath         string   `xml:"NewURLPath"`
	NewType            string   `xml:"NewType"`
	NewIPv4Address     string   `xml:"NewIPv4Address"`
	NewIPv6Address     string   `xml:"NewIPv6Address"`
	NewIPv6InterfaceID string   `xml:"NewIPv6InterfaceID"`
	NewMACAddress      string   `xml:"NewMACAddress"`
	NewHostName        string   `xml:"NewHostName"`
}

type SetServiceByIndexResponse struct {
	XMLName xml.Name `xml:"SetServiceByIndexResponse"`
}

func (client *ServiceClient) SetServiceByIndex(in *SetServiceByIndexRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetServiceByIndexRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[SetServiceByIndexRequest]{
			In: in,
		},
	}
	out := &SetServiceByIndexResponse{}
	soapResponse := &tr064.SOAPResponse[SetServiceByIndexResponse]{
		Body: &tr064.SOAPResponseBody[SetServiceByIndexResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetServiceByIndex", soapRequest, soapResponse)
}

type DeleteServiceByIndexRequest struct {
	XMLName      xml.Name `xml:"u:DeleteServiceByIndexRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewIndex     uint32   `xml:"NewIndex"`
}

type DeleteServiceByIndexResponse struct {
	XMLName xml.Name `xml:"DeleteServiceByIndexResponse"`
}

func (client *ServiceClient) DeleteServiceByIndex(in *DeleteServiceByIndexRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[DeleteServiceByIndexRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[DeleteServiceByIndexRequest]{
			In: in,
		},
	}
	out := &DeleteServiceByIndexResponse{}
	soapResponse := &tr064.SOAPResponse[DeleteServiceByIndexResponse]{
		Body: &tr064.SOAPResponseBody[DeleteServiceByIndexResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "DeleteServiceByIndex", soapRequest, soapResponse)
}
