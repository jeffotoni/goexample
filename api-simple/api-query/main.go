package main

import (
    "log"

    "fmt"
    "net/http"
    "strconv"
    "strings"
)

//http://localhost:8080/products?filters=1&filters=2&filters=3&filters=4
func main() {
    getProductsHandler := http.HandlerFunc(getProducts)
    http.Handle("/products", getProductsHandler)
    http.ListenAndServe(":8080", nil)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()
    filters, present := query["filters"] //filters=["1", "2", "3", "4"]
    if !present || len(filters) == 0 {
        fmt.Println("filters not present")
        return
    }

    var sint []int
    for _, v := range filters {
        intv, err := strconv.Atoi(v)
        if err != nil {
            log.Println(err)
            continue
        }
        sint = append(sint, intv)
    }

    // int
    fmt.Printf("%t", sint)

    w.WriteHeader(200)
    w.Write([]byte(strings.Join(filters, ",")))
}
