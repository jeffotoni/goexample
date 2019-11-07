package main_test

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

/*func BenchmarkConstBuilder(b *testing.B) {
    for n := 0; n < b.N; n++ {
        var b strings.Builder
        b.WriteString("my_string")
        b.WriteString(longStr)
    }
}*/
