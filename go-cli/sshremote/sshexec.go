// Go in action
// @jeffotoni
// 2019-01-11

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"golang.org/x/crypto/ssh"
	yaml "gopkg.in/yaml.v2"
)

type ConfigYaml struct {
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
	User    string `yaml:"user"`
	Port    int    `yaml:"port"`
	Type    string `yaml:"type"`
	File    string `yaml:"file"`
	Env     string `yaml:"env"`
}

// GetConfig Method responsible for struct
// of the yaml it returns the Config object
// then the conf conf config definition
// is made in yaml.Unmarshal and will receive &confg
func GetYaml(path_yaml string) *ConfigYaml {

	// Config structure instance
	//var conf = &ConfigYaml{}
	var conf *ConfigYaml

	var yamlByte []byte
	var err error

	// Reading file and loading in buffer byte
	if yamlByte, err = ioutil.ReadFile(path_yaml); err != nil {
		log.Println("Error: ", err)
	}

	// Doing magic yaml
	if err := yaml.Unmarshal(yamlByte, &conf); err != nil {
		log.Println("Error", err)
	}

	return conf
}

func PublicKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}
	return ssh.PublicKeys(key)
}

func PublicKey(Key string) ssh.AuthMethod {
	buffer := []byte(Key)
	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}
	return ssh.PublicKeys(key)
}

func (c *ConfigYaml) Config() *ssh.ClientConfig {
	if c.Type == "file" {
		config := &ssh.ClientConfig{
			User: c.User,
			Auth: []ssh.AuthMethod{
				PublicKeyFile(c.File),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
	} else {
		config := &ssh.ClientConfig{
			User: c.User,
			Auth: []ssh.AuthMethod{
				PublicKey(c.Key),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
	}
	return config
}

func (cf *ConfigYaml) ConnectSession() (*ssh.Session, error) {
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

	examplePtr := flag.String("file", "server-1.yaml", " Help:")
	flag.Parse()

	if len(os.Args) < 2 {
		flag.PrintDefaults()
		return
	}
	fileYaml := *examplePtr

	// get Yaml
	cf := GetYaml(fileYaml)

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
