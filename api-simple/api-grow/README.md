# API Growth

Um simples exemplo de uma API de um CRUD porém que grava seus dados em MEMÓRIA. O objetivo 
é enender a construção de uma API utilizando Go. A imagem gerada deste projeto não poderá 
passar de 6Mb isto mesmo tem que ser pequena e o mais rápida possível 😍.

Abaixo os comandos para compilar e executar o growth você mesmo.

### Docker Build

```bash
$ docker build --no-cache -f Dockerfile -t jeffotoni/apigrow:latest .
```
Depois de fazer build do projeto vamos conferir sua dimensão

```bash
$ docker images | grep jeffotoni
jeffotoni/apigrow  latest  c931a510e393  10 minutes ago   4.94MB
```
Prontinho, agora vamos executar e testar nossa apigrow ❤️

```bash
$ docker run --rm -it -p 8080:8080 jeffotoni/apigrow
```
Agora vamos testar nossa API 🦾

#### POST
```bash
$ curl -i -XPOST localhost:8080/api/v1/growth -d @1mb-growth_json.json
{"msg":"In progress"}
```
#### GET
```bash
$ curl -i -XGET localhost:8080/api/v1/growth/post/status
{"msg":"complete","test value"":183.26, "count":42450}
```
#### GET
```bash
$ curl -i -XGET localhost:8080/api/v1/growth/brz/ngdp_r/2002
{"Country":"BRZ","Indicator":"NGDP_R","Value":183.26,"Year":2002}
```
#### PUT
```bash
$ curl -i -XPUT localhost:8080/api/v1/growth/brz/ngdp_r/2002 \
-d '{"value":333.98}'
```
#### GET
```bash
$ curl -i -XGET localhost:8080/api/v1/growth/brz/ngdp_r/2002
{"Country":"BRZ","Indicator":"NGDP_R","Value":333.98,"Year":2002}
```
#### DELETE
```bash
$ curl -i -XPUT localhost:8080/api/v1/growth/brz/ngdp_r/2002 
```
