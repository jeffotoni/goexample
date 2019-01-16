// Go in action
// @jeffotoni
// 2019-01-16

package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type SPoints struct {
	Time     string
	Sequence string
	Line     string
}

type DataInflux struct {
	Name    string   `json:"name"`
	Columns []string `json:"columns"`
	//Points  [][]interface{} `json:"points"`
	Points []SPoints `json:"points"`
}

// {
// "name": "log_lines",
// "columns": ["time", "sequence_number", "line"],
// "points": [
// [1400425947368, 1, "this line is first"],
// [1400425947368, 2, "and this is second"]
// ]
// }

func main() {

	//fmt.Println("start write influxdb")
	var df = DataInflux{}
	df.Name = "logs_app"
	df.Columns = append(df.Columns, "time")
	df.Columns = append(df.Columns, "sequence")
	df.Columns = append(df.Columns, "line")

	for i := 0; i < 3; i++ {
		sp := SPoints{
			"140042594736" + strconv.Itoa(i),
			"100" + strconv.Itoa(i),
			strconv.Itoa(i),
		}

		df.Points = append(df.Points, sp)
	}

	b, _ := json.Marshal(df)
	//fmt.Println(df)
	fmt.Println(string(b))
}
