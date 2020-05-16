## Resumo kafka

Vamos usar o confluent como plataforma e subir o kafka,zookeeper, kafka broker, Registry, ksql e rest-proxy.


Aqui todo o serviço que encontra-se em nosso docker-compose.yaml irá baixar as imagens e fazer seu start.
```bash

$ docker-compose up --build

$ docker-compose ps

```

Agora que o seriço está rodando podemos ir no browser e acessar [localhost:9021](http://localhost:9021)
Neste ambiente poderemos visualizar todo arsenal que o kafka disponibiliza de forma visual.

Caso queira usar o kafka bash basta instalar usando o comando abaixo.

Poderá encontrar aqui: [kafka-shell](https://github.com/devshawn/kafka-shell).

## Usando kafka-shell

```bash

$ pip3 install kafka-shell

```

Em nosso exemplo estamos usando a plataforma da confluent, e em seus brokers e zookeeper já possuem o kafka-sell, vou apresentar abaixo como executa-los.
Como estamos usando docker-compose para subir todo serviço do kafka iremos usar docker-compose exec ou docker exec.


### Criando Topico
```bash

$ docker-compose exec broker kafka-topics --create --topic mytopic-test-2020 --partitions 1 --replication-factor 1 --if-not-exists --zookeeper zookeeper:2181
Created topic meu-topico-legal-jeff.

```

### Listando todos meus topicos
```bash

$ docker-compose exec broker kafka-topics --list --zookeeper zookeeper:2181

```

### Describe um tópico 
```bash

$ docker-compose exec broker kafka-topics --describe meu-topico-legal-jeff --zookeeper zookeeper:2181

```

### Producer mensanges

```bash

$ docker-compose exec broker  \
  bash -c "seq 100 | kafka-console-producer --request-required-acks 1 --broker-list localhost:9092 --topic meu-topico-legal-jeff && echo 'Produced 100 messages.'"

```

### Consumer mensagens

```bash

$ docker-compose exec broker  \
  kafka-console-consumer --bootstrap-server localhost:9092 --topic meu-topico-legal-jeff --from-beginning --max-messages 100

```

## ZOOKEEPER

O Zookeeper é um sistema centralizador e de gerenciamento para qualquer tipo de sistema distribuído. Sistema distribuído são diferentes módulos de software executando em diferentes nós / clusters (podem estar em locais geograficamente distantes), mas executando como um sistema. O Zookeeper facilita a comunicação entre os nós, compartilhando configurações entre os nós, mantém o controle de qual nó é líder, qual nó se junta / sai etc. O Zookeeper é quem mantém os sistemas distribuídos sãos e mantém a consistência. O Zookeeper é basicamente uma plataforma de orquestração.

O Zookeeper em si é um sistema distribuído que consiste em vários nós em um conjunto. O Zookeeper é um serviço centralizado para manter esses metadados.

O Zookeeper também desempenha um papel vital para servir a muitos outros propósitos, como detecção de líder, gerenciamento de configuração, sincronização, detecção de quando um novo nó entra ou sai do cluster, etc.


Kafka usa o Zookeeper para o seguinte:

### Elegendo um controlador
O controlador é um dos intermediários e é responsável por manter o relacionamento líder / seguidor para todas as partições. Quando um nó é desligado, é o controlador que instrui outras réplicas a se tornarem líderes de partição para substituir os líderes de partição no nó que está desaparecendo. O Zookeeper é usado para eleger um controlador, verifique se há apenas um e escolha um novo se ele travar.

### Associação ao cluster
Quais corretores estão ativos e fazem parte do cluster? isso também é gerenciado através do ZooKeeper.

### Configuração de tópico
Quais tópicos existem, quantas partições cada um possui, onde estão as réplicas, quem é o líder preferencial, quais substituições de configuração são definidas para cada tópico

Cotas - quantos dados cada cliente tem permissão para ler e gravar

ACLs - quem tem permissão para ler e gravar em qual tópico (consumidor antigo de alto nível) - Quais grupos de consumidores existem, quem são seus membros e qual é o último deslocamento que cada grupo obteve de cada partição.


## Broker

Um broker é o componente responsável por receber as requisições de producers e consumers, armazenar as mensagens e executar a replicação das mesmas.
Os brokers são gerenciados por outro componente o zookeeper. Este componente é bastante utilizado para controlar os diferentes integrantes de um cluster.
Além das tarefas descritas acima, os brokers também realizam outras tarefas, como gerenciar os líderes de cada partição, realizar a limpeza de dados ou a compactação das mensagens.
Pretendo escrever em detalhes cada um destes tópicos avançados.


## Log

Um log pode ser descrito como uma sequência temporal de mensagens, onde as novas mensagens sempre são adicionadas no final do log. Desta forma, uma mensagem enviada em t0 sempre estará posicionada antes de uma mensagem enviada em t1.

Cada mensagem dentro do log possui algumas informações:

1. Timestamp: data-hora da inserção
2. Offset: índice da mensagem na partição
3. Key: chave da mensagem
4. Value: a mensagem propriamente dita chamado de payload

Todas as mensagens dentro de uma partição serão um conjunto chave / valor.


## Partições

A primeira coisa a entender é que uma partição de tópico é a unidade de paralelismo em Kafka.

Kafka sempre fornece os dados de uma única partição para um thread do consumidor. Assim, o grau de paralelismo no consumidor (dentro de um grupo de consumidores) é limitado pelo número de partições sendo consumidas. Portanto, em geral, quanto mais partições houver em um cluster Kafka, maior será a taxa de transferência possível.


## Qual é o numero de Partições que deveriamos criar para nosso cenário?

Uma fórmula aproximada para escolher o número de partições é baseada na taxa de transferência. Você mede o tempo todo que pode obter em uma única partição para produção (chame de p ) e consumo (chame de c ). 

MAX(t/p, t/c)

t: taxa de transferência desejada
p: taxa de transferência do producer
c: taxa de transferência do consumer

Embora seja possível aumentar o número de partições ao longo do tempo, é preciso ter cuidado se as mensagens forem produzidas com chaves. Ao publicar uma mensagem com chave, o Kafka mapeia deterministicamente a mensagem para uma partição com base no hash da chave. Isso garante que as mensagens com a mesma chave sejam sempre roteadas para a mesma partição.

Em geral, mais partições em um cluster Kafka levam a uma taxa de transferência mais alta. No entanto, é preciso estar ciente do impacto potencial de ter muitas partições no total ou por broker em coisas como disponibilidade e latência. 


## PRODUCERS

### Acks = 0
o produtor não aguarda nenhum tipo de resposta do cluster. É o modo com throughput mais elevado. É importante levar em conta que nesse modo a perda de dados é possível uma vez que o produtor não aguarda nenhum tipo de sinal do cluster.

### Acks = 1
, o produtor aguarda por um ok do líder da partição. Sendo assim sabemos que ao menos 1 broker recebeu a mensagem. Já é uma configuração bem mais segura que Acks=0, mas não é 100% segura uma vez que o broker líder pode cair antes que a replicação seja realizada, e o produtor não seria notificado nesse cenário.

### Acks = -1
O produtor aguarda o retorno até que o líder e todas as réplicas recebam a mensagem. É o modo mais seguro, 


## CONSUMER

Não adianta ter mais consumidores do que partições. Caso o grupo 1 possuísse 5 consumidores, 3 deles ficariam ociosos pois o Kafka não conseguiria mandar mensagens de uma mesma partição à mais de um consumidor do mesmo grupo.


## Estratégias de commit de offsets

### No máximo uma vez
Neste modo, o consumidor commita o offset para o Kafka assim que recebe a mensagem.
Mensagens podem ser perdidas, mas nunca processadas com duplicação.

### Pelo menos uma vez
O offset é commitado após o processamento da mensagem
Mensagens nunca serão perdidas, mas podem ser processadas com duplicação.

### Exatamente uma vez
Uma mensagem tem a garantia de ser enviada uma única vez para um determinado consumidor.


## bash curl


### List Info Topics
```bash

$ curl "http://localhost:8082/topics" | jq

```

### List Info Topic específico
```bash

$ curl "http://localhost:8082/topics/topicgo1" | jq

```

### List Info Partitions Topic
```bash

$ curl "http://localhost:8082/topics/topicgo1/partitions" | jq

```

### Produce JSON Menssage

```bash
$ curl -X POST -H "Content-Type: application/vnd.kafka.json.v2+json" \
      -H "Accept: application/vnd.kafka.v2+json" \
      --data '{"records":[{"value":{"msg":"success 4"}}]}' "http://localhost:8082/topics/topicgo1"

```

### Create a Consumer

```bash
$ curl -X POST -H "Content-Type: application/vnd.kafka.v2+json" \
      --data '{"name": "my_consumer_instance", "format": "json", "auto.offset.reset": "earliest"}' \
      http://localhost:8082/consumers/go_json_consumer

```

# Out
```json
  {
  	"instance_id":"my_consumer_instance",
	"base_uri":"http://rest-proxy:8082/consumers/go_json_consumer/instances/my_consumer_instance"
  } 
```

### Subscription Consumer
```bash
$ curl -X POST -H "Content-Type: application/vnd.kafka.v2+json" --data '{"topics":["topicgo1"]}' \
 http://localhost:8082/consumers/go_json_consumer/instances/my_consumer_instance/subscription
 ```

### Consume JSON Menssage

```bash
$ curl -X GET -H "Accept: application/vnd.kafka.json.v2+json" \
  http://localhost:8082/consumers/go_json_consumer/instances/my_consumer_instance/records

```

### Delete
```bash
curl -X DELETE -H "Content-Type: application/vnd.kafka.v2+json" \
      http://localhost:8082/consumers/go_json_consumer/instances/my_consumer_instance

```