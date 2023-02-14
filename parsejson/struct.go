/*
* Golang reflect update struct to tag
*
* @package     main
* @author      @jeffotoni
* @size        2017
 */

package main

import (
	"errors"
	"fmt"
	"reflect"
)

// const tag
const tagName = "default"

// struct config
type Config struct {
	Domain  string `default:"s3go.gov"`
	Schema  string `default:"https"`
	Port    string `default:"9002"`
	Db      string `default:"postgresql"`
	Cluster string `default:"10"`
	Passwd  string `default:"x37c$%2"`
	Login   string `default:"postgres"`
}

func Default(struc interface{}) (err error) {

	structref := reflect.TypeOf(struc)

	if structref.Kind() != reflect.Ptr {

		err = errors.New("Not a pointer")
		return
	}

	elemField := structref.Elem()
	if elemField.Kind() != reflect.Struct {
		err = errors.New("Not a struct")
		return
	}

	//refValue := reflect.ValueOf(s).Elem()
	for i := 0; i < elemField.NumField(); i++ {

		field := elemField.Field(i)
		// value := refValue.Field(i)
		// kind := field.Type.Kind()
		tagVal := field.Tag.Get(tagName)

		// just string
		reflect.ValueOf(struc).Elem().Field(i).SetString(tagVal)
	}

	return
}

func main() {

	c := Config{}

	err := Default(&c)

	if err != nil {

		fmt.Println("error: ", err)
		return
	}

	fmt.Println("Domain: ", c.Domain)
	fmt.Println("Schema: ", c.Schema)
	fmt.Println("Port: ", c.Port)
	fmt.Println("Db: ", c.Db)
	fmt.Println("Cluster: ", c.Cluster)
	fmt.Println("Passwd: ", c.Passwd)
	fmt.Println("Login: ", c.Login)

}
