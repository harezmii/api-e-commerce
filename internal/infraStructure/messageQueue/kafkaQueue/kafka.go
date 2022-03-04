package kafkaQueue

import (
	"api/internal/entity"
	"api/pkg/config"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewKafkaProducer() entity.Kafka {
	cfg := config.GetConf()
	k := entity.Kafka{
		Topic:  cfg.Kafka.Topic,
		Config: cfg.Kafka.Conf,
	}

	return k
}
func Producer(key string, message string) {
	k := NewKafkaProducer()
	c := k.Config

	/*
			The \\ sign is used because of an error when unmarshaling with viper from the config file.
		    The resulting error is due to the .(dot) sign.
	*/

	configKafka := &kafka.ConfigMap{
		"bootstrap.servers": c["bootstrap\\servers"],
		"security.protocol": c["security\\protocol"],
		"sasl.mechanisms":   c["sasl\\mechanisms"],
		"sasl.username":     c["sasl\\username"],
		"sasl.password":     c["sasl\\password"],
		"client.id":         c["client\\id"],
	}
	producer, producerError := kafka.NewProducer(configKafka)
	if producerError != nil {
		fmt.Println("Producer Error")
	}
	topic := k.Topic
	deliveryChan := make(chan kafka.Event)
	writeError := producer.Produce(&kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny}, Key: []byte(key), Value: []byte(message)}, deliveryChan)
	if writeError != nil {
		fmt.Println("Write Error")
	}
	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	}
	close(deliveryChan)
}
