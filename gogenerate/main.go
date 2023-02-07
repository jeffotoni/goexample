// File: main.go
package main

//go:generate echo "Generating file_gen.go"
//go:generate sh -c "echo 'package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello, World!\")\n}\n' > file_gen.go"
//go:generate go run file_gen.go
func main() {
	// Use the generated file here.
}
