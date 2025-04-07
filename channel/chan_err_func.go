package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

type Consumer struct {
	Concurrency int
}

func (w *Consumer) Start(ctx context.Context) <-chan error {
	errCh := make(chan error)
	sem := make(chan struct{}, w.Concurrency)
	var wg sync.WaitGroup

	go func() {
		for i := 0; i < 5; i++ {
			select {
			case <-ctx.Done():
				errCh <- errors.New("worker canceled via context")
				break
			default:
				sem <- struct{}{}
				wg.Add(1)

				go func(msgID int) {
					defer func() {
						<-sem
						wg.Done()
					}()

					if msgID%2 == 0 {
						errCh <- fmt.Errorf("error processing msg %d", msgID)
						return
					}
					fmt.Printf("✅ Processed with success: msg %d\n", msgID)
				}(i)
			}
		}

		wg.Wait()
		close(errCh)
	}()

	return errCh
}

func main() {
	ctx := context.Background()
	worker := &Consumer{Concurrency: 2}

	errCh := worker.Start(ctx)
	for err := range errCh { // Read the errors returned from the channel
		fmt.Printf("❌ Error received: %s\n", err.Error())
	}
}
