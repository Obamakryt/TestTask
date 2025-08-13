package main

import (
	"GOtestprogect/internal/task"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	Data := task.NewData()

	DataStorage := task.DataStorage{Data}

	mux := http.NewServeMux()
	mux.HandleFunc("/tasks")
	mux.HandleFunc("tasks/", task.HandlerAddTaks)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		log.Println("Server started on :8080")
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	close()
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Shutdown Failed: %v", err)
	}
	log.Println("Server close")
}g
