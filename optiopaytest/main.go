package main

import (
	"badoo/_packages/log"

	"os"
	"runtime"
	"runtime/pprof"

	"github.com/optiopay/kafka"
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

	brokerConf := kafka.NewBrokerConf("marko")
	broker, err := kafka.Dial([]string{"kafka1:9092", "kafka2:9092", "kafka3:9092"}, brokerConf)
	if err != nil {
		log.Fatalf("Error while dialing to kafka: %s", err)
	}

	consumerConf := kafka.NewConsumerConf("laccess_users", 32)
	consumerConf.RetryLimit = 0
	consumerConf.RetryWait = 0
	consumerConf.StartOffset = kafka.StartOffsetOldest

	consumer, err := broker.BatchConsumer(consumerConf)
	if err != nil {
		log.Fatalf("Error while creating kafka consumer: %s", consumer)
	}

	for {
		_, err := consumer.ConsumeBatch()
		if err != nil {
			if err == kafka.ErrNoData {
				break
			}
		}
	}

	runtime.GC()
}
