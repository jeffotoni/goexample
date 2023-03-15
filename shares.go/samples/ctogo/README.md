# Criando lib em C para executar em Go

Iremos criar uma Lib em C uma biblioteca .h e um .c para que possamos abrir e executar em Go.

Um detalhe é que iremos gerar uma lib statica em C com um exemplo bem simples para que possamos entender todo o processo.

#### Arquivos em C

##### soma.c
```c
#include "soma.h"

int Soma(int x, int y) {
    return x + y;
}

```
Agora vamos criar nossa biblioteca .h
##### soma.h

```c
#ifndef SOMA_H
#define SOMA_H

int Soma(int x, int y);

#endif // SOMA_H

```

Vamos agora compilar e deixarmos pronto para usarmos em nosso programa em Go.

```bash
$ export LD_LIBRARY_PATH=.
$ gcc -c -o libsoma.a soma.c
$ ar rcs libsoma.a libsoma.o
$ ls -lh lib*
-rw-rw-r-- 1 jeffotoni jeffotoni 1,4K mar 15 01:38 libsoma.a
-rw-rw-r-- 1 jeffotoni jeffotoni 1,3K mar 15 01:38 libsoma.o

```
Agora que compilamos tudo e temos libsoma.a, podemos fazer a chamada em nosso programa Go.
Podemos apagar a libsoma.o, ela só será utilizada para gera nossa libsoma.a então podemos remove-la para não confundir.

```bash
$ rm -f libsoma.o
```
#### nosso main.go


```go
package main

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -lsoma
// #include "soma.h"
import "C"
import "fmt"

// go build -o gsoma main.go
func main() {
	x := 15
	y := 20
	result := C.Soma(C.int(x), C.int(y))
	fmt.Printf("Resultado: %d\n", result)
}

```


```bash
$ go run main
Resultado: 35
```

Ou podemos compilar desta forma.


```bash
$ go build -o gsoma main.go
$ ./gsoma
Resultado: 35
```

O arquivo binário é dinâmico ou seja nosso CGO_ENABLED é 1 ou seja usamos CGO para comunicar com C. libs C etc. Vamos da uma conferida.

```bash
$ ldd gsoma
  	linux-vdso.so.1 (0x00007ffd664a2000)
	libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007f863da00000)
	/lib64/ld-linux-x86-64.so.2 (0x00007f863ddb3000)

```
