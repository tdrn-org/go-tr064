// generated from spec version: 1.0
package igdconn

import (
	"encoding/xml"
	"github.com/tdrn-org/go-tr064"
)

type ServiceClient struct {
	TR064Client *tr064.Client
	Service     tr064.ServiceDescriptor
}

type SetConnectionTypeRequest struct {
	XMLName           xml.Name `xml:"u:SetConnectionTypeRequest"`
	XMLNameSpace      string   `xml:"xmlns:u,attr"`
	NewConnectionType string   `xml:"NewConnectionType"`
}

type SetConnectionTypeResponse struct {
	XMLName xml.Name `xml:"SetConnectionTypeResponse"`
}

func (client *ServiceClient) SetConnectionType(in *SetConnectionTypeRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetConnectionTypeRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[SetConnectionTypeRequest]{
			In: in,
		},
	}
	out := &SetConnectionTypeResponse{}
	soapResponse := &tr064.SOAPResponse[SetConnectionTypeResponse]{
		Body: &tr064.SOAPResponseBody[SetConnectionTypeResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetConnectionType", soapRequest, soapResponse)
}

type GetConnectionTypeInfoRequest struct {
	XMLName      xml.Name `xml:"u:GetConnectionTypeInfoRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetConnectionTypeInfoResponse struct {
	XMLName                    xml.Name `xml:"GetConnectionTypeInfoResponse"`
	NewConnectionType          string   `xml:"NewConnectionType"`
	NewPossibleConnectionTypes string   `xml:"NewPossibleConnectionTypes"`
}

func (client *ServiceClient) GetConnectionTypeInfo(out *GetConnectionTypeInfoResponse) error {
	in := &GetConnectionTypeInfoRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetConnectionTypeInfoRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetConnectionTypeInfoRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetConnectionTypeInfoResponse]{
		Body: &tr064.SOAPResponseBody[GetConnectionTypeInfoResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetConnectionTypeInfo", soapRequest, soapResponse)
}

type GetAutoDisconnectTimeRequest struct {
	XMLName      xml.Name `xml:"u:GetAutoDisconnectTimeRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetAutoDisconnectTimeResponse struct {
	XMLName               xml.Name `xml:"GetAutoDisconnectTimeResponse"`
	NewAutoDisconnectTime uint32   `xml:"NewAutoDisconnectTime"`
}

func (client *ServiceClient) GetAutoDisconnectTime(out *GetAutoDisconnectTimeResponse) error {
	in := &GetAutoDisconnectTimeRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetAutoDisconnectTimeRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetAutoDisconnectTimeRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetAutoDisconnectTimeResponse]{
		Body: &tr064.SOAPResponseBody[GetAutoDisconnectTimeResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetAutoDisconnectTime", soapRequest, soapResponse)
}

type GetIdleDisconnectTimeRequest struct {
	XMLName      xml.Name `xml:"u:GetIdleDisconnectTimeRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetIdleDisconnectTimeResponse struct {
	XMLName               xml.Name `xml:"GetIdleDisconnectTimeResponse"`
	NewIdleDisconnectTime uint32   `xml:"NewIdleDisconnectTime"`
}

func (client *ServiceClient) GetIdleDisconnectTime(out *GetIdleDisconnectTimeResponse) error {
	in := &GetIdleDisconnectTimeRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetIdleDisconnectTimeRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetIdleDisconnectTimeRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetIdleDisconnectTimeResponse]{
		Body: &tr064.SOAPResponseBody[GetIdleDisconnectTimeResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetIdleDisconnectTime", soapRequest, soapResponse)
}

type RequestConnectionRequest struct {
	XMLName      xml.Name `xml:"u:RequestConnectionRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type RequestConnectionResponse struct {
	XMLName xml.Name `xml:"RequestConnectionResponse"`
}

func (client *ServiceClient) RequestConnection() error {
	in := &RequestConnectionRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[RequestConnectionRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[RequestConnectionRequest]{
			In: in,
		},
	}
	out := &RequestConnectionResponse{}
	soapResponse := &tr064.SOAPResponse[RequestConnectionResponse]{
		Body: &tr064.SOAPResponseBody[RequestConnectionResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "RequestConnection", soapRequest, soapResponse)
}

type RequestTerminationRequest struct {
	XMLName      xml.Name `xml:"u:RequestTerminationRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type RequestTerminationResponse struct {
	XMLName xml.Name `xml:"RequestTerminationResponse"`
}

func (client *ServiceClient) RequestTermination() error {
	in := &RequestTerminationRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[RequestTerminationRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[RequestTerminationRequest]{
			In: in,
		},
	}
	out := &RequestTerminationResponse{}
	soapResponse := &tr064.SOAPResponse[RequestTerminationResponse]{
		Body: &tr064.SOAPResponseBody[RequestTerminationResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "RequestTermination", soapRequest, soapResponse)
}

type ForceTerminationRequest struct {
	XMLName      xml.Name `xml:"u:ForceTerminationRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type ForceTerminationResponse struct {
	XMLName xml.Name `xml:"ForceTerminationResponse"`
}

func (client *ServiceClient) ForceTermination() error {
	in := &ForceTerminationRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[ForceTerminationRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[ForceTerminationRequest]{
			In: in,
		},
	}
	out := &ForceTerminationResponse{}
	soapResponse := &tr064.SOAPResponse[ForceTerminationResponse]{
		Body: &tr064.SOAPResponseBody[ForceTerminationResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "ForceTermination", soapRequest, soapResponse)
}

type GetStatusInfoRequest struct {
	XMLName      xml.Name `xml:"u:GetStatusInfoRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetStatusInfoResponse struct {
	XMLName                xml.Name `xml:"GetStatusInfoResponse"`
	NewConnectionStatus    string   `xml:"NewConnectionStatus"`
	NewLastConnectionError string   `xml:"NewLastConnectionError"`
	NewUptime              uint32   `xml:"NewUptime"`
}

func (client *ServiceClient) GetStatusInfo(out *GetStatusInfoResponse) error {
	in := &GetStatusInfoRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetStatusInfoRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetStatusInfoRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetStatusInfoResponse]{
		Body: &tr064.SOAPResponseBody[GetStatusInfoResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetStatusInfo", soapRequest, soapResponse)
}

type GetNATRSIPStatusRequest struct {
	XMLName      xml.Name `xml:"u:GetNATRSIPStatusRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetNATRSIPStatusResponse struct {
	XMLName          xml.Name `xml:"GetNATRSIPStatusResponse"`
	NewRSIPAvailable bool     `xml:"NewRSIPAvailable"`
	NewNATEnabled    bool     `xml:"NewNATEnabled"`
}

func (client *ServiceClient) GetNATRSIPStatus(out *GetNATRSIPStatusResponse) error {
	in := &GetNATRSIPStatusRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetNATRSIPStatusRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetNATRSIPStatusRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetNATRSIPStatusResponse]{
		Body: &tr064.SOAPResponseBody[GetNATRSIPStatusResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetNATRSIPStatus", soapRequest, soapResponse)
}

type GetGenericPortMappingEntryRequest struct {
	XMLName             xml.Name `xml:"u:GetGenericPortMappingEntryRequest"`
	XMLNameSpace        string   `xml:"xmlns:u,attr"`
	NewPortMappingIndex uint16   `xml:"NewPortMappingIndex"`
}

type GetGenericPortMappingEntryResponse struct {
	XMLName                   xml.Name `xml:"GetGenericPortMappingEntryResponse"`
	NewRemoteHost             string   `xml:"NewRemoteHost"`
	NewExternalPort           uint16   `xml:"NewExternalPort"`
	NewProtocol               string   `xml:"NewProtocol"`
	NewInternalPort           uint16   `xml:"NewInternalPort"`
	NewInternalClient         string   `xml:"NewInternalClient"`
	NewEnabled                bool     `xml:"NewEnabled"`
	NewPortMappingDescription string   `xml:"NewPortMappingDescription"`
	NewLeaseDuration          uint32   `xml:"NewLeaseDuration"`
}

func (client *ServiceClient) GetGenericPortMappingEntry(in *GetGenericPortMappingEntryRequest, out *GetGenericPortMappingEntryResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[GetGenericPortMappingEntryRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetGenericPortMappingEntryRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetGenericPortMappingEntryResponse]{
		Body: &tr064.SOAPResponseBody[GetGenericPortMappingEntryResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetGenericPortMappingEntry", soapRequest, soapResponse)
}

type GetSpecificPortMappingEntryRequest struct {
	XMLName         xml.Name `xml:"u:GetSpecificPortMappingEntryRequest"`
	XMLNameSpace    string   `xml:"xmlns:u,attr"`
	NewRemoteHost   string   `xml:"NewRemoteHost"`
	NewExternalPort uint16   `xml:"NewExternalPort"`
	NewProtocol     string   `xml:"NewProtocol"`
}

type GetSpecificPortMappingEntryResponse struct {
	XMLName                   xml.Name `xml:"GetSpecificPortMappingEntryResponse"`
	NewInternalPort           uint16   `xml:"NewInternalPort"`
	NewInternalClient         string   `xml:"NewInternalClient"`
	NewEnabled                bool     `xml:"NewEnabled"`
	NewPortMappingDescription string   `xml:"NewPortMappingDescription"`
	NewLeaseDuration          uint32   `xml:"NewLeaseDuration"`
}

func (client *ServiceClient) GetSpecificPortMappingEntry(in *GetSpecificPortMappingEntryRequest, out *GetSpecificPortMappingEntryResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[GetSpecificPortMappingEntryRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetSpecificPortMappingEntryRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetSpecificPortMappingEntryResponse]{
		Body: &tr064.SOAPResponseBody[GetSpecificPortMappingEntryResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetSpecificPortMappingEntry", soapRequest, soapResponse)
}

type AddPortMappingRequest struct {
	XMLName                   xml.Name `xml:"u:AddPortMappingRequest"`
	XMLNameSpace              string   `xml:"xmlns:u,attr"`
	NewRemoteHost             string   `xml:"NewRemoteHost"`
	NewExternalPort           uint16   `xml:"NewExternalPort"`
	NewProtocol               string   `xml:"NewProtocol"`
	NewInternalPort           uint16   `xml:"NewInternalPort"`
	NewInternalClient         string   `xml:"NewInternalClient"`
	NewEnabled                bool     `xml:"NewEnabled"`
	NewPortMappingDescription string   `xml:"NewPortMappingDescription"`
	NewLeaseDuration          uint32   `xml:"NewLeaseDuration"`
}

type AddPortMappingResponse struct {
	XMLName xml.Name `xml:"AddPortMappingResponse"`
}

func (client *ServiceClient) AddPortMapping(in *AddPortMappingRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[AddPortMappingRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[AddPortMappingRequest]{
			In: in,
		},
	}
	out := &AddPortMappingResponse{}
	soapResponse := &tr064.SOAPResponse[AddPortMappingResponse]{
		Body: &tr064.SOAPResponseBody[AddPortMappingResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "AddPortMapping", soapRequest, soapResponse)
}

type DeletePortMappingRequest struct {
	XMLName         xml.Name `xml:"u:DeletePortMappingRequest"`
	XMLNameSpace    string   `xml:"xmlns:u,attr"`
	NewRemoteHost   string   `xml:"NewRemoteHost"`
	NewExternalPort uint16   `xml:"NewExternalPort"`
	NewProtocol     string   `xml:"NewProtocol"`
}

type DeletePortMappingResponse struct {
	XMLName xml.Name `xml:"DeletePortMappingResponse"`
}

func (client *ServiceClient) DeletePortMapping(in *DeletePortMappingRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[DeletePortMappingRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[DeletePortMappingRequest]{
			In: in,
		},
	}
	out := &DeletePortMappingResponse{}
	soapResponse := &tr064.SOAPResponse[DeletePortMappingResponse]{
		Body: &tr064.SOAPResponseBody[DeletePortMappingResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "DeletePortMapping", soapRequest, soapResponse)
}

type GetExternalIPAddressRequest struct {
	XMLName      xml.Name `xml:"u:GetExternalIPAddressRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetExternalIPAddressResponse struct {
	XMLName              xml.Name `xml:"GetExternalIPAddressResponse"`
	NewExternalIPAddress string   `xml:"NewExternalIPAddress"`
}

func (client *ServiceClient) GetExternalIPAddress(out *GetExternalIPAddressResponse) error {
	in := &GetExternalIPAddressRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetExternalIPAddressRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetExternalIPAddressRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetExternalIPAddressResponse]{
		Body: &tr064.SOAPResponseBody[GetExternalIPAddressResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetExternalIPAddress", soapRequest, soapResponse)
}

type X_AVM_DE_GetExternalIPv6AddressRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM_DE_GetExternalIPv6AddressRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetExternalIPv6AddressResponse struct {
	XMLName                xml.Name `xml:"X_AVM_DE_GetExternalIPv6AddressResponse"`
	NewExternalIPv6Address string   `xml:"NewExternalIPv6Address"`
	NewPrefixLength        uint8    `xml:"NewPrefixLength"`
	NewValidLifetime       uint32   `xml:"NewValidLifetime"`
	NewPreferedLifetime    uint32   `xml:"NewPreferedLifetime"`
}

func (client *ServiceClient) X_AVM_DE_GetExternalIPv6Address(out *X_AVM_DE_GetExternalIPv6AddressResponse) error {
	in := &X_AVM_DE_GetExternalIPv6AddressRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetExternalIPv6AddressRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetExternalIPv6AddressRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetExternalIPv6AddressResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetExternalIPv6AddressResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM_DE_GetExternalIPv6Address", soapRequest, soapResponse)
}

type X_AVM_DE_GetIPv6PrefixRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM_DE_GetIPv6PrefixRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetIPv6PrefixResponse struct {
	XMLName             xml.Name `xml:"X_AVM_DE_GetIPv6PrefixResponse"`
	NewIPv6Prefix       string   `xml:"NewIPv6Prefix"`
	NewPrefixLength     uint8    `xml:"NewPrefixLength"`
	NewValidLifetime    uint32   `xml:"NewValidLifetime"`
	NewPreferedLifetime uint32   `xml:"NewPreferedLifetime"`
}

func (client *ServiceClient) X_AVM_DE_GetIPv6Prefix(out *X_AVM_DE_GetIPv6PrefixResponse) error {
	in := &X_AVM_DE_GetIPv6PrefixRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetIPv6PrefixRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetIPv6PrefixRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetIPv6PrefixResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetIPv6PrefixResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM_DE_GetIPv6Prefix", soapRequest, soapResponse)
}

type X_AVM_DE_GetDNSServerRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM_DE_GetDNSServerRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetDNSServerResponse struct {
	XMLName           xml.Name `xml:"X_AVM_DE_GetDNSServerResponse"`
	NewIPv4DNSServer1 string   `xml:"NewIPv4DNSServer1"`
	NewIPv4DNSServer2 string   `xml:"NewIPv4DNSServer2"`
}

func (client *ServiceClient) X_AVM_DE_GetDNSServer(out *X_AVM_DE_GetDNSServerResponse) error {
	in := &X_AVM_DE_GetDNSServerRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetDNSServerRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetDNSServerRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetDNSServerResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetDNSServerResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM_DE_GetDNSServer", soapRequest, soapResponse)
}

type X_AVM_DE_GetIPv6DNSServerRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM_DE_GetIPv6DNSServerRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetIPv6DNSServerResponse struct {
	XMLName           xml.Name `xml:"X_AVM_DE_GetIPv6DNSServerResponse"`
	NewIPv6DNSServer1 string   `xml:"NewIPv6DNSServer1"`
	NewValidLifetime1 uint32   `xml:"NewValidLifetime1"`
	NewIPv6DNSServer2 string   `xml:"NewIPv6DNSServer2"`
	NewValidLifetime  uint32   `xml:"NewValidLifetime"`
}

func (client *ServiceClient) X_AVM_DE_GetIPv6DNSServer(out *X_AVM_DE_GetIPv6DNSServerResponse) error {
	in := &X_AVM_DE_GetIPv6DNSServerRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetIPv6DNSServerRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetIPv6DNSServerRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetIPv6DNSServerResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetIPv6DNSServerResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM_DE_GetIPv6DNSServer", soapRequest, soapResponse)
}
