/*
* Golang in action
*
* @package     main
* @author      @jeffotoni
* @size        2018
 */

package main

import (
	"log"
	"os"
	"time"
)

func watchFile(filePath string) error {
	initialStat, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	for {
		stat, err := os.Stat(filePath)
		if err != nil {
			return err
		}

		if stat.Size() != initialStat.Size() || stat.ModTime() != initialStat.ModTime() {
			break
		}

		time.Sleep(1 * time.Second)
	}

	return nil
}

func main() {

	// criando func watcher
	doneChan := make(chan bool)

	for {
		go func(doneChan chan bool) {
			defer func() {
				doneChan <- true
			}()

			err := watchFile("/home/netcatc/Downloads")
			if err != nil {

				log.Println(err)
			}

			log.Println("File has been changed")

		}(doneChan)

		<-doneChan
	}
}
