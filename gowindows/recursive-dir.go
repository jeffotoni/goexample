/*
* Golang recursive dir
*
* @package     main
* @author      @jeffotoni
* @size        2018
 */

package main

import (
	"bufio"
	"fmt"
	"github.com/satori/go.uuid"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var visited int

func main() {

	// criar pasta a partir dos nomes dos arquivos
	// encontrados em suas respectivas pastas ano/mes/colaborador

	/*
		AUGUSTO DE FREITAS 8001
		ALICIA SILVA DO AMPARO SANTOS 6035
		ALYSSON TADEU XAVIER DA PAZ 6750
		ANA RODRIGUES DA SILVA NETA 4206
		ANA ROSA LUCAS 3386
		ANDERSON SOUZA ARANTES 11770
		APARECIDA ROSA DO COUTO 5483
		BRUNO OTAVIO TEIXEIRA VENANCIO 960
		CELIO MARCOS DA SIILVA 15877
		CHRISTIANE ALESSANDRA DIAS LIMA 4218
	*/

	// dir/to/walk/skip
	subDirToSkip := "skip"

	// criando func watcher
	doneChan := make(chan bool)

	// lendo o diretorio de origem
	// deixar dinamico
	// ler de um file
	configDir := "./config"

	var dirOrigem, dirDestino string

	var pathVetor []string

	if IsFile(configDir) {

		file, err := os.Open(configDir)

		if err != nil {
			log.Println("error ao ler o config: ", err)
		}

		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() { // internally, it advances token based on sperator

			// path origem
			// path destino
			pathVetor = append(pathVetor, scanner.Text())
		}

		if pathVetor[0] != "" && pathVetor[1] != "" {

			fmt.Println("ok")
			dirOrigem = pathVetor[0]
			dirDestino = pathVetor[1]

			fmt.Println(pathVetor[0])
			fmt.Println(pathVetor[1])

		} else {

			log.Println("error, o config tem que possuir os paths de origem e destino!")
			os.Exit(0)
		}

	} else {

		fmt.Println("Configure seu path por favor no arquivo config")
		os.Exit(0)
	}

	for {

		go func(doneChan chan bool, dirOrigem, subDirToSkip string) {

			defer func() {
				doneChan <- true
			}()

			visited++
			err := filepath.Walk(dirOrigem, func(path string, info os.FileInfo, err error) error {

				if err != nil {
					log.Printf("prevent panic by handling failure accessing a path %q: %v\n", dirOrigem, err)
					return err
				}

				if info.IsDir() && info.Name() == subDirToSkip {

					log.Printf("skipping a dir without errors: %+v \n", info.Name())
					return filepath.SkipDir

				}

				if IsFile(path) {

					// files
					ext, _ := DetectType(path)
					log.Printf("file : %q - ext: %s\n", path, ext)

					// criawr diretorios com nome dos arquivos
					// se o diretorio existir so copiar para dentro dele
					// o arquivo
					// ao criar diretorio, copia o arquivo para dentro dele

					nV := strings.Split(path, "/")

					nomeFileCopy := nV[len(nV)-1]

					pathCopyDir := dirDestino + "/" + nomeFileCopy

					log.Println("COPIAR TO: " + pathCopyDir + "\n")

					// criar diretorio
					if _, err := os.Stat(pathCopyDir); os.IsNotExist(err) {

						os.MkdirAll(pathCopyDir, os.ModePerm)
						log.Println("criado com sucesso")
					}

					// arquivo a ser copiado e seu diretorio
					pathCopyDir = pathCopyDir + "/" + nomeFileCopy

					if !IsFile(pathCopyDir) {

						// copiando o arquivo para dentro do novo diretorio
						CopyFile(path, pathCopyDir)
						log.Println("copiado com sucesso", pathCopyDir)

					} else {

						// muda o nome e copia da mesma forma
						// copiando o arquivo para dentro do novo diretorio

						// gerando uuid para arquivo
						cryptUuid := genUUIDv4()

						// arquivo a ser copiado e seu diretorio
						pathCopyDir = pathCopyDir + "-" + cryptUuid

						CopyFile(path, pathCopyDir)
						log.Println("copiado com sucesso com md5: ", pathCopyDir)
					}

				} else {

					// diretorios
					log.Printf("diretorio: %q\n", path)
				}

				return nil
			})

			if err != nil {
				log.Printf("error walking the path %q: %v\n", dirOrigem, err)
			}

			log.Printf("wait 5 secs, visitou %d\n", visited)
			time.Sleep(1 * time.Second)

		}(doneChan, dirOrigem, subDirToSkip)

		<-doneChan

		fmt.Println("Finalizamos com sucesso toda organização da estrutura!")
		break
	}
}

// it is dir or file
func IsFile(path string) bool {

	if fi, err := os.Stat(path); err == nil {
		if fi.Mode().IsRegular() {
			return true
		}
	}
	return false
}

// it is dir or file
func IsDir(path string) bool {

	fileInfo, _ := os.Stat(path)
	return fileInfo.IsDir()
}

// Detected types of files
func DetectType(pathFile string) (contentType string, err error) {

	file, err := os.Open(pathFile)

	if err != nil {
		return
	}

	defer file.Close()

	// // Only the
	// first 512 bytes
	// are used to sniff the content type.
	buffer := make([]byte, 512)

	// Always returns a valid content-type
	// and "application/octet-stream"
	// if no others seemed to match.
	//contentType := http.DetectContentType(buffer)
	n, err := file.Read(buffer)

	if err != nil && err != io.EOF {
		return
	}

	contentType = http.DetectContentType(buffer[:n])

	return
}

// CopyFile copies a file from src to dst. If src and dst files exist, and are
// the same, then return success. Otherise, attempt to create a hard link
// between the two files. If that fail, copy the file contents from src to dst.
func CopyFile(src, dst string) (err error) {

	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}

	dfi, err := os.Stat(dst)

	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return
		}
	}
	if err = os.Link(src, dst); err == nil {
		return
	}

	err = copyFileContents(src, dst)
	return
}

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func copyFileContents(src, dst string) (err error) {

	in, err := os.Open(src)
	if err != nil {
		return
	}

	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return
	}

	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}

// gerando semente unica
// uuid
func genUUIDv4() (id string) {

	idX, err := uuid.NewV4()

	if err != nil {

		log.Println("Error gerar Uuid", err)
		//return
	}

	// convert para string
	id = fmt.Sprintf("%s", idX)

	return
}

//func Md5Gerar() {
//h := md5.New()
//times := time.Now().UnixNano() / 1000000
//fmt.Println(times)
//timess := strconv.FormatInt(times, 10)
//io.WriteString(h, timess+path)
//cryptmd5 := fmt.Sprintf("%x", h.Sum(nil))
//}
