package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	f, err := os.Create("memprofile")
	if err != nil {
		log.Fatalf("Error while opening 'memprofile' file: %s", err)
	}
	defer f.Close()

	runtime.GC()

	defer func() {
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatalf("Error while writing heap profile: %s", err)
		}
	}()

	// insert here

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":               "kafka1:9092,kafka2:9092,kafka3:9092",
		"group.id":                        "marko",
		"session.timeout.ms":              6000,
		"go.events.channel.enable":        true,
		"go.application.rebalance.enable": true,
		"default.topic.config":            kafka.ConfigMap{"auto.offset.reset": "earliest"}})

	if err != nil {
		log.Fatalf("Failed to create consumer: %s\n", err)
	}

	err = c.Subscribe("laccess_users", nil)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while subscribing to topics\n")
		os.Exit(1)
	}

	topic := "laccess_users"
	var messages uint64

	ch := c.Events()

	for ev := range ch {
		switch e := ev.(type) {
		case *kafka.Message:
			//log.Printf("%% Message on %s", e.TopicPartition)
			messages++
		case kafka.PartitionEOF:
			log.Printf("%% Reached %v\n", e)
			log.Printf("got %v messages", messages)
			runtime.GC()
			os.Exit(0)
		case kafka.Error:
			log.Printf("%% Error: %v\n", e)
			break
		case kafka.AssignedPartitions:
			if err := c.Assign([]kafka.TopicPartition{kafka.TopicPartition{
				Topic:     &topic,
				Partition: 32,
				Offset:    kafka.OffsetBeginning,
			}}); err != nil {
				log.Fatal("Error while assigning partition and offset.")
			}
		}
	}
}
