package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	ExampleShutdown()
}

// This example demonstrates the use of shutdown service gracefully
func ExampleShutdown() {
	// Generate a new context
	ctx := NewContext()

	// Run service with this context
	go func(ctx context.Context) {
		if err := ExitWaitGroupAdd(ctx, 1); err != nil {
			return
		}
		defer ExitWaitGroupDone(ctx)

		otherEvent := make(chan struct{})
	FOR_LOOP:
		for {
			select {
			case <-ctx.Done():
				break FOR_LOOP
			case <-otherEvent:
				// ...
			}
		}
	}(ctx)

	// Wait interrupt signal and shutdown gracefully
	if err := WaitAndShutdown(ctx, time.Second*5, func(timeout time.Duration) error {
		log.Println("close")
		return nil
	}); err != nil {
		log.Println("Shutdown error:", err)
		return
	}
}

// ContextKey context key type
type ContextKey string

// CloseFunc close function
type CloseFunc func(time.Duration) error

const (
	// ContextCancel cancel func value name in context
	ContextCancel ContextKey = "shutdown/cancel"
	// ContextExitWaitGroup waitgroup object name in context
	ContextExitWaitGroup ContextKey = "shutdown/exitWaitGroup"
)

var (
	// ErrNoCancel no cancel value error
	ErrNoCancel = errors.New("no cancel value in context")
	// ErrCancelValueTypeError cancel value in context type error
	ErrCancelValueTypeError = errors.New("cancel value in context type error")
	// ErrNoExitWaitGroup no exitWaitGroup value in context
	ErrNoExitWaitGroup = errors.New("no exitWaitGroup value in context")
	// ErrExitWaitGroupValueTypeError exitWaitGroup value in context type error
	ErrExitWaitGroupValueTypeError = errors.New("exitWaitGroup value in context type error")
	// ErrShutdownTimeout shutdown timeout
	ErrShutdownTimeout = errors.New("shutdown timeout")
	// TerminationSignals 退出信号量.
	TerminationSignals = []os.Signal{syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT}
)

// NewContext create new context
func NewContext() (ctx context.Context) {
	var exitWaitGroup sync.WaitGroup
	var ctxBase context.Context
	var cancel context.CancelFunc
	ctxBase, cancel = context.WithCancel(context.Background())
	ctx = context.WithValue(context.WithValue(ctxBase, ContextExitWaitGroup, &exitWaitGroup), ContextCancel, cancel)
	return
}

// Shutdown shutdown gracefully
func Shutdown(ctx context.Context, shutdownTimeout time.Duration, closeFunc CloseFunc) error {
	value := ctx.Value(ContextCancel)
	if value == nil {
		return ErrNoCancel
	}
	cancel, ok := value.(context.CancelFunc)
	if !ok {
		return ErrCancelValueTypeError
	}
	value = ctx.Value(ContextExitWaitGroup)
	if value == nil {
		return ErrNoExitWaitGroup
	}
	exitWaitGroup, ok := value.(*sync.WaitGroup)
	if !ok {
		return ErrExitWaitGroupValueTypeError
	}
	exitSignal := make(chan struct{})
	go func() {
		cancel()
		if closeFunc != nil {
			closeFunc(shutdownTimeout)
		}
		exitWaitGroup.Wait()
		exitSignal <- struct{}{}
	}()
	select {
	case <-exitSignal:
		fmt.Println("Exit successfully")
	case <-time.After(shutdownTimeout):
		return ErrShutdownTimeout
	}
	if err := ctx.Err(); err != context.Canceled {
		return err
	}
	return nil
}

// WaitAndShutdown shutdown gracefully
func WaitAndShutdown(ctx context.Context, shutdownTimeout time.Duration, closeFunc CloseFunc) error {
	return WaitTerminationSignal(shutdownTimeout, func(timeout time.Duration) error {
		return Shutdown(ctx, timeout, closeFunc)
	})
}

// ExitWaitGroupAdd waitgroup counter adds delta
func ExitWaitGroupAdd(ctx context.Context, i int) error {
	value := ctx.Value(ContextExitWaitGroup)
	if value == nil {
		return ErrNoExitWaitGroup
	}
	exitWaitGroup, ok := value.(*sync.WaitGroup)
	if !ok {
		return ErrExitWaitGroupValueTypeError
	}
	exitWaitGroup.Add(i)
	return nil
}

// ExitWaitGroupDone waitgroup down
func ExitWaitGroupDone(ctx context.Context) error {
	value := ctx.Value(ContextExitWaitGroup)
	if value == nil {
		return ErrNoExitWaitGroup
	}
	exitWaitGroup, ok := value.(*sync.WaitGroup)
	if !ok {
		return ErrExitWaitGroupValueTypeError
	}
	exitWaitGroup.Done()
	return nil
}

// WaitTerminationSignal 等待退出信号
func WaitTerminationSignal(timeout time.Duration, fn CloseFunc) error {
	terminationSignalsCh := make(chan os.Signal, 1)
	signal.Notify(terminationSignalsCh, TerminationSignals...)
	defer func() {
		signal.Stop(terminationSignalsCh)
		close(terminationSignalsCh)
	}()
	<-terminationSignalsCh
	return fn(timeout)
}
