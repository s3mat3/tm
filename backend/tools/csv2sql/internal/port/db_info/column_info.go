/*
 column_info_model.go
 Copyright © 2026 s3mat3
 This code is licensed under the MIT License, see the LICENSE file for details
 Author s3mat3
*/
// Package xxx package brief here.
package db_info

type ColumnInfo struct {
	Name string			// Column name
	Disp string			// Column display name (optional)
	PK bool				// Primary key (optional)
	FK string			// Foreign key, style tablename.columnname (optional)
	Type string			// Column data type with size
	Constraint string	// Column constraint (optional)
	Comment string		// Column comment (optional)
}

// colmun
func NewColumnInfo(
	n string,
	d string,
	p bool,
	f string,
	t string,
	c string,
	r string) *ColumnInfo {
	return &ColumnInfo {
		Name: n,
			Disp: d,
			PK: p,
			FK: f,
			Type: t,
			Constraint: c,
			Comment: r,
	}
} //<-- NewColumnInfo ends here.
//<-- column_info_model.go ends here.
