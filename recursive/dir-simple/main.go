package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

func main() {

	LerDir("./data/")
}

func Duracao(seg string) (duration time.Duration) {

	duration, err := time.ParseDuration(seg)

	if err != nil {
		log.Fatalf("error on parse FileExpirationTime: %v", err)
	}

	return
}

func LerFile(Pathf string) []os.FileInfo {

	files, err := ioutil.ReadDir(Pathf)

	if err != nil {
		log.Fatalf("error reading output directory: %v", err)
	}

	return files

}

func LerDir(Pathf string) {

	duration := Duracao("10s")
	files := LerFile(Pathf)

	for _, file := range files {
		if file.ModTime().Add(duration).Before(time.Now()) {
			err = os.Remove(path.Join(Pathf, file.Name()))
			if err != nil {
				log.Fatalf("error deleting file: %v", err)
			} else {
				log.Println("removendo file: ", file.Name())
			}
		}
	}
}
