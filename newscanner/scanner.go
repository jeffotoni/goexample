package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const TOTAL_TWITTES = 280

func main() {

	file, err := os.Open("./text.txt")
	if err != nil {
		log.Println(err)
		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	strtext := ""
	j := 0
	for scanner.Scan() {

		linhatwitt := scanner.Text()
		total := len(linhatwitt)

		for i := 0; i < len(linhatwitt); i++ {
			if i <= TOTAL_TWITTES {
				strtext = strtext + string(linhatwitt[i])
			} else {
				break
			}
		}

		strprint := fmt.Sprintf("[%d/%d]", j, total) + " " + strtext
		fmt.Println(strprint)

		strtext = ""
		j++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
