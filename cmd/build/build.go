//go:build tools
// +build tools

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

package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/tdrn-org/go-tr064"
)

// Used via go:generate to perform build tasks.
func main() {
	switch os.Args[1] {
	case "generate":
		generate()
	}
}

var specNamePattern = regexp.MustCompile(`^/(.+)\.xml$`)

// Generate from generate.conf (if available)
func generate() {
	confFile, err := os.Open("./generate.conf")
	if err != nil {
		return
	}
	conf := bufio.NewReader(confFile)
	for {
		line, err := conf.ReadString('\n')
		if errors.Is(err, io.EOF) {
			return
		} else if err != nil {
			log.Fatal(err)
		}
		if strings.HasPrefix(line, "#") {
			continue
		}
		line = strings.Trim(line, " \r\n")
		parsedUrl, err := url.Parse(line)
		if err != nil {
			log.Fatal(err)
		}
		baseUrl := *parsedUrl
		baseUrl.User = nil
		baseUrl.Path = "/"
		match := specNamePattern.FindStringSubmatch(parsedUrl.Path)
		if match == nil {
			log.Fatal("Unexpected URL: '", parsedUrl.Redacted(), "'")
		}
		spec := tr064.ServiceSpec(match[1])
		tr064.Generate(&baseUrl, spec, ".")
	}
}
