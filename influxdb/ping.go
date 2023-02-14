// Go in action
// @jeffotoni
// 2019-01-16

package main

import (
	"fmt"

	client "github.com/influxdata/influxdb1-client/v2"
)

const (
	Host     = "http://localhost:8086"
	MyDB     = "usermetrics"
	username = ""
	password = ""
)

// Ping the cluster using the HTTP client
func Ping() {
	// Make client
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: Host,
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	defer c.Close()

	ms, version, err := c.Ping(0)
	if err != nil {
		fmt.Println("Error pinging InfluxDB Cluster: ", err.Error())
	}

	fmt.Println("time: ", ms, " version:", version)
}

func main() {
	Ping()
}
