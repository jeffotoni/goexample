# API Growth

Um simples exemplo de uma API executando um CRUD porém ele grava os dados em MEMÓRIA 
legal não é 😍?. O objetivo é entender a construção de uma API Go usando somente a strand library. 
A imagem gerada deste projeto não poderá passar de 6Mb isto mesmo tem que 
ser pequena e executar o mais rápido possível 😍.

Abaixo os comandos para compilar e executar o growth você mesmo.

O arquivo json utilizado tem 3Mb mais de 40k de registros e sua estrutura é um vetor com os seguintes campos:
```bash
[
   {
      "Country":"BRZ",
      "Indicator":"NGDP_R",
      "Value":183.26,
      "Year":2002
   },
   {
      "Country":"AFG",
      "Indicator":"NGDP_R",
      "Value":198.736,
      "Year":2003
   }
]
```
O json acima irá possuir mais de 40k de grupos de Growth.
Vamos armazenar este nosso coleguinha e memória e brincar com ele fazendo GET, PUT, DELETE.

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
