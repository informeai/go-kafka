package main

import (
	"log"

	"github.com/informeai/go-kafka/internal"
)

func main() {
	kafka := internal.NewKafka()
	if err := kafka.CreateProducers(200); err != nil {
		log.Fatal(err)
	}
  if err := kafka.ConnectProducers([]string{"kafka:9092"}); err != nil {
		log.Fatal(err)
	}
  if err := kafka.ExecuteProducers("message"); err != nil{
    log.Fatal(err)
  }
}
