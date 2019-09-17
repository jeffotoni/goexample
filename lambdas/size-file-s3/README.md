# Lambda 

Este lambda foi desenvolvido em Go então para que possa compilar você precisa instalar em sua máquina mais detalhes aqui: [golang.org](https://golang.org) é muito simples instalar e usa-lo.

Caso não queira instalar tem como fazer tudo via docker, [golang docker](https://github.com/jeffotoni/goworkshopdevops/blob/master/README.md#installation-docker)

Este lambda irá escutar uma bucket S3 receber os dados do arquivo e se o tamanho do arquivo não for o que você definiu ele excluir o arquivo do bucket.
Edita o sizeobject.go na linha 18 SIZE, lá basta subistituir o valor da variável para o tamanho que deseja em bytes.
Para tudo funcionar corretamente você precisa criar o lambda lá na AWS e depois criar a politica para acesso ao seu bucket pelo lambda, tem um arquivo [funcao-acl-politica-bucket.acl] de exemplo da política de acesso você pode utiliza-lo como exemplo e melhorar conforme sua necessidade.

Deixei um Makefile para ajudar, para para rodar make deploy você precisa do awscli instalado e permissão para executa-lo.

## make build
Somente build, não há necessidade de awscli, ele compila o programa e gera o zip para que você possa enviar para AWS LAMBDA manualmente.
Você precisa editar o Makefile e colocar o nome do lambda na variavel: LAMBDA_NOME e mais nada.
Detalhe você precisa do zip instalado em seu Linux.

```bash
$ make build
```

## make deploy
Este comando ele faz build e envia para S3 e depois atualiza sua função Lambda dinâmicamente, vc não precisa entrar no console aws para fazer isto este deploy faz isto para você.
Mas para isto funcionar corretamente vc precisa configurar suas credenciais corretamente em sua máquina e instalar o awscli.
Abre o Makefile e configura as variaveis LAMBADA_NOME e BUCKET_DEPLOY que é o nome do Bucket que deseja jogar o zip para fazer deploy.
Assim que tudo finalizar ele remove o arquivo zip e binário gerado de seu diretório.

```bash
$ make deploy
```

