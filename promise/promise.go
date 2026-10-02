// Package promise provides a small Promise-like abstraction built with Go channels.
package promise

import (
	"fmt"
	"sync"
)

// Promise represents a computation that will finish with a value or an error.
//
// A Promise is completed exactly once. The done channel is closed when the
// result is available, so multiple consumers can wait for the same result.
type Promise struct {
	done chan struct{}

	once  sync.Once
	value any
	err   error
}

// New starts work in a goroutine and returns its Promise immediately.
//
// The executor must call exactly one of resolve or reject. If it calls both,
// only the first call has an effect. A panic inside the executor is converted
// into an error and rejects the Promise.
func New(executor func(resolve func(any), reject func(error))) *Promise {
	p := &Promise{done: make(chan struct{})}

	resolve := func(value any) {
		p.once.Do(func() {
			p.value = value
			close(p.done)
		})
	}

	reject := func(err error) {
		p.once.Do(func() {
			p.err = err
			close(p.done)
		})
	}

	go func() {
		defer func() {
			if recovered := recover(); recovered != nil {
				reject(fmt.Errorf("promise executor panic: %v", recovered))
			}
		}()

		executor(resolve, reject)
	}()

	return p
}

// Await blocks until the Promise is fulfilled or rejected.
func (p *Promise) Await() (any, error) {
	<-p.done
	return p.value, p.err
}

// Then creates a new Promise whose value is produced by onFulfilled.
// Errors skip onFulfilled and propagate to the returned Promise.
func (p *Promise) Then(onFulfilled func(any) any) *Promise {
	return New(func(resolve func(any), reject func(error)) {
		value, err := p.Await()
		if err != nil {
			reject(err)
			return
		}

		if onFulfilled == nil {
			resolve(value)
			return
		}

		resolve(onFulfilled(value))
	})
}

// Catch creates a new Promise that handles a rejection from p.
// A fulfilled Promise passes through unchanged.
func (p *Promise) Catch(onRejected func(error) any) *Promise {
	return New(func(resolve func(any), reject func(error)) {
		value, err := p.Await()
		if err == nil {
			resolve(value)
			return
		}

		if onRejected == nil {
			reject(err)
			return
		}

		resolve(onRejected(err))
	})
}

// Finally creates a new Promise that runs onFinally regardless of the result.
// The original value or error is preserved unless onFinally panics.
func (p *Promise) Finally(onFinally func()) *Promise {
	return New(func(resolve func(any), reject func(error)) {
		value, err := p.Await()

		if onFinally != nil {
			onFinally()
		}

		if err != nil {
			reject(err)
			return
		}

		resolve(value)
	})
}
