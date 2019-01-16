// Go in action
// @jeffotoni
// 2019-01-16

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// user_steps,user=user-app,screen=list-user,name=jeffotoni idle=20.4 busy=130.5 system=30.4 1465839830100400200
// 																							 20190116015538
//   |    -------------------- --------------  |
//   |             |             |             |
//   |             |             |             |
// +-----------+--------+-+---------+-+---------+
// |measurement|,tag_set| |field_set| |timestamp|
// +-----------+--------+-+---------+-+---------+

func main() {

	filename := "./user_steps.influxdb"
	// write file
	var line string
	f, _ := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	for i := 1; i < 2; i++ {
		now := time.Now()
		nanos := now.UnixNano()
		line = `user_steps,user=user-app,screen=screen-` + strconv.Itoa(i) + `,name=jeffotoni idle=` + fmt.Sprintf("%f", float64(i)*3.4) + ` busy=` + fmt.Sprintf("%f", float64(i)*10.4) + ` system=` + fmt.Sprintf("%f", float64(i)*23.4) + ` ` + fmt.Sprintf("%d", nanos)
		_, _ = f.WriteString(line + "\n")
	}
}
