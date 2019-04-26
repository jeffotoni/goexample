package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

import "github.com/nu7hatch/gouuid"

type Login struct {
	Uuid   string
	User   string
	Status int
}

func (a *Login) SetLogin() {
	fmt.Println("Uuid: ", a.Uuid, "User: ", a.User)
}

func randomString1(size int) string {
	var alpha = "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"
	buf := make([]byte, size)
	for i := 0; i < size; i++ {
		buf[i] = alpha[rand.Intn(len(alpha))]
	}
	return string(buf)
}

func randomString2(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func MyUuid() (uuid string) {

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {

	User := strings.ToLower(randomString1(20))
	u, err := uuid.NewV4()

	if err != nil {
		fmt.Println(err)
		return
	}

	l := Login{Uuid: u.String(), User: User, Status: 10}
	l.SetLogin()
}
