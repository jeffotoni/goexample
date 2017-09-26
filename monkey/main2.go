package main

import (
	  "fmt"
	    "unsafe"
    )

    func a() int { return 1 }

    func main() {
	      f := a
	        fmt.Printf("0x%x\n", *(*uintptr)(unsafe.Pointer(&f)))
	}
