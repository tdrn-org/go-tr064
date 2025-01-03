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

package mock_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064/mock"
)

func TestMockStartStop(t *testing.T) {
	tr064Mock := mock.Start("testdata")
	defer tr064Mock.Shutdown()
	require.NotNil(t, tr064Mock)
	err := tr064Mock.Ping()
	require.NoError(t, err)
	err = tr064Mock.SecurePing()
	require.NoError(t, err)
}

func TestServiceMockFromFile(t *testing.T) {
	tr064Mock := mock.Start("testdata", mock.ServiceMockFromFile("/mocknotfound", "testdata/notfound.xml"), mock.ServiceMockFromFile("/mockping", "testdata/ping.xml"))
	defer tr064Mock.Shutdown()
	response, err := http.Post(tr064Mock.Server().JoinPath("/mocknotfound").String(), "text/xml", nil)
	require.NoError(t, err)
	require.Equal(t, http.StatusNotFound, response.StatusCode)
	response, err = http.Post(tr064Mock.Server().JoinPath("/mockping").String(), "text/xml", nil)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode)
}
