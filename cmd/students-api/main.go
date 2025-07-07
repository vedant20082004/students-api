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
	"github.com/vedant20082004/students-api/internal/http/handlers/student"
	"github.com/vedant20082004/students-api/internal/storage/sqlite"
)

func main() {

	// load config
	cfg :=config.MustLoad()


	// logger setup (if not using inbuilt)
	// database setup
	storage, err := sqlite.New(cfg)
	if err != nil{
		log.Fatal(err)
	}

	slog.Info("storage initialized", slog.String("env", cfg.Env),slog.String("version", "1.0.0"))



	// setupr router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students",student.New(storage))


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
	
		
	err1 := server.Shutdown(ctx)
	
	if err1 != nil {
		slog.Error("Failed to shut down ", slog.String("ERROR", err1.Error()))
	}

	slog.Info("Server shutdown gracefully")

	// fmt.Println("Server Started")	


}