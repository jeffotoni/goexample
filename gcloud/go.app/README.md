# Go.app example

O objetivo deste repo é somente exemplificar e colocar em prática como subir um app feito em  Go no GKE(google Kubernetes).

Então todo nosso material você irá precisar possuir uma conta no console.cloud.google.com.

Abaixo tem os passos de como vamos subir um app feito em Go para o GKE do google.

 - Podemos criar nosso GKE pelo console cloud.google ou pelo gclood usando linha de comando e o arquivo é o create.cluster.gke.sh.

 - Pode criar nosso GKE pela api e temos um arquivo com a chamada da api ele é o create.cluster.gke.api.sh.

 - Vamos utilizar o Register do Google para armazenar nossa imagem de nossa API GO.


### Build go.app
```bash
$ CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w"
```

### Docker Build go.app Localmente
```bash
$ docker build --no-cache -f Dockerfile -t jeffotoni/go.app .
```

### Docker Build go.app Google
```bash
$ docker build --no-cache -f Dockerfile -t gcr.io/projeto-eng1/go.app:latest .
```

### Docker run go.app
```bash
$ docker run -p 8080:8080 --rm --name go.app jeffotoni/go.app
```

### Auth Configure-Docker Google Register

```bash
$ gcloud auth configure-docker
```

### Docker Push go.app Google
```bash
$ docker push gcr.io/projeto-eng1/go.app:latest
```

## GKE

Para subirmos nosso projeteo para GKE precisamos criar nosso cluster
Antes precisaremos instalar nosso kubectl e utiliza-lo em conjunto com gcloud.


### Criando cluster GKE

Se entrar dentro do arquivo create.cluster.gke.sh irá ver todos os comandos que usamos para criar um cluster utilizando gcloud.

Deixei como Default os seriços que iremos subir no Kubernetes.

```bash
sh create.cluster.gke.sh
```

### Set Credentials

Agora vamos setar nossas credentials do nosso app feito em Go.

```bash
$ gcloud container clusters get-credentials go-app --zone southamerica-east1-c --project projeto-eng1
```

### Implantar 

Pronto agora vamos implamntar e subir nosso app feito em Go no cluster.

```bash
$ kubectl apply -f deployment.yaml
```

Caso queira deletar nosso app do cluster

```bash
$ kubectl delete -f deployment.yaml
```

### Forward
Antes de expor a porta 80 e um IP público, vamos fazer um forward e testar nosso app feito em Go.

Lista seus pods e seleciona
```bash
$ kubectl get pods
```

Agora com o pod selecionado passa como parametro para port-forward

```bash
$ kubectl port-forward go.app-7d495cf6f7-4rzgm 8080
```
```bash
$ curl localhost:8080/ping
```

### Implantar Service Cargas de Trabalho Export IP

Agora vamos expor o IP publico para fazermos nossas chamadas na nossa api exemplo feita em Go.

```bash
$ kubectl apply -f gke.service.ip.yaml
```

Para testar podemos usar o hey um programa feito em Go equivalente ao wrk ou ab.

```bash
$ hey -z 20s -n 1000 -c 30 -m GET "http://localhost/ping"
```

Outra forma de testar e entrando na pasta mock, lá tem um client para fazer chamadas em seu app.Go, basta passar o -host

```bash
$ go run main <IP-REMOTE/LOCAL>
```