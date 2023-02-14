package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	//"log"

	"math/big"
)

func createHash(key string) []byte {
	hasher := md5.New()
	io.WriteString(hasher, key)
	hash := hasher.Sum(nil)
	dst := make([]byte, hex.EncodedLen(len(hash)))
	hex.Encode(dst, hash)

	return dst
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher(createHash(passphrase))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)

	return ciphertext
}

func decrypt(data []byte, passphrase string) []byte {
	key := createHash(passphrase)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func encryptFile(filename string, data []byte, passphrase string) {
	f, _ := os.Create(filename)
	defer f.Close()
	f.Write(encrypt(data, passphrase))
}

func decryptFile(filename string, passphrase string) []byte {
	data, _ := ioutil.ReadFile(filename)
	return decrypt(data, passphrase)
}

func main() {
	fmt.Println("Starting the application...")
	ciphertext := encrypt([]byte("100000_00000000"), "password")

	bi := big.NewInt(0)
	bi.SetBytes(ciphertext)
	fmt.Println(" Big int: ", bi.String())

	f := fmt.Sprintf("%x", bi)
	fmt.Println(" bytes  : ", bi.Bytes())

	bi2 := big.NewInt(0)
	bi2.SetString(f, 16)

	fmt.Println(" Big int: ", bi2.String())

	//log.Println(ciphertext)

	fmt.Printf("Encrypted: %x\n", ciphertext)
	plaintext := decrypt(bi.Bytes(), "password")
	fmt.Printf("Decrypted: %s\n", plaintext)

}
