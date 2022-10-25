package main

import (
	"fmt"
	"time"
)

func main() {

	loc, _ := time.LoadLocation("America/Sao_Paulo")

	t, err := time.ParseInLocation("2006-01-02 15:04:05", "2020-05-22 12:27:00", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	t = t.Add(time.Second * time.Duration(3))
	unix := t.Unix()
	fmt.Println("t: ", unix)

	tm := time.Unix(1590556209, 0)
	fmt.Println("t1: ", tm)
}
