/*
 table_info.go
 Copyright © 2026 s3mat3
 This code is licensed under the MIT License, see the LICENSE file for details
 Author s3mat3
*/
// Package xxx package brief here.
package db_info

type TableInfo struct {
	Name string				// Table name
	Comment string			// Table comment (optional)
	Columns []*ColumnInfo	// Table column definations
}

func NewTableInfo(n string, c string) *TableInfo {
	return &TableInfo{
		Name: n,
		Comment: c,
	}
}

func (t *TableInfo) AppendColumnInfo(c *ColumnInfo) {
	t.Columns = append(t.Columns, c)
}
//<-- table_info.go ends here.
