// Go in action
// @jeffotoni
// 2019-01-16

package main

import (
	"fmt"

	"github.com/influxdata/influxdb1-client/v2"
)

const (
	Host     = "http://localhost:8086"
	MyDB     = "usermetrics"
	username = ""
	password = ""
)

func Query(measurement string) {

	// Make client
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: Host,
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	defer c.Close()

	q := client.Query{
		Command:  "select * from " + measurement,
		Database: MyDB,
	}
	if response, err := c.Query(q); err == nil && response.Error() == nil {
		fmt.Println(response.Results)
	}
}

func main() {
	Query("user_steps")
}
