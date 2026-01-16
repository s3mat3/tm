/*
 converter.go
 Copyright © 2026 s3mat3
 This code is licensed under the MIT License, see the LICENSE file for details
 Author s3mat3
*/
// Package converter convert from file to internal model.
package converter

import (
	"backend/tools/csv2sql/internal/port/db_info"
)

type Converter interface {
	Convert() (*db_info.DBInfo, error)
}

//<-- converter.go ends here.
