package main

import (
	"fmt"
	"time"
)

var doneChan chan int

type Promise struct {
	Result chan string
	Error  chan error
}

func NewPromise(f func(chan string, chan error)) *Promise {
	r := make(chan string, 1)
	e := make(chan error, 1)
	go f(r, e)

	return &Promise{
		Result: r,
		Error:  e,
	}
}

func (p *Promise) then(r func(string), e func(error)) {
	select {
	case result := <-p.Result:
		r(result)
		break
	case err := <-p.Error:
		e(err)
	}
}

func (p *Promise) Then(r func(string), e func(error)) {
	go p.then(r, e)
}

func Func1(r chan string, e chan error) {
	<-time.Tick(time.Second * 1)
	r <- "Promise"
}

func main() {

	doneChan = make(chan int, 2)
	var p = NewPromise(Func1)

	p.Then(func(result string) {
		result = result + " then concatenando seu resultado."
		fmt.Println(result)
		doneChan <- 1
	}, func(err error) { fmt.Println(err) })

	<-doneChan
}
