/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size        2017
 */

package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func Blowfish(password string) string {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {

		panic(err)
	}

	return string(bytes)
}

//
//
//
func CheckBlowfish(password, hash string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil

}

func main() {

	password := "1234567890#$"

	hashBlo := Blowfish(password)

	if CheckBlowfish(password, hashBlo) {

		fmt.Println("ok, password correct!")

	} else {

		fmt.Println("ok, password incorrect!")
	}

}
