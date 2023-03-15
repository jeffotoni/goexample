package main

import "fmt"
import "github.com/rainycape/dl"
import "log"

//import x "doubler"

func main() {

	fmt.Println("test lib")

	lib, err := dl.Open("./libdoubler.so", 0)
	if err != nil {

		log.Fatalln(err)
	}

	//fmt.Println(lib.DoubleIt(100))
	defer lib.Close()

	//var DoubleIt func([]byte, uint, int, ...interface{}) int
	var DoubleIt func(y int) int

	if err := lib.Sym("DoubleIt", &DoubleIt); err != nil {
		log.Println(err)
	}

	//buf := make([]byte, 200)
	fmt.Println(DoubleIt(100))

	lib2, err := dl.Open("./goshared/main1.so", 0)
	if err != nil {

		log.Fatalln(err)
	}

	//fmt.Println(lib.DoubleIt(100))
	defer lib2.Close()

	//var DoubleIt func([]byte, uint, int, ...interface{}) int
	var Soma func(x, y int) int

	if err := lib2.Sym("Soma", &Soma); err != nil {
		log.Println(err)
	}

	//buf := make([]byte, 200)
	fmt.Println(Soma(100, 400))

	//s := string(buf[:bytes.IndexByte(buf, 0)])
	//fmt.Println(s)

	//fmt.Println(x.DoubleIt(10))
}
