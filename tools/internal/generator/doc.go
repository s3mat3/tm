/*
 doc.go
 Copyright © 2026 s3mat3
 This code is licensed under the MIT License, see the LICENSE file for details
 Author s3mat3
*/

package generator

import (
	"github.com/s3mat3/tm/tools/internal/port/db_info"
)

type DOC struct {}

func NewDOC() (*DOC) {
	return &DOC{}
}


// Generate markdown for db tables description
func (d *DOC) Generate(db *db_info.DBInfo) (string, error){
	var c = "## Table description for database " + db.Name + "\n\n"
	for _, v := range db.Tables {
		var ts = "### Table " + v.Name + "\n\n"
		ts += "table description: " + v.Comment + "\n\n"
		ts += "| column | PK | FK | type | display | comment |\n"
		ts += "|--------|:--:|:--:|------|---------|---------|\n"
		cl := ""
		for _, c := range v.Columns {
			b := ""
			if c.PK {
				b = "\u2713"
			}
			cl += "| " + c.Name + " | " + b + " | " + c.FK + " | " + c.Type + " | " + " | " + c.Disp + " | " + c.Comment + " |\n"
		}
		ts += cl + "\n"
		c += ts
	}
	pu := NewDiagram()
	dia,_ := pu.Generate(db)
	c += "### Table relation\n\n"
	c += dia
	return c, nil
}

//<-- doc.go ends here.
