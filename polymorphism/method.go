/*
* Example sync.Mutex
*
* @package     main
* @author      @jeffotoni
* @size        01/08/2017
*
 */

package main

import (
	"fmt"
)

type Tipo struct {
	name string
}

func (valor *Tipo) MetodoA() {
	fmt.Printf("%s\n", valor.name)
}

func (valor Tipo) MetodoB(nameX string) {
	fmt.Printf("%s\n", valor.name)
	fmt.Println("name: ", nameX)
}

func main() {

	//
	//
	//
	valor := Tipo{"valueType"}

	//
	//
	//
	valor.MetodoB("jeff")

	//
	//
	//
	Tipo.MetodoB(valor, "jeff")

	//
	//
	//
	ponteiro := &Tipo{"ponteiro"}

	//
	//
	//
	(*Tipo).MetodoA(ponteiro)

	//
	//
	//
	ponteiro.MetodoA()

}
