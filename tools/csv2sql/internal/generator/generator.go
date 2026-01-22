/*
 generator.go
 Copyright © 2026 s3mat3
 This code is licensed under the MIT License, see the LICENSE file for details
 Author s3mat3
*/
// Package generator Generates various formatted strings from internal DBInfo format
package generator

import(
	"github.com/s3mat3/tm/tools/csv2sql/internal/port/db_info"
)

type Generator interface {
	Generate(*db_info.DBInfo) (string, error)
}
//<-- generator.go ends here.
