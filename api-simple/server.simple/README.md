# Example MyServerExample

## Executar

```bash

$ go run main.go

```

## Makefile

```bash

$ make build

```

### /ping
```bash
$ curl -i -XPOST \
-H "X-Key-Token: your-key-here" \
localhost:8080/ping 

```

### /robo/relatorio/
```bash

$ curl -i -XGET \
-H "X-Key-Token: your-key-here" \
-H "Content-Type: application/json" \
localhost:8080/robo/relatorio/1 

```