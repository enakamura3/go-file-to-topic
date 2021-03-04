# go file reader -> send to topic

- [go file reader -> send to topic](#go-file-reader---send-to-topic)
  - [About](#about)
  - [Getting Started](#getting-started)
  - [File reader](#file-reader)
  - [Kafka Producer](#kafka-producer)
  - [Kafka](#kafka)

## About

Go lang project to read content from a CSV file and post a message on Kafka Topic

## Getting Started

1. Configure [Kafka](#kafka)
2. Export variables:

```sh
export FILE_NAME=$(pwd)/resources/username.csv
export TOPIC=USERS
export KAFKA_ADDRESS=localhost:9092
```
3. Execute: 

```sh
go run main.go

2021/03/04 16:56:18 total registers read: 5
2021/03/04 16:56:18 connected to kafka server
2021/03/04 16:56:18 message sent: {booker12 9012 Rachel Booker}
2021/03/04 16:56:18 message sent: {grey07 2070 Laura Grey}
2021/03/04 16:56:18 message sent: {johnson81 4081 Craig Johnson}
2021/03/04 16:56:18 message sent: {jenkins46 9346 Mary Jenkins}
2021/03/04 16:56:18 message sent: {smith79 5079 Jamie Smith}
2021/03/04 16:56:18 Successfully produced record to topic USERS partition [2] @ offset 4
2021/03/04 16:56:18 Successfully produced record to topic USERS partition [2] @ offset 5
2021/03/04 16:56:18 Successfully produced record to topic USERS partition [0] @ offset 6
2021/03/04 16:56:18 Successfully produced record to topic USERS partition [0] @ offset 7
2021/03/04 16:56:18 Successfully produced record to topic USERS partition [0] @ offset 8
```

## File reader

Article to read:

> https://medium.com/swlh/processing-16gb-file-in-seconds-go-lang-3982c235dfa2


## Kafka Producer

Golang kafka producer example:

> https://github.com/confluentinc/examples/blob/6.1.0-post/clients/cloud/go/producer.go

Confluent document:

> https://docs.confluent.io/platform/current/tutorials/examples/clients/docs/go.html

## Kafka

We will use Kafka on docker using:

> https://github.com/wurstmeister/kafka-docker

Follow the instructions on:

> https://github.com/wurstmeister/kafka-docker#pre-requisites

1. Install docker-compose
2. On `docker-compose.yml` file, change the ip address in `KAFKA_ADVERTISED_HOST_NAME` variable  
3. On `docker-compose.yml` file, set a local port __9092__ on kafka container

```yaml
    ports:
      - "9092:9092"
```

4. Start: 

```sh
docker-compose up -d
```

Create a topic:

```sh
docker exec -ti kafka-docker_kafka_1 /opt/kafka/bin/kafka-topics.sh --bootstrap-server localhost:9092 --create --topic USERS --partitions 3 --replication-factor 1

Created topic USERS.
```

List topics:

```sh
docker exec -ti kafka-docker_kafka_1 /opt/kafka/bin/kafka-topics.sh --bootstrap-server localhost:9092 --list

USERS
```

See topic details:

```sh
docker exec -ti kafka-docker_kafka_1 /opt/kafka/bin/kafka-topics.sh --bootstrap-server localhost:9092 --topic USERS --describe

Topic: USERS    PartitionCount: 3       ReplicationFactor: 1    Configs: segment.bytes=1073741824
        Topic: USERS    Partition: 0    Leader: 1001    Replicas: 1001  Isr: 1001
        Topic: USERS    Partition: 1    Leader: 1001    Replicas: 1001  Isr: 1001
        Topic: USERS    Partition: 2    Leader: 1001    Replicas: 1001  Isr: 1001
```

