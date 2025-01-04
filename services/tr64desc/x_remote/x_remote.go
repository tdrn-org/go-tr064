// generated from spec version: 1.0
package x_remote

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
	XMLName               xml.Name `xml:"GetInfoResponse"`
	NewEnabled            bool     `xml:"NewEnabled"`
	NewPort               string   `xml:"NewPort"`
	NewUsername           string   `xml:"NewUsername"`
	NewLetsEncryptEnabled bool     `xml:"NewLetsEncryptEnabled"`
	NewLetsEncryptState   string   `xml:"NewLetsEncryptState"`
}

func (client *ServiceClient) GetInfo(out *GetInfoResponse) error {
	in := &GetInfoRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetInfo", soapRequest, soapResponse)
}

type SetConfigRequest struct {
	XMLName      xml.Name `xml:"u:SetConfigRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewEnabled   bool     `xml:"NewEnabled"`
	NewPort      string   `xml:"NewPort"`
	NewUsername  string   `xml:"NewUsername"`
	NewPassword  string   `xml:"NewPassword"`
}

type SetConfigResponse struct {
	XMLName xml.Name `xml:"SetConfigResponse"`
}

func (client *ServiceClient) SetConfig(in *SetConfigRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := tr064.NewSOAPRequest(in)
	out := &SetConfigResponse{}
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "SetConfig", soapRequest, soapResponse)
}

type SetEnableRequest struct {
	XMLName      xml.Name `xml:"u:SetEnableRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewEnabled   bool     `xml:"NewEnabled"`
}

type SetEnableResponse struct {
	XMLName xml.Name `xml:"SetEnableResponse"`
	NewPort string   `xml:"NewPort"`
}

func (client *ServiceClient) SetEnable(in *SetEnableRequest, out *SetEnableResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "SetEnable", soapRequest, soapResponse)
}

type SetLetsEncryptEnableRequest struct {
	XMLName               xml.Name `xml:"u:SetLetsEncryptEnableRequest"`
	XMLNameSpace          string   `xml:"xmlns:u,attr"`
	NewLetsEncryptEnabled bool     `xml:"NewLetsEncryptEnabled"`
}

type SetLetsEncryptEnableResponse struct {
	XMLName xml.Name `xml:"SetLetsEncryptEnableResponse"`
}

func (client *ServiceClient) SetLetsEncryptEnable(in *SetLetsEncryptEnableRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := tr064.NewSOAPRequest(in)
	out := &SetLetsEncryptEnableResponse{}
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "SetLetsEncryptEnable", soapRequest, soapResponse)
}

type GetDDNSInfoRequest struct {
	XMLName      xml.Name `xml:"u:GetDDNSInfoRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetDDNSInfoResponse struct {
	XMLName         xml.Name `xml:"GetDDNSInfoResponse"`
	NewEnabled      bool     `xml:"NewEnabled"`
	NewProviderName string   `xml:"NewProviderName"`
	NewUpdateURL    string   `xml:"NewUpdateURL"`
	NewDomain       string   `xml:"NewDomain"`
	NewStatusIPv4   string   `xml:"NewStatusIPv4"`
	NewStatusIPv6   string   `xml:"NewStatusIPv6"`
	NewUsername     string   `xml:"NewUsername"`
	NewMode         string   `xml:"NewMode"`
	NewServerIPv4   string   `xml:"NewServerIPv4"`
	NewServerIPv6   string   `xml:"NewServerIPv6"`
}

func (client *ServiceClient) GetDDNSInfo(out *GetDDNSInfoResponse) error {
	in := &GetDDNSInfoRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetDDNSInfo", soapRequest, soapResponse)
}

type GetDDNSProvidersRequest struct {
	XMLName      xml.Name `xml:"u:GetDDNSProvidersRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetDDNSProvidersResponse struct {
	XMLName         xml.Name `xml:"GetDDNSProvidersResponse"`
	NewProviderList string   `xml:"NewProviderList"`
}

func (client *ServiceClient) GetDDNSProviders(out *GetDDNSProvidersResponse) error {
	in := &GetDDNSProvidersRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetDDNSProviders", soapRequest, soapResponse)
}

type SetDDNSConfigRequest struct {
	XMLName         xml.Name `xml:"u:SetDDNSConfigRequest"`
	XMLNameSpace    string   `xml:"xmlns:u,attr"`
	NewEnabled      bool     `xml:"NewEnabled"`
	NewProviderName string   `xml:"NewProviderName"`
	NewUpdateURL    string   `xml:"NewUpdateURL"`
	NewDomain       string   `xml:"NewDomain"`
	NewUsername     string   `xml:"NewUsername"`
	NewMode         string   `xml:"NewMode"`
	NewServerIPv4   string   `xml:"NewServerIPv4"`
	NewServerIPv6   string   `xml:"NewServerIPv6"`
	NewPassword     string   `xml:"NewPassword"`
}

type SetDDNSConfigResponse struct {
	XMLName xml.Name `xml:"SetDDNSConfigResponse"`
}

func (client *ServiceClient) SetDDNSConfig(in *SetDDNSConfigRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := tr064.NewSOAPRequest(in)
	out := &SetDDNSConfigResponse{}
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "SetDDNSConfig", soapRequest, soapResponse)
}
