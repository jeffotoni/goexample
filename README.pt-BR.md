# goexample

[![Go](https://img.shields.io/badge/Go-exemplos-00ADD8?style=flat-square&logo=go&logoColor=white)](https://go.dev/)
[![Concorrência](https://img.shields.io/badge/concorrência-goroutines%20%26%20canais-00ADD8?style=flat-square)](./goroutine/)
[![APIs](https://img.shields.io/badge/APIs-http%20e%20web-0A7BBB?style=flat-square)](./api/)
[![WebAssembly](https://img.shields.io/badge/WebAssembly-experimentos-654FF0?style=flat-square&logo=webassembly&logoColor=white)](./webassembly/)
[![Ebitengine](https://img.shields.io/badge/Ebitengine-experimentos-111111?style=flat-square)](./ebiten/)
[![gRPC](https://img.shields.io/badge/gRPC-exemplos-244c5a?style=flat-square)](./grpc/)
[![Docker](https://img.shields.io/badge/Docker-exemplos-2496ED?style=flat-square&logo=docker&logoColor=white)](./docker/)
[![Kubernetes](https://img.shields.io/badge/Kubernetes-exemplos-326CE5?style=flat-square&logo=kubernetes&logoColor=white)](./kubernetes/)
[![AWS](https://img.shields.io/badge/AWS-exemplos-232F3E?style=flat-square&logo=amazonaws&logoColor=white)](./aws-ec2/)
[![Kafka](https://img.shields.io/badge/Kafka-exemplos-231F20?style=flat-square&logo=apachekafka&logoColor=white)](./kafka/)
[![Redis](https://img.shields.io/badge/Redis-exemplos-DC382D?style=flat-square&logo=redis&logoColor=white)](./redis/)

A documentação em inglês está disponível em [`README.md`](./README.md).

## Sobre este repositório

Este é um caderno de estudos em Go mantido ao longo do tempo. Ele reúne pequenos experimentos, testes, protótipos e ideias de implementação criados durante o aprendizado da linguagem e a preparação de projetos maiores.

Alguns exemplos têm poucas linhas. Outros são experimentos completos com dependências temporárias, servidores, clientes, benchmarks ou serviços externos.

O repositório não é uma biblioteca única e não deve ser compilado como um projeto só. Cada diretório tem seu próprio objetivo e deve ser lido e executado de forma independente.

A árvore atual se parece mais com um laboratório pessoal do que com um repositório de aplicação convencional.

## Visão geral do repositório

Os números são aproximados e usam o sinal de mais de propósito. O repositório continua crescendo com novos experimentos.

| Escopo | Tamanho |
| --- | --- |
| Diretórios na raiz | 190+ |
| Arquivos de código Go | 990+ |
| READMEs locais | 50+ |
| Áreas principais | Linguagem, concorrência, web, cloud, dados, performance, gráficos e sistemas |

## O que você encontra aqui

| Área | Exemplos |
| --- | --- |
| Fundamentos da linguagem | [`variable`](./variable/), [`functions`](./functions/), [`interface`](./interface/), [`generics`](./generics/), [`pointer`](./pointer/), [`parsejson`](./parsejson/) e [`gogenerate`](./gogenerate/) |
| Concorrência | [`goroutine`](./goroutine/), [`channel`](./channel/), [`context`](./context/), [`patterns`](./patterns/), [`sync`](./sync/), [`race`](./race/) e [`memoization`](./memoization/) |
| Web e APIs | [`api`](./api/), [`api-simple`](./api-simple/), [`gin-gonic`](./gin-gonic/), [`echo`](./echo/), [`websocket`](./websocket/), [`http2`](./http2/), [`socket`](./socket/), [`grpc`](./grpc/) e [`swagger`](./swagger/) |
| Cloud e mensageria | [`aws-ec2`](./aws-ec2/), [`aws-s3`](./aws-s3/), [`aws-sqs`](./aws-sqs/), [`aws-sns`](./aws-sns/), [`google.cloud`](./google.cloud/), [`kafka`](./kafka/), [`queue.examples`](./queue.examples/) e [`kubernetes`](./kubernetes/) |
| Dados e persistência | [`redis`](./redis/), [`mongodb`](./mongodb/), [`postgres`](./postgres/), [`sqlite3`](./sqlite3/), [`boltdb`](./boltdb/), [`cache-memory`](./cache-memory/), [`crypt`](./crypt/) e [`marshal`](./marshal/) |
| Algoritmos e performance | [`benchmark`](./benchmark/), [`dijkstra`](./dijkstra/), [`performance`](./performance/), [`gc`](./gc/), [`memory`](./memory/), [`pprof.go`](./pprof.go) e [`ordenacao.quicksort`](./ordenacao.quicksort/) |
| Gráficos e mídia | [`ebiten`](./ebiten/), [`webassembly`](./webassembly/), [`gocv.video`](./gocv.video/), [`go-qrcode`](./go-qrcode/), [`image`](./image/) e [`upload`](./upload/) |
| Segurança e arquitetura | [`jwt`](./jwt/), [`authbasic`](./authbasic/), [`go-tls-ssl`](./go-tls-ssl/), [`circuit-breaker`](./circuit-breaker/), [`patherns`](./patherns/) e [`mocktest`](./mocktest/) |

## Por onde começar

- [`goroutine`](./goroutine/) reúne muitos experimentos de concorrência.
- [`channel`](./channel/) explora criação, direção, buffer e comunicação entre canais.
- [`patterns`](./patterns/) contém exemplos de fan in, fan out, worker, semaphore e pipeline.
- [`api`](./api/) e [`api-simple`](./api-simple/) reúnem experimentos com HTTP e APIs web.
- [`benchmark`](./benchmark/) contém testes focados em performance.
- [`webassembly`](./webassembly/) contém experimentos de Go no navegador.
- [`crypt`](./crypt/), [`jwt`](./jwt/) e [`go-tls-ssl`](./go-tls-ssl/) cobrem experimentos relacionados à segurança.
- [`queue.examples`](./queue.examples/), [`kafka`](./kafka/) e [`google.cloud`](./google.cloud/) exploram mensageria e integrações cloud.

## Como o repositório está organizado

O repositório não versiona arquivos `go.mod` ou `go.sum` de propósito. Cada exemplo mantém seu código e pode criar um módulo temporário quando precisa de dependências externas.

Antes de executar um exemplo, confira os arquivos locais. O script auxiliar cria um módulo temporário isolado e o remove quando o comando termina.

```bash
./scripts/run-example.sh caminho/do/exemplo
```

Também é possível passar outro comando depois do diretório do exemplo.

```bash
./scripts/run-example.sh caminho/do/exemplo go test ./...
```

Para um exemplo que use apenas a biblioteca padrão, também é possível desativar o modo de módulos diretamente.

```bash
cd caminho/do/pacote
GO111MODULE=off go test .
```

## Política de dependências

Os manifestos de dependência são locais e temporários de propósito. Isso inclui exemplos em Go, Node.js, PHP e Java. O repositório fica focado no código e não mantém dependências antigas de experimentos no dependency graph.

O script auxiliar executa `go mod tidy` dentro de um diretório temporário para exemplos Go. Ele nunca grava um arquivo de módulo no repositório. As outras linguagens devem seguir o mesmo princípio de ambiente local.

## Estilo do projeto

Os exemplos priorizam clareza e experimentação em vez de uma arquitetura única. Alguns diretórios são modernos, outros são históricos, e alguns mostram uma abordagem de baixo nível de propósito.

Leia o código e o README local antes de considerar um exemplo pronto para produção. Serviços externos podem exigir credenciais, containers locais ou configurações adicionais.

## Contribuição e aprendizado

Este repositório começou como uma forma pessoal de aprender Go praticando. Se algum exemplo ajudar no seu entendimento, use-o como ponto de partida e adapte-o ao seu projeto.

O próximo passo mais valioso para qualquer experimento costuma ser um teste pequeno que registre o comportamento estudado.
