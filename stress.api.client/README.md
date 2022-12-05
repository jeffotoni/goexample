# Spring Boot vs Go

Projeto para observarmos e analisarmos a execução do Spring Boot vs Go.
Você poderá utilizar docker, docker-compose para subir as aplicações locais se desejar é uma possibilidade.


Para conseguir rodar e executar os serviços terá que instalar:

	- Java
	- Maven
	- springboot
	- Go
	- k6

**Neste repo irá encontrar 4 serviços:**
	- Um mock (um server que retorna json) onde é um server que responde na porta 3000 e retorna um json
	- Um server feito em Go
 	- Um server feito em Spring Boot
 	- k6 serviço que irá fazer nosso test de stress

Nosso objetivo é analisar a execução dos servers e capturar o seu consumo de memória, cpu, requests etc.

### Mock server

No diretório mock.server teremos um server feito em Go que irá ser responsável por responder na porta 3000 e retornar um JSON.

```bash
$ go run main.go
Run Server Mock 0.0.0.0:3000
[GET] /v1/customer

```

```bash
$ curl -i -XGET http://localhost:3000/v1/customer
HTTP/1.1 201 
Location: /v1/client
Engine: Spring Boot
Content-Type: application/json
Content-Length: 3616
Date: Sat, 03 Dec 2022 14:39:02 GMT

[
    {
        "createdAt": "2022-11-04T19:10:17.305Z",
        "name": "BillyStoltenberg",
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1164.jpg",
        "id": "1"
    },
    {
        "createdAt": "2022-11-05T09:01:38.207Z",
        "name": "JodiKertzmann",
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/502.jpg",
        "id": "2"
    },
   ....
]
```


### Spring Boot executar

Neste diretório, irá encontrar nosso serve feito em Spring Boot para receber uma requisição e comunicar com mock.

Com spring boot podemos executar utilizando run do mvn ou gerar um JAR com mvn package.

```bash
$ cd spring.boot
$ mvn spring-boot:run
[INFO] --- spring-boot-maven-plugin:2.7.5:run (default-cli) @ sclient ---
[INFO] Attaching agents: []

  .   ____          _            __ _ _
 /\\ / ___'_ __ _ _(_)_ __  __ _ \ \ \ \
( ( )\___ | '_ | '_| | '_ \/ _` | \ \ \ \
 \\/  ___)| |_)| | | | | || (_| |  ) ) ) )
  '  |____| .__|_| |_|_| |_\__, | / / / /
 =========|_|==============|___/=/_/_/_/
 :: Spring Boot ::                (v2.7.5)

2022-12-03 11:36:34.604  INFO 1923713 --- [           main] c.jeffotoni.sclient.SclientApplication   : Starting SclientApplication using Java 17.0.1 on jeffotoni with PID 1923713 (/spring.boot/target/classes started by jeffotoni in /stress.api.client/spring.boot)
2022-12-03 11:36:34.605  INFO 1923713 --- [           main] c.jeffotoni.sclient.SclientApplication   : No active profile set, falling back to 1 default profile: "default"
2022-12-03 11:36:35.046  INFO 1923713 --- [           main] o.s.b.w.embedded.tomcat.TomcatWebServer  : Tomcat initialized with port(s): 8080 (http)

```
Ou podem gerar o **JAR.**

```bash
$ cd spring.boot
$ mvn package
INFO] Results:
[INFO] 
[INFO] Tests run: 1, Failures: 0, Errors: 0, Skipped: 0
[INFO] 
[INFO] 
[INFO] --- maven-jar-plugin:3.2.2:jar (default-jar) @ sclient ---
[INFO] 
[INFO] --- spring-boot-maven-plugin:2.7.5:repackage (repackage) @ sclient ---
[INFO] Replacing main artifact with repackaged archive
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time:  3.283 s
[INFO] Finished at: 2022-12-03T11:56:57-03:00
[INFO] ------------------------------------------------------------------------
```

```bash
$ java -jar target/sclient-0.0.1-SNAPSHOT.jar
[INFO] --- spring-boot-maven-plugin:2.7.5:run (default-cli) @ sclient ---
[INFO] Attaching agents: []

  .   ____          _            __ _ _
 /\\ / ___'_ __ _ _(_)_ __  __ _ \ \ \ \
( ( )\___ | '_ | '_| | '_ \/ _` | \ \ \ \
 \\/  ___)| |_)| | | | | || (_| |  ) ) ) )
  '  |____| .__|_| |_|_| |_\__, | / / / /
 =========|_|==============|___/=/_/_/_/
 :: Spring Boot ::                (v2.7.5)

2022-12-03 11:36:34.604  INFO 1923713 --- [           main] c.jeffotoni.sclient.SclientApplication   : Starting SclientApplication using Java 17.0.1 on jeffotoni with PID 1923713 (/spring.boot/target/classes started by jeffotoni in /stress.api.client/spring.boot)
2022-12-03 11:36:34.605  INFO 1923713 --- [           main] c.jeffotoni.sclient.SclientApplication   : No active profile set, falling back to 1 default profile: "default"
2022-12-03 11:36:35.046  INFO 1923713 --- [           main] o.s.b.w.embedded.tomcat.TomcatWebServer  : Tomcat initialized with port(s): 8080 (http)

```

```bash
$ curl -i -XGET http://localhost:8080/v1/client
HTTP/1.1 201 
Location: /v1/client
Engine: Spring Boot
Content-Type: application/json
Content-Length: 3616
Date: Sat, 03 Dec 2022 14:39:02 GMT

[
    {
        "createdAt": "2022-11-04T19:10:17.305Z",
        "name": "BillyStoltenberg",
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1164.jpg",
        "id": "1"
    },
    {
        "createdAt": "2022-11-05T09:01:38.207Z",
        "name": "JodiKertzmann",
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/502.jpg",
        "id": "2"
    },
   ....
]
```

