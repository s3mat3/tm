/*
 db_info.go
 Copyright © 2026 s3mat3
 This code is licensed under the MIT License, see the LICENSE file for details
 Author s3mat3
*/
// Package db_info DTO for tble generator
package db_info

type DBInfo struct {
	Name string			// Database name
	Comment string		// Database comment (optional)
	Tables []*TableInfo	// Database table definations
}

func NewDBInfo(n string, c string) * DBInfo {
	return &DBInfo{
		Name: n,
		Comment: c,
	}
}

func (d *DBInfo) AppendTableInfo(t *TableInfo) {
	d.Tables = append(d.Tables, t)
}
//<-- db_info.go ends here.
