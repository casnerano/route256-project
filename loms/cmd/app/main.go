package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"route256/loms/internal/config"
	"route256/loms/internal/server"
	"route256/loms/internal/service/order"
)

func main() {
	configuration, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	s, err := server.New(configuration.Server)
	if err != nil {
		log.Fatal(fmt.Errorf("failed initialization server: %w", err))
	}

	go func() {
		if err = s.Run(); err != http.ErrServerClosed {
			log.Fatal(fmt.Errorf("failed run server: %w", err))
		}
	}()

	cancelUnpaidWorker := order.NewCancelUnpaidWorker(
		time.Duration(configuration.Order.CancelUnpaidTimeout) * time.Second,
	)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = cancelUnpaidWorker.Run(ctx); err != nil {
			log.Fatal(fmt.Errorf("failed cancel unpaid worker: %w", err))
		}
	}()

	<-ctx.Done()
	wg.Wait()

	if err = s.Shutdown(); err != nil {
		log.Fatal(fmt.Errorf("failed shutdown server: %w", err))
	}
}
