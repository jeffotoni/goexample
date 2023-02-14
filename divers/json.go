/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size		   2017
 */

package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var ff interface{}

	b := []byte(`{"Name":"s3 Google","Age":26,"Parents":["Igrind","Aws Amazon"]}`)
	err := json.Unmarshal(b, &ff)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("------------------ start ---------------")

	fmt.Println(ff)

	m := ff.(map[string]interface{})

	for k, v := range m {

		switch vv := v.(type) {

		case string:
			fmt.Println(k, "is string", vv)

		case int:
			fmt.Println(k, "is int", vv)

		case []interface{}:

			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)

			}

		default:
			fmt.Println(k, "is of a type I don't know how to handle")

		}
	}

	//var mm interface{}

	var m2 = map[string]interface{}{

		"Name": "Jeff Otoni",
		"Age":  32,
		"Parents": []interface{}{
			"Pull",
			"Aws",
		},
	}

	fmt.Println("----------------------------------------")
	fmt.Println(m2)

	for k, v := range m2 {

		switch vv := v.(type) {

		case string:
			fmt.Println(k, "is string", vv)

		case int:
			fmt.Println(k, "is int", vv)

		case []interface{}:

			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)

			}

		default:
			fmt.Println(k, "is of a type I don't know how to handle")

		}
	}

	fmt.Println("----------------- end -----------------")

}
