package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Hello(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {

			fmt.Println("error: ", err)

		}

		body = []byte(strings.Trim(string(body), "\n"))

		//src := body

		dst := new(bytes.Buffer)

		// quando nao conheco a struct
		var data map[string]interface{}

		// quando nao conheco a struct
		err = json.Unmarshal(body, &data)
		if err != nil {

			fmt.Println("Unmarshal", err)
		}

		fmt.Println(data)

		err = json.Compact(dst, body)

		if err != nil {

			fmt.Println("json.compact: ", err)
		}

		err = json.NewDecoder(r.Body).Decode(data)

		if err != nil {

			fmt.Println("Error json.newDecoder: ", err)
		}

		err = json.Unmarshal([]byte(dst.String()), &data)

		fmt.Println(data)

		fmt.Println(dst)

		//r.ParseForm()
		w.Header().Set("Server", "A Go Web Server")
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(200)

		fmt.Fprintf(w, "%q", dst)

	} else if r.Method == "GET" {

		http.Error(w, "POST only", http.StatusMethodNotAllowed)
	}
}

func main() {

	//src := []byte(`{
	//      "Name":"Adam Ng",
	//      "Age":36,
	//      "Job":"CEO"
	///     }`)

	http.HandleFunc("/hello", Hello)
	http.ListenAndServe(":8189", nil)

}
