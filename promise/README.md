# Promise in Go

[![Go](https://img.shields.io/badge/Go-1.19%2B-00ADD8?style=flat-square&logo=go&logoColor=white)](https://go.dev/)
[![Concurrency](https://img.shields.io/badge/concurrency-channels-00ADD8?style=flat-square)](https://go.dev/tour/concurrency/1)
[![Tests](https://img.shields.io/badge/tests-passing-2ea44f?style=flat-square)](./promise_test.go)
[![Status](https://img.shields.io/badge/status-experimental-f0ad4e?style=flat-square)](./promise.go)

This directory contains a small Promise-like abstraction inspired by JavaScript.

The goal is educational. It explores how a future result can be represented in Go with a goroutine, a completion channel, and `sync.Once`.

The Portuguese version is available in [`README.pt-BR.md`](./README.pt-BR.md).

## What a JavaScript Promise represents

A Promise represents the future result of an operation. It starts as `pending` and ends once as either `fulfilled` or `rejected`.

```text
pending   -> still running
fulfilled -> completed successfully
rejected  -> completed with an error
```

A Promise is not a thread. It is an object used to observe a result that will be available later. The runtime decides how the underlying operation runs.

## How JavaScript runs Promise callbacks

JavaScript normally executes application code on a main thread. The runtime can handle timers, network operations, and other external work. When that work finishes, Promise callbacks are placed in the microtask queue.

The event loop runs microtasks after the current call stack becomes empty and before the next regular task.

```javascript
console.log("1");

setTimeout(() => console.log("timer"), 0);

Promise.resolve().then(() => console.log("promise"));

console.log("2");
```

The output is

```text
1
2
promise
timer
```

The executor passed to `new Promise` runs immediately. A callback registered with `then` runs later through the microtask queue.

See [`flow.svg`](./flow.svg) for a visual overview.

## JavaScript example

```javascript
function factorial(n) {
  if (!Number.isInteger(n) || n < 0) {
    throw new Error("n must be a non-negative integer");
  }

  let result = 1;
  for (let current = 2; current <= n; current++) {
    result *= current;
  }

  return result;
}

function factorialAsync(n) {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      try {
        resolve(factorial(n));
      } catch (error) {
        reject(error);
      }
    }, 100);
  });
}

Promise.all([
  factorialAsync(5),
  factorialAsync(10),
  factorialAsync(15),
])
  .then(results => console.log(results))
  .catch(error => console.error(error))
  .finally(() => console.log("finished"));
```

This code creates three operations, waits for all of them, keeps the input order, and sends failures to `catch`.

## The Go proposal

Go already provides goroutines, channels, `select`, `sync.WaitGroup`, and `context`. A Promise is therefore an additional abstraction, not a replacement for Go concurrency primitives.

The chosen implementation uses one completion channel.

```text
New
  └─ creates done and starts a goroutine

executor
  ├─ resolve(value) -> stores the value and closes done
  └─ reject(error)  -> stores the error and closes done

Await
  └─ waits for done and returns the value or error
```

The channel is closed instead of receiving the result directly. Closing it broadcasts completion, so several consumers can wait for the same Promise. `sync.Once` ensures that only the first call to `resolve` or `reject` wins.

## Current API

```go
package main

import (
  "fmt"
  "github.com/your/module/promise"
)

func main() {
  p := promise.New(func(resolve func(any), reject func(error)) {
    resolve(2)
  })

  result := p.
    Then(func(value any) any {
      return value.(int) * 3
    }).
    Catch(func(err error) any {
      fmt.Println("error:", err)
      return 0
    })

  value, err := result.Await()
  if err != nil {
    panic(err)
  }

  fmt.Println(value)
}
```

The package currently provides

- `New`
- `Await`
- `Then`
- `Catch`
- `Finally`

An executor panic is converted into an error. A Promise can be consumed more than once, including after it has already completed.

## Files

- [`promise.go`](./promise.go) contains the implementation.
- [`promise_test.go`](./promise_test.go) covers success, rejection, chaining, recovery, finalization, multiple consumers, panic handling, and error propagation.
- [`flow.svg`](./flow.svg) shows the JavaScript Promise flow.
- [`README.pt-BR.md`](./README.pt-BR.md) contains the Portuguese documentation.

The repository also contains older experiments in [`goroutine`](../goroutine)

- [`promise-fact-chebyrash.go`](../goroutine/promise-fact-chebyrash.go)
- [`promise-sum-chebyrash.go`](../goroutine/promise-sum-chebyrash.go)
- [`promise-proprio.go`](../goroutine/promise-proprio.go)
- [`promise-ex1.go`](../goroutine/promise-ex1.go)

## Limitations

This is a study project. It does not yet provide `Promise.all`, `Race`, cancellation, timeouts, generic types, or Promise adoption when a callback returns another Promise.

`Await` blocks a goroutine. The abstraction does not reproduce the JavaScript event loop or its microtask scheduling.

## Run the tests

This directory does not use a Go module, so run the package with module mode disabled from inside `promise`.

```bash
cd promise
GO111MODULE=off go test .
GO111MODULE=off go test -race .
GO111MODULE=off go vet .
```

The next useful step is to compare this implementation with direct goroutine and channel code before adding more features.
