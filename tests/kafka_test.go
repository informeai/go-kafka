package tests

import (
	"log"
	"testing"

	"github.com/informeai/go-kafka/internal"
)

func TestNewKafka(t *testing.T) {
	kafka := internal.NewKafka()
	if kafka == nil {
		t.Errorf("TestNewKafka: expect(!nil) got(%v)\n", kafka)
	}
}

func TestKafkaCreateProducers(t *testing.T) {
	kafka := internal.NewKafka()
	if kafka == nil {
		t.Errorf("TestKafkaCreateProducers: expect(!nil) got(%v)\n", kafka)
	}
	if err := kafka.CreateProducers(1); err != nil {
		t.Errorf("TestKafkaCreateProducers: expect(nil) got(%s)\n", err.Error())
	}
  log.Printf("kafka: %+v\n",kafka)
}

func TestKafkaCreateConsumers(t *testing.T) {
	kafka := internal.NewKafka()
	if kafka == nil {
		t.Errorf("TestKafkaCreateConsumers: expect(!nil) got(%v)\n", kafka)
	}
	if err := kafka.CreateConsumers(1); err != nil {
		t.Errorf("TestKafkaCreateConsumers: expect(nil) got(%s)\n", err.Error())
	}
  log.Printf("kafka: %+v\n",kafka)
}
