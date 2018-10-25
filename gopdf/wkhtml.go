package main

import (
	//"bytes"
	"fmt"
	. "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5010"
	}
	http.HandleFunc("/hello", hello)
	{
		println("run server port:" + port)
	}

	http.HandleFunc("/", headerHtmltoPdf)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func headerHtmltoPdf(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	done := gerarHtmltoPdf()

	if done == "done" {

		w.Write([]byte("Pdf gerado com sucesso"))
	} else {
		w.Write([]byte("Pdf deu erro ao gerar"))
	}

}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}

func gerarHtmltoPdf() string {

	// Create new PDF generator
	pdfg, err := NewPDFGenerator()
	if err != nil {
		log.Println("err: ", err)
	}

	//SetPath("./wkhtmlpdf")

	// Set global options
	//pdfg.Dpi.Set(300)
	//pdfg.Orientation.Set(OrientationLandscape)
	//pdfg.Grayscale.Set(true)
	// Create a new input page from an URL
	// page := NewPage("https://godoc.org/github.com/SebastiaanKlippert/go-wkhtmltopdf")

	//<img src="http://api.qrserver.com/v1/create-qr-code/?data=HelloWorld" alt="img" height="42" width="42"></img>
	htmlStr := `<html><body><h1 style="color:red;">This is an html
   from pdf to test color<h1></body></html>`

	pdfg.AddPage(NewPageReader(strings.NewReader(htmlStr)))
	//page := NewPageReader(bytes.NewReader([]byte(htmlStr)))
	//page.NoBackground.Set(true)
	//page.DisableExternalLinks.Set(false)
	//pdfg.AddPage(page)

	// set dpi of the content
	pdfg.Dpi.Set(350)

	// set margins to zero at all direction
	pdfg.MarginBottom.Set(0)
	pdfg.MarginTop.Set(0)
	pdfg.MarginLeft.Set(0)
	pdfg.MarginRight.Set(0)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Write buffer contents to file on disk
	//err = pdfg.WriteFile("/tmp/simplesamplex103939_1.pdf")
	byteFile = pdfg.Bytes()

	fmt.Println(byteFile)

	// if err != nil {
	//  log.Fatal(err)
	// }

	fmt.Println("Done")

	return "done"
	// Output: Done
	//
}
