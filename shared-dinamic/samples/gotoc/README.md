# Criando lib em Go para executar em C

Iremos criar uma Lib em Go uma biblioteca .go em seguida vamos gerar arquivos para serem consumidos em C, iremos gerar .h e um .so para que possamos abrir e executar em C.

#### Arquivo em G

```go
package main

import "C"

//export Soma
func Soma(a, b int) int {
	return a + b
}

func main() {}

```
Agora vamos gerar nossos arquivos soma.h e soma.so apartir do nossa lib Go soma.go

```bash
$ go build -o soma.so -buildmode=c-shared soma.go
$ ls -lh lib*
-rw-rw-r-- 1 jeffotoni jeffotoni   98 mar 15 02:05 soma.go
-rw-rw-r-- 1 jeffotoni jeffotoni 1,7K mar 15 02:06 soma.h
-rw-rw-r-- 1 jeffotoni jeffotoni 1,3M mar 15 02:06 soma.so
```
Agora vamos criar nosso arquivo e C e fazer a chamda desta lib gerada em Go.

```c
#include <stdio.h>
#include "soma.h"

int main() {
    int result = Soma(20, 15);
    printf("Resultado: %d\n", result);
    return 0;
}
```

Vamos agora compilar nosso programa em C que seja chama soma.c

```bash
$ gcc -o csoma soma.c soma.so
$ ./csoma
Resultado: 35
```
Caso tenham interesse podem abrir o arquivo soma.h para conferir seu conteúdo.
E também poderá visualizar o arquivo .so com o comando nm.

```bash
$ nm soma.so | grep -e "T Soma"
000000000006c230 T Soma

```