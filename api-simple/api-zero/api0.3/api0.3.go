// curl -i -XPOST -H "Content-type:application/json" \
// -d '{"id":3,"name":"carlos lima o foda","location":"Brasil","age":24}' \
// localhost:3000/api/v1/user
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type User struct {
	ID       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Location string `json:"location,omitempty"`
	Age      int    `json:"age,omitempty"`
}

type msgError struct {
	Msg string `json:"msg"`
}

var msgE msgError

func main() {

	http.HandleFunc("/api/v1/user",
		func(w http.ResponseWriter, r *http.Request) {
			var p User
			err := json.NewDecoder(r.Body).Decode(&p)
			if err != nil {
				msgE.Msg = "Error body!"
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(&msgE)
				return
			}

			if p.Age < 0 {
				msgE.Msg = "Campo Age obrigatorio!"
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(&msgE)
				return
			}
			//fmt.Println(p)
			//p.Ip = c.RemoteAddr
			time.Sleep(time.Millisecond * 50)
			msgE.Msg = "success"
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(&msgE)
		})

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
