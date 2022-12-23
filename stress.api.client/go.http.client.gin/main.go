package main

import (
  "bytes"
  "context"
  "fmt"
  "io"
  "io/ioutil"
  "log"
  "net/http"
  "os"
  "strconv"
  "strings"
  "time"

  "github.com/gin-gonic/gin"
)

var client = &http.Client{Transport: &http.Transport{
  DisableKeepAlives: true,
  //MaxIdleConns:      10,
}}

var (
  Domain = os.Getenv("DOMAIN")
  key    = "key.pem"
  cert   = "cert.pem"

  cfsslkey  = "./cfssl/server-key.pem"
  cfsslcert = "./cfssl/server.pem"

  serverCrt = "./certs/server.crt"
  serverKey = "./certs/server.key"
)

func main() {
  //r := gin.Default()
  r := gin.New()

  r.Use(func(c *gin.Context) {})
  r.NoRoute(func(c *gin.Context) {
    c.AbortWithStatus(http.StatusNotFound)
  })

  // Ping handler
  r.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "pong")
  })

  r.POST("/v1/client/post", func(c *gin.Context) {
    c.Header("Content-Type", "application/json")
    c.Header("Engine", "Go")
    c.Header("Location", "/v1/client/post")
    c.Header("Date", time.Now().Format("2006-01-02T15:04:05.000Z"))

    rawData := c.Request.Body
    defer rawData.Close()

    body, err := ioutil.ReadAll(rawData)
    if err != nil {
      c.String(http.StatusBadRequest, "Error reading request body: "+err.Error())
      return
    }

    // start := time.Now()
    body, code, err := AdapterConnect("post", body)
    if err != nil {
      log.Println("Error Server connect:", err, " code:", code)
      c.String(http.StatusBadRequest, "Error reading request body: "+err.Error())
      return
    }
    // end := time.Now().Sub(start)
    // log.Println("Service Adapter [POST] timeTotal:", end.String())
    length := strconv.Itoa(len(body))
    c.Header("Content-Length", length)
    c.String(http.StatusOK, string(body))
  })

  // log.Fatal(r.RunTLS("0.0.0.0:443", "server.pem", "server-key.pem"))
  log.Fatal(r.Run())
}

func AdapterConnect(method string, bodyPost []byte) (body []byte, code int, err error) {

  // var client2 = &http.Client{
  //  Transport: &http.Transport{
  //    DisableKeepAlives: true,
  //    MaxIdleConns:      2,
  //    // MaxIdleConnsPerHost: 10,
  //    // MaxConnsPerHost:     10,
  //    // IdleConnTimeout: time.Duration(time.Millisecond * 10),
  //  }}

  ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
  defer cancel()

  if len(Domain) == 0 {
    Domain = "http://127.0.0.1:3000"
  }

  // send POST

  var Url string = Domain + "/v1/customer"
  var req = &http.Request{}

  // http2.ConfigureTransport(client.Transport)

  if strings.ToLower(method) == "get" {
    Url = Url + "/get"
    req, err = http.NewRequestWithContext(ctx, "GET", Url, nil)
    if err != nil {
      fmt.Printf("Error %s", err)
      return
    }
  } else if strings.ToLower(method) == "post" {
    bodysend := bytes.NewBuffer(bodyPost)
    Url = Url + "/post"
    req, err = http.NewRequestWithContext(ctx, "POST", Url, bodysend)
    if err != nil {
      fmt.Printf("Error %s", err)
      return
    }
  }

  req.Header.Set("Content-Type", "application/json")
  resp, err := client.Do(req)
  if err != nil {
    fmt.Printf("Error %s", err)
    return
  }
  defer resp.Body.Close()
  code = resp.StatusCode

  body, err = io.ReadAll(resp.Body)
  if err != nil {
    fmt.Printf("Error %s", err)
    return
  }

  return
}
