package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"route256/loms/internal"
	"sync"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	app, err := internal.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = app.RunGRPCServer(); err != nil {
			log.Fatal(fmt.Errorf("failed run grpc server: %w", err))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = app.RunHTTPServer(); err != nil && err != http.ErrServerClosed {
			log.Fatal(fmt.Errorf("failed run http server: %w", err))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = app.RunCancelUnpaidOrderWorker(ctx); err != nil {
			log.Fatal(fmt.Errorf("failed run cancel unpaid worker: %w", err))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = app.RunOutboxWorker(ctx); err != nil {
			log.Fatal(fmt.Errorf("failed run order status outbox worker: %w", err))
		}
	}()

	<-ctx.Done()

	if err = app.Shutdown(); err != nil {
		log.Fatal(fmt.Errorf("failed shutdown server: %w", err))
	}

	wg.Wait()
}
