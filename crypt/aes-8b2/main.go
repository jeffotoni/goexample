// You can edit this code!
// Click here and start typing.
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"golang.org/x/crypto/pbkdf2"
)

var pmid string = `{
    "callProviderOut":{
       "body":{
          "object":[
             {
                "Assesoria":"2901",
                "CodBarras":"846000000006649901090112004515699439308132512994",
                "CodBarrasRemcamp":"846500000001779901093111850200671608208132512996",
                "Custcode":"1.265026350",
                "Desconto":"000",
                "DescontoRemcamp":"040",
                "DtValidadeCamp":"28/02/2022",
                "DtVencimento":"20/07/2021",
                "Fatura":"4515699433",
                "IdFatura":"None",
                "Valor":"64.99"
             },
             {
                "Assesoria":"2901",
                "CodBarras":"846800000008649901090112004497806788808132512993",
                "CodBarrasRemcamp":"846500000001779901093111850200671608208132512996",
                "Custcode":"1.265026350",
                "Desconto":"000",
                "DescontoRemcamp":"040",
                "DtValidadeCamp":"28/02/2022",
                "DtVencimento":"20/06/2021",
                "Fatura":"4497806788",
                "IdFatura":"None",
                "Valor":"64.99"
             }
          ]
       },
       "header":{
          "Content-Type":[
             "application/json"
          ],
          "Server":[
             "Jetty(6.1.26)"
          ],
          "httpStatus":"200"
       }
    },
    "inputHeader":{
       "Accept":"*/*",
       "Accept-Encoding":"gzip, deflate, br",
       "Authorization":"Basic",
       "Cache-Control":"no-cache",
       "Clientid":"URAEVA",
       "Connection":"keep-alive",
       "Content-Type":"application/json",
       "Messageid":"24686EF4F0E44692A258BE44",
       "Postman-Token":"0c26099f-10dd-487d-bae2-6040f2eb45e6",
       "User-Agent":"PostmanRuntime/7.28.4",
       "X-Forwarded-For":"172.18.0.1",
       "X-Forwarded-Host":"localhost",
       "X-Forwarded-Port":"8000",
       "X-Forwarded-Proto":"http",
       "X-Real-Ip":"172.18.0.1",
       "execInfo":"{\"ID\":\"rPreviousPlanInvoicesV1#3e49d1c3-2c82-49ca-a730-93f5df762ab0\",\"Sync\":true,\"StartAt\":\"2022-03-14T16:16:42.3008667-03:00\",\"Timeout\":30000000000}"
    },
    "inputProviderOut":{
       "cpf":"1234",
       "origemAtendimento":"5678"
    },
    "inputUriParams":{
       "channel":"5678",
       "socialSecNo":"1234"
    },
    "inputValidate":{
       "valid":true
    }
 }`

type InputURI struct {
	InputURIParams InputURIParams `json:"inputUriParams"`
}
type InputURIParams struct {
	Channel     string `json:"channel"`
	SocialSecNo string `json:"socialSecNo"`
}

var (
	KeyPmid = "95803820"
)

func main() {
	out := map[string]string{"cpf": "1234"}
	fmt.Println(`map out:`, out)
	fmt.Println(out["cpf"])
	b, _ := json.Marshal(out)
	fmt.Println(string(b))
	var i InputURI
	json.Unmarshal([]byte(pmid), &i)
	fmt.Println(i)
	fmt.Println(`valor de cpf:`, i.InputURIParams.SocialSecNo)
	var my = make(map[string]interface{}, 1)
	my["cpf"] = i.InputURIParams.SocialSecNo
	fmt.Println(my)

	keyHexStr, nonceHexStr, err := genNonce()
	key, nonce, err := validNonce(i.InputURIParams.SocialSecNo, keyHexStr, nonceHexStr)
	if err != nil {
		log.Println(err)
		return
	}
	plainText := "Este é o texto plano a ser cifrado jeffotoni!!"
	cipherText, err := encrypt(key, nonce, plainText)
	if err != nil {
		log.Println(err)
		return
	}
	println("success crypt")
	fmt.Printf("%x\n", cipherText)
	println("....................................................")
	println("decrypt:")
	plainDecryt, err := decrypt(key, nonce, cipherText)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(plainDecryt)
}

func validNonce(cpf, keyHexStr, nonceHexStr string) ([]byte, []byte, error) {
	// key, err := hex.DecodeString(keyHexStr)
	// if err != nil {
	// 	return nil, nil, err
	// }
	key, _ := deriveKey(cpf, []byte(KeyPmid))

	nonce, err := hex.DecodeString(nonceHexStr)
	if err != nil {
		return nil, nil, err
	}

	return key, nonce, nil
}

func genNonce() (string, string, error) {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return "", "", err
	}

	// Never use more than 2^32 random nonces with a given key because of
	// the risk of a repeat.
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", "", err
	}

	return fmt.Sprintf("%x", key), fmt.Sprintf("%x", nonce), nil

}

func deriveKey(passphrase string, salt []byte) ([]byte, []byte) {
	if salt == nil {
		salt = make([]byte, 8)
		// http://www.ietf.org/rfc/rfc2898.txt
		// Salt.
		rand.Read(salt)
	}
	return pbkdf2.Key([]byte(passphrase), salt, 1000, 32, sha256.New), salt
}

func encrypt(publicKey, nonce []byte, plaintext string) (ciphertext string, err error) {

	block, err := aes.NewCipher(publicKey)
	if err != nil {
		log.Println(err)
		return
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Println(err)
		return
	}

	ciphertextBtye := aesgcm.Seal(nil, nonce, []byte(plaintext), nil)
	return fmt.Sprintf("%x", ciphertextBtye), nil
}

func decrypt(key []byte, nonce []byte, cipherHexStr string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	cipherText, err := hex.DecodeString(cipherHexStr)
	if err != nil {
		return "", err
	}

	plainText, err := aesgcm.Open(nil, nonce, []byte(cipherText), nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}
