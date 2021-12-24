package main

import (
    "database/sql"
    "fmt"

    _ "github.com/godror/godror"
)

// lib go oracle
// https://github.com/sijms/go-ora

// example oracle
// https://golangrepo.com/repo/mattn-go-oci8-go-database-drivers

// oracle com ODBC
// https://stackoverflow.com/questions/30043488/connecting-to-oracle-db-in-go

// Lib Oracle
// https://pkg.go.dev/gopkg.in/goracle.v2

// https://pkg.go.dev/github.com/godror/godror#section-documentation
func main() {
    // db, err := sql.Open("godror", `user="scott" password="tiger" connectString="dbhost:1521/orclpdb1"`)
    db, err := sql.Open("godror", "<your username>/<your password>@service_name")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    rows, err := db.Query("select sysdate from dual")
    if err != nil {
        fmt.Println("Error running query")
        fmt.Println(err)
        return
    }
    defer rows.Close()

    var thedate string
    for rows.Next() {

        rows.Scan(&thedate)
    }
    fmt.Printf("The date is: %s\n", thedate)
}
