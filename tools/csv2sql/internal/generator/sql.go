/*
 sql.go
 Copyright © 2026 s3mat3
 This code is licensed under the MIT License, see the LICENSE file for details
 Author s3mat3
*/

package generator

import (
	"strings"

	"github.com/s3mat3/tm/tools/csv2sql/internal/port/db_info"
)

type SQL struct {}

func NewSQL() (*SQL){
	return &SQL{}
}

func (s *SQL) Generate(db *db_info.DBInfo) (string, error) {
	var c = "--------------------\n-- Tables for " + db.Name + "\n--------------------\n\n"
	for _, v := range db.Tables {
		var sql = "--------------------\n-- Table: " + v.Name + "\n"
		if v.Comment != "" {
			sql += "-- Description: " + v.Comment + "\n"
		}
		sql += "--------------------\n"
		sql += "DROP TABLE IF EXISTS " + v.Name + " CASCADE;\n"
		sql += "CREATE TABLE IF NOT EXISTS " + v.Name + "(\n"
		var pk []string
		var fk = map[string]string{}
		var cc = map[string]string{}
		// var fi []string
		for _, c := range v.Columns {
			if c.PK {
				pk = append(pk, c.Name)
			}
			if c.FK != "" {
				fk[c.Name] = c.FK
			}
			sql += "\t" + c.Name + " " + c.Type
			if c.Constraint != "" {
				sql += " CONSTRAINT " + c.Constraint
			}
			if c.Comment != "" {
				sql += ", -- " + c.Comment
				cc[c.Name] = c.Comment
			}
			sql += "\n"
		}
		// create table constraint
		if len(pk) > 0 { // primary key constraint
			var ps = ""
			for _, p := range pk {
				ps += p + ","
			}
			sql += "\tPRIMARY KEY (" + strings.Trim(ps, ",") + "),\n"
		}
		if len(fk) > 0 { // foreign key constraint
			for k, v := range fk {
				fs := strings.Split(v, ".")
				sql += "\tFOREIGN KEY (" + k + ") REFERENCES " + fs[0] + " (" + fs[1] + ") ON UPDATE CASCADE ON DELETE CASCADE,\n"
				delete(fk, k)
			}
		}
		sql = strings.TrimRight(sql, "\n")
		sql = strings.TrimRight(sql, ",")
		sql += "\n);\n"
		// create table comment
		if v.Comment != "" {
			sql += "COMMENT ON TABLE " + v.Name + " IS " + "'" + v.Comment + "';\n"
		}
		//crete column comment
		if len(cc) > 0 {
			for k, val := range cc {
				sql += "COMMENT ON COLUMN " + v.Name + "." + k + " IS " + "'" + val + "';\n"
			}
		}
		c += sql
	}
	return c, nil
}
//<-- sql.go ends here.
