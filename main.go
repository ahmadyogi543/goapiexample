package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	server := NewAPIServer(":8080")
	server.Handler = server.Routes()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Fatalln("server:", err)
		}
	}()

	log.Println("server: listening on :8080")
	if err := server.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			log.Println("server: closed under request")
		} else {
			log.Fatalln("server: closed unexpectedly")
		}
	}
	log.Println("server: shutdown complete")
}
