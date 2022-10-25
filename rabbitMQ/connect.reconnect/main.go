// Back-End in Go server
// @jeffotoni

package main

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

var (
	once sync.Once
	mux  sync.Mutex
	conn *amqp.Connection
	err  error
)

func main() {

	println("testando reconnect RabbitMQ")
}

func Connect(host string) (*amqp.Connection, error) {
	once.Do(func() {
		if conn != nil {
			return
		}
		conn, err = amqp.Dial(host)
		if err != nil {
			return
		}
		go reconnect(host, conn)
		return
	})
	if conn == nil {
		return nil, errors.New("Error nao existe conexao estabelecida!")
	}
	return conn, nil
}

func reconnect(host string, conn *amqp.Connection) {
WATCH:

	conErr := <-conn.NotifyClose(make(chan *amqp.Error))
	if conErr != nil {
		var err error
		log.Println("CRITICAL: Connection dropped, reconnecting")
		for i := 1; i <= 7200; i++ {
			mux.Lock()
			conn, err = amqp.Dial(host)
			mux.Unlock()

			if err == nil {
				log.Println("INFO: Reconnected")
				goto WATCH
			}
			time.Sleep(600 * time.Millisecond)
		}
		log.Println("CRITICAL: Failed to reconnect:", err)
	} else {
		log.Println("INFO: Connection dropped normally, will not reconnect")
	}
}
