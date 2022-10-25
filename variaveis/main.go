package main

import "fmt"

type Ponto struct {
    x, y float64
}

func main() {
    p := Ponto{2, 3}
    fmt.Printf("p\t%#v\n\n", p)
    fmt.Printf("\t(%p)\n\n", p)
    //var pp *Ponto
    pp := new(Ponto)
    fmt.Printf("pp\t%#v\n", pp)
    fmt.Printf("\t(%p)\n\n", pp)
    pp = &p
    fmt.Printf("pp\t%#v\n", pp)
    fmt.Printf("\t(%p)\n\n", pp)
    fmt.Printf("*pp\t%#v\n\n", *pp)
}
