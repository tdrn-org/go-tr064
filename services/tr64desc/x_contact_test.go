// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/tr64desc/x_contact"
	"log"
	"net/http"
	"testing"
)

var x_contactMock = &mock.ServiceMock{
	Path:       "/upnp/control/x_contact",
	HandleFunc: x_contactHandler,
}

func TestX_AVM_DE_OnTel(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", x_contactMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server(), tr064.TR064Spec("tr64desc"))
	client.Debug = true
	serviceClient := &x_contact.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceName: "X_AVM_DE_OnTel",
			ServiceType: "urn:dslforum-org:service:X_AVM-DE_OnTel:1",
			ServiceId:   "urn:X_AVM-DE_OnTel-com:serviceId:X_AVM-DE_OnTel1",
			ServiceUrl:  "/upnp/control/x_contact",
		},
	}
	{
		out := &x_contact.GetInfoResponse{}
		require.NoError(t, serviceClient.GetInfo(out))
	}
	{
		in := &x_contact.SetEnableRequest{}
		require.NoError(t, serviceClient.SetEnable(in))
	}
	{
		in := &x_contact.SetConfigRequest{}
		require.NoError(t, serviceClient.SetConfig(in))
	}
	{
		in := &x_contact.GetInfoByIndexRequest{}
		out := &x_contact.GetInfoByIndexResponse{}
		require.NoError(t, serviceClient.GetInfoByIndex(in, out))
	}
	{
		in := &x_contact.SetEnableByIndexRequest{}
		require.NoError(t, serviceClient.SetEnableByIndex(in))
	}
	{
		in := &x_contact.SetConfigByIndexRequest{}
		require.NoError(t, serviceClient.SetConfigByIndex(in))
	}
	{
		in := &x_contact.DeleteByIndexRequest{}
		require.NoError(t, serviceClient.DeleteByIndex(in))
	}
	{
		out := &x_contact.GetNumberOfEntriesResponse{}
		require.NoError(t, serviceClient.GetNumberOfEntries(out))
	}
	{
		out := &x_contact.GetCallListResponse{}
		require.NoError(t, serviceClient.GetCallList(out))
	}
	{
		out := &x_contact.GetPhonebookListResponse{}
		require.NoError(t, serviceClient.GetPhonebookList(out))
	}
	{
		in := &x_contact.GetPhonebookRequest{}
		out := &x_contact.GetPhonebookResponse{}
		require.NoError(t, serviceClient.GetPhonebook(in, out))
	}
	{
		in := &x_contact.AddPhonebookRequest{}
		require.NoError(t, serviceClient.AddPhonebook(in))
	}
	{
		in := &x_contact.DeletePhonebookRequest{}
		require.NoError(t, serviceClient.DeletePhonebook(in))
	}
	{
		in := &x_contact.GetPhonebookEntryRequest{}
		out := &x_contact.GetPhonebookEntryResponse{}
		require.NoError(t, serviceClient.GetPhonebookEntry(in, out))
	}
	{
		in := &x_contact.GetPhonebookEntryUIDRequest{}
		out := &x_contact.GetPhonebookEntryUIDResponse{}
		require.NoError(t, serviceClient.GetPhonebookEntryUID(in, out))
	}
	{
		in := &x_contact.SetPhonebookEntryRequest{}
		require.NoError(t, serviceClient.SetPhonebookEntry(in))
	}
	{
		in := &x_contact.SetPhonebookEntryUIDRequest{}
		out := &x_contact.SetPhonebookEntryUIDResponse{}
		require.NoError(t, serviceClient.SetPhonebookEntryUID(in, out))
	}
	{
		in := &x_contact.DeletePhonebookEntryRequest{}
		require.NoError(t, serviceClient.DeletePhonebookEntry(in))
	}
	{
		in := &x_contact.DeletePhonebookEntryUIDRequest{}
		require.NoError(t, serviceClient.DeletePhonebookEntryUID(in))
	}
	{
		in := &x_contact.GetCallBarringEntryRequest{}
		out := &x_contact.GetCallBarringEntryResponse{}
		require.NoError(t, serviceClient.GetCallBarringEntry(in, out))
	}
	{
		in := &x_contact.GetCallBarringEntryByNumRequest{}
		out := &x_contact.GetCallBarringEntryByNumResponse{}
		require.NoError(t, serviceClient.GetCallBarringEntryByNum(in, out))
	}
	{
		out := &x_contact.GetCallBarringListResponse{}
		require.NoError(t, serviceClient.GetCallBarringList(out))
	}
	{
		in := &x_contact.SetCallBarringEntryRequest{}
		out := &x_contact.SetCallBarringEntryResponse{}
		require.NoError(t, serviceClient.SetCallBarringEntry(in, out))
	}
	{
		in := &x_contact.DeleteCallBarringEntryUIDRequest{}
		require.NoError(t, serviceClient.DeleteCallBarringEntryUID(in))
	}
	{
		out := &x_contact.GetDECTHandsetListResponse{}
		require.NoError(t, serviceClient.GetDECTHandsetList(out))
	}
	{
		in := &x_contact.GetDECTHandsetInfoRequest{}
		out := &x_contact.GetDECTHandsetInfoResponse{}
		require.NoError(t, serviceClient.GetDECTHandsetInfo(in, out))
	}
	{
		in := &x_contact.SetDECTHandsetPhonebookRequest{}
		require.NoError(t, serviceClient.SetDECTHandsetPhonebook(in))
	}
	{
		out := &x_contact.GetNumberOfDeflectionsResponse{}
		require.NoError(t, serviceClient.GetNumberOfDeflections(out))
	}
	{
		in := &x_contact.GetDeflectionRequest{}
		out := &x_contact.GetDeflectionResponse{}
		require.NoError(t, serviceClient.GetDeflection(in, out))
	}
	{
		out := &x_contact.GetDeflectionsResponse{}
		require.NoError(t, serviceClient.GetDeflections(out))
	}
	{
		in := &x_contact.SetDeflectionEnableRequest{}
		require.NoError(t, serviceClient.SetDeflectionEnable(in))
	}
}

