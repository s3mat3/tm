/*
 file_exporter.go
 Copyright © 2026 s3mat3
 This code is licensed under the MIT License, see the LICENSE file for details
 Author s3mat3
*/

package port

import (
	"github.com/s3mat3/tm/tools/internal/port/file"
	"fmt"
	"os"
	"strings"
)

func is_exist_dir(d string) bool {
	t, err := os.Stat(d)
	if os.IsNotExist(err) || !t.IsDir() {
		return false
	}
	return true
}

type FileExporter struct {
	info file.Info
	base string
}
// NewFileExporter is constructor
//
// - b is base directory.
// - n is file name without path.
func NewFileExporter(b string, n string) *FileExporter {
	t := strings.TrimRight(b, "/") + "/"
	return &FileExporter{
		info: file.Info{Path: n},
		base: t,
	}
}

func (f *FileExporter) Export(s string) (int, error) {
	if !is_exist_dir(f.base) {
		os.MkdirAll(f.base, 0700)
	}
	p := f.base + f.info.Path
	wp, err := os.OpenFile(p, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644)
	if err != nil {
		return -1, fmt.Errorf("can not create file")
	}
	defer wp.Close()
	n, err := wp.WriteString(s)
	if err != nil {
		return -1, fmt.Errorf("can not write file")
	}
	return n, nil
}
//<-- file_exporter.go ends here.
