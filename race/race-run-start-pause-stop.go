package main

import "fmt"
import "time"
import "sync"

// var Lock sync.Mutex

func prepare() (start, pause, play, quit, wait func()) {
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
				fmt.Println(i)
				i++
				time.Sleep(250 * time.Millisecond)
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
		close(ch)
		chWork = ch
		chWorkBackup = ch

		// chControl
		chControl = make(chan struct{})

		// wg
		wg = sync.WaitGroup{}
		wg.Add(1)

		go routine()
	}

	pause = func() {

		//Lock.Lock()
		chWork = nil
		//Lock.Unlock()

		chControl <- struct{}{}
		fmt.Println("pause")
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
	time.Sleep(1 * time.Second)
}

func main() {
	start, pause, play, quit, wait := prepare()

	sleep()
	start()
	fmt.Println("start() called")

	sleep()
	pause()

	sleep()
	play()

	sleep()
	pause()

	sleep()
	play()

	sleep()
	quit()

	wait()
	fmt.Println("done")
}
