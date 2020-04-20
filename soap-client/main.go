package main

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	//soapAction := "urn:consultaCEP"
	httpMethod := "POST"
	//username := ""
	//password := ""

	url := fmt.Sprintf("%s%s%s",
		"https://apps.correios.com.br",
		"/SigepMasterJPA/AtendeClienteService/AtendeCliente",
		"",
	)

	payload := []byte(strings.TrimSpace(`
    <x:Envelope xmlns:x="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cli="http://cliente.bean.master.sigep.bsb.correios.com.br/">
    <x:Body>
        <cli:consultaCEP>
            <cep>04167001</cep>
        </cli:consultaCEP>
    </x:Body>
</x:Envelope>
`,
	))

	req, err := http.NewRequest(httpMethod, url, bytes.NewReader(payload))
	if err != nil {
		log.Fatal("Error on creating request object. ", err.Error())
		return
	}

	req.Header.Set("Content-type", "text/xml; charset=utf-8")
	//req.Header.Set("SOAPAction", soapAction)
	//req.SetBasicAuth(username, password)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Error on dispatching request. ", err.Error())
		return
	}

	type Envelope struct {
		XMLName xml.Name `xml:"Envelope"`
		Text    string   `xml:",chardata"`
		Soap    string   `xml:"soap,attr"`
		Body    struct {
			Text                string `xml:",chardata"`
			ConsultaCEPResponse struct {
				Text   string `xml:",chardata"`
				Ns2    string `xml:"ns2,attr"`
				Return struct {
					Text         string `xml:",chardata"`
					Bairro       string `xml:"bairro"`
					Cep          string `xml:"cep"`
					Cidade       string `xml:"cidade"`
					Complemento2 string `xml:"complemento2"`
					End          string `xml:"end"`
					Uf           string `xml:"uf"`
				} `xml:"return"`
			} `xml:"consultaCEPResponse"`
		} `xml:"Body"`
	}

	result := new(Envelope)
	err = xml.NewDecoder(res.Body).Decode(result)
	if err != nil {
		log.Fatal("Error on unmarshaling xml. ", err.Error())
		return
	}

	users := result.Body.ConsultaCEPResponse.Return
	fmt.Println("users:", users)
	fmt.Println("cep:", users.Cep)
	fmt.Println("cidade:", users.Cidade)
	fmt.Println("bairro:", users.Bairro)

}
