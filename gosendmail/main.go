package main

import (
	"log"

	gse "github.com/jeffotoni/gosendmail"
)

func main() {
	email := gse.New()
	if email == nil {
		log.Println("Error New() check the required fields: EMAIL_HOST,EMAiL_USERNAME,EMAIL_PASSWORD,EMAIL_PORT")
		return
	}
	m := gse.NewMessage("Vamos testar pkg sendmail", "Pkg feito em Go para enviar email SMTP")
	m.To = []string{"<your-mail@gmail.com>"}
	//m.CC = []string{""}
	//m.BCC = []string{"bc@gmail.com"}
	m.AttachFile("./love.png")
	if err := email.Send(m); err != nil {
		log.Println("Error when sending:", err.Error())
		return
	}
	log.Println("Sent with success")
}
