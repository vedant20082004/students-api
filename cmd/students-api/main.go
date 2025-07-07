package main

import (
	"context"
	// "fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vedant20082004/students-api/internal/config"
)

func main() {

	// load config
	cfg :=config.MustLoad()


	// logger setup (if not using inbuilt)
	// database setup

	// setupr router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("WELCOME TO STUDENTS API"))
	})


	// setup server

	server := http.Server{
		Addr: cfg.Addr,
		Handler: router,

	}

	slog.Info("Server Started addr : ", slog.String("address" , cfg.Addr))
	// fmt.Printf("Server Started addr : %s", cfg.HttpServer.Addr)	

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func(){

		err := server.ListenAndServe()

		if err != nil {
			log.Fatal("Failed to start server")
		}
	}()

	<-done

	slog.Info("Shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(),5*time.Second)
		defer cancel()
	
		
	err := server.Shutdown(ctx)
	
	if err != nil {
		slog.Error("Failed to shut down ", slog.String("ERROR", err.Error()))
	}

	slog.Info("Server shutdown gracefully")

	// fmt.Println("Server Started")	


}