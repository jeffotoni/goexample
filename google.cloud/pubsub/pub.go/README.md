# Pub/Sub Google Cloud

Este exemplos é utilizando Pub/Sub do Google, ele publica uma mensagem em um tópico.

Para este módulo funcionar você precisa liberar suas credenciais, e existe 3 formas de fazer isto. Sugiro que crie um novo usuário IAM com permissões que deseja e baixe as credenciais para usar com a variavel de ambiente GOOGLE_APPLICATION_CREDENTIALS onde vc irá apontar o path exemplo:

```bash
$ export GOOGLE_APPLICATION_CREDENTIALS="/path/credentials.json"
```

Qualquer dúvida pode visitar aqui [authentication](https://cloud.google.com/docs/authentication/production#command-line)


