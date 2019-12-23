package main

import "fmt"
import "github.com/jeffotoni/codenation.dev/aula01/gomodinit/project2/pkg/math"
import "github.com/jeffotoni/codenation.dev/aula01/gomodinit/project2/pkg/util"
import . "github.com/jeffotoni/gcolor"

func main() {
	fmt.Println("Projeto 2 github..")
	math.Sum(3, 4)
	util.Lower("JEFFOTONI")
	Yellow.Cprintln("Yellow color!!")
}
