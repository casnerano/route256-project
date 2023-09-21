package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"sync"
	"syscall"

	"route256/loms/internal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	app, err := internal.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err = app.RunServer(); err != http.ErrServerClosed {
			log.Fatal(fmt.Errorf("failed run server: %w", err))
		}
	}()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = app.RunCancelUnpaidOrderWorker(ctx); err != nil {
			log.Fatal(fmt.Errorf("failed cancel unpaid worker: %w", err))
		}
	}()

	<-ctx.Done()
	wg.Wait()

	if err = app.ShutdownServer(); err != nil {
		log.Fatal(fmt.Errorf("failed shutdown server: %w", err))
	}
}
