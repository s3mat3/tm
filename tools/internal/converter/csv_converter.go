/*
 csv.go
 Copyright © 2026 s3mat3
 This code is licensed under the MIT License, see the LICENSE file for details
 Author s3mat3
*/

package converter

import (
	"github.com/s3mat3/tm/tools/internal/port/db_info"
	"github.com/s3mat3/tm/tools/internal/reader"
	"errors"
	"fmt"
	"strings"
)

type CSVConverter struct {
	reader *reader.LineReader
}

func NewCSVConverter(r *reader.LineReader) (*CSVConverter) {
	return &CSVConverter {
		reader: r,
	}
}

// is_skip_line detects empty, "//", and "--" first fields
func is_skip_line(l []string) bool {
	if l[0] == "" || l[0][0:2] == "//" || l[0][0:2] == "--" {
		return true
	}
	return false
}

func change_t_to_true(c string) bool {
	if c == "T" || c == "t" {
		return true
	}
	return false
}

// split_commas Split a string by commas.
func split_commas(l string) []string {
	var r []string
	c := strings.Split(l, ",")
	for _, f := range c {
		x := strings.Trim(f, " \"'")
		r = append(r, x)
	}
	return r
}
// Convert is CSV formated string to internal struct.
//
// # CSV format
//
// The expected file structure is comma-separated CSV, and columns can be double-, or unquoted.
// Enter the database name (database:NAME) in the first column of line 1.
// Enter the table name (table:NAME) in the first column of line X (excluding lines 1 and 2).
// Use the // notation at the beginning of line X's first column to ignore that column as a comment line.
// Enter the table structure below the table name.
// Enter column information in the table column column for the table structure.
//
// ## Column meaning
//
// - Column 1: Column name (alphanumeric, for SQL) (Required)
// - Column 2: Column display name (Optional)
// - Column 3: Primary key flag: Enter t or T to specify the column as a primary_key (Optional)
// - Column 4: Foreign reference key: Enter the referenced table name and column name (Optional)
// - Column 5: Data type: Enter the data type of this column (Required)
// - Column 6: Enter the SQL constraints for this column (Optional)
// - Column 7: Column comment  (Optional)
func (c *CSVConverter) Convert() (*db_info.DBInfo, error) {
	var db *db_info.DBInfo
	// search database name
	for {
		l, err := c.reader.Read()
		if err != nil {
			return nil, fmt.Errorf("no database name cause: %s", err)
		}
		ln := split_commas(l)
		if is_skip_line(ln) {
			continue
		}
		arr := strings.Split(ln[0], ":")
		if len(arr) > 1 && arr[0] == "database" {
			var dc = ""
			if len(ln) > 1 && ln[1] != "" {
				dc = ln[1]
			}
			db = db_info.NewDBInfo(arr[1], dc)
			break
		}
	}
	var tb *db_info.TableInfo
	// create db_info.Table struct
	for {
		l, err := c.reader.Read()
		if errors.Is(err, reader.NML) { // maybe NML
			break
		}
		ln := split_commas(l)
		if is_skip_line(ln) {
			continue
		}
		arr := strings.Split(ln[0], ":")
		if len(arr) == 2 && arr[0] == "table" {
			var tc = ""
			if len(ln) > 1 && ln[1] != "" {
				tc = ln[1]
			}
			tb = db_info.NewTableInfo(arr[1], tc)
			continue
		} else if arr[0] == "end-table" { // check end of table definetion
			// Updte tables
			db.AppendTableInfo(tb)
			continue
		} else { //
			c := db_info.NewColumnInfo(
				ln[0],
				ln[1],
				change_t_to_true(ln[2]),
				ln[3],
				ln[4],
				ln[5],
				ln[6])
			tb.AppendColumnInfo(c)
		}
	}
	return db, nil
}

//<-- csv.go ends here.
