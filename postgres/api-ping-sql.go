package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

// Back-End in Go server
// @jeffotoni
// 2019-01-04

func firstHandler(w http.ResponseWriter, r *http.Request) {
    err := db.Ping()
    if err != nil {
        log.Fatal(err)
    }
    rows, err := db.Query("SELECT id, created_at, updated_at FROM script WHERE updated_at = $1", 3)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var created_at, updated_at, id int
    for rows.Next() {
        err := rows.Scan(&id, &created_at, &updated_at)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Fprintf("%s %s %s", id, created_at, updated_at)
    }
}

var r = mux.NewRouter()
var db *sql.DB

func main() {
    db, err := sql.Open("postgres", "user=youruser host=localhost dbname=yourdb sslmode=verify-full")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    r.HandleFunc("/ping", firstHandler)

    http.Handle("/", r)

    http.ListenAndServe(":8080", nil)
}
