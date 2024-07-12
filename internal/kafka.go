package internal

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/IBM/sarama"
	"github.com/brianvoe/gofakeit/v7"
)

// Message is struct for message
type Message struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// Producer is struct for producers
type Producer struct {
	conn sarama.SyncProducer
}

// NewProducer return instance of producer
func NewProducer() *Producer {
	return &Producer{}
}

// ConnectProducer is connect in kafka
func (p *Producer) ConnectProducer(urls []string) error {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	conn, err := sarama.NewSyncProducer(urls, config)
	if err != nil {
		return err
	}
	p.conn = conn
	return nil
}

// ProduceMessage execut send the msg to kafka
func (p *Producer) ProduceMessage(topic string) error {
	defer p.conn.Close()

	ticker := time.NewTicker(time.Millisecond * 500)
	for {
		select {
		case <-ticker.C:
			message := Message{
				Name:  gofakeit.Name(),
				Email: gofakeit.Email(),
				Phone: gofakeit.Phone(),
			}
			bytesMessage, err := p.marshaler(message)
			if err != nil {
				return err
			}
			msg := &sarama.ProducerMessage{
				Topic: topic,
				Value: sarama.StringEncoder(bytesMessage),
			}
			partition, offset, err := p.conn.SendMessage(msg)
			if err != nil {
				return err
			}
			fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
		default:
			continue
		}
	}
}

// marshaler execute marshal for message
func (p *Producer) marshaler(message Message) ([]byte, error) {
	bytesMessage, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}
	return bytesMessage, nil
}

// Consumer is struct for consumers
type Consumer struct {
	conn sarama.Consumer
}

// NewConsumer return instance of consumer
func NewConsumer() *Consumer {
	return &Consumer{}
}

// ConnectConsumer is connect in kafka
func (c *Consumer) ConnectConsumer(urls []string) error {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	conn, err := sarama.NewConsumer(urls, config)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

// ConsumeMessage return messages from Kafka
func (c *Consumer) ConsumeMessage(topic string) error {
	cons, err := c.conn.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		return err
	}
	defer cons.Close()
	for msg := range cons.Messages() {
		var message Message
		if err := json.Unmarshal(msg.Value, &message); err != nil {
			return err
		}
		fmt.Printf("msg: %+v\n", message)
	}
	return nil
}

// Kafka is struct for kafka broker
type Kafka struct {
	Producers []*Producer
	Consumers []*Consumer
}

// NewKafka return instance of kafka
func NewKafka() *Kafka {
	return &Kafka{Producers: make([]*Producer, 0), Consumers: make([]*Consumer, 0)}
}

// CreateProducers generate producers the kafka
func (k *Kafka) CreateProducers(quantity int) error {
	for i := 0; i < quantity; i++ {
		k.Producers = append(k.Producers, NewProducer())
	}
	return nil
}

// CreateConsumers generate consumers the kafka
func (k *Kafka) CreateConsumers(quantity int) error {
	for i := 0; i < quantity; i++ {
		k.Consumers = append(k.Consumers, NewConsumer())
	}
	return nil
}

// ConnectProducers connect all producers in broker
func (k *Kafka) ConnectProducers(urls []string) error {
	for _, p := range k.Producers {
		if err := p.ConnectProducer(urls); err != nil {
			return err
		}
	}
	return nil
}

// ConnectConsumers connect all consumers in broker
func (k *Kafka) ConnectConsumers(urls []string) error {
	for _, c := range k.Consumers {
		if err := c.ConnectConsumer(urls); err != nil {
			return err
		}
	}
	return nil
}

// ExecuteProducers execute logic for all producers in broker
func (k *Kafka) ExecuteProducers(topic string) (err error) {
	for _, p := range k.Producers {
		if err = p.ProduceMessage(topic); err != nil {
			return
		}
	}
	return nil
}

// ExecuteConsumers execute logic for all consumers in broker
func (k *Kafka) ExecuteConsumers(topic string) error {
	for _, c := range k.Consumers {
		if err := c.ConsumeMessage(topic); err != nil {
			return err
		}
	}
	return nil
}
