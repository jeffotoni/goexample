/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size        2017
 */

package main

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	acessekey = "2928x9339"
	PathLocal = "uploads"
)

func main() {

	router := mux.NewRouter()

	router. // multipart/form-data  =namefile=@file.tar
		HandleFunc("/upload-form", uploadFormFile).Methods("POST")

	router. // multipart/form-data =namefile[]=@file.tar
		HandleFunc("/upload-form-multi", uploadFormMuilt).Methods("POST")

	router. // --data-binary "@file1.jpg"
		HandleFunc("/upload-bin", uploadBinary)

	//
	listenAndServer(router)
}

func listenAndServer(router *mux.Router) {

	//
	conf := &http.Server{Handler: router, Addr: ":8081"}

	go conf.ListenAndServe()

	//
	go func() { http.ListenAndServe(":8082", router) }()

	//
	go func() { http.ListenAndServe(":8083", nil) }()

	//
	func() { http.ListenAndServe(":8084", nil) }()
}

func GetName(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	nameX := "Jefferson"

	//server
	fmt.Println("Id: ", vars["id"])

	// client
	fmt.Fprintln(w, "", 200, `{"msg":"`+nameX+`"}`)
}

// This method uploadFormFile only receives files coming in
// the multipart / form-data format, ie comes from a form
// sent by our client
func uploadFormFile(w http.ResponseWriter, r *http.Request) {

	//sizeUpload := r.ContentLength / 1048576 // 1Mb convert to Mb
	file, handler, _ := r.FormFile("fileupload")
	r.ParseMultipartForm(10485760)
	defer file.Close()

	pathUserAcess := PathLocal + "/" + acessekey + "/" + handler.Filename
	existPath, _ := os.Stat(PathLocal + "/" + acessekey)
	if existPath == nil { os.MkdirAll(PathLocal+"/"+acessekey, 0777) }

	f, _ := os.OpenFile(pathUserAcess, os.O_WRONLY|os.O_CREATE, 0777)
	defer f.Close()

	fmt.Println("Copy: ", pathUserAcess)
	sizef, _ := io.Copy(f, file)

	fmt.Fprintln(w, "", 500, `{"msg":"ok upload size recebido"}`, sizef, "Bytes")
}

func uploadFormMuilt(w http.ResponseWriter, r *http.Request) {

	errup := r.ParseMultipartForm(100000)

	if errup != nil {
		log.Printf("Error: Content-type or submitted format is incorrect to upload MultiForm  %s\n", errup)
		return
	}

	multi := r.MultipartForm
	uploadFormFileMulti(w, r, multi.File["fileupload[]"])
}

// This method uploadFormFile only receives files coming in
// the multipart / form-data format, ie comes from a form
// sent by our client
func uploadFormFileMulti(w http.ResponseWriter, r *http.Request, files []*multipart.FileHeader) {

	if len(r.MultipartForm.File["fileupload[]"]) > 0 {
		fmt.Println("size array: ", len(files))
		var uploadSize int64
		for i, _ := range files {
			file, err := files[i].Open()
			defer file.Close()
			pathUserAcess := PathLocal + "/" + acessekey + "/" + files[i].Filename

			existPath, _ := os.Stat(PathLocal + "/" + acessekey)
			if existPath == nil { os.MkdirAll(PathLocal+"/"+acessekey, 0777) }
			f, _ := os.Create(pathUserAcess)
			defer f.Close()

			fmt.Println("Copy: ", pathUserAcess)
			sizef, _ := io.Copy(f, file)
			if err == nil {
				uploadSize += sizef
			}
		}
	}
}

func uploadBinary(w http.ResponseWriter, r *http.Request) {

	pathUpKeyUser := PathLocal + "/" + acessekey
	nameFileUp := r.Header.Get("Name-File")

	pathUpKeyUserFull := pathUpKeyUser + "/" + nameFileUp

	existPath, _ := os.Stat(pathUpKeyUserFull)
	if existPath == nil { os.MkdirAll(pathUpKeyUser, 0777) }

	ff, _ := os.OpenFile(pathUpKeyUserFull, os.O_WRONLY|os.O_CREATE, 0777)
	defer ff.Close()

	fmt.Println("Copy: ", pathUpKeyUserFull)
	sizef, _ := io.Copy(ff, r.Body)

	fmt.Fprintln(w, "", 500, `{"msg":"ok upload size recebido"}`, sizef, "Bytes")
}
