package main

import (
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type sendS3 struct {
	Path string
}

func sendFileDO(info sendS3, wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	println(info.Path)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	err := filepath.Walk("/tmp/teste1", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		wg.Add(1)
		go sendFileDO(
			sendS3{
				Path: path,
			},
			&wg,
		)
		return nil
	})
	if err != nil {
		log.Println(err)
	}

	wg.Wait()
}
