package main

import (
    "log"
    "net/http"
)

func main() {

    http.Handle("/", http.HandlerFunc(handler))
    http.ListenAndServe(":8090", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

    keys, ok := r.URL.Query()["key"]

    if !ok || len(keys[0]) < 1 {
        log.Println("Url Param 'key' is missing")
        return
    }

    // Query()["key"] will return an array of items,
    // we only want the single item.
    key := keys[0]

    log.Println("Url Param 'key' is: " + string(key))
}
