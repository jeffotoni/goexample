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
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

var (

	//Object of type Config that
	//we will use to access our struct
	newFile  *os.File
	fileInfo os.FileInfo
	err      error
	DirConf  = "config"
	NameConf = "config.yaml"
	pconf    = DirConf + "/" + NameConf
	returns  string
)

type Config struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Server      string `yaml:"server"`
	Port        string `yaml:"port"`
}

// Method responsible for assembling struct from yaml it
// returns the Config object, so the conf conf config definition
// does not exist, now the method is of the Config type
// then the yaml.Unmarshal method will receive confg only without &conf
func (conf *Config) GetConf() *Config {

	TestStructConfig()

	var yamlByte []byte
	var err error

	if yamlByte, err = ioutil.ReadFile(pconf); err != nil {
		log.Println("Error: ", err)
	}

	// Unmarshal recebe o arquivo em formato de byte e atribui os valores passados aos campos da estrutura.
	// Se houver um erro ele exibe uma mensagem de erro na tela informando o erro
	if err := yaml.Unmarshal(yamlByte, conf); err != nil {
		log.Println("Error", err)
	}

	return conf
}

// GetConfig Method responsible for struct
// of the yaml it returns the Config object
// then the conf conf config definition
// is made in yaml.Unmarshal and will receive &confg
func GetConfig() *Config {

	// Config structure instance
	var conf *Config

	// Tests whether the file exists
	// otherwise it creates
	TestStructConfig()

	var yamlByte []byte
	var err error

	// Reading file and loading in buffer byte
	if yamlByte, err = ioutil.ReadFile(pconf); err != nil {
		log.Println("Error: ", err)
	}

	// Doing magic yaml
	if err := yaml.Unmarshal(yamlByte, &conf); err != nil {
		log.Println("Error", err)
	}

	return conf
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
var TxtConfig = `
# Config server

title: Test Yaml  here
description: Test description here...
server:  localhost
port:  8080

`

// func main
func main() {

	// I declare only 1 object
	// and make the instance
	yamlc := GetConfig()

	fmt.Println("Testing Title ", yamlc.Title)
	fmt.Println("Testing Server ", yamlc.Server)

	// Another form of istanciar the object,
	// one declares and I can pass parameters to the structure
	yamlC := Config{}
	yamlcc := yamlC.GetConf()

	fmt.Println("Testing Descr. ", yamlcc.Description)
	fmt.Println("Testing Port ", yamlcc.Port)

}
