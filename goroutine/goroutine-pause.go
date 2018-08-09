package main

import "fmt"
import "time"
import "sync"

func prepare() (pause func(t time.Duration), start, play, quit, wait func()) {
	var (
		chWork       <-chan struct{}
		chWorkBackup <-chan struct{}
		chControl    chan struct{}
		wg           sync.WaitGroup
	)

	routine := func() {

		defer wg.Done()

		i := 0
		for {
			select {
			case <-chWork:
				fmt.Println("i: ", i)
				i++
				time.Sleep(850 * time.Millisecond)
			case _, ok := <-chControl:
				if ok {
					continue
				}
				return
			}
		}
	}

	start = func() {
		// chWork, chWorkBackup
		ch := make(chan struct{})
		defer close(ch)

		chWork = ch
		chWorkBackup = ch

		// chControl
		chControl = make(chan struct{})

		// wg
		wg = sync.WaitGroup{}
		wg.Add(1)

		go routine()
	}

	pause = func(t time.Duration) {
		chWork = nil
		chControl <- struct{}{}
		fmt.Println("pause: ", t)
		time.Sleep(t * time.Second)
	}

	play = func() {
		fmt.Println("play")
		chWork = chWorkBackup
		chControl <- struct{}{}
	}

	quit = func() {
		chWork = nil
		close(chControl)
		fmt.Println("quit")
	}

	wait = func() {
		wg.Wait()
	}

	return
}

func sleep() {
	time.Sleep(3 * time.Second)
}

func main() {

	//_, start, play, _, _ := prepare()
	pause, start, play, quit, wait := prepare()

	j := 0

	for {

		sleep()
		start()
		fmt.Println("start() called")

		pause(5)
		play()

		sleep()
		pause(10)

		sleep()
		play()

		sleep()
		quit()

		wait()

		fmt.Println("done: ", j)

		j++
	}
}
