package promise

import (
	"errors"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
)

func TestPromiseResolvesValue(t *testing.T) {
	p := New(func(resolve func(any), reject func(error)) {
		resolve(42)
	})

	value, err := p.Await()
	if err != nil {
		t.Fatalf("Await() error = %v", err)
	}
	if value != 42 {
		t.Fatalf("Await() value = %v, want 42", value)
	}
}

func TestPromiseRejectsError(t *testing.T) {
	wantErr := errors.New("falha simulada")
	p := New(func(resolve func(any), reject func(error)) {
		reject(wantErr)
	})

	_, err := p.Await()
	if !errors.Is(err, wantErr) {
		t.Fatalf("Await() error = %v, want %v", err, wantErr)
	}
}

func TestPromiseCompletesOnlyOnce(t *testing.T) {
	p := New(func(resolve func(any), reject func(error)) {
		resolve("primeiro")
		resolve("segundo")
		reject(errors.New("não deveria vencer"))
	})

	value, err := p.Await()
	if err != nil {
		t.Fatalf("Await() error = %v", err)
	}
	if value != "primeiro" {
		t.Fatalf("Await() value = %v, want primeiro", value)
	}
}

func TestPromiseThenChainsValues(t *testing.T) {
	p := New(func(resolve func(any), reject func(error)) {
		resolve(2)
	})

	result := p.
		Then(func(value any) any {
			return value.(int) * 3
		}).
		Then(func(value any) any {
			return value.(int) + 1
		})

	value, err := result.Await()
	if err != nil {
		t.Fatalf("Await() error = %v", err)
	}
	if value != 7 {
		t.Fatalf("Await() value = %v, want 7", value)
	}
}

func TestPromiseCatchRecoversError(t *testing.T) {
	wantErr := errors.New("erro original")
	p := New(func(resolve func(any), reject func(error)) {
		reject(wantErr)
	})

	result := p.Catch(func(err error) any {
		if !errors.Is(err, wantErr) {
			t.Errorf("Catch() error = %v, want %v", err, wantErr)
		}
		return "recuperado"
	})

	value, err := result.Await()
	if err != nil {
		t.Fatalf("Await() error = %v", err)
	}
	if value != "recuperado" {
		t.Fatalf("Await() value = %v, want recuperado", value)
	}
}

func TestPromiseFinallyRunsAndPreservesValue(t *testing.T) {
	var called atomic.Bool
	p := New(func(resolve func(any), reject func(error)) {
		resolve("ok")
	})

	value, err := p.Finally(func() {
		called.Store(true)
	}).Await()

	if err != nil {
		t.Fatalf("Await() error = %v", err)
	}
	if value != "ok" {
		t.Fatalf("Await() value = %v, want ok", value)
	}
	if !called.Load() {
		t.Fatal("Finally() callback was not called")
	}
}

func TestPromiseFinallyRunsOnError(t *testing.T) {
	wantErr := errors.New("falha")
	var called atomic.Bool
	p := New(func(resolve func(any), reject func(error)) {
		reject(wantErr)
	})

	_, err := p.Finally(func() {
		called.Store(true)
	}).Await()

	if !errors.Is(err, wantErr) {
		t.Fatalf("Await() error = %v, want %v", err, wantErr)
	}
	if !called.Load() {
		t.Fatal("Finally() callback was not called after rejection")
	}
}

func TestPromiseAllowsMultipleConsumers(t *testing.T) {
	p := New(func(resolve func(any), reject func(error)) {
		resolve("mesmo resultado")
	})

	const consumers = 8
	values := make(chan any, consumers)
	var wg sync.WaitGroup
	wg.Add(consumers)

	for i := 0; i < consumers; i++ {
		go func() {
			defer wg.Done()
			value, err := p.Await()
			if err != nil {
				t.Errorf("Await() error = %v", err)
				return
			}
			values <- value
		}()
	}

	wg.Wait()
	close(values)

	for value := range values {
		if value != "mesmo resultado" {
			t.Fatalf("consumer value = %v, want mesmo resultado", value)
		}
	}
}

func TestPromiseConvertsExecutorPanicToError(t *testing.T) {
	p := New(func(resolve func(any), reject func(error)) {
		panic("quebrou")
	})

	_, err := p.Await()
	if err == nil {
		t.Fatal("Await() error = nil, want panic error")
	}
	if !strings.Contains(err.Error(), "quebrou") {
		t.Fatalf("Await() error = %q, want panic message", err)
	}
}

func TestPromisePropagatesErrorThroughThen(t *testing.T) {
	wantErr := errors.New("não passou")
	called := false

	p := New(func(resolve func(any), reject func(error)) {
		reject(wantErr)
	})

	_, err := p.Then(func(value any) any {
		called = true
		return value
	}).Await()

	if !errors.Is(err, wantErr) {
		t.Fatalf("Await() error = %v, want %v", err, wantErr)
	}
	if called {
		t.Fatal("Then() callback was called after rejection")
	}
}
