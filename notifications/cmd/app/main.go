package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"route256/notifications/internal/config"
	"route256/notifications/pkg/kafka"
	"sync"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	consumerGroupHandler := kafka.NewConsumerGroupHandler()

	consumerGroup, err := kafka.NewConsumerGroup(
		c.OrderStatus.Brokers,
		"order_status_group",
		c.OrderStatus.Topics,
		consumerGroupHandler,
	)
	if err != nil {
		log.Fatal(err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		consumerGroup.Run(ctx)
	}()

	<-consumerGroupHandler.Ready()
	log.Println("Order status consumer group is started!")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	var isRunning = true

	for isRunning {
		select {
		case <-ctx.Done():
			log.Println("Terminating: context deadlined.")
			isRunning = false
		case <-sigterm:
			log.Println("Terminating: via signal.")
			isRunning = false
		}
	}

	cancel()
	wg.Wait()

	if err = consumerGroup.Close(); err != nil {
		log.Fatalf("Error closing consumer group: %v", err)
	}
}
