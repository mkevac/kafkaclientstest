package main

import (
	"badoo/_packages/log"

	"os"
	"runtime"
	"runtime/pprof"

	"github.com/Shopify/sarama"
)

func main() {

	f, err := os.Create("memprofile")
	if err != nil {
		log.Fatal("Error while opening file 'memprofile'")
	}
	defer f.Close()

	runtime.GC()

	defer func() {
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatalf("Error while writing memory profile: %s", err)
		}
	}()

	config := sarama.NewConfig()

	consumer, err := sarama.NewConsumer([]string{"kafka1:9092", "kafka2:9092", "kafka3:9092"}, config)
	if err != nil {
		log.Fatalf("Error while creating sarama consumer: %s", err)
	}

	partitionConsumer, err := consumer.ConsumePartition("laccess_users", 32, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("Error while creating sarama partition consumer: %s", err)
	}

	var consumed int

	messages := partitionConsumer.Messages()

	_ = <-messages

	maxOffset := partitionConsumer.HighWaterMarkOffset()

	log.Infof("max offset is %d", maxOffset)

	for message := range messages {
		consumed++
		if message.Offset+1 == maxOffset {
			break
		}
	}

	runtime.GC()

	log.Infof("Consumed %d messages", consumed)
}
