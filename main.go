package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/braintree/manners"
	"github.com/sl1pm4t/sems-healthz/handlers"
)

const version = "0.1.1"

func main() {
	var (
		healthAddr = flag.String("http", "0.0.0.0:4480", "Health service address.")
	)
	flag.Parse()

	log.Println("Starting server...v", version)
	log.Printf("Health service listening on %s", *healthAddr)

	errChan := make(chan error, 10)

	// setup the HTTP endpoints
	hmux := http.NewServeMux()
	hmux.HandleFunc("/healthz", handlers.HealthzHandler)
	hmux.HandleFunc("/readiness", handlers.ReadinessHandler)

	// create HTTP server, and attach logging handler
	healthServer := manners.NewServer()
	healthServer.Addr = *healthAddr
	healthServer.Handler = handlers.LoggingHandler(hmux)

	go func() {
		errChan <- healthServer.ListenAndServe()
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case err := <-errChan:
			if err != nil {
				log.Fatal(err)
			}
		case s := <-signalChan:
			log.Println(fmt.Sprintf("Captured %v. Exiting...", s))
			handlers.SetReadinessStatus(http.StatusServiceUnavailable)
			os.Exit(0)
		}
	}
}
