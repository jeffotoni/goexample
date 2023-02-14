package main

import (
  "bytes"
  "io"
  "io/ioutil"
  "log"
  "mime/multipart"
  "net/http"
  "net/textproto"
  "os"
  "path"
)

func main() {
  fileDir, _ := os.Getwd()
  fileName := "jeff.png"
  filePath := path.Join(fileDir, fileName)

  file, _ := os.Open(filePath)
  defer file.Close()

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  //part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))

  mh := make(textproto.MIMEHeader)
  mh.Set("Content-Type", "image/png")
  mh.Set("Content-Disposition", "form-data; name=\"file\"; filename=\"jeff.png\"")
  part, err := writer.CreatePart(mh)
  if nil != err {
    panic(err.Error())
    return
  }

  io.Copy(part, file)
  writer.Close()

  r, _ := http.NewRequest("POST", "http://localhost:8282", body)
  r.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, err := client.Do(r)
  if err != nil {
    log.Println(err.Error())
    return
  }
  defer resp.Body.Close()
  bodyMsg, _ := ioutil.ReadAll(resp.Body)
  log.Println(string(bodyMsg))

}
