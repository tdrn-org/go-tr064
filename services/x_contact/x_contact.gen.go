// generated from spec version: 1.0
package x_contact

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
	XMLName        xml.Name `xml:"GetInfoResponse"`
	NewEnable      bool     `xml:"NewEnable"`
	NewStatus      string   `xml:"NewStatus"`
	NewLastConnect string   `xml:"NewLastConnect"`
	NewUrl         string   `xml:"NewUrl"`
	NewServiceId   string   `xml:"NewServiceId"`
	NewUsername    string   `xml:"NewUsername"`
	NewName        string   `xml:"NewName"`
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

type SetEnableRequest struct {
	XMLName      xml.Name `xml:"u:SetEnableRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewEnable    bool     `xml:"NewEnable"`
}

type SetEnableResponse struct {
	XMLName xml.Name `xml:"SetEnableResponse"`
}

func (client *ServiceClient) SetEnable(in *SetEnableRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetEnableRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[SetEnableRequest]{
			In: in,
		},
	}
	out := &SetEnableResponse{}
	soapResponse := &tr064.SOAPResponse[SetEnableResponse]{
		Body: &tr064.SOAPResponseBody[SetEnableResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetEnable", soapRequest, soapResponse)
}

type SetConfigRequest struct {
	XMLName      xml.Name `xml:"u:SetConfigRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewEnable    bool     `xml:"NewEnable"`
	NewUrl       string   `xml:"NewUrl"`
	NewServiceId string   `xml:"NewServiceId"`
	NewUsername  string   `xml:"NewUsername"`
	NewPassword  string   `xml:"NewPassword"`
	NewName      string   `xml:"NewName"`
}

type SetConfigResponse struct {
	XMLName xml.Name `xml:"SetConfigResponse"`
}

