// Go in action
// @jeffotoni
// 2019-01-16

package main

import (
	"fmt"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
)

const (
	Host     = "http://localhost:8086"
	MyDB     = "usermetrics"
	username = ""
	password = ""
)

func InsertOnly() {
	// Make client
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: Host,
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	defer c.Close()

	// Create a new point batch
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})

	// Create a point and add to batch
	tags := map[string]string{"user": "user-app"}
	fields := map[string]interface{}{
		"idle":   10.1,
		"busy":   53.3,
		"system": 46.6,
	}
	pt, err := client.NewPoint("user_steps", tags, fields, time.Now())
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	bp.AddPoint(pt)

	// Write the batch
	c.Write(bp)
}

func main() {

	InsertOnly()
}
