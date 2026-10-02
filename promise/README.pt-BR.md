# Promise em Go - um exercício inspirado no JavaScript

[![Go](https://img.shields.io/badge/Go-1.19%2B-00ADD8?style=flat-square&logo=go&logoColor=white)](https://go.dev/)
[![Concorrência](https://img.shields.io/badge/concorrência-canais-00ADD8?style=flat-square)](https://go.dev/tour/concurrency/1)
[![Testes](https://img.shields.io/badge/testes-passando-2ea44f?style=flat-square)](./promise_test.go)
[![Status](https://img.shields.io/badge/status-experimental-f0ad4e?style=flat-square)](./promise.go)

Este diretório nasce de uma brincadeira antiga: tentar reproduzir, em Go, uma ideia que ficou muito popular no JavaScript - o objeto `Promise`.

A proposta não é afirmar que Go precisa de Promises, nem fingir que os dois runtimes funcionam da mesma maneira. O objetivo é investigar uma pergunta interessante:

> **É possível criar em Go uma API com o comportamento conceitual de uma Promise do JavaScript, usando goroutines, canais, mutexes e outras primitivas da linguagem?**

## 1. O que é uma Promise no JavaScript?

Uma `Promise` é um objeto que representa o resultado futuro de uma operação. Esse resultado ainda pode não existir quando a Promise é criada, mas ela informa se a operação:

- ainda está em andamento;
- terminou com sucesso;
- terminou com erro.

Uma Promise possui três estados conceituais:

```text
pending   -> ainda não terminou
fulfilled -> terminou com sucesso
rejected  -> terminou com erro
```

Depois que uma Promise passa de `pending` para `fulfilled` ou `rejected`, ela fica estabelecida. A transição acontece uma única vez; uma segunda tentativa de resolver ou rejeitar não muda o resultado já definido.

Uma forma comum de criar uma Promise é:

```javascript
const promessa = new Promise((resolve, reject) => {
  // inicia alguma operação

  if (deuCerto) {
    resolve(resultado);
    return;
  }

  reject(new Error("algo deu errado"));
});
```

Depois, o código consumidor registra o que deve acontecer:

```javascript
promessa
  .then(resultado => {
    console.log("sucesso:", resultado);
  })
  .catch(erro => {
    console.error("erro:", erro);
  });
```

### Promise não é uma thread

Uma Promise não cria, por si só, uma thread nem executa uma função em paralelo. Ela é principalmente uma representação e uma API para observar uma operação que terminará no futuro.

Em um navegador, a operação pode ser realizada por APIs do ambiente, como rede, temporizadores ou arquivos. No Node.js, muitas operações dependem do runtime e do `libuv`. O JavaScript normalmente executa o código da aplicação em uma thread principal, enquanto o ambiente pode cuidar de operações externas e avisar quando elas terminarem.

Por isso, estas duas ideias devem ser separadas:

1. **Promise:** como representar sucesso, erro e continuação futura;
2. **mecanismo assíncrono:** como a operação realmente é executada.

No Go, uma goroutine é mais próxima do mecanismo de execução concorrente. Uma Promise em Go seria uma abstração adicional sobre goroutines, canais ou outras formas de sincronização.

## 2. Como a engine do JavaScript participa

Para entender o comportamento de uma Promise, é útil separar os componentes do runtime:

| Componente | Responsabilidade |
| --- | --- |
| Call stack | Executa o código JavaScript atual, de forma ordenada. |
| APIs do ambiente | Fazem trabalhos externos, como timer, rede ou I/O. |
| Task queue | Recebe tarefas como callbacks de timers e eventos. |
| Microtask queue | Recebe continuations de `Promise.then`, `catch` e `finally`. |
| Event loop | Verifica quando a stack está livre e escolhe o próximo trabalho. |

O executor passado ao construtor da Promise é chamado imediatamente:

```javascript
console.log("A");

const p = new Promise(resolve => {
  console.log("executor");
  resolve("resultado");
});

p.then(valor => console.log(valor));

console.log("B");
```

A saída conceitual é:

```text
A
executor
B
resultado
```

O executor é síncrono, mas o callback do `then` é agendado para a fila de microtasks. Ele não interrompe a execução atual para aparecer no meio do código.

### Microtasks e tasks

Quando o JavaScript termina a execução da stack atual, a engine drena as microtasks antes de buscar a próxima task. Essa ordem explica muitos resultados que parecem estranhos à primeira vista:

```javascript
console.log("1");

setTimeout(() => console.log("timer"), 0);

Promise.resolve().then(() => console.log("promise"));

console.log("2");
```

A saída esperada é:

```text
1
2
promise
timer
```

O timer e a Promise não são executados enquanto o código síncrono atual ainda está na call stack. Porém, a continuação da Promise entra em uma fila de microtasks, que é processada antes da task do timer.

### Fluxo mental em ASCII

```text
┌──────────────────────┐
│ código JavaScript    │
│ na call stack        │
└──────────┬───────────┘
           │ cria Promise e registra then/catch
           v
┌──────────────────────┐       operação externa
│ Promise: pending     │ ─────────────────────────┐
└──────────┬───────────┘                          │
           │                                      v
           │                           ┌─────────────────────┐
           │                           │ timer / rede / I/O  │
           │                           └──────────┬──────────┘
           │                                      │ termina
           │                                      v
           │                           resolve(...) ou reject(...)
           │                                      │
           v                                      v
┌──────────────────────┐       ┌─────────────────────────────┐
│ fulfilled / rejected │ ────> │ fila de microtasks          │
└──────────────────────┘       │ callbacks then/catch/finally │
                               └──────────────┬──────────────┘
                                              │ stack livre
                                              v
                                   callbacks são executados
```

Um desenho equivalente está em [`flow.svg`](./flow.svg).

## 3. Exemplo completo em JavaScript

O exemplo abaixo simula o cálculo de um fatorial assíncrono. O cálculo é simples, mas o exemplo ajuda a visualizar o contrato da Promise: iniciar uma operação, resolver com um valor ou rejeitar com um erro, e observar o resultado depois.

```javascript
function fatorial(n) {
  if (!Number.isInteger(n) || n < 0) {
    throw new Error("n deve ser um inteiro não negativo");
  }

  let resultado = 1;
  for (let atual = 2; atual <= n; atual++) {
    resultado *= atual;
  }

  return resultado;
}

function fatorialAsync(n) {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      try {
        resolve(fatorial(n));
      } catch (erro) {
        reject(erro);
      }
    }, 100);
  });
}

console.log("iniciando");

Promise.all([
  fatorialAsync(5),
  fatorialAsync(10),
  fatorialAsync(15),
])
  .then(resultados => {
    console.log("resultados:", resultados);
  })
  .catch(erro => {
    console.error("falhou:", erro.message);
  })
  .finally(() => {
    console.log("finalizado");
  });

console.log("continua executando enquanto os resultados não chegam");
```

### O que esse código faz?

1. `fatorialAsync` retorna imediatamente uma Promise no estado `pending`.
2. `setTimeout` agenda o trabalho para depois da execução síncrona atual.
3. Quando o timer dispara, o cálculo é realizado.
4. `resolve` leva o valor para a Promise, ou `reject` leva o erro.
5. `Promise.all` aguarda todas as Promises.
6. O `then` recebe um array com os resultados na mesma ordem das Promises de entrada.
7. O `catch` trata a falha do conjunto.
8. O `finally` é executado ao final, independentemente do sucesso ou erro.

É importante observar que `Promise.all` não necessariamente executa os trabalhos em paralelo por conta própria. Ele apenas coordena Promises já criadas. O paralelismo ou a concorrência dependem da operação e do runtime.

## 4. A pergunta: conseguimos fazer algo parecido em Go?

Sim, conseguimos criar uma API com conceitos semelhantes:

```text
New        -> cria uma operação pendente
resolve    -> completa com sucesso
reject     -> completa com erro
Then       -> registra continuação de sucesso
Catch      -> registra continuação de erro
Finally    -> registra continuação de encerramento
Await      -> espera pelo encerramento
All        -> combina várias operações
```

Mas “mesmo comportamento” precisa ser entendido com cuidado. Podemos reproduzir o **contrato da abstração**, mas não o modelo inteiro da engine JavaScript. Em Go:

- uma goroutine pode executar código simultaneamente em vários processadores;
- um canal pode bloquear diretamente quem envia ou recebe;
- não existe uma microtask queue padrão equivalente à do JavaScript;
- o scheduler do Go é diferente do event loop do JavaScript;
- `Await` provavelmente bloqueará uma goroutine, em vez de apenas registrar uma continuação;
- cancelamento normalmente deve ser modelado explicitamente com `context.Context`.

Uma API genérica poderia começar com uma ideia semelhante a esta:

```go
type Promise[T any] struct {
	// estado interno protegido por mutex,
	// ou resultado comunicado por canal.
}

func New[T any](work func(resolve func(T), reject func(error))) *Promise[T]

func (p *Promise[T]) Then(func(T) T) *Promise[T]
func (p *Promise[T]) Catch(func(error) T) *Promise[T]
func (p *Promise[T]) Finally(func()) *Promise[T]
func (p *Promise[T]) Await() (T, error)
```

O código acima é apenas uma proposta de contrato. Ainda seria necessário decidir como lidar com transformação de tipos, callbacks assíncronos, pânico, cancelamento, timeout, múltiplos consumidores e vazamento de goroutines.

## 5. Possíveis implementações internas

### Opção A: canais

Um canal pode comunicar um resultado único:

```go
type result[T any] struct {
	value T
	err   error
}

type Promise[T any] struct {
	result <-chan result[T]
}
```

Vantagens dessa abordagem:

- combina naturalmente com o modelo de concorrência do Go;
- `Await` pode usar uma simples operação de leitura;
- a comunicação de resultado fica explícita.

Riscos:

- callbacks registrados depois que o resultado chegou precisam ser tratados;
- múltiplos consumidores podem disputar o mesmo valor se o canal não for planejado;
- uma goroutine pode ficar presa se ninguém consumir o resultado;
- o fechamento do canal e o valor zero precisam ser tratados com cuidado.

### Opção B: estado protegido por mutex

Outra possibilidade é guardar o estado e a lista de callbacks em uma estrutura protegida:

```text
Promise
 ├─ estado: pending | fulfilled | rejected
 ├─ valor ou erro
 ├─ callbacks de sucesso
 └─ callbacks de erro
```

Quando `resolve` ou `reject` é chamado, a estrutura muda de estado uma única vez. Os callbacks já registrados são executados, e callbacks adicionados depois podem ser executados imediatamente ou agendados.

Essa opção se aproxima mais da ideia de vários consumidores observando o mesmo resultado, mas exige mais sincronização e mais decisões de projeto.

### Opção C: combinação de canal e estado

Uma implementação prática pode usar:

- canal ou `sync.Once` para garantir a conclusão única;
- mutex para proteger callbacks e estado;
- `context.Context` para cancelamento;
- goroutines para executar trabalho e callbacks;
- `sync.WaitGroup` ou `errgroup` para combinações do tipo `All`.

O desafio é evitar que a API esconda custos importantes. Uma chamada que parece simples pode criar goroutines, manter referências vivas e bloquear recursos.

## 6. O que os experimentos atuais demonstram

### Fatoriais concorrentes

Em [`promise-fact-chebyrash.go`](../goroutine/promise-fact-chebyrash.go), três cálculos são transformados em Promises:

```go
factorial1 := findFactorialPromise(5)
factorial2 := findFactorialPromise(10)
factorial3 := findFactorialPromise(15)
```

Depois, cada resultado recebe um `Then` e a aplicação usa `promise.AwaitAll` para aguardar todos. Conceitualmente, isso se parece com:

```javascript
Promise.all([
  fatorialAsync(5),
  fatorialAsync(10),
  fatorialAsync(15),
]);
```

### Resolução e rejeição

Em [`promise-sum-chebyrash.go`](../goroutine/promise-sum-chebyrash.go), o trabalho chama `resolve` quando a soma dá certo e `reject` quando encontra um erro. O exemplo também explora encadeamento, `Catch` e `Await`.

### Implementação própria com canais

Em [`promise-proprio.go`](../goroutine/promise-proprio.go), a estrutura contém canais separados para resultado e erro:

```go
type Promise struct {
	Result chan string
	Error  chan error
}
```

O método `Then` inicia uma goroutine que usa `select` para esperar pelo resultado ou pelo erro. É uma demonstração direta e didática, mas ainda não representa todo o contrato de uma Promise JavaScript: por exemplo, não há encadeamento de valores, controle explícito de conclusão única, cancelamento, `Finally` ou coordenação robusta de múltiplos consumidores.

## 7. Vantagens de uma Promise em Go

- **API familiar para quem vem do JavaScript:** facilita comparar os dois modelos;
- **encadeamento expressivo:** operações sucessivas podem ser representadas com `Then`, `Catch` e `Finally`;
- **composição:** uma operação como `All` pode coordenar várias tarefas;
- **separação entre produtor e consumidor:** quem inicia a operação não precisa controlar diretamente todos os consumidores;
- **exercício de concorrência:** força a estudar goroutines, canais, sincronização, memória e tratamento de erros;
- **contrato único:** diferentes fontes assíncronas podem expor a mesma forma de consumo.

## 8. Desvantagens e cuidados

- **duplicação de conceitos do Go:** goroutines, canais, `select`, `WaitGroup` e `context` já resolvem grande parte dos problemas;
- **sem equivalência perfeita com o JavaScript:** uma API igual não produz o mesmo runtime;
- **risco de esconder bloqueios:** `Await` pode bloquear uma goroutine sem que isso fique evidente para quem lê o código;
- **risco de vazamento:** goroutines e callbacks podem permanecer vivos se a operação nunca terminar ou for abandonada;
- **cancelamento não automático:** rejeitar uma Promise não necessariamente cancela o trabalho que já começou;
- **tipagem e encadeamento:** permitir que um `Then` transforme `T` em outro tipo exige uma API genérica mais sofisticada;
- **tratamento de pânico:** é preciso decidir se um `panic` será convertido em erro e em qual goroutine isso será recuperado;
- **overhead:** mutexes, canais, alocações e goroutines podem custar mais do que uma chamada direta;
- **manutenção:** uma abstração que imita outra linguagem pode ficar menos idiomática para desenvolvedores Go.

## 9. Proposta para a implementação deste diretório

O caminho mais seguro é implementar por etapas, com testes para cada contrato:

1. **Estado básico:** `pending`, `fulfilled` e `rejected`;
2. **Conclusão única:** apenas a primeira chamada de `resolve` ou `reject` vence;
3. **Espera:** `Await() (T, error)` sem corrida de dados;
4. **Callbacks:** `Then`, `Catch` e `Finally` para callbacks registrados antes e depois da conclusão;
5. **Encadeamento:** resultado de um callback alimenta a próxima Promise;
6. **Composição:** `All`, talvez `Race`, com preservação da ordem e propagação de erro;
7. **Cancelamento:** integração com `context.Context`;
8. **Robustez:** timeout, pânico, callbacks que bloqueiam e operações que nunca terminam;
9. **Exemplos:** reproduzir o fatorial, a soma e outros pequenos exercícios;
10. **Comparação:** medir a diferença entre usar a Promise experimental e usar diretamente goroutines/canais.

## 10. Implementação prática atual

A primeira versão deste diretório escolhe a alternativa mais simples: um canal de conclusão.

```text
New()
  └─ cria done e inicia uma goroutine

executor
  ├─ resolve(valor) -> guarda o valor e fecha done
  └─ reject(erro)   -> guarda o erro e fecha done

Await()
  └─ espera <-done e devolve valor/erro
```

O canal `done` é fechado, em vez de receber o valor diretamente. Isso permite que vários consumidores aguardem a mesma Promise, inclusive consumidores registrados depois que ela terminou. `sync.Once` garante que apenas a primeira chamada de `resolve` ou `reject` vença.

Exemplo mínimo:

```go
p := promise.New(func(resolve func(any), reject func(error)) {
	go buscarUsuario(resolve, reject)
})

resultado := p.
	Then(func(valor any) any {
		usuario := valor.(Usuario)
		return usuario.Nome
	}).
	Catch(func(err error) any {
		fmt.Println("erro:", err)
		return "desconhecido"
	})

nome, err := resultado.Await()
if err != nil {
	panic(err)
}

fmt.Println(nome)
```

A implementação está em [`promise.go`](./promise.go) e os casos cobertos estão em [`promise_test.go`](./promise_test.go). Ela é propositalmente pequena: não implementa `Promise.all`, cancelamento ou adoção de outra Promise retornada por um callback. Esses pontos ficam para uma evolução posterior.

Para executar os testes a partir da raiz deste repositório:

```bash
GO111MODULE=off go test -v
```

## 11. Conclusão

A brincadeira é possível e é um bom exercício: Go consegue representar uma operação futura, notificar sucesso ou erro, encadear transformações e aguardar várias operações.

O ponto mais importante é não confundir a aparência da API com o mecanismo por baixo dela. No JavaScript, Promise, microtasks e event loop trabalham juntos em um runtime orientado a eventos. No Go, a base é formada por goroutines, canais, scheduler, sincronização e contexto.

Portanto, a implementação em Go deve ser apresentada como uma **Promise inspirada no JavaScript**, útil para estudo e experimentação. Ela pode reproduzir muitos comportamentos observáveis, mas não deve esconder as diferenças fundamentais entre concorrência em Go e assincronicidade no JavaScript.
