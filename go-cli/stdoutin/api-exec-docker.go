// Go in action
// @jeffotoni
// 2019-01-16

package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type Tpdf struct {
	File string `json:"file"`
}

type TpdfOut struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Output  string `json:"output"`
}

// exec handler
func handlerExec(w http.ResponseWriter, r *http.Request) {

	if strings.ToUpper(r.Method) == http.MethodGet {

		content_type := strings.ToLower(r.Header.Get("Content-Type"))
		if content_type != "application/json" {
			w.WriteHeader(http.StatusOK)
			jsonstr := `{"status":"error","message":"Error Content-Type is is required"}`
			io.WriteString(w, jsonstr)
			return
		}

		//set header..
		w.Header().Add("Content-Type", "application/json")

		//docker run --rm -i jeffotoni/pdftotext < /tmp/multa.pdf
		//////////?

		// docker run --rm -i jeffotoni/pdftotext < /tmp/multa.pdf
		// docker run --rm -i jeffotoni/pdftotext < $PATH_PDF
		// depois executa
		//
		//

		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {

			w.WriteHeader(http.StatusOK)
			jsonstr := `{"status":"error","message":"Error ReadAll is is required"}`
			io.WriteString(w, jsonstr)
			return
		}

		ljson := Tpdf{}
		err = json.Unmarshal(body, &ljson)
		if err != nil {
			w.WriteHeader(http.StatusOK)
			jsonstr := `{"status":"error","message":"Error Unmarshal is is required"}`
			io.WriteString(w, jsonstr)
			return
		}

		if len(ljson.File) <= 0 {
			w.WriteHeader(http.StatusOK)
			jsonstr := `{"status":"error","message":"Error name file pdf and path"}`
			io.WriteString(w, jsonstr)
			return
		}

		pdf := ljson.File
		cmd := exec.Command("sh", "docker-pdf-totext.sh")
		newEnv := append(os.Environ(), "PATH_PDF="+pdf)
		cmd.Env = newEnv
		out, err := cmd.CombinedOutput()
		if err != nil {
			w.WriteHeader(http.StatusOK)
			jsonstr := `{"status":"error","message":"Error Command in pdf and path failed!"}`
			io.WriteString(w, jsonstr)
			return
		}
		//out64 := Base64Encb(out)
		out64 := string(out)

		tjson := TpdfOut{Message: "pdf content as text", Status: "success", Output: out64}
		b, err := json.Marshal(&tjson)
		if err != nil {
			w.WriteHeader(http.StatusOK)
			jsonstr := `{"status":"error","message":"Marshal error"}`
			io.WriteString(w, jsonstr)
			return
		}

		//fmt.Println("aqui: ", string(b))
		w.WriteHeader(http.StatusOK)
		jsonstr := string(b)
		io.WriteString(w, jsonstr)
		return

		// go func() {

		// }()

	} else {
		w.WriteHeader(http.StatusOK)
		jsonstr := `{"status":"error","message":"Only Method GET"}`
		io.WriteString(w, jsonstr)
		return
	}
}

func Base64Encb(btext []byte) string {
	sEnc := b64.URLEncoding.EncodeToString(btext)
	return sEnc
}

func main() {

	mux := http.NewServeMux()

	// handler ping
	mux.Handle("/api/v1/pdf", http.HandlerFunc(handlerExec))

	//fmt.Println(cfg.Host)
	// Create the HTML Server
	ApiServer := &http.Server{
		Addr:    "0.0.0.0:5011",
		Handler: mux,
		//ReadTimeout:    5200 * time.Millisecond,
		//WriteTimeout:   4000 * time.Millisecond,
		MaxHeaderBytes: 1 << 30, // 1G
	}
	log.Println("Run Server Load Pdf: 5011")
	ApiServer.ListenAndServe()
}
