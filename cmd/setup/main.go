package main

import (
	"GOtestprogect/internal/logger"
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
	Service := task.TaskService{DataStorage}
	DataService := task.DataService{&Service}
	LoggerChan := logger.AsLogger()
	Handler := task.Handler{LoggerChan, DataService}

	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			Handler.HandlerAddTaks(w, r)
		case http.MethodGet:
			Handler.HandlerTakeTasks(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/tasks/", Handler.HandlerTaskByID)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	
	go func() {
		log.Println("Server started")
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Shutdown Failed: %v", err)
	}
	log.Println("Server close")
	close(LoggerChan)

}
