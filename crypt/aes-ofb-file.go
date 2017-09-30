/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size        2017
 */

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	KEY = "AES256Key-32Characters1234567890"
)

func EncryptOFBFile(fileName string) string {

	key := []byte(KEY)

	inFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}
	defer inFile.Close()

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	// If the key is unique for each ciphertext, then it's ok to use a zero
	// IV.
	var iv [aes.BlockSize]byte
	stream := cipher.NewOFB(block, iv[:])

	fileNameCry := fileName + ".cry"

	outFile, err := os.OpenFile(fileNameCry, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer outFile.Close()

	writer := &cipher.StreamWriter{S: stream, W: outFile}

	// Copy the input file to the output file, encrypting as we go.
	if _, err := io.Copy(writer, inFile); err != nil {
		fmt.Println(err)
	}

	return fileNameCry

	// Note that this example is simplistic in that it omits any
	// authentication of the encrypted data. If you were actually to use
	// StreamReader in this manner, an attacker could flip arbitrary bits in
	// the decrypted result.
}

//
//
//
func DecryptOFBFile(fileName string) string {

	key := []byte(KEY)

	inFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer inFile.Close()

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	// If the key is unique for each ciphertext, then it's ok to use a zero
	// IV.
	var iv [aes.BlockSize]byte
	stream := cipher.NewOFB(block, iv[:])

	fileNameDecrypt := strings.Trim(fileName, ".cry")
	outFile, err := os.OpenFile(fileNameDecrypt, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer outFile.Close()

	reader := &cipher.StreamReader{S: stream, R: inFile}
	// Copy the input file to the output file, decrypting as we go.
	if _, err := io.Copy(outFile, reader); err != nil {
		fmt.Println(err)
	}

	return "done: " + fileNameDecrypt

	// Note that this example is simplistic in that it omits any
	// authentication of the encrypted data. If you were actually to use
	// StreamReader in this manner, an attacker could flip arbitrary bits in
	// the output.
}

//
//
//
func main() {

	// multiple 16 len text
	fileCry := EncryptOFBFile("mozilla.pdf")

	fmt.Println(fileCry)

	fileDecry := DecryptOFBFile(fileCry)
	fmt.Println(fileDecry)
}
