package messaging

import "fmt"

type KafkaPublisher struct {
	broker string
}

func NewKafkaPublisher(broker string) *KafkaPublisher {
	fmt.Println("Initialising Kafka")
	return &KafkaPublisher{
		broker: broker,
	}
}

func (k *KafkaPublisher) Publish(content any) {
	fmt.Println("Publishing message ==>", content)
}

func (k *KafkaPublisher) Close() {
	fmt.Println("Closing down kafka")
}
