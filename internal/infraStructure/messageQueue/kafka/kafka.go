package kafka

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

func Producer() {
	// to produce messages
	//topic := "fzb4lrgg-"
	//partition := 0

	// EOF hata tls.config den dolayı
	mechanism, err := scram.Mechanism(scram.SHA256, "fzb4lrgg", "UlOeCHTz4pYCSG_w2FsH844Up4I1PUgb")
	if err != nil {

	}
	dialer := &kafka.Dialer{
		TLS:           &tls.Config{},
		SASLMechanism: mechanism,
		DualStack:     true,
	}
	w := kafka.NewWriter(kafka.WriterConfig{Brokers: []string{"glider-01.srvs.cloudkafka.com:9094",
		"glider-02.srvs.cloudkafka.com:9094",
		"glider-03.srvs.cloudkafka.com:9094"}, Dialer: dialer})

	writeError := w.WriteMessages(context.Background(), kafka.Message{Topic: "ds", Partition: 0, Key: []byte("key-a"), Value: []byte("merhaba")})
	if writeError != nil {
		fmt.Println("Write Error :", writeError.Error())
	}
	fmt.Println("Success")
}
