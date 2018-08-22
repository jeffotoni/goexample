package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func Hello(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {

			fmt.Println("error: ", err)

		}

		// convert para string, forca remocao do \n
		// e mantem o byte para o body
		body = []byte(strings.TrimRight(string(body), "\r\n"))

		//src := body

		// buffer
		dst := new(bytes.Buffer)

		// quando nao conheco a struct
		var data map[string]interface{}

		// quando nao conheco a struct
		err = json.Unmarshal(body, &data)
		if err != nil {

			fmt.Println("Unmarshal", err)
		}

		fmt.Println("Unmarshal", data)

		err = json.Compact(dst, body)

		if err != nil {

			fmt.Println("json.compact: ", err)
		}

		fmt.Println("Dst body: ", dst)

		//criando json newEncoder
		enc := json.NewEncoder(os.Stdout)

		// de encode Data
		err = enc.Encode(data)

		fmt.Println("Encode: ", data)

		//err = json.NewDecoder(r.Body).Decode(data)

		if err != nil {

			fmt.Println("Error json.newDecoder: ", err)
		}

		// unmarshal pode usar com map -> quando nao
		// sabemos qual struct pode vim
		err = json.Unmarshal([]byte(strings.TrimRight(dst.String(), "\n")), &data)

		fmt.Println("Unmarshal2: ", data)
		fmt.Println("vetor map: ", data["password"])

		read_line := strings.TrimRightFunc(fmt.Sprintf("%s", data["password"]), func(c rune) bool {
			//In windows newline is \r\n
			return c == '\r' || c == '\n'
		})
		// json.compact
		fmt.Println("Data map read_line: ", read_line)

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
