package main

import (
    "bytes"
    "crypto/tls"
    "errors"
    "fmt"
    "net"
    "net/smtp"
    "text/template"
)

type loginAuth struct {
    username, password string
}

func LoginAuth(username, password string) smtp.Auth {
    return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
    return "LOGIN", []byte(a.username), nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
    if more {
        switch string(fromServer) {
        case "Username:":
            return []byte(a.username), nil
        case "Password:":
            return []byte(a.password), nil
        default:
            return nil, errors.New("Unknown from server")
        }
    }
    return nil, nil
}

func main() {

    // Sender data.
    from := "O365 logging name"
    password := "O365 logging pasword"

    // Receiver email address.
    to := []string{
        "receiver email",
    }

    // smtp server configuration.
    smtpHost := "smtp.office365.com"
    smtpPort := "587"

    conn, err := net.Dial("tcp", "smtp.office365.com:587")
    if err != nil {
        println(err)
    }

    c, err := smtp.NewClient(conn, smtpHost)
    if err != nil {
        println(err)
    }

    tlsconfig := &tls.Config{
        ServerName: smtpHost,
    }

    if err = c.StartTLS(tlsconfig); err != nil {
        println(err)
    }

    auth := LoginAuth(from, password)

    if err = c.Auth(auth); err != nil {
        println(err)
    }

    t, _ := template.New("email").Parse(`<html>
                        <body>
                            <h3>Name:</h3><span>{{.Name}}</span><br/><br/>
                            <h3>Email:</h3><span>{{.Message}}</span><br/>
                        </body>
                        </html>`)

    var body bytes.Buffer

    mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
    body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

    t.Execute(&body, struct {
        Name    string
        Message string
    }{
        Name:    "Hasan Yousef",
        Message: "This is a test message in a HTML template",
    })

    // Sending email.
    err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Email Sent!")
}
