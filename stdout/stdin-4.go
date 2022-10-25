// Go in Action
// @jeffotoni
// 2019-03-25

package main

import (
        "bufio"
        "fmt"
        "io/ioutil"
        "os"
)

func main() {
        fi, _ := os.Stdin.Stat() // get the FileInfo struct describing the standard input.

        if (fi.Mode() & os.ModeCharDevice) == 0 {
                fmt.Println("data is from pipe")
                // do things for data from pipe

                bytes, _ := ioutil.ReadAll(os.Stdin)
                str := string(bytes)
                fmt.Println(str)

        } else {
                fmt.Println("data is from terminal")
                // do things from data from terminal

                ConsoleReader := bufio.NewReader(os.Stdin)
                fmt.Println("Enter your name : ")

                input, err := ConsoleReader.ReadString('\n')

                if err != nil {
                        fmt.Println(err)
                        os.Exit(1)
                }

                fmt.Println("Your name is : ", input)

        }

}
