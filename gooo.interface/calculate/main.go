package main

import "fmt"

// go build -gcflags="-m -l"
type Income interface {
	calculate() int
	source() string
}

type FixedBilling struct {
	projectName  string
	biddedAmount int
}

type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

// func (i Income) Interf() {
// 	fmt.Println("no possible interface:", i)
// }

func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

// func (fb *FixedBilling) calculate(t *int) int {
// 	x := *t
// 	return fb.biddedAmount * x
// }

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

func calculateNetIncome(ic []Income) {
	var netincome int = 0
	for _, income := range ic {
		//fmt.Printf("Income From %s = $%d\n", income.source(), income.calculate())
		netincome += income.calculate()
	}
	println(netincome)
}

type Advertisement struct {
	adName     string
	CPC        int
	noOfClicks int
}

func (a Advertisement) calculate() int {
	return a.CPC * a.noOfClicks
}

func (a Advertisement) source() string {
	return a.adName
}

func Math() *int {
	x := 10
	return &x
}

func Math2(x *int) *int {
	return x
}

func Math3(e []int) []int {
	return e
}

func Math4(m map[string]int) int {
	return 10
}

func Math5(b []byte) int {
	return 10
}

func main() {
	a := make([]int, 1e5, 1e5)
	a = append(a, 10)

	var b []*int
	var e []int
	var c *int
	var cx int = 200
	c = &cx
	b = append(b, c)
	e = append(e, 100)
	e = append(e, 200)

	d := Math2(c)
	fmt.Println("d:", d, " b:", b)

	var bb = []byte("jef")
	var bbb []byte
	var bbb2 = make([]byte, 8)
	var m = map[string]int{
		"jeff1": 10,
		"jeff2": 11,
		"jeff3": 12,
	}

	fmt.Println(e, c)
	fmt.Println(a)
	fmt.Println(Math3(e))

	fmt.Println(bb)
	fmt.Println(m)
	fmt.Println(Math4(m))
	fmt.Println(Math5(bbb))
	fmt.Println(Math5(bbb2))

	project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
	project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
	project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
	//var t int = 2
	project1.projectName = ""
	//println(project1.calculate(&t))
	project2.projectName = ""
	project3.projectName = ""

	fmt.Println(project1, project2, project3)
	incomeStreams := []Income{project1, project2, project3}
	calculateNetIncome(incomeStreams)
	println(Math())
}
