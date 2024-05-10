package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	weather "weather-tracker"
	handler2 "weather-tracker/internal/handler"
)

func main() {
	srv := new(weather.Server)
	handler := new(handler2.Handler)

	go func() {
		if err := srv.Run("8080", handler.InitRoutes()); err != nil {
			log.Fatalf("error with running server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<- quit

	fmt.Println("Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("error with shutting down server: %s", err.Error())
	}
}