func x_contactHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetInfo":
		x_contact_GetInfo(w)
	case "SetEnable":
		x_contact_SetEnable(w)
	case "SetConfig":
		x_contact_SetConfig(w)
	case "GetInfoByIndex":
		x_contact_GetInfoByIndex(w)
	case "SetEnableByIndex":
		x_contact_SetEnableByIndex(w)
	case "SetConfigByIndex":
		x_contact_SetConfigByIndex(w)
	case "DeleteByIndex":
		x_contact_DeleteByIndex(w)
	case "GetNumberOfEntries":
		x_contact_GetNumberOfEntries(w)
	case "GetCallList":
		x_contact_GetCallList(w)
	case "GetPhonebookList":
		x_contact_GetPhonebookList(w)
	case "GetPhonebook":
		x_contact_GetPhonebook(w)
	case "AddPhonebook":
		x_contact_AddPhonebook(w)
	case "DeletePhonebook":
		x_contact_DeletePhonebook(w)
	case "GetPhonebookEntry":
		x_contact_GetPhonebookEntry(w)
	case "GetPhonebookEntryUID":
		x_contact_GetPhonebookEntryUID(w)
	case "SetPhonebookEntry":
		x_contact_SetPhonebookEntry(w)
	case "SetPhonebookEntryUID":
		x_contact_SetPhonebookEntryUID(w)
	case "DeletePhonebookEntry":
		x_contact_DeletePhonebookEntry(w)
	case "DeletePhonebookEntryUID":
		x_contact_DeletePhonebookEntryUID(w)
	case "GetCallBarringEntry":
		x_contact_GetCallBarringEntry(w)
	case "GetCallBarringEntryByNum":
		x_contact_GetCallBarringEntryByNum(w)
	case "GetCallBarringList":
		x_contact_GetCallBarringList(w)
	case "SetCallBarringEntry":
		x_contact_SetCallBarringEntry(w)
	case "DeleteCallBarringEntryUID":
		x_contact_DeleteCallBarringEntryUID(w)
	case "GetDECTHandsetList":
		x_contact_GetDECTHandsetList(w)
	case "GetDECTHandsetInfo":
		x_contact_GetDECTHandsetInfo(w)
	case "SetDECTHandsetPhonebook":
		x_contact_SetDECTHandsetPhonebook(w)
	case "GetNumberOfDeflections":
		x_contact_GetNumberOfDeflections(w)
	case "GetDeflection":
		x_contact_GetDeflection(w)
	case "GetDeflections":
		x_contact_GetDeflections(w)
	case "SetDeflectionEnable":
		x_contact_SetDeflectionEnable(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func x_contact_GetInfo(w http.ResponseWriter) {
	out := x_contact.GetInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_contact_SetEnable(w http.ResponseWriter) {
	out := x_contact.SetEnableResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_contact_SetConfig(w http.ResponseWriter) {
	out := x_contact.SetConfigResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_contact_GetInfoByIndex(w http.ResponseWriter) {
	out := x_contact.GetInfoByIndexResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_contact_SetEnableByIndex(w http.ResponseWriter) {
	out := x_contact.SetEnableByIndexResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_contact_SetConfigByIndex(w http.ResponseWriter) {
	out := x_contact.SetConfigByIndexResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_contact_DeleteByIndex(w http.ResponseWriter) {
	out := x_contact.DeleteByIndexResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_contact_GetNumberOfEntries(w http.ResponseWriter) {
	out := x_contact.GetNumberOfEntriesResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_contact_GetCallList(w http.ResponseWriter) {
	out := x_contact.GetCallListResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_contact_GetPhonebookList(w http.ResponseWriter) {
	out := x_contact.GetPhonebookListResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_contact_GetPhonebook(w http.ResponseWriter) {
	out := x_contact.GetPhonebookResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_contact_AddPhonebook(w http.ResponseWriter) {
	out := x_contact.AddPhonebookResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_contact_DeletePhonebook(w http.ResponseWriter) {
	out := x_contact.DeletePhonebookResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_contact_GetPhonebookEntry(w http.ResponseWriter) {
	out := x_contact.GetPhonebookEntryResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_contact_GetPhonebookEntryUID(w http.ResponseWriter) {
	out := x_contact.GetPhonebookEntryUIDResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_contact_SetPhonebookEntry(w http.ResponseWriter) {
	out := x_contact.SetPhonebookEntryResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_contact_SetPhonebookEntryUID(w http.ResponseWriter) {
	out := x_contact.SetPhonebookEntryUIDResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_contact_DeletePhonebookEntry(w http.ResponseWriter) {
	out := x_contact.DeletePhonebookEntryResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_contact_DeletePhonebookEntryUID(w http.ResponseWriter) {
	out := x_contact.DeletePhonebookEntryUIDResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_contact_GetCallBarringEntry(w http.ResponseWriter) {
	out := x_contact.GetCallBarringEntryResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_contact_GetCallBarringEntryByNum(w http.ResponseWriter) {
	out := x_contact.GetCallBarringEntryByNumResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_contact_GetCallBarringList(w http.ResponseWriter) {
	out := x_contact.GetCallBarringListResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_contact_SetCallBarringEntry(w http.ResponseWriter) {
	out := x_contact.SetCallBarringEntryResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_contact_DeleteCallBarringEntryUID(w http.ResponseWriter) {
	out := x_contact.DeleteCallBarringEntryUIDResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_contact_GetDECTHandsetList(w http.ResponseWriter) {
	out := x_contact.GetDECTHandsetListResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_contact_GetDECTHandsetInfo(w http.ResponseWriter) {
	out := x_contact.GetDECTHandsetInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_contact_SetDECTHandsetPhonebook(w http.ResponseWriter) {
	out := x_contact.SetDECTHandsetPhonebookResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_contact_GetNumberOfDeflections(w http.ResponseWriter) {
	out := x_contact.GetNumberOfDeflectionsResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_contact_GetDeflection(w http.ResponseWriter) {
	out := x_contact.GetDeflectionResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_contact_GetDeflections(w http.ResponseWriter) {
	out := x_contact.GetDeflectionsResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_contact_SetDeflectionEnable(w http.ResponseWriter) {
	out := x_contact.SetDeflectionEnableResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
