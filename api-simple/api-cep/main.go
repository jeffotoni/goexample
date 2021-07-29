package main

import (
    "io"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/{cep}", Cep).Methods("GET").Schemes("http")
    r.HandleFunc("/", Notfound)
    log.Println("Run Server")
    http.ListenAndServe("0.0.0.0:8090", r)
}

func Notfound(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)
    log.Println("notfound:", r.URL)
    return
}

func Cep(w http.ResponseWriter, r *http.Request) {
    log.Println("call", r.UserAgent(), " - ", r.URL)
    w.Header().Set("Content-Type", "application/json")
    mcep := mux.Vars(r)
    cep, ok := mcep["cep"]
    if !ok {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("cep is required"))
        return
    }
    if len(cep) > 0 {
        client := &http.Client{}
        urlStr := "https://viacep.com.br/ws/" + cep + "/json"
        r, _ := http.NewRequest("GET", urlStr, nil)
        r.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
        r.Header.Add("Content-Type", "application/json")
        resp, err := client.Do(r)
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte("cep is invalid:" + err.Error()))
            return
        }

        if resp.StatusCode != 200 {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte("StatusCode is invalid:" + resp.Status))
            return
        }

        defer resp.Body.Close()

        body, err := io.ReadAll(resp.Body)
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte("ReadAll is invalid:" + err.Error()))
            return
        }
        w.WriteHeader(http.StatusOK)
        w.Write(body)
        return
    }

    w.WriteHeader(http.StatusBadRequest)
    return
}
