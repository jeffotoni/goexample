// Go in action
// @jeffotoni
// 2019-03-19

package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "path"
    "time"
)

func main() {
    log.Printf("Running on :8080 ...")
    err := http.ListenAndServe("127.0.0.1:8080",
        http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            filePath := "/tmp/test.avi"
            file, err := os.Open(filePath)
            if err != nil {
                log.Printf("%s not found\n", filePath)
                fmt.Fprint(w, "<html><body style='font-size:100px'>Not Found</body></html>")
                return
            }

            defer file.Close()
            fileStat, err := os.Stat(filePath)
            if err != nil {
                fmt.Println(err)
            }
            fmt.Printf("serve %s\n", filePath)
            _, filename := path.Split(filePath)
            t := fileStat.ModTime()
            fmt.Printf("time %+v\n", t)

            http.ServeContent(w, r, filename, time.Time{}, file)
            //strings.NewReader("I am some content.\n")
        }))
    log.Fatal(err)
}
