/*
* Golang config with json
*
* @package     main
* @author      @jeffotoni
* @size        2017
 */

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Our config that will be given merge,
// it that the user will access
type Config struct {
	Ping        string
	Title       string
	Description string

	Database database
}

// Defining the database structure
// that is a vector in our json
type database struct {
	Ping      string
	Server    string
	Name      string
	Port      string
	ConnecMax string
	Enabled   string
}

// Defining our json struct,
// it that we will apply the Unmarshal
type confjson struct {
	Ping        string
	Title       string
	Description string

	Database []string `json:"Database"`
}

var (

	//Object of type Config that
	//we will use to access our struct
	newFile  *os.File
	fileInfo os.FileInfo
	err      error
	DirConf  = "config"
	NameConf = "config.json"
	pconf    = DirConf + "/" + NameConf
	returns  string
)

func main() {

	// Starting our json object
	jconf := GetConfig()

	// all object
	fmt.Println(jconf)

	// // start json
	fmt.Println(jconf.Ping)
	fmt.Println(jconf.Title)
	fmt.Println(jconf.Description)

	// // Only one field
	fmt.Println(jconf.Database)
	fmt.Println(jconf.Database.Ping)
	fmt.Println(jconf.Database.Server)
	fmt.Println(jconf.Database.Name)

}

func GetConfig() Config {

	// Testing file exist, if there is no create
	TestStructConfig()

	conf, errj := DecodeFileJson()
	if errj != nil {

		log.Fatal("Decode File error", errj)
	}

	return conf

}

// Reads info from config file
func DecodeFileJson() (Config, error) {

	// JUST confjson
	var config confjson

	// Merge Config
	var C Config

	// Encoding our json string
	errj := JsonDecode(&config)

	C.Ping = config.Ping
	C.Title = config.Title
	C.Description = config.Description

	if len(config.Database) > 0 {

		// to convert
		newVet := make(map[string]string)

		for key, _ := range config.Database {

			vet1 := strings.Split(config.Database[key], "=")
			newVet[strings.ToLower(vet1[0])] = vet1[1]

			switch strings.ToLower(vet1[0]) {

			case "ping":
				C.Database.Ping = newVet["ping"]
			case "server":
				C.Database.Server = newVet["server"]

			case "name":
				C.Database.Name = newVet["name"]

			default:

			}
		}
	}

	return C, errj
}

func JsonDecode(CJson interface{}) error {

	// read file config.json
	content, err := ioutil.ReadFile(pconf)
	if err != nil {

		// CREATE FILE HERE
		fmt.Print(err)
	}

	configJson := []byte(string(content))

	err = json.Unmarshal(configJson, &CJson)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	return err
}

// TestStructConfig responsible for testing and creating the
// configuration files of our webservice
// Testing if there is config / config.gcfg if
// it does not find in the current path,
// create directory and its default configuration file.
func TestStructConfig() {

	// Taking the default content
	// from our config
	content := []byte(TxtConfig)

	//Our config patch
	fileInfo, err = os.Stat(pconf)

	if fileInfo == nil {

		// NOT EXIST
		// create
		os.MkdirAll(DirConf, 0777)
		newFile, err = os.Create(pconf)

		if err != nil {

			log.Fatal("Need to create the web service config!", err)
			os.Exit(1)

		} else {

			fmt.Println("Path config created successfully..")
		}

		_, err := newFile.Write(content)

		if err != nil {

			log.Fatal("Need to save the contents in the config!", err)
			os.Exit(1)
		} else {

			fmt.Println("Config.gcfg content saved successfully..")
		}

		newFile.Close()
	}
}

// Our config case has no structure created
// the system will dynamically create
// in the current directory in config/config.gcfg
var TxtConfig = `{
"ping":"ok",
"title":"Own example",
"description":"Testing our json config",

"database":[
		"ping=ok",
		"server=192.168.12.14", 
		"name=Redis",
		"port=3459",
		"connectMax=10000",
		"enabled=true"
		]
}
		`
