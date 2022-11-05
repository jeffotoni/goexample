# Pattern Singleton

Sempre me deparo com necessidade de implementar o Pattern Singleton em meus projetos, mas em Golang existem algumas particularidades que temos que tomar cuidado. Neste artigo, vou apresentar duas formas de implementar o Singleton usando Golang, a forma “Not Thread Safe” e a forma “Thread Safe”. O objetivo é apresentar de forma prática e técnica as formas de implementação e quando temos que implementar o patterns singleton.

[Post Singleton](https://medium.com/golang-issue/pattern-singleton-com-golang-issue-1-938d1debe626)


```go
/*
* Example DriverPg Go
* @package     main
* @author      @jeffotoni
* @size        10/09/2018
 */

package main

import (
	"fmt"
	"time"
)

type DriverPg struct {
	conn string
}

// variavel Global
var instance *DriverPg

// fazendo uma chamada do
// metodo, antes de qualquer
// chamada
var instanceNew = *Connect() 

... 
```