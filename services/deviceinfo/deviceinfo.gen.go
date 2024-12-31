// generated from spec version: 1.0
package deviceinfo

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
	NewManufacturerName string   `xml:"NewManufacturerName"`
	NewManufacturerOUI  string   `xml:"NewManufacturerOUI"`
	NewModelName        string   `xml:"NewModelName"`
	NewDescription      string   `xml:"NewDescription"`
	NewProductClass     string   `xml:"NewProductClass"`
	NewSerialNumber     string   `xml:"NewSerialNumber"`
	NewSoftwareVersion  string   `xml:"NewSoftwareVersion"`
	NewHardwareVersion  string   `xml:"NewHardwareVersion"`
	NewSpecVersion      string   `xml:"NewSpecVersion"`
	NewProvisioningCode string   `xml:"NewProvisioningCode"`
	NewUpTime           uint32   `xml:"NewUpTime"`
	NewDeviceLog        string   `xml:"NewDeviceLog"`
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

type SetProvisioningCodeRequest struct {
	XMLName             xml.Name `xml:"u:SetProvisioningCodeRequest"`
	XMLNameSpace        string   `xml:"xmlns:u,attr"`
	NewProvisioningCode string   `xml:"NewProvisioningCode"`
}

type SetProvisioningCodeResponse struct {
	XMLName xml.Name `xml:"SetProvisioningCodeResponse"`
}

func (client *ServiceClient) SetProvisioningCode(in *SetProvisioningCodeRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetProvisioningCodeRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[SetProvisioningCodeRequest]{
			In: in,
		},
	}
	out := &SetProvisioningCodeResponse{}
	soapResponse := &tr064.SOAPResponse[SetProvisioningCodeResponse]{
		Body: &tr064.SOAPResponseBody[SetProvisioningCodeResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetProvisioningCode", soapRequest, soapResponse)
}

type GetDeviceLogRequest struct {
	XMLName      xml.Name `xml:"u:GetDeviceLogRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetDeviceLogResponse struct {
	XMLName      xml.Name `xml:"GetDeviceLogResponse"`
	NewDeviceLog string   `xml:"NewDeviceLog"`
}

func (client *ServiceClient) GetDeviceLog(out *GetDeviceLogResponse) error {
	in := &GetDeviceLogRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetDeviceLogRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetDeviceLogRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetDeviceLogResponse]{
		Body: &tr064.SOAPResponseBody[GetDeviceLogResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetDeviceLog", soapRequest, soapResponse)
}

type GetSecurityPortRequest struct {
	XMLName      xml.Name `xml:"u:GetSecurityPortRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetSecurityPortResponse struct {
	XMLName         xml.Name `xml:"GetSecurityPortResponse"`
	NewSecurityPort uint16   `xml:"NewSecurityPort"`
}

func (client *ServiceClient) GetSecurityPort(out *GetSecurityPortResponse) error {
	in := &GetSecurityPortRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetSecurityPortRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetSecurityPortRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetSecurityPortResponse]{
		Body: &tr064.SOAPResponseBody[GetSecurityPortResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetSecurityPort", soapRequest, soapResponse)
}

type X_AVM_DE_GetDeviceLogPathRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetDeviceLogPathRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetDeviceLogPathResponse struct {
	XMLName          xml.Name `xml:"X_AVM-DE_GetDeviceLogPathResponse"`
	NewDeviceLogPath string   `xml:"NewDeviceLogPath"`
}

func (client *ServiceClient) X_AVM_DE_GetDeviceLogPath(out *X_AVM_DE_GetDeviceLogPathResponse) error {
	in := &X_AVM_DE_GetDeviceLogPathRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetDeviceLogPathRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetDeviceLogPathRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetDeviceLogPathResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetDeviceLogPathResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetDeviceLogPath", soapRequest, soapResponse)
}