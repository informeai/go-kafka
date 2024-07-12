package main

import (
	"log"

	"github.com/informeai/go-kafka/internal"
)

func main() {
	kafka := internal.NewKafka()
	if err := kafka.CreateConsumers(100); err != nil {
		log.Fatal(err)
	}
  if err := kafka.ConnectConsumers([]string{"kafka:9092"}); err != nil {
		log.Fatal(err)
	}
  if err := kafka.ExecuteConsumers("message"); err != nil{
    log.Fatal(err)
  }
}
