package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"route256/cart/internal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	app, err := internal.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err = app.RunGRPCServer(); err != http.ErrServerClosed {
			log.Fatal(fmt.Errorf("failed run server: %w", err))
		}
	}()

	go func() {
		if err = app.RunHTTPServer(); err != http.ErrServerClosed {
			log.Fatal(fmt.Errorf("failed run server: %w", err))
		}
	}()

	<-ctx.Done()

	if err = app.Shutdown(); err != nil {
		log.Fatal(fmt.Errorf("failed shutdown server: %w", err))
	}
}
