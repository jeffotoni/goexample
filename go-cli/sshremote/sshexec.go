// Go in action
// @jeffotoni
// 2019-01-11

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"golang.org/x/crypto/ssh"
)

// const KEY_PEM = `-----BEGIN RSA PRIVATE KEY-----

// -----END RSA PRIVATE KEY-----`

func PublicKey(Key string) ssh.AuthMethod {
	buffer := []byte(Key)
	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}
	return ssh.PublicKeys(key)
}

type ConfHost struct {
	Host, User, Password, Key string
	Port                      int
}

func (c *ConfHost) Config() *ssh.ClientConfig {

	// SSH client config
	config := &ssh.ClientConfig{
		User: c.User,
		Auth: []ssh.AuthMethod{
			PublicKey(c.Key),
			//ssh.Password(c.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return config
}

func (cf *ConfHost) ConnectSession() (*ssh.Session, error) {

	config := cf.Config()

	// Connect to host
	connect, err := ssh.Dial("tcp", cf.Host+":"+strconv.Itoa(cf.Port), config)
	if err != nil {
		return nil, fmt.Errorf("Failed to dial: %s", err)
	}
	//defer connect.Close()
	session, err := connect.NewSession()
	if err != nil {
		return nil, fmt.Errorf("Failed to create session: %s", err)
	}

	return session, nil
}

func main() {

	// read from file, settings recipe
	cf := ConfHost{User: "xxxxx", Port: 22, Host: "xxxxx", Key: os.Getenv("KEY_AWS")}

	// connect ssh host
	sess, err := cf.ConnectSession()
	if err != nil {
		log.Println("session: ", err)
	}

	// StdinPipe for commands
	stdin, err := sess.StdinPipe()
	if err != nil {
		log.Println("session: ", err)
	}

	// Uncomment to store output in variable
	//var b bytes.Buffer
	//sess.Stdout = &b
	//sess.Stderr = &b

	// Enable system stdout
	// Comment these if you uncomment to store in variable
	sess.Stdout = os.Stdout
	sess.Stderr = os.Stderr

	// Start remote shell
	err = sess.Shell()
	if err != nil {
		log.Println("shell: ", err)
	}

	// send the commands
	commands := []string{
		"echo '################ Exec ###############'",
		"sudo -i",
		"apt update",
		"apt list --upgradable",
		"exit",
	}

	for _, cmd := range commands {
		_, _ = fmt.Fprintf(stdin, "%s\n", cmd)
	}

	// Wait for sess to finish
	err = sess.Wait()
	if err != nil {
		log.Fatal(err)
	}

	sess.Close()

	// Uncomment to store in variable
	//fmt.Println(b.String())
}
