/*
* Golang config with toml
*
* @package     main
* @author      @jeffotoni
* @size        2017
 */

package main

import (
	"encoding/xml"
	"fmt"
)

//
//
//
type Email struct {
	Where string `xml:"where,attr"`
	Addr  string
}

//
//
//
type Address struct {
	City, State string
}

//
//
//
type Result struct {
	XMLName xml.Name `xml:"Person"`
	Name    string   `xml:"FullName"`
	Phone   string
	Email   []Email
	Groups  []string `xml:"Team>Value"`
	Address
}

func main() {

	v := Result{Name: "none", Phone: "none"}

	err := xml.Unmarshal([]byte(body), &v)

	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println("XMLName:", v.XMLName)
	fmt.Println("Name:", v.Name)

	fmt.Println("Email:", v.Email)
	fmt.Println("Email:", v.Email[0])
	fmt.Println("Email:", v.Email[1])

	fmt.Println("Team:", v.Groups)
	fmt.Println("Address:", v.Address)
}

var body = `<?xml version="1.0" standalone="no"?>
		<Person>
			<FullName>Jefferson O. Lima</FullName>
			<Company>s3 company</Company>
			<Email where="devel">
				<Addr>mail@example.com</Addr>
			</Email>
			<Email where='production'>
				<Addr>opa@work.com</Addr>
			</Email>
			<Team>
				<Value>Dev-Google</Value>
				<Value>Dev-Facebook</Value>
			</Team>
			<City>Belo Horizonte</City>
			<State>Minas Gerais</State>
		</Person>
`
