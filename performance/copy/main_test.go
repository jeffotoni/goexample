package main

import (
	//"fmt"
	"github.com/jeffotoni/archiviobrasilone/apicorenew/pkg/util"
	"io"
	"os"
	"strings"
	"testing"
)

func BenchmarkCopy(b *testing.B) {
	for n := 0; n < b.N; n++ {
		io.Copy(os.Stdout, strings.NewReader("\nStart1 Sync....\n"))
	}
}

func BenchmarkPrintln(b *testing.B) {

	for n := 0; n < b.N; n++ {
		util.Printnew("\nStart2 Sync....\n")
	}
}
