package main

import (
	"log"

	gse "github.com/jeffotoni/gosendmail"
)

func main() {
	email := gse.New2()
	if email == nil {
		log.Println("Error New() check the required fields: EMAIL_HOST,EMAiL_USERNAME,EMAIL_PASSWORD,EMAIL_PORT")
		return
	}
	m := gse.NewMessage("Title here", "Body message here.")
	m.To = []string{"<your-email>@gmail.com"}
	m.CC = []string{"<copy1-email>@gmail.com", "<copy2-email>@gmail.com"}
	m.BCC = []string{"<bc-email>@gmail.com"}
	m.AttachFile("./love.png")
	m.AttachFile("./love.png")
	err := email.Send(m)
	if err != nil {
		log.Println("Error when sending:", err.Error())
		return
	}
	log.Println("Sent with success")
}
