/*
* Golang config with gcfg
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

	gcfg "gopkg.in/gcfg.v1"
)

// My type struct to be able to recover
// in the .gcfg config
type Config struct {
	Section section
}

type section struct {
	Domain   string
	Port     string
	Database string
	User     string
	Password string
	Ping     string
}

var (

	//Object of type Config that
	//we will use to access our struct
	cfg      Config
	newFile  *os.File
	fileInfo os.FileInfo
	err      error
	DirConf  = "config"
	NameConf = "config.gcfg"
	pconf    = DirConf + "/" + NameConf
	returns  string
)

func main() {

	gcfg := GetConfig()

	fmt.Println("Testing config Ping: ", gcfg.Section.Ping)
	fmt.Println("Testing config Domain: ", gcfg.Section.Domain)
	fmt.Println("Testing config Port: ", gcfg.Section.Port)
	fmt.Println("Testing config Database: ", gcfg.Section.Database)
	fmt.Println("Testing config User: ", gcfg.Section.User)
	fmt.Println("Testing config Password: ", gcfg.Section.Password)
}

// Method GetConfig responsible for creating our instance
// to access the objects of our config file
// singleton object Config
func GetConfig() *Config {

	if cfg.Section.Ping == "ok" {

		return &cfg

	} else {

		// Testing if there is config / config.gcfg if
		// it does not find in the current path,
		// create directory and its default configuration file.
		TestStructConfig()

		// This is when our executable opens the config and most
		// importantly our path is not absolute so it will look for
		// the file in the path where our server is running, ie the
		// system has to create the config if it does not exist,
		// since it is Something that can be in any directory
		// to stay more flexible, stay where it runs.
		err = gcfg.ReadFileInto(&cfg, pconf)

		// Making a simple test with our instance,
		// if it finds Instance Section will be all ok.
		if cfg.Section.Ping != "ok" {

			fmt.Println("Error reading file config.gcfg ", err)
			os.Exit(1)
		}

		return &cfg
	}
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

	cfg := GetConfig()

	msgerr := "Error reading the config file!"

	if cfg.Section.Ping == "ok" {

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
var TxtConfig = `#
# @autor jeffotoni
# @package config
# @date 17/04/2017
# This is a configuration config system
# Lines database postgres config
#

[section]
; config datadbase postgresql
ping		= ok
domain 		= localhost
port 		= 5432
database	= ukkobox # database name
user		= ukkobox # database user
password	= pass123 # password

`
