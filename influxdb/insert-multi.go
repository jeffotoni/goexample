// Go in action
// @jeffotoni
// 2019-01-16

package main

import (
	"fmt"
	"math/rand"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
)

const (
	Host     = "http://localhost:8086"
	MyDB     = "usermetrics"
	username = ""
	password = ""
)

var (
	err error
)

// Write InfluxDb
func InsertMulti() {

	//Make client
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	defer c.Close()

	Size := 1000
	rand.Seed(42)
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "us",
	})

	for i := 0; i < Size; i++ {

		users := []string{"jefferson", "hulk", "robin", "brenton"}
		// where
		tags := map[string]string{
			"user":   "user-app",
			"screen": fmt.Sprintf("screen-%d", rand.Intn(Size)),
			"name":   users[rand.Intn(len(users))],
		}

		idle := rand.Float64() * 100.0
		fields := map[string]interface{}{
			"idle":   idle,
			"busy":   360.0 - idle,
			"system": 1.3 * idle,
		}

		timein := time.Now().Local().Add(time.Hour*time.Duration(-24) +
			time.Minute*time.Duration(23) +
			time.Second*time.Duration(34))

		pt, err := client.NewPoint(
			"user_steps",
			tags,
			fields,
			//time.Now(),
			timein,
		)
		if err != nil {
			println("Error:", err.Error())
			continue
		}
		bp.AddPoint(pt)
	}

	err = c.Write(bp)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
}

func main() {

	InsertMulti()
}
