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
	"github.com/influxdata/influxdb/client/v2"
	"log"
	"net/http"
	"time"
)

const (
	MyDB     = "mydb"
	username = ""
	password = ""
)

var (
	ic  client.Client
	err error
)

func hello(res http.ResponseWriter, req *http.Request) {

	res.Write([]byte(fmt.Sprintf("The first hello !")))
}

func connect() client.Client {

	// If you
	// have not
	// connected
	if ic == nil {

		fmt.Println("connect one")
		ic, err = client.NewHTTPClient(client.HTTPConfig{
			Addr:     "http://localhost:8086",
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

func insert(influx client.Client) {

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})

	if err != nil {
		log.Fatal(err)
	}

	// Create a point and add to batch
	tags := map[string]string{"cpu": "cpu-total"}
	fields := map[string]interface{}{
		"idle":   40.1,
		"system": 52.3,
		"user":   41.6,
	}

	pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	// Write the batch
	if err := influx.Write(bp); err != nil {
		log.Fatal(err)
	}
}

func main() {

	influx := connect()
	insert(influx)

	http.HandleFunc("/", hello)

	log.Println("Start listening...")
	if err := http.ListenAndServe(":9002", nil); err != nil {
		log.Fatal(err)
	}
}
