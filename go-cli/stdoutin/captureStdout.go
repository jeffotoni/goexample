package main

import (
        "bufio"
        "fmt"
        "os"
)

func main() {

        //fmt.Println("teste jefferson")
        RedirectOutput("10")

}

func RedirectOutput(id string) {
        oldStdout := os.Stdout
        readFile, writeFile, err := os.Pipe()
        if err != nil {
                fmt.Println("err os.Pipe", err)
                //return err
        }

        os.Stdout = writeFile

        go func() {
                scanner := bufio.NewScanner(readFile)
                for scanner.Scan() {
                        line := scanner.Text()

                        // Log the stdout line to my event logger
                        //event.Log(event.Event{Id: id, Msg: line})
                        //fmt.Println(line)
                        fmt.Sprintf("%s", line)
                }
        }()

        fmt.Printf("This will be logged to our event logger\n")

        // Reset the output again
        writeFile.Close()
        os.Stdout = oldStdout
}
