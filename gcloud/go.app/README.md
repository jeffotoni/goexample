# Go.app example

O objetivo deste repo é somente exemplificar e colocar em prática como subir um app feito em  Go no GKE(google Kubernetes).

Então todo nosso material você irá precisar possuir uma conta no console.cloud.google.com.

Abaixo tem os passos de como vamos subir um app feito em Go para o GKE do google.

 - Podemos criar nosso GKE pelo console cloud.google ou pelo gclood usando linha de comando e o arquivo é o create.cluster.gke.sh.

 - Pode criar nosso GKE pela api e temos um arquivo com a chamada da api ele é o create.cluster.gke.api.sh.


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

### Credentials google
```bash
$ gcloud container clusters get-credentials go-app --zone us-central1-c --project projeto-eng1
```

## GKE

Para subirmos nosso projeteo para GKE precisamos criar nosso cluster
Antes precisaremos instalar nosso kubectl e utiliza-lo em conjunto com gcloud.


```bash
$ kubectl info
```