func (client *ServiceClient) SetConfig(in *SetConfigRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetConfigRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[SetConfigRequest]{
			In: in,
		},
	}
	out := &SetConfigResponse{}
	soapResponse := &tr064.SOAPResponse[SetConfigResponse]{
		Body: &tr064.SOAPResponseBody[SetConfigResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetConfig", soapRequest, soapResponse)
}

type GetInfoByIndexRequest struct {
	XMLName      xml.Name `xml:"u:GetInfoByIndexRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewIndex     uint16   `xml:"NewIndex"`
}

type GetInfoByIndexResponse struct {
	XMLName        xml.Name `xml:"GetInfoByIndexResponse"`
	NewEnable      bool     `xml:"NewEnable"`
	NewStatus      string   `xml:"NewStatus"`
	NewLastConnect string   `xml:"NewLastConnect"`
	NewUrl         string   `xml:"NewUrl"`
	NewServiceId   string   `xml:"NewServiceId"`
	NewUsername    string   `xml:"NewUsername"`
	NewName        string   `xml:"NewName"`
}

func (client *ServiceClient) GetInfoByIndex(in *GetInfoByIndexRequest, out *GetInfoByIndexResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[GetInfoByIndexRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetInfoByIndexRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetInfoByIndexResponse]{
		Body: &tr064.SOAPResponseBody[GetInfoByIndexResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetInfoByIndex", soapRequest, soapResponse)
}

type SetEnableByIndexRequest struct {
	XMLName      xml.Name `xml:"u:SetEnableByIndexRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewIndex     uint16   `xml:"NewIndex"`
	NewEnable    bool     `xml:"NewEnable"`
}

type SetEnableByIndexResponse struct {
	XMLName xml.Name `xml:"SetEnableByIndexResponse"`
}

func (client *ServiceClient) SetEnableByIndex(in *SetEnableByIndexRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetEnableByIndexRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[SetEnableByIndexRequest]{
			In: in,
		},
	}
	out := &SetEnableByIndexResponse{}
	soapResponse := &tr064.SOAPResponse[SetEnableByIndexResponse]{
		Body: &tr064.SOAPResponseBody[SetEnableByIndexResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetEnableByIndex", soapRequest, soapResponse)
}

type SetConfigByIndexRequest struct {
	XMLName      xml.Name `xml:"u:SetConfigByIndexRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewIndex     uint16   `xml:"NewIndex"`
	NewEnable    bool     `xml:"NewEnable"`
	NewUrl       string   `xml:"NewUrl"`
	NewServiceId string   `xml:"NewServiceId"`
	NewUsername  string   `xml:"NewUsername"`
	NewPassword  string   `xml:"NewPassword"`
	NewName      string   `xml:"NewName"`
}

type SetConfigByIndexResponse struct {
	XMLName xml.Name `xml:"SetConfigByIndexResponse"`
}

func (client *ServiceClient) SetConfigByIndex(in *SetConfigByIndexRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetConfigByIndexRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[SetConfigByIndexRequest]{
			In: in,
		},
	}
	out := &SetConfigByIndexResponse{}
	soapResponse := &tr064.SOAPResponse[SetConfigByIndexResponse]{
		Body: &tr064.SOAPResponseBody[SetConfigByIndexResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetConfigByIndex", soapRequest, soapResponse)
}

type DeleteByIndexRequest struct {
	XMLName      xml.Name `xml:"u:DeleteByIndexRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewIndex     uint16   `xml:"NewIndex"`
}

type DeleteByIndexResponse struct {
	XMLName xml.Name `xml:"DeleteByIndexResponse"`
}

func (client *ServiceClient) DeleteByIndex(in *DeleteByIndexRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[DeleteByIndexRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[DeleteByIndexRequest]{
			In: in,
		},
	}
	out := &DeleteByIndexResponse{}
	soapResponse := &tr064.SOAPResponse[DeleteByIndexResponse]{
		Body: &tr064.SOAPResponseBody[DeleteByIndexResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "DeleteByIndex", soapRequest, soapResponse)
}

type GetNumberOfEntriesRequest struct {
	XMLName      xml.Name `xml:"u:GetNumberOfEntriesRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetNumberOfEntriesResponse struct {
	XMLName                 xml.Name `xml:"GetNumberOfEntriesResponse"`
	NewOnTelNumberOfEntries uint16   `xml:"NewOnTelNumberOfEntries"`
}

func (client *ServiceClient) GetNumberOfEntries(out *GetNumberOfEntriesResponse) error {
	in := &GetNumberOfEntriesRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetNumberOfEntriesRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetNumberOfEntriesRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetNumberOfEntriesResponse]{
		Body: &tr064.SOAPResponseBody[GetNumberOfEntriesResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetNumberOfEntries", soapRequest, soapResponse)
}

type GetCallListRequest struct {
	XMLName      xml.Name `xml:"u:GetCallListRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetCallListResponse struct {
	XMLName        xml.Name `xml:"GetCallListResponse"`
	NewCallListURL string   `xml:"NewCallListURL"`
}

func (client *ServiceClient) GetCallList(out *GetCallListResponse) error {
	in := &GetCallListRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetCallListRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetCallListRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetCallListResponse]{
		Body: &tr064.SOAPResponseBody[GetCallListResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetCallList", soapRequest, soapResponse)
}

type GetPhonebookListRequest struct {
	XMLName      xml.Name `xml:"u:GetPhonebookListRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetPhonebookListResponse struct {
	XMLName          xml.Name `xml:"GetPhonebookListResponse"`
	NewPhonebookList string   `xml:"NewPhonebookList"`
}

func (client *ServiceClient) GetPhonebookList(out *GetPhonebookListResponse) error {
	in := &GetPhonebookListRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetPhonebookListRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetPhonebookListRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetPhonebookListResponse]{
		Body: &tr064.SOAPResponseBody[GetPhonebookListResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetPhonebookList", soapRequest, soapResponse)
}

type GetPhonebookRequest struct {
	XMLName        xml.Name `xml:"u:GetPhonebookRequest"`
	XMLNameSpace   string   `xml:"xmlns:u,attr"`
	NewPhonebookID uint16   `xml:"NewPhonebookID"`
}

type GetPhonebookResponse struct {
	XMLName             xml.Name `xml:"GetPhonebookResponse"`
	NewPhonebookName    string   `xml:"NewPhonebookName"`
	NewPhonebookExtraID string   `xml:"NewPhonebookExtraID"`
	NewPhonebookURL     string   `xml:"NewPhonebookURL"`
}

func (client *ServiceClient) GetPhonebook(in *GetPhonebookRequest, out *GetPhonebookResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[GetPhonebookRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetPhonebookRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetPhonebookResponse]{
		Body: &tr064.SOAPResponseBody[GetPhonebookResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetPhonebook", soapRequest, soapResponse)
}

type AddPhonebookRequest struct {
	XMLName             xml.Name `xml:"u:AddPhonebookRequest"`
	XMLNameSpace        string   `xml:"xmlns:u,attr"`
	NewPhonebookExtraID string   `xml:"NewPhonebookExtraID"`
	NewPhonebookName    string   `xml:"NewPhonebookName"`
}

type AddPhonebookResponse struct {
	XMLName xml.Name `xml:"AddPhonebookResponse"`
}

func (client *ServiceClient) AddPhonebook(in *AddPhonebookRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[AddPhonebookRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[AddPhonebookRequest]{
			In: in,
		},
	}
	out := &AddPhonebookResponse{}
	soapResponse := &tr064.SOAPResponse[AddPhonebookResponse]{
		Body: &tr064.SOAPResponseBody[AddPhonebookResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "AddPhonebook", soapRequest, soapResponse)
}

type DeletePhonebookRequest struct {
	XMLName             xml.Name `xml:"u:DeletePhonebookRequest"`
	XMLNameSpace        string   `xml:"xmlns:u,attr"`
	NewPhonebookID      uint16   `xml:"NewPhonebookID"`
	NewPhonebookExtraID string   `xml:"NewPhonebookExtraID"`
}

type DeletePhonebookResponse struct {
	XMLName xml.Name `xml:"DeletePhonebookResponse"`
}

func (client *ServiceClient) DeletePhonebook(in *DeletePhonebookRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[DeletePhonebookRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[DeletePhonebookRequest]{
			In: in,
		},
	}
	out := &DeletePhonebookResponse{}
	soapResponse := &tr064.SOAPResponse[DeletePhonebookResponse]{
		Body: &tr064.SOAPResponseBody[DeletePhonebookResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "DeletePhonebook", soapRequest, soapResponse)
}

type GetPhonebookEntryRequest struct {
	XMLName             xml.Name `xml:"u:GetPhonebookEntryRequest"`
	XMLNameSpace        string   `xml:"xmlns:u,attr"`
	NewPhonebookID      uint16   `xml:"NewPhonebookID"`
	NewPhonebookEntryID uint32   `xml:"NewPhonebookEntryID"`
}

type GetPhonebookEntryResponse struct {
	XMLName               xml.Name `xml:"GetPhonebookEntryResponse"`
	NewPhonebookEntryData string   `xml:"NewPhonebookEntryData"`
}

func (client *ServiceClient) GetPhonebookEntry(in *GetPhonebookEntryRequest, out *GetPhonebookEntryResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[GetPhonebookEntryRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetPhonebookEntryRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetPhonebookEntryResponse]{
		Body: &tr064.SOAPResponseBody[GetPhonebookEntryResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetPhonebookEntry", soapRequest, soapResponse)
}

type GetPhonebookEntryUIDRequest struct {
	XMLName                   xml.Name `xml:"u:GetPhonebookEntryUIDRequest"`
	XMLNameSpace              string   `xml:"xmlns:u,attr"`
	NewPhonebookID            uint16   `xml:"NewPhonebookID"`
	NewPhonebookEntryUniqueID uint32   `xml:"NewPhonebookEntryUniqueID"`
}

type GetPhonebookEntryUIDResponse struct {
	XMLName               xml.Name `xml:"GetPhonebookEntryUIDResponse"`
	NewPhonebookEntryData string   `xml:"NewPhonebookEntryData"`
}

func (client *ServiceClient) GetPhonebookEntryUID(in *GetPhonebookEntryUIDRequest, out *GetPhonebookEntryUIDResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[GetPhonebookEntryUIDRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetPhonebookEntryUIDRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetPhonebookEntryUIDResponse]{
		Body: &tr064.SOAPResponseBody[GetPhonebookEntryUIDResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetPhonebookEntryUID", soapRequest, soapResponse)
}

type SetPhonebookEntryRequest struct {
	XMLName               xml.Name `xml:"u:SetPhonebookEntryRequest"`
	XMLNameSpace          string   `xml:"xmlns:u,attr"`
	NewPhonebookID        uint16   `xml:"NewPhonebookID"`
	NewPhonebookEntryID   uint32   `xml:"NewPhonebookEntryID"`
	NewPhonebookEntryData string   `xml:"NewPhonebookEntryData"`
}

type SetPhonebookEntryResponse struct {
	XMLName xml.Name `xml:"SetPhonebookEntryResponse"`
}

func (client *ServiceClient) SetPhonebookEntry(in *SetPhonebookEntryRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetPhonebookEntryRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[SetPhonebookEntryRequest]{
			In: in,
		},
	}
	out := &SetPhonebookEntryResponse{}
	soapResponse := &tr064.SOAPResponse[SetPhonebookEntryResponse]{
		Body: &tr064.SOAPResponseBody[SetPhonebookEntryResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetPhonebookEntry", soapRequest, soapResponse)
}

type SetPhonebookEntryUIDRequest struct {
	XMLName               xml.Name `xml:"u:SetPhonebookEntryUIDRequest"`
	XMLNameSpace          string   `xml:"xmlns:u,attr"`
	NewPhonebookID        uint16   `xml:"NewPhonebookID"`
	NewPhonebookEntryData string   `xml:"NewPhonebookEntryData"`
}

type SetPhonebookEntryUIDResponse struct {
	XMLName                   xml.Name `xml:"SetPhonebookEntryUIDResponse"`
	NewPhonebookEntryUniqueID uint32   `xml:"NewPhonebookEntryUniqueID"`
}

func (client *ServiceClient) SetPhonebookEntryUID(in *SetPhonebookEntryUIDRequest, out *SetPhonebookEntryUIDResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetPhonebookEntryUIDRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[SetPhonebookEntryUIDRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[SetPhonebookEntryUIDResponse]{
		Body: &tr064.SOAPResponseBody[SetPhonebookEntryUIDResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetPhonebookEntryUID", soapRequest, soapResponse)
}

type DeletePhonebookEntryRequest struct {
	XMLName             xml.Name `xml:"u:DeletePhonebookEntryRequest"`
	XMLNameSpace        string   `xml:"xmlns:u,attr"`
	NewPhonebookID      uint16   `xml:"NewPhonebookID"`
	NewPhonebookEntryID uint32   `xml:"NewPhonebookEntryID"`
}

type DeletePhonebookEntryResponse struct {
	XMLName xml.Name `xml:"DeletePhonebookEntryResponse"`
}

func (client *ServiceClient) DeletePhonebookEntry(in *DeletePhonebookEntryRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[DeletePhonebookEntryRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[DeletePhonebookEntryRequest]{
			In: in,
		},
	}
	out := &DeletePhonebookEntryResponse{}
	soapResponse := &tr064.SOAPResponse[DeletePhonebookEntryResponse]{
		Body: &tr064.SOAPResponseBody[DeletePhonebookEntryResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "DeletePhonebookEntry", soapRequest, soapResponse)
}

type DeletePhonebookEntryUIDRequest struct {
	XMLName                   xml.Name `xml:"u:DeletePhonebookEntryUIDRequest"`
	XMLNameSpace              string   `xml:"xmlns:u,attr"`
	NewPhonebookID            uint16   `xml:"NewPhonebookID"`
	NewPhonebookEntryUniqueID uint32   `xml:"NewPhonebookEntryUniqueID"`
}

type DeletePhonebookEntryUIDResponse struct {
	XMLName xml.Name `xml:"DeletePhonebookEntryUIDResponse"`
}

func (client *ServiceClient) DeletePhonebookEntryUID(in *DeletePhonebookEntryUIDRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[DeletePhonebookEntryUIDRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[DeletePhonebookEntryUIDRequest]{
			In: in,
		},
	}
	out := &DeletePhonebookEntryUIDResponse{}
	soapResponse := &tr064.SOAPResponse[DeletePhonebookEntryUIDResponse]{
		Body: &tr064.SOAPResponseBody[DeletePhonebookEntryUIDResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "DeletePhonebookEntryUID", soapRequest, soapResponse)
}

type GetCallBarringEntryRequest struct {
	XMLName             xml.Name `xml:"u:GetCallBarringEntryRequest"`
	XMLNameSpace        string   `xml:"xmlns:u,attr"`
	NewPhonebookEntryID uint32   `xml:"NewPhonebookEntryID"`
}

type GetCallBarringEntryResponse struct {
	XMLName               xml.Name `xml:"GetCallBarringEntryResponse"`
	NewPhonebookEntryData string   `xml:"NewPhonebookEntryData"`
}

func (client *ServiceClient) GetCallBarringEntry(in *GetCallBarringEntryRequest, out *GetCallBarringEntryResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[GetCallBarringEntryRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetCallBarringEntryRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetCallBarringEntryResponse]{
		Body: &tr064.SOAPResponseBody[GetCallBarringEntryResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetCallBarringEntry", soapRequest, soapResponse)
}

type GetCallBarringEntryByNumRequest struct {
	XMLName      xml.Name `xml:"u:GetCallBarringEntryByNumRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewNumber    string   `xml:"NewNumber"`
}

type GetCallBarringEntryByNumResponse struct {
	XMLName               xml.Name `xml:"GetCallBarringEntryByNumResponse"`
	NewPhonebookEntryData string   `xml:"NewPhonebookEntryData"`
}

func (client *ServiceClient) GetCallBarringEntryByNum(in *GetCallBarringEntryByNumRequest, out *GetCallBarringEntryByNumResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[GetCallBarringEntryByNumRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetCallBarringEntryByNumRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetCallBarringEntryByNumResponse]{
		Body: &tr064.SOAPResponseBody[GetCallBarringEntryByNumResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetCallBarringEntryByNum", soapRequest, soapResponse)
}

type GetCallBarringListRequest struct {
	XMLName      xml.Name `xml:"u:GetCallBarringListRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetCallBarringListResponse struct {
	XMLName         xml.Name `xml:"GetCallBarringListResponse"`
	NewPhonebookURL string   `xml:"NewPhonebookURL"`
}

func (client *ServiceClient) GetCallBarringList(out *GetCallBarringListResponse) error {
	in := &GetCallBarringListRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetCallBarringListRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetCallBarringListRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetCallBarringListResponse]{
		Body: &tr064.SOAPResponseBody[GetCallBarringListResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetCallBarringList", soapRequest, soapResponse)
}

type SetCallBarringEntryRequest struct {
	XMLName               xml.Name `xml:"u:SetCallBarringEntryRequest"`
	XMLNameSpace          string   `xml:"xmlns:u,attr"`
	NewPhonebookEntryData string   `xml:"NewPhonebookEntryData"`
}

type SetCallBarringEntryResponse struct {
	XMLName                   xml.Name `xml:"SetCallBarringEntryResponse"`
	NewPhonebookEntryUniqueID uint32   `xml:"NewPhonebookEntryUniqueID"`
}

func (client *ServiceClient) SetCallBarringEntry(in *SetCallBarringEntryRequest, out *SetCallBarringEntryResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetCallBarringEntryRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[SetCallBarringEntryRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[SetCallBarringEntryResponse]{
		Body: &tr064.SOAPResponseBody[SetCallBarringEntryResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetCallBarringEntry", soapRequest, soapResponse)
}

type DeleteCallBarringEntryUIDRequest struct {
	XMLName                   xml.Name `xml:"u:DeleteCallBarringEntryUIDRequest"`
	XMLNameSpace              string   `xml:"xmlns:u,attr"`
	NewPhonebookEntryUniqueID uint32   `xml:"NewPhonebookEntryUniqueID"`
}

type DeleteCallBarringEntryUIDResponse struct {
	XMLName xml.Name `xml:"DeleteCallBarringEntryUIDResponse"`
}

func (client *ServiceClient) DeleteCallBarringEntryUID(in *DeleteCallBarringEntryUIDRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[DeleteCallBarringEntryUIDRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[DeleteCallBarringEntryUIDRequest]{
			In: in,
		},
	}
	out := &DeleteCallBarringEntryUIDResponse{}
	soapResponse := &tr064.SOAPResponse[DeleteCallBarringEntryUIDResponse]{
		Body: &tr064.SOAPResponseBody[DeleteCallBarringEntryUIDResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "DeleteCallBarringEntryUID", soapRequest, soapResponse)
}

type GetDECTHandsetListRequest struct {
	XMLName      xml.Name `xml:"u:GetDECTHandsetListRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetDECTHandsetListResponse struct {
	XMLName       xml.Name `xml:"GetDECTHandsetListResponse"`
	NewDectIDList string   `xml:"NewDectIDList"`
}

func (client *ServiceClient) GetDECTHandsetList(out *GetDECTHandsetListResponse) error {
	in := &GetDECTHandsetListRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetDECTHandsetListRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetDECTHandsetListRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetDECTHandsetListResponse]{
		Body: &tr064.SOAPResponseBody[GetDECTHandsetListResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetDECTHandsetList", soapRequest, soapResponse)
}

type GetDECTHandsetInfoRequest struct {
	XMLName      xml.Name `xml:"u:GetDECTHandsetInfoRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewDectID    uint16   `xml:"NewDectID"`
}

type GetDECTHandsetInfoResponse struct {
	XMLName        xml.Name `xml:"GetDECTHandsetInfoResponse"`
	NewHandsetName string   `xml:"NewHandsetName"`
	NewPhonebookID uint16   `xml:"NewPhonebookID"`
}

func (client *ServiceClient) GetDECTHandsetInfo(in *GetDECTHandsetInfoRequest, out *GetDECTHandsetInfoResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[GetDECTHandsetInfoRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetDECTHandsetInfoRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetDECTHandsetInfoResponse]{
		Body: &tr064.SOAPResponseBody[GetDECTHandsetInfoResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetDECTHandsetInfo", soapRequest, soapResponse)
}

type SetDECTHandsetPhonebookRequest struct {
	XMLName        xml.Name `xml:"u:SetDECTHandsetPhonebookRequest"`
	XMLNameSpace   string   `xml:"xmlns:u,attr"`
	NewDectID      uint16   `xml:"NewDectID"`
	NewPhonebookID uint16   `xml:"NewPhonebookID"`
}

type SetDECTHandsetPhonebookResponse struct {
	XMLName xml.Name `xml:"SetDECTHandsetPhonebookResponse"`
}

func (client *ServiceClient) SetDECTHandsetPhonebook(in *SetDECTHandsetPhonebookRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetDECTHandsetPhonebookRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[SetDECTHandsetPhonebookRequest]{
			In: in,
		},
	}
	out := &SetDECTHandsetPhonebookResponse{}
	soapResponse := &tr064.SOAPResponse[SetDECTHandsetPhonebookResponse]{
		Body: &tr064.SOAPResponseBody[SetDECTHandsetPhonebookResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetDECTHandsetPhonebook", soapRequest, soapResponse)
}

type GetNumberOfDeflectionsRequest struct {
	XMLName      xml.Name `xml:"u:GetNumberOfDeflectionsRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetNumberOfDeflectionsResponse struct {
	XMLName                xml.Name `xml:"GetNumberOfDeflectionsResponse"`
	NewNumberOfDeflections uint16   `xml:"NewNumberOfDeflections"`
}

func (client *ServiceClient) GetNumberOfDeflections(out *GetNumberOfDeflectionsResponse) error {
	in := &GetNumberOfDeflectionsRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetNumberOfDeflectionsRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetNumberOfDeflectionsRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetNumberOfDeflectionsResponse]{
		Body: &tr064.SOAPResponseBody[GetNumberOfDeflectionsResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetNumberOfDeflections", soapRequest, soapResponse)
}

type GetDeflectionRequest struct {
	XMLName         xml.Name `xml:"u:GetDeflectionRequest"`
	XMLNameSpace    string   `xml:"xmlns:u,attr"`
	NewDeflectionId uint16   `xml:"NewDeflectionId"`
}

type GetDeflectionResponse struct {
	XMLName               xml.Name `xml:"GetDeflectionResponse"`
	NewEnable             bool     `xml:"NewEnable"`
	NewType               string   `xml:"NewType"`
	NewNumber             string   `xml:"NewNumber"`
	NewDeflectionToNumber string   `xml:"NewDeflectionToNumber"`
	NewMode               string   `xml:"NewMode"`
	NewOutgoing           string   `xml:"NewOutgoing"`
	NewPhonebookID        uint16   `xml:"NewPhonebookID"`
}

func (client *ServiceClient) GetDeflection(in *GetDeflectionRequest, out *GetDeflectionResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[GetDeflectionRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetDeflectionRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetDeflectionResponse]{
		Body: &tr064.SOAPResponseBody[GetDeflectionResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetDeflection", soapRequest, soapResponse)
}

type GetDeflectionsRequest struct {
	XMLName      xml.Name `xml:"u:GetDeflectionsRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetDeflectionsResponse struct {
	XMLName           xml.Name `xml:"GetDeflectionsResponse"`
	NewDeflectionList string   `xml:"NewDeflectionList"`
}

func (client *ServiceClient) GetDeflections(out *GetDeflectionsResponse) error {
	in := &GetDeflectionsRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetDeflectionsRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetDeflectionsRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetDeflectionsResponse]{
		Body: &tr064.SOAPResponseBody[GetDeflectionsResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetDeflections", soapRequest, soapResponse)
}

type SetDeflectionEnableRequest struct {
	XMLName         xml.Name `xml:"u:SetDeflectionEnableRequest"`
	XMLNameSpace    string   `xml:"xmlns:u,attr"`
	NewDeflectionId uint16   `xml:"NewDeflectionId"`
	NewEnable       bool     `xml:"NewEnable"`
}

type SetDeflectionEnableResponse struct {
	XMLName xml.Name `xml:"SetDeflectionEnableResponse"`
}

func (client *ServiceClient) SetDeflectionEnable(in *SetDeflectionEnableRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetDeflectionEnableRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[SetDeflectionEnableRequest]{
			In: in,
		},
	}
	out := &SetDeflectionEnableResponse{}
	soapResponse := &tr064.SOAPResponse[SetDeflectionEnableResponse]{
		Body: &tr064.SOAPResponseBody[SetDeflectionEnableResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetDeflectionEnable", soapRequest, soapResponse)
}
