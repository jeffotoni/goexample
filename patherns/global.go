/*
* Example Singleton C
* @package     main
* @author      @jeffotoni
* @size        10/09/2018
 */

package main

import "fmt"

var (
	conn_pg2 = 0 // variavel global estática

	// O extern define variaveis que
	// serao usadas em um arquivo
	// apesar de terem sido declaradas em outro
	Conn_pg3 = 0 // variavel extern
)

func main() {

	fmt.Println(conn_pg2)
}
