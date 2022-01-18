package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/achiku/xml"
)

// SOAPEnvelope envelope
type SOAPEnvelope struct {
	XMLName xml.Name    `xml:"http://schemas.xmlsoap.org/soap/envelope Envelope"`
	Header  *SOAPHeader `xml:",omitempty"`
	Body    SOAPBody    `xml:",omitempty"`
}

// SOAPHeader header
type SOAPHeader struct {
	XMLName xml.Name    `xml:"http://schemas.xmlsoap.org/soap/envelope Header"`
	Content interface{} `xml:",omitempty"`
}

// SOAPBody body
type SOAPBody struct {
	XMLName xml.Name    `xml:"http://schemas.xmlsoap.org/soap/envelope Body"`
	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

// SOAPFault fault
type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope Fault"`
	Code    string   `xml:"faultcode,omitempty"`
	String  string   `xml:"faultstring,omitempty"`
	Actor   string   `xml:"faultactor,omitempty"`
	Detail  string   `xml:"detail,omitempty"`
}

// UnmarshalXML unmarshal SOAPBody
func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}
	var (
		token    xml.Token
		err      error
		consumed bool
	)
Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}
		if token == nil {
			break
		}
		envelopeNameSpace := "http://schemas.xmlsoap.org/soap/envelope/"
		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError(
					"Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == envelopeNameSpace && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil
				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}
				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}
				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}
	return nil
}

// Name struct
type Name struct {
	XMLName xml.Name `xml:"http://example.com/ns Name"`
	First   string   `xml:"First,omitempty"`
	Last    string   `xml:"Last,omitempty"`
}

// Auth authorization header
type Auth struct {
	XMLName xml.Name `xml:"http://example.com/ns Auth"`
	UserID  string   `xml:"UserID"`
	Pass    string   `xml:"Pass"`
}

// Person struct
type Person struct {
	XMLName xml.Name `xml:"http://example.com/ns Person"`
	ID      int      `xml:"Id,omitempty"`
	Name    *Name
	Age     int `xml:"Age,omitempty"`
}

// AbstractResponse struct
type AbstractResponse struct {
	Code   string `xml:"Code,omitempty"`
	Detail string `xml:"Detail,omitempty"`
}

// ConcreteResponse struct
type ConcreteResponse struct {
	*AbstractResponse
	XMLName           xml.Name `xml:"http://example.com/ns ConcreteResponse"`
	AdditionalMessage string   `xml:"AdditionalMessage,omitempty"`
}

// ProcessBRequest struct
type ProcessBRequest struct {
	XMLName   xml.Name `xml:"http://example.com/ns ProcessBRequest"`
	RequestID string   `xml:"RequestId"`
}

// ProcessARequest struct
type ProcessARequest struct {
	XMLName   xml.Name `xml:"http://example.com/ns ProcessARequest"`
	RequestID string   `xml:"RequestId"`
}

// ProcessAResponse struct
type ProcessAResponse struct {
	*AbstractResponse
	XMLName xml.Name `xml:"http://example.com/ns ProcessAResponse"`
	ID      string   `xml:"Id,omitifempty"`
	Process string   `xml:"Process,omitifempty"`
}

// ProcessBResponse struct
type ProcessBResponse struct {
	*AbstractResponse
	XMLName xml.Name `xml:"http://example.com/ns ProcessBResponse"`
	ID      string   `xml:"Id,omitifempty"`
	Process string   `xml:"Process,omitifempty"`
	Amount  string   `xml:"Amount,omitifempty"`
}

func processA() ProcessAResponse {
	return ProcessAResponse{
		AbstractResponse: &AbstractResponse{
			Code:   "200",
			Detail: "success",
		},
		ID:      "100",
		Process: "ProcessAResponse",
	}
}

func processB() ProcessBResponse {
	return ProcessBResponse{
		AbstractResponse: &AbstractResponse{
			Code:   "200",
			Detail: "success",
		},
		ID:      "100",
		Process: "ProcessBResponse",
		Amount:  "10000",
	}
}

func soapActionHandler(w http.ResponseWriter, r *http.Request) {
	soapAction := r.Header.Get("soapAction")
	var res interface{}
	switch soapAction {
	case "processA":
		res = processA()
	case "processB":
		res = processB()
	default:
		res = nil
	}
	v := SOAPEnvelope{
		Body: SOAPBody{
			Content: res,
		},
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/xml")
	x, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(x)
	return
}

func soapBodyHandler(w http.ResponseWriter, r *http.Request) {
	rawbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	a := regexp.MustCompile(`<ProcessARequest xmlns="http://example.com/ns">`)
	b := regexp.MustCompile(`<ProcessBRequest xmlns="http://example.com/ns">`)

	var res interface{}
	if a.MatchString(string(rawbody)) {
		res = processA()
	} else if b.MatchString(string(rawbody)) {
		res = processB()
	} else {
		res = nil
	}
	v := SOAPEnvelope{
		Body: SOAPBody{
			Content: res,
		},
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/xml")
	x, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(x)
	return
}

// NewSOAPMux return SOAP server mux
func NewSOAPMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/dispatch/soapaction", soapActionHandler)
	mux.HandleFunc("/dispatch/soapbody", soapBodyHandler)
	return mux
}

// NewSOAPServer create i2c mock server
func NewSOAPServer(port string) *http.Server {
	mux := NewSOAPMux()
	server := &http.Server{
		Handler: mux,
		Addr:    "localhost:" + port,
	}
	return server
}

func main() {
	mux := NewSOAPMux()
	s := http.Server{
		Addr:        "0.0.0.0:9000",
		Handler:     mux,
		ReadTimeout: 8 * time.Second,
	}
	log.Println("Run Server Soap port:9000")
	s.ListenAndServe()
}
