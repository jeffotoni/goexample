package main

import (
	"fmt"
)

import (
	"github.com/jeffotoni/s3fileworkclient/pkg/gbolt"
)

func main() {

	fmt.Println("Vamos ver: ping:::  ", gbolt.Ping())
	fmt.Println("Vamos ver: ping:::  ", gbolt.Ping())
	fmt.Println("Vamos ver: ping:::  ", gbolt.Ping())
	fmt.Println("Vamos ver: ping:::  ", gbolt.Ping())
	fmt.Println("Vamos ver: ping:::  ", gbolt.Ping())
	fmt.Println("Salvando user")
	gbolt.SaveLogin("jeff.otoni@gmail.com", true)

}
