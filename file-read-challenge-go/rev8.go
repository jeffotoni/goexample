package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	start := time.Now()
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	names := make([]string, 0)
	firstNames := make([]string, 0)
	dates := make([]string, 0)
	commonName := ""
	commonCount := 0
	scanner := bufio.NewScanner(file)
	nameMap := make(map[string]int)
	dateMap := make(map[string]int)

	type entry struct {
		firstName string
		name      string
		date      string
	}

	linesChunkLen := 64 * 1024
	linesChunkPoolAllocated := int64(0)
	linesPool := sync.Pool{New: func() interface{} {
		lines := make([]string, 0, linesChunkLen)
		atomic.AddInt64(&linesChunkPoolAllocated, 1)
		return lines
	}}
	lines := linesPool.Get().([]string)[:0]

	entriesPoolAllocated := int64(0)
	entriesPool := sync.Pool{New: func() interface{} {
		entries := make([]entry, 0, linesChunkLen)
		atomic.AddInt64(&entriesPoolAllocated, 1)
		return entries
	}}
	mutex := &sync.Mutex{}
	wg := sync.WaitGroup{}

	scanner.Scan()
	for {
		lines = append(lines, scanner.Text())
		willScan := scanner.Scan()
		if len(lines) == linesChunkLen || !willScan {
			linesToProcess := lines
			wg.Add(len(linesToProcess))
			go func() {
				entries := entriesPool.Get().([]entry)[:0]
				for _, text := range linesToProcess {
					// get all the names
					entry := entry{}
					split := strings.SplitN(text, "|", 9)
					name := strings.TrimSpace(split[7])
					entry.name = name

					// extract first names
					if name != "" {
						startOfName := strings.Index(name, ", ") + 2
						if endOfName := strings.Index(name[startOfName:], " "); endOfName < 0 {
							entry.firstName = name[startOfName:]
						} else {
							entry.firstName = name[startOfName : startOfName+endOfName]
						}
						if strings.HasSuffix(entry.firstName, ",") {
							entry.firstName = strings.Replace(entry.firstName, ",", "", -1)
						}
					}
					// extract dates
					chars := strings.TrimSpace(split[4])[:6]
					entry.date = chars[:4] + "-" + chars[4:6]
					entries = append(entries, entry)
				}
				mutex.Lock()
				for _, entry := range entries {
					if entry.firstName != "" {
						firstNames = append(firstNames, entry.firstName)
						nameCount := nameMap[entry.firstName]
						nameMap[entry.firstName] = nameCount + 1
						if nameCount+1 > commonCount {
							commonName = entry.firstName
							commonCount = nameCount + 1
						}
					}
					names = append(names, entry.name)
					dates = append(dates, entry.date)
					dateMap[entry.date]++
				}
				entriesPool.Put(entries)
				linesPool.Put(linesToProcess)
				wg.Add(-len(entries))
				mutex.Unlock()
			}()
			lines = linesPool.Get().([]string)[:0]
		}
		if !willScan {
			break
		}
	}
	wg.Wait()

	// report c2: names at index
	fmt.Printf("Name: %s at index: %v\n", names[0], 0)
	fmt.Printf("Name: %s at index: %v\n", names[432], 432)
	fmt.Printf("Name: %s at index: %v\n", names[43243], 43243)
	fmt.Printf("Name time: %v\n", time.Since(start))

	// report c1: total number of lines
	fmt.Printf("Total file line count: %v\n", len(names))
	fmt.Printf("Line count time: %v\n", time.Since(start))

	// report c3: donation frequency
	for k, v := range dateMap {
		fmt.Printf("Donations per month and year: %v and donation count: %v\n", k, v)
	}
	fmt.Printf("Donations time: %v\n", time.Since(start))

	// report c4: most common firstName
	fmt.Printf("The most common first name is: %s and it occurs: %v times.\n", commonName, commonCount)
	fmt.Printf("Most common name time: %v\n", time.Since(start))
}
