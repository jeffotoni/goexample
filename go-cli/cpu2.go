package main

import (
	"bytes"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type Process struct {
	pid int
	cpu float64
}

func main() {
	cmd := exec.Command("ps", "aux")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	processes := make([]*Process, 0)
	for {
		line, err := out.ReadString('\n')
		if err != nil {
			break
		}
		tokens := strings.Split(line, " ")
		ft := make([]string, 0)
		for _, t := range tokens {
			if t != "" && t != "\t" {
				ft = append(ft, t)
			}
		}
		log.Println(len(ft), ft)
		pid, err := strconv.Atoi(ft[1])
		if err != nil {
			continue
		}
		cpu, err := strconv.ParseFloat(ft[2], 64)
		if err != nil {
			log.Fatal(err)
		}
		processes = append(processes, &Process{pid, cpu})
	}
	for _, p := range processes {
		log.Println("Process ", p.pid, " takes ", p.cpu, " % of the CPU")
	}
}
