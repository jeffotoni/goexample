package gostring

import (
	//"fmt"
	"strconv"
	"strings"
	"sync"
)

var b strings.Builder
var strs string
var wait sync.WaitGroup

func Gostring(n int) string {
	//for i := 0; i < 1000; i++ {
	wait.Add(1)
	go func() {
		b.WriteString(string(n))
		wait.Done()
	}()
	//}
	wait.Wait()
	//fmt.Println(len(b.String()))
	return b.String()
}

func Gostring2(n int) string {
	b.Grow(100)
	b.WriteString(strconv.Itoa(n))
	return b.String()
}

func Gostring3(n int) string {
	b.Grow(100)
	b.WriteString(string(n))
	return b.String()
}

func Gostring4(n int) string {
	strs += string(n)
	return str
}
