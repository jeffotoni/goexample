// Go in action
// @jeffotoni
// 2019-01-16

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/influxdata/influxdb1-client/v2"
)

const (
	Host     = "http://localhost:8086"
	MyDB     = "usermetrics"
	username = ""
	password = ""
)

var (
	ic  client.Client
	err error
)

func Connect() client.Client {

	// If you
	// have not
	// connected
	if ic == nil {

		fmt.Println("connect one")
		ic, err = client.NewHTTPClient(client.HTTPConfig{
			Addr:     Host,
			Username: username,
			Password: password,
		})
	}

	if err != nil {
		log.Fatal(err)
	}

	startingTime := time.Now().UTC()
	time.Sleep(10 * time.Millisecond)
	endingTime := time.Now().UTC()

	var duration time.Duration = endingTime.Sub(startingTime)

	_, _, errx := ic.Ping(duration)

	if errx != nil {
		log.Fatal(err)
	}

	return ic
}

func main() {

	con := Connect()
	fmt.Println(con)
}
