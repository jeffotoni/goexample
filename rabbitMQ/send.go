/**
 * @AUTOR @jeffotoni
 * @DATE 05/06/2018
 * @DESCRIPTION Send rabbitmq
 */

package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"strconv"
	"time"
)

// Global
// Connect Rabbitmq
var (
	Conn    *amqp.Connection
	C2Error error
)

// Connect Rabbitmq
// Singleton
func isConnect(Conn *amqp.Connection) (*amqp.Connection, error) {

	if Conn == nil {

		Conn, C2Error = amqp.Dial("amqp://guest:guest@192.0.0.4:5672/")
	}

	return Conn, C2Error
}

// main
func main() {

	// contador
	var i int

	// loop
	for {

		// convertendo int to string
		i++
		str := strconv.Itoa(i)
		////////////////////////

		// test connec exist
		// se nao connect
		if Conn == nil {

			fmt.Println("main Connect first")
			time.Sleep(time.Second * 1)
			Conn, C2Error = isConnect(Conn)

		} else {

			fmt.Println("ok ja esta connectado utilizando Conn existe!!!")
			time.Sleep(time.Second * 1)
		}

		// se ocorreu algum erro
		// tente connectar novamente
		if Conn == nil || C2Error != nil {

			fmt.Println("Tenta Novamente, Error ao connect: " + str)
			time.Sleep(time.Second * 2)
			continue
		}

		defer Conn.Close()

		ch, err := Conn.Channel()

		// testando error
		if err != nil {

			fmt.Println("Error connecao caiu ou channel deu algum problema, reiniciar!")
			time.Sleep(time.Second * 1)
			continue
		}

		defer ch.Close()

		q, err := ch.QueueDeclare(
			"hello", // name
			false,   // durable
			false,   // delete when unused
			false,   // exclusive
			false,   // no-wait
			nil,     // arguments
		)

		if err != nil {

			fmt.Println("Error connecao caiu, ou channel deu algum problema reiniciar!")
			time.Sleep(time.Second * 1)
			continue
		}

		body := "hello " + str

		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})

		log.Printf(" [x] Sent %s", body)

		if err != nil {

			fmt.Println("Error connecao caiu, ou channel deu algum problema reiniciar!")
			time.Sleep(time.Second * 1)
			continue
		}
	}
}
