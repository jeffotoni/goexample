# Google Cloud

Este repositório são exemplos práticos e os primeiros passos ao usar Google cloud.

Antes de avançarmos é preciso criar uma conta e ativa-la para que possa brincar com todas as funcionalidades do Google Cloud.

Até o momento o que foi disponibilizado e testado foi:

	 - go.app 
	 Este é uma api em Go que foi criada para testar e aplicar no GKE, e neste repo contém o passo a passo de como foi feito.
	 Docker, deploy, kubectl etc.

	 - pub.go
	 Este é um publish do Pub/Sub do Google, ele envia mensagens para tópicos para o Pub/Sub.

	 - sub.go
	 Este é o consumer o nosso subscription, objetivo é consumier os tópicos.

	 - cloud functions
	 Este é o nosso Cloud Function, ele está recebendo um HTTP REQUEST e publicando os dados em um Pub/Sub.


### Exemplos do sdk do Google

Aqui vc irá encontrar diversos exemplos em Go que são bem interessantes para iniciar e dar os primeiros passos.

### Cloud Functions

Aqui está um link para que possa da uma conferida passo a passo de algo básico para que possa fazer seu start quando for trabalhar com cloud functions.

[cloud-functions](https://codelabs.developers.google.com/codelabs/cloud-functions-go-http/#1)

```bash
$ curl -LO https://github.com/GoogleCloudPlatform/golang-samples/archive/master.zip
```


