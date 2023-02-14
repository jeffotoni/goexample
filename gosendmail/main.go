package main

import (
	"crypto/tls"
	"fmt"

	gomail "gopkg.in/mail.v2"
)

func main() {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "user@gmail.com")

	// Set E-Mail receivers
	m.SetHeader("To", "youremail@gmail.com")

	// Set E-Mail subject
	m.SetHeader("Subject", "Gomail test subject")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", "Test de envio de email...")

	// attach
	m.Attach("./love.png")

	// Settings for SMTP server
	d := gomail.NewDialer("s3.smtp.com", 25, "user", "pass")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return
	}

	return
}

// package main

// import (
// 	"fmt"
// 	//"net/smtp"
// 	xmail "github.com/farmerx/mail"
// 	smtp "golang.org/x/net/smtp"
// )

// func main() {

// 	email := xmail.NewEMail(`{"port":25}`)
// 	email.From = `user@gmail.com`
// 	email.Host = `s3.smtp.com`
// 	email.Port = int(25) // [587 NTLM AUTH] [465，994]
// 	email.Username = `user`
// 	email.Secure = `` // SSL，TSL
// 	email.Password = `pass`
// 	authType := `LOGIN`
// 	switch authType {
// 	case ``:
// 		email.Auth = nil
// 	case `LOGIN`:
// 		email.Auth = xmail.LoginAuth(email.Username, email.Password)
// 	case `CRAM-MD5`:
// 		email.Auth = smtp.CRAMMD5Auth(email.Username, email.Password)
// 	case `PLAIN`:
// 		email.Auth = smtp.PlainAuth(email.Identity, email.Username, email.Password, email.Host)
// 	case `NTLM`:
// 		email.Auth = xmail.NTLMAuth(email.Host, email.Username, email.Password, NTLMVersion1)
// 	default:
// 		email.Auth = smtp.PlainAuth(email.Identity, email.Username, email.Password, email.Host)
// 	}

// 	email.To = []string{`youremail@gmail.com`}
// 	email.Subject = `send mail success lib farmerx test test 2021... `
// 	email.Text = "Quero testar：\r\n   isto é 您好，附件中是您订阅的报表，请注意查收。"
// 	//email.AttachFile(reportFile)
// 	if err := email.Send(); err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	println("success send")
// }

// package main

// import (
// 	"crypto/tls"
// 	"fmt"
// 	"log"
// 	"net"
// 	"net/mail"
// 	"net/smtp"
// )

// // StartTLS Email Example
// func main() {
// 	from := mail.Address{"", "user@gmail.com"}
// 	to := mail.Address{"", "youremail@gmail.com"}
// 	subj := "This is the email subject test2..."
// 	body := "This is an example body.\n With two lines."

// 	// Setup headers
// 	headers := make(map[string]string)
// 	headers["From"] = from.String()
// 	headers["To"] = to.String()
// 	headers["Subject"] = subj

// 	// Setup message
// 	message := ""
// 	for k, v := range headers {
// 		message += fmt.Sprintf("%s: %s\r\n", k, v)
// 	}
// 	message += "\r\n" + body

// 	// Connect to the SMTP Server
// 	servername := "s3.smtp.com:25"

// 	host, _, _ := net.SplitHostPort(servername)

// 	auth := smtp.PlainAuth("", "user", "pass", host)

// 	// TLS config
// 	tlsconfig := &tls.Config{
// 		InsecureSkipVerify: true,
// 		ServerName:         host,
// 	}

// 	c, err := smtp.Dial(servername)
// 	if err != nil {
// 		log.Println("Dial:", err)
// 		return
// 	}

// 	c.StartTLS(tlsconfig)

// 	// Auth
// 	if err = c.Auth(auth); err != nil {
// 		log.Println("Auth:", err)
// 		return
// 	}

// 	// To && From
// 	if err = c.Mail(from.Address); err != nil {
// 		log.Println("Mail:", err)
// 		return
// 	}

// 	if err = c.Rcpt(to.Address); err != nil {
// 		log.Println("Rcpt:", err)
// 		return
// 	}

// 	// Data
// 	w, err := c.Data()
// 	if err != nil {
// 		log.Println("Data:", err)
// 		return
// 	}

