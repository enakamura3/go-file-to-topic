package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/enakamura3/go-file-to-topic/models"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {
	FILENAME := os.Getenv("FILE_NAME")

	file, err := os.Open(FILENAME)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	users := []models.User{}

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count++
		if count == 1 {
			continue // skip first line
		}
		line := scanner.Text()
		fields := strings.Split(line, ";")
		if len(fields) > 1 {
			user := models.User{Username: fields[0], Identifier: fields[1], FirstName: fields[2], LastName: fields[3]}
			users = append(users, user)
		}
	}
	log.Println("total registers read:", count-1)
	sendMessages(&users)
}

func sendMessages(users *[]models.User) {

	TOPIC := os.Getenv("TOPIC")
	KAFKAADDRESS := os.Getenv("KAFKA_ADDRESS")

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": KAFKAADDRESS})
	if err != nil {
		log.Panic(err)
	}
	defer p.Close()
	log.Println("connected to kafka server")

	// Go-routine to handle message delivery reports and
	// possibly other event types (errors, stats, etc)
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					log.Printf("Successfully produced record to topic %s partition [%d] @ offset %v\n",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
				}
			}
		}
	}()

	for _, user := range *users {
		userm, err := json.Marshal(&user)
		if err != nil {
			log.Fatal(err)
		}
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &TOPIC, Partition: kafka.PartitionAny},
			Key:            []byte(user.Identifier),
			Value:          userm,
		}, nil)
		log.Println("message sent:", user)
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}
