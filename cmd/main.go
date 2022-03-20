package main

import (
	"context"
	"github.com/smallretardedfish/cakes/internal/repository"
	"github.com/smallretardedfish/cakes/internal/server"
	"github.com/smallretardedfish/cakes/internal/service"
	"github.com/smallretardedfish/cakes/internal/transport/handler"
	"log"
	"os"
	"os/signal"
	"time"
)

func run() error {

	repo := repository.NewMemoryStorage()
	service := service.NewUserService(repo)
	handler := handler.NewUserHandler(service)
	r := handler.InitRoutes()
	srv := server.NewHTTPServer(":8080", r)

	log.Println("starting http server on port :8080")
	log.Println("app can be interrupted by pressing Ctl+C")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	go func() {
		<-interrupt
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalln(err)
		}
	}()
	if err := srv.Run(); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Println(err)
	}
	log.Println("Good Bye!")
}
