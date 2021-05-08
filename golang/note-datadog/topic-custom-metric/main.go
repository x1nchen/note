package main

import (
	"log"
	"time"

	"github.com/DataDog/datadog-go/statsd"
)

func main() {
	statsdClient, err := statsd.New("127.0.0.1:8125")
	defer statsdClient.Close()

	if err != nil {
		log.Fatal(err)
	}

	for {
		err := statsdClient.SimpleServiceCheck("followme.mt4-agent.connect", statsd.Ok)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(10 * time.Second)
		log.Println("service check emit")
	}
}
