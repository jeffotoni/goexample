package main

import (
	"bufio"
	//"flag"
	"fmt"
	"github.com/pborman/getopt/v2"
	"io/ioutil"
	"os"
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

func main() {

	var (
		opts                                                    func(getopt.Option)
		usage                                                   func()
		err                                                     error
		filename                                                string
		filenames                                               []string
		tBytes, tWords, tLines, tChars, i, value                int
		info                                                    os.FileInfo
		bytesFlag, wordsFlag, linesFlag, charsFlag, maxLineFlag *bool
	)

	bytesFlag = getopt.BoolLong("bytes", 'c', "mostra a quantidade de bytes")
	wordsFlag = getopt.BoolLong("words", 'w', "mostra a quantidade de palavras")
	linesFlag = getopt.BoolLong("lines", 'l', "mostra a quantidade de linhas")
	charsFlag = getopt.BoolLong("chars", 'm', "mostra a quantidade de caracteres")
	maxLineFlag = getopt.BoolLong("max-line-length", 'L', "emite o comprimento da linha mais longa")

	getopt.BoolLong("help", 'h', "exibe ajuda e sai")
	getopt.BoolLong("version", 'v', "exibe a versão e sai")

	getopt.SetUsage(usage)
	getopt.Parse()

	usage = func() {
		fmt.Println(HELP)
		os.Exit(0)
	}

	if !*bytesFlag && !*wordsFlag && !*linesFlag && !*charsFlag && !*maxLineFlag {
		*bytesFlag, *charsFlag, *linesFlag = false, true, true
	}

	opts = func(opt getopt.Option) {

		switch opt.Name() {
		case "-h", "--help":
			getopt.Usage()
		case "-v", "--version":
			fmt.Println(VERSION)
			os.Exit(0)
		}
	}

	getopt.Visit(opts)

	fmt.Println(getopt.NArgs())

	//os.Exit(0)

	if len(os.Args) < 2 || getopt.NArgs() == 0 {
		getopt.Usage()
	}

	filenames = append(filenames, getopt.Args()...)

	for _, filename := range filenames {

		info, err = os.Stat(filename)

		if err != nil {
			fmt.Fprintf(os.Stderr, "erro: %s\n", err.Error())
			os.Exit(1)
		} else if info.IsDir() {
			fmt.Fprintf(os.Stderr, "%s: %s: é um diretório\n", os.Args[0], filename)
			os.Exit(1)
		}

	}

	f := make([]wcST, len(filenames))

	opts = func(opt getopt.Option) {

		switch opt.Name() {
		case "--bytes", "-c":
			f[i].readBytes()
			value = f[i].bytes
			tBytes += value
		case "--words", "-w":
			f[i].readWords()
			value = f[i].words
			tWords += value
		case "--lines", "-l":
			f[i].readLines()
			value = f[i].lines
			tLines += value
		case "--chars", "-m":
			f[i].readChars()
			value = f[i].chars
			tChars += value
		case "--max-line-length", "-L":
			f[i].readLongLine()
			value = f[i].longLine
		}

		fmt.Printf("%d ", value)
	}

	for i, filename = range filenames {

		f[i].File, err = os.Open(filename)

		if err != nil {
			continue
		}

		defer f[i].Close()

		fmt.Println("::", *bytesFlag)
		fmt.Println(*charsFlag)
		fmt.Println(*linesFlag)
		fmt.Println(*wordsFlag)

		getopt.Visit(opts)
		fmt.Printf("%s\n", f[i].Name())
	}

	if getopt.NArgs() > 1 {

		opts = func(opt getopt.Option) {
			switch opt.Name() {
			case "--bytes", "-c":
				fmt.Printf("%d ", tBytes)
			case "--words", "-w":
				fmt.Printf("%d ", tWords)
			case "--lines", "-l":
				fmt.Printf("%d ", tLines)
			case "--chars", "-m":
				fmt.Printf("%d ", tChars)
			}

		}

		getopt.Visit(opts)
		fmt.Println("Total")

	} else {

		fmt.Println("aqui..")
	}

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

func (fs *wcST) readWords() error {

	var (
		scanner *bufio.Scanner
		word    int
		err     error
	)

	fs.Seek(0, os.SEEK_SET)

	scanner = bufio.NewScanner(fs.File)
	scanner.Split(bufio.ScanWords)

	for ; scanner.Scan(); word++ {
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	fs.words = word

	return nil
}

func (fs *wcST) readLines() error {

	var (
		scanner *bufio.Scanner
		line    int
		err     error
	)

	fs.Seek(0, os.SEEK_SET)

	scanner = bufio.NewScanner(fs.File)

	for ; scanner.Scan(); line++ {
	}

	if err = scanner.Err(); err != nil {
		return nil
	}

	fs.lines = line

	return nil
}

func (fs *wcST) readChars() error {

	var (
		scanner *bufio.Scanner
		char    int
		err     error
	)

	fs.Seek(0, os.SEEK_SET)

	scanner = bufio.NewScanner(fs.File)
	scanner.Split(bufio.ScanRunes)

	for ; scanner.Scan(); char++ {
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	fs.chars = char

	return nil
}

func (fs *wcST) readLongLine() error {

	var (
		max     int
		cur     int
		scanner *bufio.Scanner
		err     error
	)

	fs.Seek(0, os.SEEK_SET)
	scanner = bufio.NewScanner(fs.File)

	for scanner.Scan() {
		if cur = len(scanner.Text()); cur > max {
			max = cur
		}
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	fs.longLine = max

	return nil
}
