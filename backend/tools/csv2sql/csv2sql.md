## csv2sql

This is a tool that converts table structures created in spreadsheet applications such as Excel and output in
 CSV format into SQL that can actually generate tables in PostgreSQL.

### build

go build -ldflags="-s -w" -trimpath -o some-path/csv2sql backend/tools/csv2sql/cmd/main.go

### usage

some-path/csv2sql -in tools/csv2sql/sample/table.csv -out tools/csv2sql/sample/out/

### Sample output

[genarated SQL](./sample/out/testdb.sql)

[genarated MD](./sample/out/testdb.md)
