package main

import (
	"log"
	"github.com/thoraf20/content-monitoring-system/content-analysis/config"
	"github.com/thoraf20/content-monitoring-system/content-analysis/consumer"
)

func main() {
	config.LoadConfig() // loads from env or config file

	log.Println("Starting Content Moderation Service...")

	// Start consuming from Redis Stream
	consumer.StartStreamConsumer()
}
