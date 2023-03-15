# Criando plugin Go .so e chamando em Go

Iremos criar uma Lib em Go em seguida vamos gerar arquivo .so para serem consumidos em Go.

#### Arquivo em Go

```go
package main

import "C"

//export Soma
func Soma(a, b int) int {
	return a + b
}

func main() {}

```
Agora vamos gerar nossos arquivo .so e soma.so apartir do nossa lib Go soma.go

```bash
$ go build -buildmode=plugin -o soma_plugin.so soma_plugin.go
$ ls -lh soma*
-rw-rw-r-- 1 jeffotoni jeffotoni  161 mar 15 11:32 soma_plugin.go
-rw-rw-r-- 1 jeffotoni jeffotoni 1,6M mar 15 11:33 soma_plugin.so
```

Agora vamos criar nosso programa Go para usar nosso plugin

```go
package main

import (
	"fmt"
	"plugin"
)

func main() {
	plug, err := plugin.Open("soma_plugin.so")
	if err != nil {
		fmt.Println("Erro ao carregar o plugin:", err)
		return
	}

	somaFunc, err := plug.Lookup("Soma")
	if err != nil {
		fmt.Println("Erro ao buscar a função 'Soma':", err)
		return
	}

	soma, ok := somaFunc.(func(int, int) int)
	if !ok {
		fmt.Println("Erro ao fazer type assertion para a função 'Soma'")
		return
	}

	resultado := soma(20, 15)
	fmt.Println("Resultado da soma: ", resultado)
}

```

Vamos agora executar nosso programa em Go.

```bash
$ go run main.go
Resultado da soma:  35
```

Vamos agora compilar nosso programa em Go.

```bash
$ go build -o gsoma main.go
$ ./gsoma
Resultado da soma:  35
```

Caso tenham interesse podem abrir o arquivo soma.h para conferir seu conteúdo.
E também poderá visualizar o arquivo .so com o comando nm.

```bash
$ nm soma.so | grep -e "T Soma"
0000000000097250 T Soma

```