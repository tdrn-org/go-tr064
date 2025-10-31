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

package tr064_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
)

func TestServices(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata")
	defer tr064Mock.Stop(t.Context())
	// Actual test
	client := tr064.NewClient(tr064Mock.Server())
	client.Debug = true
	services, err := client.Services(tr064.DefaultServiceSpec)
	require.NoError(t, err)
	require.Len(t, services, 2)
}

func TestServicesByType(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata")
	defer tr064Mock.Stop(t.Context())
	// Actual test
	client := tr064.NewClient(tr064Mock.Server())
	client.Debug = true
	services, err := client.ServicesByType(tr064.DefaultServiceSpec, "MockPing")
	require.NoError(t, err)
	require.Len(t, services, 1)
}

func TestGet(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata")
	defer tr064Mock.Stop(t.Context())
	// Actual test
	client := tr064.NewClient(tr064Mock.Server())
	client.Debug = true
	response, err := client.Get(tr064.DefaultServiceSpec.Path() + "?query=1234")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode)
}
