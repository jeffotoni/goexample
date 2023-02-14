/*
* Golang notify
*
* @package     main
* @author      @jeffotoni
* @size        2018
 */

package main

import (
	"fmt"
	"github.com/howeyc/fsnotify"
	"log"
)

func main() {

	fmt.Println("autocomplete..")
	log.Println("log....")

	// via syscall sistema operacional
	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		log.Println("opa")
		log.Fatal(err)
	}

	done := make(chan bool)

	// Process events
	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				log.Println("event:", ev)
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Watch("/home/netcatc")
	if err != nil {
		log.Fatal(err)
	}

	err = watcher.Watch("/home/netcatc")
	if err != nil {
		log.Fatal(err)
	}
	// Hang so program doesn't exit
	<-done

	/* ... do stuff ... */
	watcher.Close()
}
