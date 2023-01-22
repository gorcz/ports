package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorcz/ports/internal/router"
)

const httpServerPort = 8080

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	registerOSSignalHandler(ctx, cancel)

	serviceRouter := router.NewRouter()
	startHTTPServer(ctx, serviceRouter, httpServerPort)

	log.Println("HTTP server stopped gracefully")
}

func registerOSSignalHandler(ctx context.Context, handler func()) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT)
	signal.Notify(signals, syscall.SIGTERM)

	go func() {
		select {
		case <-ctx.Done():
			return
		case sig := <-signals:
			log.Printf("%s signal received\n", sig.String())
			handler()
		}
	}()
}
