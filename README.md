# goexample

[![Go](https://img.shields.io/badge/Go-examples-00ADD8?style=flat-square&logo=go&logoColor=white)](https://go.dev/)
[![Concurrency](https://img.shields.io/badge/concurrency-goroutines%20%26%20channels-00ADD8?style=flat-square)](./goroutine/)
[![APIs](https://img.shields.io/badge/APIs-http%20and%20web-0A7BBB?style=flat-square)](./api/)
[![WebAssembly](https://img.shields.io/badge/WebAssembly-experiments-654FF0?style=flat-square&logo=webassembly&logoColor=white)](./webassembly/)
[![Ebitengine](https://img.shields.io/badge/Ebitengine-experiments-111111?style=flat-square)](./ebiten/)
[![gRPC](https://img.shields.io/badge/gRPC-examples-244c5a?style=flat-square)](./grpc/)
[![Docker](https://img.shields.io/badge/Docker-examples-2496ED?style=flat-square&logo=docker&logoColor=white)](./docker/)
[![Kubernetes](https://img.shields.io/badge/Kubernetes-examples-326CE5?style=flat-square&logo=kubernetes&logoColor=white)](./kubernetes/)
[![AWS](https://img.shields.io/badge/AWS-examples-232F3E?style=flat-square&logo=amazonaws&logoColor=white)](./aws-ec2/)
[![Kafka](https://img.shields.io/badge/Kafka-examples-231F20?style=flat-square&logo=apachekafka&logoColor=white)](./kafka/)
[![Redis](https://img.shields.io/badge/Redis-examples-DC382D?style=flat-square&logo=redis&logoColor=white)](./redis/)

The Portuguese documentation is available in [`README.pt-BR.md`](./README.pt-BR.md).

## About this repository

This is a long-running Go notebook. It contains small experiments, tests, prototypes, and implementation ideas built while learning the language and preparing larger projects.

Some examples are only a few lines long. Others are complete experiments with temporary dependencies, servers, clients, benchmarks, or external services.

The repository is not a single library and it is not meant to be built as one project. Each directory has its own purpose and should be read and run independently.

The current tree is closer to a personal laboratory than to a conventional application repository.

## Repository snapshot

These numbers are approximate and intentionally use a plus sign. The repository keeps growing as new experiments are added.

| Scope | Size |
| --- | --- |
| Top level directories | 190+ |
| Go source files | 990+ |
| Local README files | 50+ |
| Main areas | Language, concurrency, web, cloud, data, performance, graphics, and systems |

## What you can find here

| Area | Examples |
| --- | --- |
| Language foundations | [`variable`](./variable/), [`functions`](./functions/), [`interface`](./interface/), [`generics`](./generics/), [`pointer`](./pointer/), [`parsejson`](./parsejson/), and [`gogenerate`](./gogenerate/) |
| Concurrency | [`goroutine`](./goroutine/), [`channel`](./channel/), [`context`](./context/), [`patterns`](./patterns/), [`sync`](./sync/), [`race`](./race/), and [`memoization`](./memoization/) |
| Web and APIs | [`api`](./api/), [`api-simple`](./api-simple/), [`gin-gonic`](./gin-gonic/), [`echo`](./echo/), [`websocket`](./websocket/), [`http2`](./http2/), [`socket`](./socket/), [`grpc`](./grpc/), and [`swagger`](./swagger/) |
| Cloud and messaging | [`aws-ec2`](./aws-ec2/), [`aws-s3`](./aws-s3/), [`aws-sqs`](./aws-sqs/), [`aws-sns`](./aws-sns/), [`google.cloud`](./google.cloud/), [`kafka`](./kafka/), [`queue.examples`](./queue.examples/), and [`kubernetes`](./kubernetes/) |
| Data and persistence | [`redis`](./redis/), [`mongodb`](./mongodb/), [`postgres`](./postgres/), [`sqlite3`](./sqlite3/), [`boltdb`](./boltdb/), [`cache-memory`](./cache-memory/), [`crypt`](./crypt/), and [`marshal`](./marshal/) |
| Algorithms and performance | [`benchmark`](./benchmark/), [`dijkstra`](./dijkstra/), [`performance`](./performance/), [`gc`](./gc/), [`memory`](./memory/), [`pprof.go`](./pprof.go), and [`ordenacao.quicksort`](./ordenacao.quicksort/) |
| Graphics and media | [`ebiten`](./ebiten/), [`webassembly`](./webassembly/), [`gocv.video`](./gocv.video/), [`go-qrcode`](./go-qrcode/), [`image`](./image/), and [`upload`](./upload/) |
| Security and architecture | [`jwt`](./jwt/), [`authbasic`](./authbasic/), [`go-tls-ssl`](./go-tls-ssl/), [`circuit-breaker`](./circuit-breaker/), [`patherns`](./patherns/), and [`mocktest`](./mocktest/) |

## A few starting points

- [`goroutine`](./goroutine/) contains many concurrency experiments.
- [`channel`](./channel/) focuses on channel creation, direction, buffering, and communication.
- [`patterns`](./patterns/) contains fan in, fan out, worker, semaphore, and pipeline examples.
- [`api`](./api/) and [`api-simple`](./api-simple/) contain HTTP and web API experiments.
- [`benchmark`](./benchmark/) contains focused performance tests.
- [`webassembly`](./webassembly/) contains browser oriented Go experiments.
- [`crypt`](./crypt/), [`jwt`](./jwt/), and [`go-tls-ssl`](./go-tls-ssl/) cover security related experiments.
- [`queue.examples`](./queue.examples/), [`kafka`](./kafka/), and [`google.cloud`](./google.cloud/) explore messaging and cloud integrations.

## How the repository is organized

The repository intentionally does not commit `go.mod` or `go.sum` files. Each example owns its source code and can create a temporary module when external dependencies are needed.

Before running an example, enter its directory and check its local files. The helper script creates an isolated temporary module and removes it when the command finishes.

```bash
./scripts/run-example.sh path/to/example
```

You can pass another command after the example directory.

```bash
./scripts/run-example.sh path/to/example go test ./...
```

For a standard library example, module mode can also be disabled directly.

```bash
cd path/to/package
GO111MODULE=off go test .
```

## Dependency policy

Dependency files are deliberately local and temporary. This keeps the repository focused on source code and prevents old experiment manifests from keeping stale dependency alerts alive.

The helper script runs `go mod tidy` inside a temporary directory. It never writes a module file into the repository.

## Project style

The examples favor clarity and experimentation over a single shared architecture. Some directories are modern, some are historical, and some intentionally show a low level approach.

Read the local code and README before assuming that an example is production ready. External services may require credentials, local containers, or additional configuration.

## Contributions and learning

This repository started as a personal way to learn Go through practice. If an example helps you understand a concept, use it as a starting point and adapt it to your own project.

The most valuable next step for any experiment is usually a small test that records the behavior being explored.
