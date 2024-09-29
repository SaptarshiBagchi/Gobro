package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"gobro.starter/internal/adapters/http/user"
	"gobro.starter/internal/adapters/messaging"
	"gobro.starter/internal/ports"
)

func SetupApp() {

	router := mux.NewRouter()

	//setup message broker adapter
	var kafkaPublisher ports.PublisherConfig = messaging.NewKafkaPublisher("localhost")
	//Setup messaging broker port
	publisher := ports.GetMessagingInstance(kafkaPublisher)

	//setup user routes
	user.SetupUserRoutes(router, publisher)

	// Create the HTTP server
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Graceful shutdown channel
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt) // or use syscall.SIGTERM for termination signal

	// Start the server in a goroutine
	go func() {
		fmt.Println("Starting the server on port 8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %s\n", err)
		}
	}()

	// Wait for an interrupt signal
	<-stop
	fmt.Println("Shutting down the server...")

	// Create a context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Gracefully shutdown the server
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %s\n", err)
	}
	//Close down the adapters as well for smoother shutfown
	defer kafkaPublisher.Close()

	fmt.Println("Server exited gracefully")
}
