package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os/signal"
    "route256/cart/internal/server"
    "syscall"
)

func main() {
    ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
    defer cancel()

    s, err := server.New()
    if err != nil {
        log.Fatal(fmt.Errorf("failed initialization server: %w", err))
    }

    go func() {
        if err = s.Run(); err != http.ErrServerClosed {
            log.Fatal(fmt.Errorf("failed run server: %w", err))
        }
    }()

    <-ctx.Done()

    if err = s.Shutdown(); err != nil {
        log.Fatal(fmt.Errorf("failed shutdown server: %w", err))
    }
}
