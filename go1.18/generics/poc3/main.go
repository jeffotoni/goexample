package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//type MyStructGeneric[T any] struct {
//	G *MyStructGeneric[T]
//}
//
//func NewMyStructGeneric[T any]() *MyStructGeneric[T] {
//	return &MyStructGeneric[T]{
//		G: &MyStructGeneric[T]{},
//	}
//}

type MyStructGeneric[T any] struct {
	field T
}

func NewMyStructGeneric[T any]() MyStructGeneric[T] {
	return MyStructGeneric[T]{}
}

type User struct {
	ID   string
	Name string
	Cpf  int
}
type Company struct {
	ID            int
	Phone         string
	CorporateName string
}

type Covid struct {
	Week   int
	Data   string
	County string
}

type Universities struct {
	Name    string
	Acronym string
	Country string
}

type RabbitMQ struct {
	Headers string
	Payload string
}

func myFunc(n any) {
	b, err := json.Marshal(&n)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
}

func main() {
	var u = &MyStructGeneric[Universities]{}
	u.field.Name = "Federal do Rio de Janeiro"
	u.field.Acronym = "UFRJ"
	u.field.Country = "Brasil"
	fmt.Println(u)

	myFunc(u.field)

	// bj, _ := json.Marshal(&u.field)
	// println("json")
	// fmt.Println(string(bj))

	var c MyStructGeneric[Covid]
	c.field.Week = 202211
	c.field.Data = "2022-03-18"
	c.field.County = "Cruzeiro do Sul"
	fmt.Println(c)
	bj, _ := json.Marshal(&c.field)
	println("json")
	fmt.Println(string(bj))

	var b = new(MyStructGeneric[User])
	b.field.ID = "987654321"
	b.field.Name = "jeffotoni"
	b.field.Cpf = 39393939393
	fmt.Println(b)
	bj, _ = json.Marshal(&b.field)
	println("json")
	fmt.Println(string(bj))

	p := NewMyStructGeneric[Company]()
	p.field.ID = 123456
	p.field.CorporateName = "GENERIC COMPANY LTDA"
	p.field.Phone = "3198765433"
	fmt.Println(p)

	bj, _ = json.Marshal(&p.field)
	println("json")
	fmt.Println(string(bj))

	r := NewMyStructGeneric[RabbitMQ]()
	r.field.Headers = "Content-Type:application/json"
	r.field.Payload = `{"name":"my name","doc":"23232"}`
	myFunc(r.field)
}
