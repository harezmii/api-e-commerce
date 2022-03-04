package kafkaQueue

import (
	"api/internal/entity"
	"api/pkg/config"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewKafkaProducer() entity.Kafka {
	k := entity.Kafka{
		Topic:  config.GetEnvironment("topic", config.STRING).(string),
		Config: config.GetEnvironment("config", config.MAP).(map[string]interface{}),
	}

	return k
}
func Producer(key string, message string) {
	k := NewKafkaProducer()
	c := k.Config

	configKafka := &kafka.ConfigMap{
		"bootstrap.servers": c["bootstrap.servers"].(string),
		"security.protocol": c["security.protocol"].(string),
		"sasl.mechanisms":   c["sasl.mechanisms"].(string),
		"sasl.username":     c["sasl.username"].(string),
		"sasl.password":     c["sasl.password"].(string),
		"client.id":         c["client.id"].(string),
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
