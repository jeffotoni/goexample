// Go in action
// @jeffotoni
// 2019-01-11

package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

var KEY_PEM = `-----BEGIN RSA PRIVATE KEY-----

-----END RSA PRIVATE KEY-----`

func PublicKeyFile() ssh.AuthMethod {
	buffer := []byte(KEY_PEM)
	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}
	return ssh.PublicKeys(key)
}

func main() {

	hostname := "server.com"
	port := "22"
	username := "centos"
	//password := ""

	// SSH client config
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			//ssh.Password(password),
			PublicKeyFile(),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to host
	client, err := ssh.Dial("tcp", hostname+":"+port, config)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Create sesssion
	sess, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer sess.Close()

	// StdinPipe for commands
	stdin, err := sess.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	// stdoutIn, err := sess.StdoutPipe()
	// if err != nil {
	// 	log.Fatal(err)
	// }

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
		log.Fatal(err)
	}

	// send the commands
	commands := []string{
		"echo '################ executando ###############'",
		// "", //coloca aqui os comandos
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

	// Uncomment to store in variable
	//fmt.Println(b.String())
}
