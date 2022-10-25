package main

import (
	"strings"
)

/*var (
	s string = "jefferson"
)*/

func main() {
	a, b := 10, 20
	var ms [2]string

	ms[0] = "jefferson"
	ms[1] = "otoni"
	// ms = append(ms, "jeff")
	//ms = append(ms, "lima")

	// fmt.Println(ms)

	// log.Println("func inline:", s)
	// fmt.Println("teste.. import..")
	// println(ConcatStr("jeff ", " is Go!"), s)
	ConcatStr(ms)
	println(Sum(a, b))

}

func ConcatStr(ms [2]string) (m2 [1]string) {
	ms2 := make([]string, 10)
	ms2 = append(ms2, ms[0])
	ms2 = append(ms2, ms[1])
	// any think
	// m2[0] = ms[0] + ms[1]
	// m2[0] = "jeff"
	m2[0] = strings.Join(ms2, " ")
	return
}

func Sum(a, b int) int {
	return a + b
}
