package main

import (
      "fmt"
)

// https://golang.org/doc/effective_go.html#allocation_make
func main() {

      arr := [10]int{}
      slice := arr[:]
      slice2 := arr[:2]
      a := make([]byte, 10)
      b := [10]byte{}

      var p *[]int = new([]int)
      *p = make([]int, 10, 10)

      // better
      v := make([]int, 10)

      fmt.Println("arr   :", arr)
      fmt.Println("slice :", slice)
      fmt.Println("slice2:", slice2)
      fmt.Println("a     :", a)
      fmt.Println("b     :", b)
      fmt.Println("p     :", p)
      fmt.Println("v     :", v)

      fmt.Println("-- MAKE --")
      a := make([]int, 0)
      aPtr := &a
      fmt.Println("pointer == nil :", *aPtr == nil)
      fmt.Printf("pointer value: %p\n\n", *aPtr)

      fmt.Println("-- COMPOSITE LITERAL --")
      b := []int{}
      bPtr := &b
      fmt.Println("pointer == nil :", *bPtr == nil)
      fmt.Printf("pointer value: %p\n\n", *bPtr)

      fmt.Println("-- NEW --")
      cPtr := new([]int)
      fmt.Println("pointer == nil :", *cPtr == nil)
      fmt.Printf("pointer value: %p\n\n", *cPtr)

      fmt.Println("-- VAR (not initialized) --")
      var d []int
      dPtr := &d
      fmt.Println("pointer == nil :", *dPtr == nil)
      fmt.Printf("pointer value: %p\n", *dPtr)

}
