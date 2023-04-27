package main

import (
	"context"
	"fmt"
	"golang-micro/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	r := router.RouterApi()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	var wg sync.WaitGroup
	var errApi = make(chan error, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		errApi <- srv.ListenAndServe()
	}()
	wg.Wait()

	var quit = make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
	select {
	case <-quit:
		fmt.Println("Got an interrupt, exiting...")
		shutdown(srv)
	case err := <-errApi:
		if err != nil {
			fmt.Println("Error while running api, exiting...: ", err)
			shutdown(srv)
		}
	}

	r.Run(":8080")
}

func shutdown(srv *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
