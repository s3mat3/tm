/*
 %fiename%
 Copyright © 2025 s3mat3
 This code is licensed under the MIT License, see the LICENSE file for details
 Author s3mat3

 Brief:
   Convert csv to SQL

 Description:
   A tool to convert database table definitions created in Excel (exported as CSV) into SQL.
*/
// Package main entry point for csv2sql.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/s3mat3/tm/tools/internal/converter"
	"github.com/s3mat3/tm/tools/internal/generator"
	"github.com/s3mat3/tm/tools/internal/port"
	"github.com/s3mat3/tm/tools/internal/port/args"
	_ "github.com/s3mat3/tm/tools/internal/port/db_info"
	"github.com/s3mat3/tm/tools/internal/reader"
)

func IsExistsInput(n string) bool {
	_, err := os.Stat(n)
	return !os.IsNotExist(err)
}


func main() {
	i := flag.String("in", "./table.csv", "Input file name")
	o := flag.String("out", "./out", "Output directory name")
	m := flag.String("mode", "all", "Output mode [all | sql | doc]")
	w := flag.Bool("drop", true, "SQL with first drop table")
	flag.Parse()
	op := args.NewOptions(i, o, m, w)
	// fmt.Printf("Start convert \n\tinput: %s \n\toutput: %s \n\tfor: %s\n", op.In, op.Out, op.Mode)
	var ip port.Importer = port.NewFileImporter(op.In)
	t, err := ip.Import()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lr := reader.NewLineReader(t)
	c := converter.NewCSVConverter(lr)
	db, err := c.Convert()
	if err != nil {
		panic(err.Error())
	}
	// //
	if op.Mode == "all" || op.Mode == "sql" {
		fname := db.Name + ".sql"
		sql := generator.NewSQL()
		s, _ := sql.Generate(db)
		exp := port.NewFileExporter(op.Out, fname)
		sn, _ := exp.Export(s)
		fmt.Printf("Write SQL from %s to %s in %d chars\n" ,op.In ,op.Out + "/" + fname, sn)
	}

	if op.Mode == "all" || op.Mode == "doc" {
		fname := db.Name + ".md"
		doc := generator.NewDOC()
		d, _ := doc.Generate(db)
		exp := port.NewFileExporter(op.Out, fname)
		dn, _ := exp.Export(d)
		fmt.Printf("Write md from %s to %s in %d chars\n" ,op.In ,op.Out + "/" + fname, dn)
	}

}

// defer func() {
// 	fmt.Println("Enter defer")
// 	fp.Close()
// 	if r := recover(); r != nil {
// 		fmt.Println(r)
// 		os.Exit(2)
// 	}
// 	fmt.Println("r is nil")
// }()

//<-- csv2sql.go ends here.
