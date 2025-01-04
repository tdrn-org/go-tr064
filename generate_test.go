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
	"io"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	tr064 "github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
)

func TestGenerate(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata")
	defer tr064Mock.Shutdown()
	// Prepare temp dir (for generate output)
	dir := &generateDir{}
	dir.Mkdir()
	defer dir.Remove()
	// Actual test
	tr064.Generate(tr064Mock.Server(), tr064.DefaultSpec, dir.TempDir)
	require.NoError(t, dir.FileNotEmpty("services/tr64desc/mockecho/mockecho.go"))
	require.NoError(t, dir.FileNotEmpty("services/tr64desc/mockecho/name.go"))
	require.NoError(t, dir.FileNotEmpty("services/tr64desc/mockecho_test.go"))
	require.NoError(t, dir.FileNotEmpty("services/tr64desc/mockping/mockping.go"))
	require.NoError(t, dir.FileNotEmpty("services/tr64desc/mockping/name.go"))
	require.NoError(t, dir.FileNotEmpty("services/tr64desc/mockping_test.go"))
}

type generateDir struct {
	TempDir string
}

func (dir *generateDir) Mkdir() {
	tempDir, err := os.MkdirTemp("", "generate")
	if err != nil {
		log.Fatal(err)
	}
	dir.TempDir = tempDir
}

func (dir *generateDir) Remove() {
	os.RemoveAll(dir.TempDir)
}

func (dir *generateDir) FileNotEmpty(file string) error {
	fileInfo, err := os.Stat(filepath.Join(dir.TempDir, file))
	if err != nil {
		return err
	}
	if fileInfo.Size() == 0 {
		return io.EOF
	}
	return nil
}
