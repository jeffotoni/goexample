package main

import (
  "bytes"
  "io"
  "mime/multipart"
  "net/http"
  "os"
  "path"
  "path/filepath"
)

func main() {
  fileDir, _ := os.Getwd()
  fileName := "jeff.png"
  filePath := path.Join(fileDir, fileName)

  file, _ := os.Open(filePath)
  defer file.Close()

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
  io.Copy(part, file)
  writer.Close()

  r, _ := http.NewRequest("POST", "http://localhost:8282", body)
  r.Header.Add("Content-Type", writer.FormDataContentType())
  client := &http.Client{}
  client.Do(r)
}
