// File: main.go
package main

//go:generate echo "Criar este file: file_gen.go"
//go:generate touch file_gen.go
//go:generate echo "Vamos colocar conteudo neste file" > file_gen.go

func main() {
    // Use the generated file here.
}
