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
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
)

func TestServices(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata")
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server(), tr064.DefaultSpec)
	client.Debug = true
	services, err := client.Services()
	require.NoError(t, err)
	require.Equal(t, 2, len(services))
}

func TestServicesByName(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata")
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server(), tr064.DefaultSpec)
	client.Debug = true
	services, err := client.ServicesByName("MockPing")
	require.NoError(t, err)
	require.Equal(t, 1, len(services))
}
