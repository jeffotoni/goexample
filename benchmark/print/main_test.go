package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

var str string = "teste now"
var writer *bufio.Writer

func BenchmarkPrintf(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fmt.Sprintf("%s %d", str, n)
	}
}

func BenchmarkPrintln(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fmt.Sprintln(str, n)
	}
}

func BenchmarkBufio(b *testing.B) {
	//writer = bufio.NewWriter(os.Stdout)
	for n := 0; n < b.N; n++ {
		writer = bufio.NewWriter(os.Stdout)
		writer.WriteString(str)
		//writer.Flush()
	}
}

func BenchmarkAreaPointer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var CC = Circle{r: float64(n)}
		_ = CC.area1()
	}
}

func BenchmarkAreaCopy(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var CC = Circle{r: float64(n)}
		_ = CC.area2()
	}
}

func BenchmarkAreaHeap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var CC = Circle{r: float64(n)}
		_ = CC.area3()
	}
}

func BenchmarkSweets1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RandomDonut1()
	}
}

func BenchmarkSweets2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RandomDonut2()
	}
}

/*func BenchmarkConstBuilder(b *testing.B) {
    for n := 0; n < b.N; n++ {
        var b strings.Builder
        b.WriteString("my_string")
        b.WriteString(longStr)
    }
}*/
