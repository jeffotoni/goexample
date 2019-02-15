// Go in action
// @jeffotoni
// 2019-01-11

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

///////////////////////////////////////////
/// paste your key here
var KEY_PEM = `-----BEGIN RSA PRIVATE KEY-----
-----END RSA PRIVATE KEY-----`

///////////////////////////////////////

func main() {

	var hostname, port, username string
	flag.StringVar(&hostname, "host", "server.com", "host to server")
	flag.StringVar(&port, "port", "22", "port to server")
	flag.StringVar(&username, "user", "ubuntu", "user to server")
	flag.Parse()

	////
	fmt.Println("connecting to your server: " + hostname + " port: " + port)
	///////////////////

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

	// Enable system stdout
	// Comment these if you uncomment to store in variable
	sess.Stdout = os.Stdout
	sess.Stderr = os.Stderr

	// Start remote shell
	err = sess.Shell()
	if err != nil {
		log.Fatal(err)
	}

	/////////////////////////////////////////////////////////
	/// send the commands here
	//
	commands := []string{
		"echo '################ Exec ###############'",
		// "", // here commands
		"ls -lh",
		"pwd",
		"netstat -plant",
		"exit",
	}
	/////////////////////////////////////////////////////////////

	// exec comands..
	for _, cmd := range commands {
		_, _ = fmt.Fprintf(stdin, "%s\n", cmd)
	}

	// Wait for sess to finish
	err = sess.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

// public.. private
func PublicKeyFile() ssh.AuthMethod {
	buffer := []byte(KEY_PEM)
	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}
	return ssh.PublicKeys(key)
}
