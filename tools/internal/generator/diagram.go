/*
 diagram.go
 Copyright © 2026 s3mat3
 This code is licensed under the MIT License, see the LICENSE file for details
 Author s3mat3
*/

package generator

import (
	"github.com/s3mat3/tm/tools/internal/port/db_info"
	"strings"
)

type Diagram struct {}

func NewDiagram() (*Diagram) {
	return &Diagram {}
}

func gen_entity_member(col *db_info.ColumnInfo) (string) {
		return col.Name + " : " + col.Type
}

func (d *Diagram) Generate(db *db_info.DBInfo) (string, error) {
	var rl = map[string][]string{}
	var uml = "```plantuml\n@startuml " + db.Name + "\n"
	// uml += "skinparam entityAttributeIconSize 0\n"
	for _, v := range db.Tables {
		uml += "entity " + "\"" + v.Name + "\"" + "{\n"
		for _, c := range v.Columns {
			u := gen_entity_member(c)
			if c.PK {
				u = "+ " + u + " [PK]"
			}
			if c.FK != "" {
				a := strings.Split(c.FK, ".")
				rl[v.Name] = append(rl[v.Name], a[0])
				u = "# " + u + " [FK] " + c.FK
			}
			uml += "\t" + u + "\n"
		}
		uml += "}\n"
	}
	for k, v := range rl {
		if len(v) > 0 {
			u := ""
			for _, fk := range v {
				u += k + " --{ " + fk + "\n"
			}
			uml += u
		}
	}
	uml += "@enduml\n```\n"
	return uml, nil
}
//<-- diagram.go ends here.
