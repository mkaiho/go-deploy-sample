package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var seed = int(time.Now().UnixNano())

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", handler)
	srv := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			// Error starting or closing listener:
			fmt.Println("Server closed with error:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	fmt.Printf("SIGNAL %v recieved.\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	fmt.Println("Shutdown.")

}

func handler(w http.ResponseWriter, r *http.Request) {
	message := rand.Intn(seed)
	fmt.Fprint(w, message)
}
