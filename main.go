package main

import (
	"context"
	"fmt"
	"league/internal/api"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/echo", api.EchoHandler)
	mux.HandleFunc("/invert", api.InvertHandler)
	mux.HandleFunc("/sum", api.SumHandler)
	mux.HandleFunc("/multiply", api.MultiplyHandler)
	mux.HandleFunc("/flatten", api.FlattenHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Graceful shutdown listener
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		fmt.Println("Shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			fmt.Printf("Server shutdown failed: %v\n", err)
		}
	}()

	fmt.Println("Server running at http://localhost:8080")
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Printf("Server error: %v\n", err)
	}
}
