package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	fmt.Println("Today : ", now.Format("Mon, Jan 2, 2006 at 3:04pm"))

	longTimeAgo := time.Date(2010, time.May, 18, 23, 0, 0, 0, time.UTC)

	// compare time with time.Equal()

	sameTime := longTimeAgo.Equal(now)

	fmt.Println("longTimeAgo equals to now ? : ", sameTime)

	// calculate the time different between today
	// and long time ago

	diff := now.Sub(longTimeAgo)

	// convert diff to days
	days := int(diff.Hours() / 24)

	fmt.Printf("18th May 2010 was %d days ago \n", days)

}
