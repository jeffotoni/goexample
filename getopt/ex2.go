package main

import (
	// "bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const HELP = `Uso: wcgo [OPÇÃO]... [ARQUIVO]...

Mostra a quantidade de linhas, palavras e bytes para cada ARQUIVO,
e uma linha com o total se mais de um ARQUIVO for especificado.
Uma palavra é uma sequência de caracteres com tamanho diferente
de zero delimitada por espaço em branco.

Se ARQUIVO não for especificado ou for -, lê a entrada padrão.

As opções abaixo podem ser usadas para selecionar qual quantidade será
mostrada, sempre na seguinte ordem: linha, palavra, byte, tamanho máximo
de linha.
  -c, --bytes            mostra a quantidade de bytes
  -m, --chars            mostra a quantidade de caracteres
  -l, --lines            mostra a quantidade de linhas
  -L, --max-line-length  emite o comprimento da linha mais longa
  -w, --words            emite a quantidade de palavras
      --help             mostra esta ajuda e sai
      --version          informa a versão e sai`

const VERSION = "wcgo: version 1.0.0"

type wcST struct {
	*os.File
	bytes    int
	lines    int
	words    int
	chars    int
	longLine int
}

var (
	//opts                                                    func(getopt.Option)
	usage                                                   func()
	err                                                     error
	filename                                                string
	filenames                                               []string
	tBytes, tWords, tLines, tChars, i, value                int
	info                                                    os.FileInfo
	bytesFlag, wordsFlag, linesFlag, charsFlag, maxLineFlag *bool

	stringCmd  string
	stringCmd2 string
	cmdIn      int
)

func main() {

	//
	//
	//
	flag.String("c", "", "empty")

	//
	//
	//
	flag.String("m", "", "empty")

	flag.String("l", "", "empty")

	flag.String("L", "", "empty")

	flag.String("w", "", "empty")

	//
	//
	//
	sizeArgs := len(os.Args)

	if sizeArgs <= 1 {

		flag.PrintDefaults()
		os.Exit(0)
	}

	for x := range os.Args {

		stringCmd = strings.Trim(os.Args[x], "-")
		stringCmd = strings.Trim(stringCmd, "/")
		stringCmd = strings.Trim(stringCmd, ".")
		stringCmd = strings.Trim(stringCmd, "-")

		stringCmd = strings.TrimSpace(stringCmd)
		stringCmd = strings.ToLower(stringCmd)

		// fmt.Println("args: ", sizeArgs, " ", x)

		//
		switch stringCmd {

		case "c":

			stringCmd2 = strings.Trim(os.Args[x+1], "-")
			stringCmd2 = strings.TrimSpace(stringCmd2)

			f := make([]wcST, 1)

			info, err = os.Stat(stringCmd2)

			if err != nil {

				fmt.Fprintf(os.Stderr, "erro: %s\n", err.Error())
				os.Exit(1)

			} else if info.IsDir() {

				fmt.Fprintf(os.Stderr, "%s: é um diretório\n", stringCmd2)
				os.Exit(1)
			}

			f[0].readBytes()
			value = f[0].bytes
			tBytes += value

			fmt.Println("Executar c: ", tBytes)

			cmdIn += 1
		}

	}

	fmt.Println(cmdIn)
}

func (fs *wcST) readBytes() error {

	var (
		cByte []byte
		err   error
	)

	fs.Seek(0, os.SEEK_SET)

	cByte, err = ioutil.ReadAll(fs.File)

	if err != nil {
		return err
	}

	fs.bytes = len(cByte)

	return nil

}
