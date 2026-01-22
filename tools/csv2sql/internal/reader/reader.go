/*
 reader.go
 Copyright © 2026 s3mat3
 This code is licensed under the MIT License, see the LICENSE file for details
 Author s3mat3
*/
// Package reader
package reader
import "errors"
type Reader interface {
	Read() (string, error)
	
}

var NML  = errors.New("NML")

//<-- reader.go ends here.
