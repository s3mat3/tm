/*
 file_importer.go
 Copyright © 2026 s3mat3
 This code is licensed under the MIT License, see the LICENSE file for details
 Author s3mat3
*/

package port

import (
	"fmt"
	"os"
	"github.com/s3mat3/tm/tools/internal/port/file"
)


func is_exist_file(p string) bool {
	_, err := os.Stat(p)
	return !os.IsNotExist(err)
}

type FileImporter struct {
	file.Info
}

func NewFileImporter(p string) *FileImporter {
	return &FileImporter{
		file.Info{Path: p},
	}
}

func (f *FileImporter) Import() (string, error) {
	// check exists
	if !is_exist_file(f.Path) {
		return "", fmt.Errorf("can not open file: %s", f.Path)
	}
	b, err := os.ReadFile(f.Path)
	if err != nil {
		return "", fmt.Errorf("fail to read file: %s cause: %s", f.Path, err)
	}
	return string(b), nil
}
//<-- file_importer.go ends here.
