// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"

	"github.com/fasthttp/websocket"
)

var addr = flag.String("addr", "localhost:3000", "http service address")
var mutex sync.Mutex

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Uso: ./client.stress [threads]")
	}
	tmp1 := os.Args[1]

	threads, _ := strconv.Atoi(tmp1)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	finish := make(chan bool, 1)

	var j int
	for i := 0; i < threads; i++ {
		go func() {
			mutex.Lock()
			j++
			mutex.Unlock()

			client := strconv.Itoa(j)
			u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws/user_" + client}
			log.Printf("connecting to %s", u.String())

			c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
			if err != nil {
				log.Fatal("dial:", err)
			}
			defer c.Close()

			done := make(chan struct{})

			go func() {
				defer close(done)
				for {
					_, message, err := c.ReadMessage()
					if err != nil {
						log.Println("read:", err)
						return
					}
					log.Printf("recv: %s", message)
				}
			}()

			//err = c.WriteMessage(websocket.TextMessage, []byte(`{"name":"jeffotoni", "code":"x39399393939"}`))
			err = c.WriteMessage(websocket.TextMessage, []byte(`0987654321`))
			if err != nil {
				log.Println("write:", err)
				return
			}

			for {
				select {
				case <-done:
					return
				case <-interrupt:
					log.Println("interrupt")
					// Cleanly close the connection by sending a close message and then
					// waiting (with timeout) for the server to close the connection.
					err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
					if err != nil {
						log.Println("write close:", err)
						return
					}
					select {
					case <-done:
					case <-time.After(time.Millisecond):
					}
					return
				}
			}
		}()
	}

	go func() {
		sig := <-interrupt
		fmt.Println()
		fmt.Println(sig)
		finish <- true
	}()

	fmt.Println("awaiting signal")
	<-finish
	fmt.Println("exiting")

}
