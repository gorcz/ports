package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

const shutdownTimeout = time.Duration(10) * time.Second

func startHTTPServer(ctx context.Context, router http.Handler, port uint16) {
	wg := new(sync.WaitGroup)

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		ReadHeaderTimeout: 1 * time.Second,
		ReadTimeout:       30 * time.Second,
		Handler:           router,
	}

	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		log.Panicf("cannot create http-server listener %s", err)
	}
	if server.Addr != ln.Addr().String() {
		server.Addr = ln.Addr().String()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		go func() {
			<-ctx.Done()

			shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
			defer cancel()

			err = server.Shutdown(shutdownCtx)
			if err != nil {
				log.Panicf("cannot close http-server %s\n", err)
			}
		}()

		log.Printf("HTTP http-server is listening on port :%d", port)
		if err = server.Serve(ln); err != nil && err != http.ErrServerClosed {
			log.Panicf("http-server error: %s", err)
		}
	}()

	wg.Wait()
}