### go.http.client server

No diretório go.http.client teremos um server feito em Go que irá ser responsável por comunicar com mock server, fazer uma chamada REST.

```bash
$ go run main.go
Run Server port 0.0.0.0:8080
[GET] /v1/client

```

```bash
$ curl -i -XGET http://localhost:8080/v1/client
HTTP/1.1 200 OK
Content-Length: 3616
Content-Type: application/json
Date: 2022-12-03T12:50:19.310Z
Engine: Go
Location: /v1/client

[
    {
        "createdAt": "2022-11-04T19:10:17.305Z",
        "name": "BillyStoltenberg",
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1164.jpg",
        "id": "1"
    },
    {
        "createdAt": "2022-11-05T09:01:38.207Z",
        "name": "JodiKertzmann",
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/502.jpg",
        "id": "2"
    },
   ....
]
```

### k6 run

O K6 é um programa para fazermos test de stress, vamos usar o k6 para vermos o resultado de ambos os serviços.

Precisa instalar o k6 e precisamos criar nosso script para que o k6 possa executar o que gostariamos.

```bash
$ cd k6
$ k6 run -d 90s -u 100 script.js

          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: script.js
     output: -

  scenarios: (100.00%) 1 scenario, 100 max VUs, 2m0s max duration (incl. graceful stop):
           * default: 100 looping VUs for 1m30s (gracefulStop: 30s)


running (1m30.0s), 000/100 VUs, 1864368 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  1m30s

     data_received..................: 7.0 GB  78 MB/s
     data_sent......................: 226 MB  2.5 MB/s
     http_req_blocked...............: avg=2.16µs  min=665ns    med=1.53µs  max=27.63ms p(90)=2.02µs  p(95)=2.3µs  
     http_req_connecting............: avg=103ns   min=0s       med=0s      max=26.27ms p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=4.75ms  min=124.11µs med=3.81ms  max=72.82ms p(90)=9.73ms  p(95)=11.95ms
       { expected_response:true }...: avg=4.75ms  min=124.11µs med=3.81ms  max=72.82ms p(90)=9.73ms  p(95)=11.95ms
     http_req_failed................: 0.00%   ✓ 0            ✗ 1864368
     http_req_receiving.............: avg=31.13µs min=8.19µs   med=20.47µs max=13.57ms p(90)=29.98µs p(95)=38.06µs
     http_req_sending...............: avg=10.33µs min=3.42µs   med=7.51µs  max=27.68ms p(90)=9.36µs  p(95)=10.66µs
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s      p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=4.71ms  min=109.5µs  med=3.77ms  max=72.58ms p(90)=9.68ms  p(95)=11.89ms
     http_reqs......................: 1864368 20714.381246/s
     iteration_duration.............: avg=4.81ms  min=144.48µs med=3.87ms  max=73.08ms p(90)=9.79ms  p(95)=12.01ms
     iterations.....................: 1864368 20714.381246/s
     vus............................: 100     min=100        max=100  
     vus_max........................: 100     min=100        max=100  
```

### Usando docker-compose

Podemos também usar o docker-compose ou docker se desejar, já existe o arquivo Dockerfile para você brincar.

Eu criei as imagens e subir para um registry e deixei público para facilitar, mas fiquem a vontade em fazer da forma que achar melhor.

**docker-compose Go**

Foi utilizado para gerar a imagem Go o seguinte docker build:
```bash
$ docker build --no-cache -t public.ecr.aws/v5p3h8r5/gserver .
$ docker run --name gserver -p 8080:8080 --rm public.ecr.aws/v5p3h8r5/gserver:latest
```

```bash
$ cd go.http.client
$ docker-compose up
reating gserver     ... done
Creating gmockserver ... done
Attaching to gmockserver, gserver
gmockserver    | 2022/12/03 18:32:43 Run Server Mock 0.0.0.0:3000
gmockserver    | 2022/12/03 18:32:43 [GET] /v1/customer
gserver        | 2022/12/03 18:32:43 Run Server port 0.0.0.0:8080
gserver        | 2022/12/03 18:32:43 [GET] /v1/client

```

**Vamos testar**
```bash
$ curl -i -XGET http://localhost:8080/v1/client

```

**docker-compose Spring Boot**

Foi utilizado para gerar a imagem Spring Boot o seguinte docker build:
```bash
$ docker build -t public.ecr.aws/v5p3h8r5/jserver
$ docker run --name jserver -p 8080:8080 --rm public.ecr.aws/v5p3h8r5/jserver:latest
```

```bash
$ cd spring.boot
$ docker-compose up
➜  spring.boot git:(master) ✗ docker-compose up
jserver is up-to-date
Starting gmockserver ... done
Attaching to jserver, gmockserver
gmockserver    | 2022/12/05 22:40:01 Run Server Mock 0.0.0.0:3000
gmockserver    | 2022/12/05 22:40:01 [GET] /v1/customer
jserver        | 
jserver        |   .   ____          _            __ _ _
jserver        |  /\\ / ___'_ __ _ _(_)_ __  __ _ \ \ \ \
jserver        | ( ( )\___ | '_ | '_| | '_ \/ _` | \ \ \ \
jserver        |  \\/  ___)| |_)| | | | | || (_| |  ) ) ) )
jserver        |   '  |____| .__|_| |_|_| |_\__, | / / / /
jserver        |  =========|_|==============|___/=/_/_/_/
jserver        |  :: Spring Boot ::                (v2.7.5)
jserver        | 

```

**Vamos testar**
```bash
$ curl -i -XGET http://localhost:8080/v1/client
```