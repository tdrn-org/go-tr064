/*
 * Copyright 2024 Holger de Carne
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package tr064

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func Generate(url *url.URL, dir string) {
	tr64descUrl := url.JoinPath(rootSpec)
	log.Println("Reading '", tr64descUrl, "'...")
	tr64desc := &tr64descDoc{}
	err := unmarshalDocument(tr64descUrl, tr64desc)
	if err != nil {
		log.Fatal("Failed to unmarshal ", tr64descUrl, " cause: ", err)
	}
	log.Println("Generating...")
	gc := &generateContext{
		url:            url,
		serviceClients: make(map[string]*bytes.Buffer),
		serviceNames:   make(map[string]string),
	}
	err = gc.generate(tr64desc, dir)
	if err != nil {
		log.Fatal("Failed to generate:", err)
	}
}

type generateContext struct {
	url            *url.URL
	serviceClients map[string]*bytes.Buffer
	serviceNames   map[string]string
	err            error
}

func (gc *generateContext) generate(tr64desc *tr64descDoc, dir string) error {
	gc.err = tr64desc.walk(gc.url, gc.generateService)
	gc.flushFiles(dir)
	return gc.err
}

func (gc *generateContext) generateService(service *serviceDoc, scpd *scpdDoc) error {
	packageName := service.scpdName()
	gc.serviceNames[service.Name()] = packageName
	buffer := gc.serviceClients[packageName]
	if buffer == nil {
		buffer = &bytes.Buffer{}
		gc.serviceClients[packageName] = buffer
		gc.emit(buffer, "// generated from %s\n", scpd.SpecVersion.signature())
		gc.emit(buffer, "package %s\n", packageName)
		gc.emit(buffer, "import (\n")
		gc.emit(buffer, "\"encoding/xml\"\n")
		gc.emit(buffer, "\"github.com/tdrn-org/go-tr064\"\n")
		gc.emit(buffer, ")\n")
		gc.emit(buffer, "\ntype ServiceClient struct {\n")
		gc.emit(buffer, "TR064Client *tr064.Client\n")
		gc.emit(buffer, "Service tr064.ServiceDescriptor\n")
		gc.emit(buffer, "}\n")
		for _, action := range scpd.ActionList.Actions {
			hasRequest := gc.generateServiceActionArgument(buffer, scpd, &action, "in")
			hasResponse := gc.generateServiceActionArgument(buffer, scpd, &action, "out")
			gc.generateServiceAction(buffer, &action, hasRequest, hasResponse)
		}
	}
	return gc.err
}

func (gc *generateContext) generateServiceActionArgument(buffer *bytes.Buffer, scpd *scpdDoc, action *actionDoc, direction string) bool {
	if gc.err != nil {
		return false
	}
	actionName := mangleName(action.Name)
	var typeName string
	if direction == "in" {
		typeName = actionName + "Request"
	} else if direction == "out" {
		typeName = actionName + "Response"
	} else {
		log.Fatal("Unexpected direction '", direction, "'")
	}
	gc.emit(buffer, "\ntype %s struct {\n", typeName)
	if direction == "in" {
		gc.emit(buffer, "XMLName xml.Name `xml:\"u:%sRequest\"`\n", action.Name)
		gc.emit(buffer, "XMLNameSpace string `xml:\"xmlns:u,attr\"`\n")
	} else {
		gc.emit(buffer, "XMLName xml.Name `xml:\"%sResponse\"`\n", action.Name)
	}
	variableCount := 0
	for _, argument := range action.ArgumentList.Arguments {
		if argument.Direction == direction {
			variable := scpd.lookupVariable(argument.RelatedStateVariable)
			if variable == nil {
				gc.err = fmt.Errorf("unknown state variable '%s'", argument.Name)
				return false
			}
			variableName := mangleName(argument.Name)
			variableType := variableType(variable)
			gc.emit(buffer, "%s %s `xml:\"%s\"`\n", variableName, variableType, argument.Name)
			variableCount++
		}
	}
	gc.emit(buffer, "}\n")
	return variableCount > 0
}

func (gc *generateContext) generateServiceAction(buffer *bytes.Buffer, action *actionDoc, hasRequest bool, hasResponse bool) {
	if gc.err != nil {
		return
	}
	actionName := mangleName(action.Name)
	requestTypeName := actionName + "Request"
	responseTypeName := actionName + "Response"
	if hasRequest && hasResponse {
		gc.emit(buffer, "\nfunc (client *ServiceClient) %s(in *%s, out *%s) error {\n", actionName, requestTypeName, responseTypeName)
	} else if hasRequest {
		gc.emit(buffer, "\nfunc (client *ServiceClient) %s(in *%s) error {\n", actionName, requestTypeName)
	} else if hasResponse {
		gc.emit(buffer, "\nfunc (client *ServiceClient) %s(out *%s) error {\n", actionName, responseTypeName)
	} else {
		gc.emit(buffer, "\nfunc (client *ServiceClient) %s() error {\n", actionName)
	}
	if hasRequest {
		gc.emit(buffer, "in.XMLNameSpace = client.Service.Type()\n")
	} else {
		gc.emit(buffer, "in := &%s{XMLNameSpace: client.Service.Type() }\n", requestTypeName)
	}
	gc.emit(buffer, "soapRequest := &tr064.SOAPRequest[%s]{\n", requestTypeName)
	gc.emit(buffer, "XMLNameSpace: tr064.XMLNameSpace,\n")
	gc.emit(buffer, "XMLEncodingStyle: tr064.XMLEncodingStyle,\n")
	gc.emit(buffer, "Body: &tr064.SOAPRequestBody[%s]{\n", requestTypeName)
	gc.emit(buffer, "In: in,\n")
	gc.emit(buffer, "},\n")
	gc.emit(buffer, "}\n")
	if !hasResponse {
		gc.emit(buffer, "out := &%s{}\n", responseTypeName)
	}
	gc.emit(buffer, "soapResponse := &tr064.SOAPResponse[%s]{\n", responseTypeName)
	gc.emit(buffer, "Body: &tr064.SOAPResponseBody[%s]{\n", responseTypeName)
	gc.emit(buffer, "Out: out,\n")
	gc.emit(buffer, "},\n")
	gc.emit(buffer, "}\n")
	gc.emit(buffer, "return client.TR064Client.InvokeService(client.Service, \"%s\", soapRequest, soapResponse)\n", action.Name)
	gc.emit(buffer, "}\n")
}

func (gc *generateContext) emit(buffer *bytes.Buffer, format string, a ...any) {
	if gc.err != nil {
		return
	}
	_, gc.err = buffer.WriteString(fmt.Sprintf(format, a...))
}

func (gc *generateContext) flushFiles(dir string) {
	if gc.err != nil {
		return
	}
	for packageName, buffer := range gc.serviceClients {
		log.Println("Writing service client '", packageName, "'...")
		packageDir := filepath.Join(dir, "services", packageName)
		code, err := format.Source(buffer.Bytes())
		if err != nil {
			gc.err = fmt.Errorf("failed to format generated code (cause: %w)", err)
			return
		}
		err = os.MkdirAll(packageDir, 0777)
		if err != nil {
			gc.err = fmt.Errorf("failed to create directory '%s' (cause: %w)", packageName, err)
			return
		}
		file := filepath.Join(packageDir, packageName+".gen.go")
		err = os.WriteFile(file, code, 0666)
		if err != nil {
			gc.err = fmt.Errorf("failed to write file '%s' (cause: %w)", file, err)
			return
		}
	}
	for serviceName, packageName := range gc.serviceNames {
		log.Println("Writing service name '", packageName, "'/'", serviceName, "'...")
		buffer := &bytes.Buffer{}
		gc.emit(buffer, "// %s\n", serviceName)
		gc.emit(buffer, "package %s\n", packageName)
		gc.emit(buffer, "const ServiceName = \"%s\"\n", serviceName)
		code, err := format.Source(buffer.Bytes())
		if err != nil {
			gc.err = fmt.Errorf("failed to format generated code (cause: %w)", err)
			return
		}
		file := filepath.Join(dir, "services", packageName, "name.gen.go")
		err = os.WriteFile(file, code, 0666)
		if err != nil {
			gc.err = fmt.Errorf("failed to write file '%s' (cause: %w)", file, err)
			return
		}
	}
}

func mangleName(name string) string {
	mangled := strings.ReplaceAll(name, "-", "_")
	mangled = strings.ReplaceAll(mangled, ".", "_")
	return mangled
}

var dataTypeMap map[string]string = map[string]string{
	"i1":       "int8",
	"i2":       "int16",
	"i4":       "int32",
	"ui1":      "uint8",
	"ui2":      "uint16",
	"ui4":      "uint32",
	"boolean":  "bool",
	"string":   "string",
	"dateTime": "string",
	"uuid":     "string",
}

func variableType(variable *stateVariableDoc) string {
	mappedType := dataTypeMap[variable.DataType]
	if mappedType != "" {
		return mappedType
	}
	log.Println("Unknown variable data type '", variable.DataType, "' resolved to any")
	return "any"
}
