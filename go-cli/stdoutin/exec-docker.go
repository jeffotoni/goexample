// Go in action
// @jeffotoni
// 2019-01-16

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	//docker run --rm -i jeffotoni/pdftotext < /tmp/multa.pdf
	//////////
	// docker run --rm -i jeffotoni/pdftotext < /tmp/multa.pdf
	// docker run --rm -i jeffotoni/pdftotext < $PATH_PDF
	// depois executa

	ppdf := flag.String("file", "", "example: -file my.pdf")
	flag.Parse()
	if len(os.Args) < 3 {
		flag.PrintDefaults()
		return
	}
	pdf := *ppdf

	//cmd := exec.Command("docker", "run", "jeffotoni/pdftotext", "< ./golang.pdf")
	cmd := exec.Command("sh", "docker-pdf-totext.sh")
	newEnv := append(os.Environ(), "PATH_PDF="+pdf)
	cmd.Env = newEnv
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("%s", out)
}
