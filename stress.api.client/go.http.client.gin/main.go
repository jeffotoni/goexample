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

func init() {
	if len(Domain) == 0 {
		Domain = "http://127.0.0.1:3000"
	}
}

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

	r.GET("/v1/client/get", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Engine", "Go")
		c.Header("Location", "/v1/client/post")
		c.Header("Date", time.Now().Format("2006-01-02T15:04:05.000Z"))
		//c.String(200, `[{"createdAt":"2022-11-04T19:10:17.305Z","name":"BillyStoltenberg","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1164.jpg","id":"1"},{"createdAt":"2022-11-05T09:01:38.207Z","name":"JodiKertzmann","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/502.jpg","id":"2"},{"createdAt":"2022-11-05T15:36:31.390Z","name":"AngelKuhlmanIV","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/75.jpg","id":"3"},{"createdAt":"2022-11-04T19:23:42.344Z","name":"KellyNolan","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/787.jpg","id":"4"},{"createdAt":"2022-11-05T06:33:55.777Z","name":"JeremySchneider","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1153.jpg","id":"5"},{"createdAt":"2022-11-05T07:29:47.957Z","name":"EvanDuBuque","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/99.jpg","id":"6"},{"createdAt":"2022-12-01T10:43:35.133Z","name":"DeloresDoyle","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/331.jpg","id":"7"},{"createdAt":"2022-12-01T22:53:22.031Z","name":"Ms.JeremyBruen","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/383.jpg","id":"8"},{"createdAt":"2022-12-02T00:59:05.444Z","name":"JoannRitchie","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/130.jpg","id":"9"},{"createdAt":"2022-12-01T21:37:21.680Z","name":"Mrs.StevenCummerata","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/771.jpg","id":"10"},{"createdAt":"2022-12-02T09:01:39.388Z","name":"HattiePfeffer","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1005.jpg","id":"11"},{"createdAt":"2022-12-01T10:11:38.831Z","name":"EdithMacejkovic","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/882.jpg","id":"12"},{"createdAt":"2022-12-02T08:34:12.087Z","name":"BettyBlock","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/216.jpg","id":"13"},{"createdAt":"2022-12-02T06:56:11.901Z","name":"ArmandoBecker","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/822.jpg","id":"14"},{"createdAt":"2022-12-01T17:22:36.489Z","name":"MargueriteWilliamson","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/205.jpg","id":"15"},{"createdAt":"2022-12-02T02:59:48.845Z","name":"LorenePaucek","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/31.jpg","id":"16"},{"createdAt":"2022-12-01T22:23:37.039Z","name":"TonyaHeidenreich","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/685.jpg","id":"17"},{"createdAt":"2022-12-01T15:55:35.929Z","name":"MissAntoniaReynolds","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/208.jpg","id":"18"},{"createdAt":"2022-12-01T14:15:27.068Z","name":"Ms.GlendaSchimmel","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/145.jpg","id":"19"},{"createdAt":"2022-12-01T17:29:53.880Z","name":"JodySchadenDDS","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/35.jpg","id":"20"}]`)
		//return
		// start := time.Now()

		//body, code, err := AdapterConnect("get", nil)

		//s1 := time.Now()
		body, code, err := AdapterConnect("get", nil)
		//fmt.Println("time:", time.Since(s1))
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	// send POST
	var Url string = Domain + "/v1/customer"
	var req = &http.Request{}

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
