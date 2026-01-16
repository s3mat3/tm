/*
 importer.go
 Copyright © 2026 s3mat3
 This code is licensed under the MIT License, see the LICENSE file for details
 Author s3mat3
*/
// Package importer Import avrious data from outside
package port


type Importer interface {
	Import() (string, error)
}
//<-- importer.go ends here.
