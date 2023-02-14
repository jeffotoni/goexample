/*
* Golang config with toml
*
* @package     main
* @author      @jeffotoni
* @size        2017
 */

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// Info from config file
type Config struct {
	Ping        string
	Title       string
	Description string
	Database    database
	Clients     clients
}

// Info from database file
type database struct {
	Ping   string
	Server string
	//Connection_max int64
	ConnMax int `toml:"conect_max"`
	Enabled bool
	Ports   []int
}

// info from clients file
type clients struct {
	Ping  string
	Data  [][]interface{}
	Hosts []string
}

var (

	//Object of type Config that
	//we will use to access our struct
	newFile  *os.File
	fileInfo os.FileInfo
	err      error
	DirConf  = "config"
	NameConf = "config.toml"
	pconf    = DirConf + "/" + NameConf
	returns  string
)

func main() {

	toml := GetConfig()

	// config data
	fmt.Println(toml.Title)
	fmt.Println(toml.Description)

	// database data
	fmt.Println(toml.Database.Server)
	fmt.Println(toml.Database.ConnMax)
	fmt.Println(toml.Database.Enabled)

	// list ports
	fmt.Println(toml.Database.Ports[0])
	fmt.Println(toml.Database.Ports[1])

	//
	fmt.Println(toml.Clients.Data)
	fmt.Println(toml.Clients.Data[0])
	fmt.Println(toml.Clients.Data[0][0])

	//
	fmt.Println(toml.Clients.Data)
	fmt.Println(toml.Clients.Data[1])
	fmt.Println(toml.Clients.Data[1][1])

	//
	fmt.Println(toml.Clients.Hosts)
	fmt.Println(toml.Clients.Hosts[0])
	fmt.Println(toml.Clients.Hosts[1])

	//connString := fmt.Sprintf("%s:%s", toml.Database.Server, toml.Database.Ports)
	//fmt.Println(toml.Database.Server)

}

// Reads info from config file
func GetConfig() Config {

	TestStructConfig()

	var config Config

	if _, err := toml.DecodeFile(pconf, &config); err != nil {
		log.Fatal(err)
	}

	return config
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

// Using TestConfig, Testing to see if Section
// Config is instantiated correctly.
func TestConfig() string {

	toml := GetConfig()

	msgerr := "Error reading the config file!"

	if toml.Ping == "ok" {

		returns = "ok"

	} else if toml.Clients.Ping == "ok" {

		returns = "ok"

	} else if toml.Database.Ping == "ok" {

		returns = "ok"

	} else {

		returns = "error"
	}

	if returns != "ok" {

		fmt.Println(msgerr)
		os.Exit(1)
	}

	return returns
}

// Our config case has no structure created
// the system will dynamically create
// in the current directory in config/config.gcfg
var TxtConfig = `# This is a TOML document example.

ping = "ok"
title = "TOML Example @jeffotoni"
description = "Testing toml config now."

[database]
ping = "ok"
server = "192.168.12.14"
ports = [ 8801, 8802, 8803 ]
conect_max = 10000
enabled = true

[clients]
ping = "ok"
data = [ ["amazon", "google"], [1, 2] ]
hosts = [
  "aws",
  "cloud"
]

`