// 	_, err = w.Write([]byte(message))
// 	if err != nil {
// 		log.Println("Write:", err)
// 		return
// 	}

// 	err = w.Close()
// 	if err != nil {
// 		log.Println("Close:", err)
// 		return
// 	}
// 	c.Quit()
// }

// package main

// import (
// 	"crypto/tls"
// 	"fmt"
// 	"log"
// 	"net/smtp"
// 	"strings"
// )

// func main() {
// 	sendMail("user", []string{"yourmail@gmail.com"}, "Hello2 World2...")
// 	fmt.Println("Email successfully sent!")
// }

// func checkErr(err error) {
// 	if err != nil {
// 		log.Println("ERROR: " + err.Error())
// 	}
// }

// func sendMail(from string, to []string, body string) {
// 	//configuração
// 	servername := "s3.smtp.com:25"             //servidor SMTP e PORTA
// 	host := "s3.smtp.com"                      //host
// 	pass := "pass"                               //senha
// 	auth := smtp.PlainAuth("", from, pass, host) //autenticação
// 	tlsConfig := &tls.Config{
// 		InsecureSkipVerify: true,
// 		ServerName:         host,
// 	}
// 	toHeader := strings.Join(to, ",")
// 	msg := "From: " + from + "\n" + "To: " + toHeader + "\n" + "Subject: Hello World\n\n" + body

// 	//conecta com o servidor SMTP
// 	conn, err := tls.Dial("tcp", servername, tlsConfig)
// 	checkErr(err)

// 	//retorna client SMTP
// 	c, err := smtp.NewClient(conn, host)
// 	checkErr(err)

// 	//autentica
// 	err = c.Auth(auth)
// 	checkErr(err)

// 	//adiciona remetente
// 	err = c.Mail(from)
// 	checkErr(err)

// 	//adiciona destinatários
// 	for _, addr := range to {
// 		err = c.Rcpt(addr)
// 		checkErr(err)
// 	}

// 	//prepara corpo do email
// 	w, err := c.Data()
// 	checkErr(err)

// 	//adiciona corpo do e-mail
// 	_, err = w.Write([]byte(msg))
// 	checkErr(err)

// 	//fecha corpo do e-mail
// 	err = w.Close()
// 	checkErr(err)

// 	//encerra conexão
// 	c.Quit()
// }

// package main

// import (
// 	"log"
// 	"net/smtp"
// )

// func main() {
// 	// Choose auth method and set it up
// 	auth := smtp.PlainAuth("", "user", "pass", "s3.smtp.com")

// 	to := []string{"yourmail@gmail.com"}
// 	msg := []byte("To: teste.. envio send email\r\n" +
// 		"Subject: Why are you not using Mailtrap yet?\r\n" +
// 		"\r\n" +
// 		"Jff... eng... Here’s the space for our great sales pitch\r\n")
// 	err := smtp.SendMail("s3.smtp.com:25", auth, "oper@yourmail.com.br", to, msg)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// package main

// import (
// 	"log"

// 	gse "github.com/jeffotoni/gosendmail"
// )

// // $ export EMAIL_HOST="smtp.<YOUR-SERVER>.com"
// // $ export EMAIL_USERNAME="<USER>"
// // $ export EMAIL_PASSWORD="<PASSWORD>"
// // $ export EMAIL_PORT="587"
// // $ export EMAIL_INSECURE=yes
// // $ export EMAIL_FROM="<your-from>"
// // $ go mod init <your-pkg>
// // $ go mod tidy
// // $ go run main

// func main() {
// 	email := gse.New()
// 	if email == nil {
// 		log.Println("Error New() check the required fields: EMAIL_HOST,EMAiL_USERNAME,EMAIL_PASSWORD,EMAIL_PORT")
// 		return
// 	}
// 	m := gse.NewMessage("Vamos testar pkg sendmail", "Pkg feito em Go para enviar email SMTP")
// 	m.To = []string{"yourmail@gmail.com"}
// 	// m.CC = []string{""}
// 	// m.BCC = []string{"bc@gmail.com"}
// 	// m.AttachFile("./love.png")
// 	if err := email.Send(m); err != nil {
// 		log.Println("Error when sending:", err.Error())
// 		return
// 	}
// 	log.Println("Sent with success")
// }
