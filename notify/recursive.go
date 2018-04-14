/*
* Golang recursive dir
*
* @package     main
* @author      @jeffotoni
* @size        2018
 */

package main

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

var visited int

func main() {

	dir := "/home/netcatc/Downloads"

	subDirToSkip := "skip" // dir/to/walk/skip

	// criando func watcher
	doneChan := make(chan bool)

	for {

		go func(doneChan chan bool, dir, subDirToSkip string) {

			defer func() {
				doneChan <- true
			}()

			visited++
			err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

				if err != nil {
					log.Printf("prevent panic by handling failure accessing a path %q: %v\n", dir, err)
					return err
				}

				if info.IsDir() && info.Name() == subDirToSkip {

					log.Printf("skipping a dir without errors: %+v \n", info.Name())
					return filepath.SkipDir

				}

				log.Printf("visited file: %q\n", path)
				return nil
			})

			if err != nil {
				log.Printf("error walking the path %q: %v\n", dir, err)
			}

			log.Printf("wait 5 secs, visitou %d\n", visited)
			time.Sleep(5 * time.Second)

		}(doneChan, dir, subDirToSkip)

		<-doneChan
	}
}
