/*
 info.go
 Copyright © 2026 s3mat3
 This code is licensed under the MIT License, see the LICENSE file for details
 Author s3mat3
*/
// Package file package brief here.
package file

type Info struct {
	Path string
}

func NewInfo(p string) *Info {
	return &Info {
		Path: p,
	}
}

//<-- info.go ends here.
