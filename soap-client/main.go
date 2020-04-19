package main

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	soapAction := "urn:consultaCEP"
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

	println(string(payload))
	req, err := http.NewRequest(httpMethod, url, bytes.NewReader(payload))
	if err != nil {
		log.Fatal("Error on creating request object. ", err.Error())
		return
	}

	req.Header.Set("Content-type", "text/xml")
	req.Header.Set("SOAPAction", soapAction)
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

	b, err := ioutil.ReadAll(res.Body)
	fmt.Println(err)
	fmt.Println(string(b))

	return

	type UserList struct {
		XMLName xml.Name
		Body    struct {
			XMLName             xml.Name
			consultaCEPResponse struct {
				XMLName xml.Name
				Return  []string `xml:"return"`
			} `xml:"consultaCEPResponse"`
		}
	}

	result := new(UserList)
	err = xml.NewDecoder(res.Body).Decode(result)
	if err != nil {
		log.Fatal("Error on unmarshaling xml. ", err.Error())
		return
	}

	users := result.Body.consultaCEPResponse.Return
	fmt.Println(strings.Join(users, ", "))

	// 	<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	//     <soap:Body>
	//         <ns2:consultaCEPResponse xmlns:ns2="http://cliente.bean.master.sigep.bsb.correios.com.br/">
	//             <return>
	//                 <bairro>Industrial</bairro>
	//                 <cep>32223030</cep>
	//                 <cidade>Contagem</cidade>
	//                 <complemento2></complemento2>
	//                 <end>Rua Senador Benedito Valadares</end>
	//                 <uf>MG</uf>
	//             </return>
	//         </ns2:consultaCEPResponse>
	//     </soap:Body>
	// </soap:Envelope>

}
