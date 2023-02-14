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

func Create() {
	tags := map[string]string{"user": "user-app"}
	fields := map[string]interface{}{
		"idle":   9.1,
		"busy":   1.3,
		"system": 3.6,
	}
	pt, err := client.NewPoint("user_steps", tags, fields, time.Now())
	if err == nil {
		fmt.Println("We created user_steps: ", pt.String())
	}
}

func main() {

	Create()
}
