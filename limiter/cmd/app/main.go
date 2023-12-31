package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"

	"golang.org/x/time/rate"
)

var (
	listenAddr = "0.0.0.0:3000"
	rpcLimit   = 10
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	var rateLimiter = rate.NewLimiter(rate.Limit(rpcLimit), rpcLimit)

	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = listener.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	go func() {
		<-ctx.Done()
		if err = listener.Close(); err != nil {
			log.Println(err)
		}
	}()

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Println("Error accepting: ", err)
			return
		}

		go func(conn net.Conn) {
			defer func() {
				err := conn.Close()
				if err != nil {
					log.Println(err)
				}
			}()

			buf := make([]byte, 1)

			for {
				_, err := conn.Read(buf)
				if err != nil {
					log.Println("Error reading: ", err)
					return
				}

				log.Println("Received message: ", buf)

				err = rateLimiter.Wait(ctx)
				if err != nil {
					log.Println("Error limiter wait: ", err)
				}

				_, err = conn.Write([]byte{1})
				if err != nil {
					log.Println("Error writing:", err)
					return
				}
			}
		}(conn)
	}
}

func init() {
	flag.StringVar(&listenAddr, "addr", listenAddr, "TCP listen address")
	flag.IntVar(&rpcLimit, "rpc", rpcLimit, "Request per seconds limit")

	flag.Parse()
}
