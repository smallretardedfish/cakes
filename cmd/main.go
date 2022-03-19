package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func CakeHandler(c *gin.Context) {
	if _, err := c.Writer.Write([]byte("cake")); err != nil {
		log.Println(err)
		return
	}
	c.Status(http.StatusOK)
}

func run() error {
	r := gin.Default()
	r.GET("/cake", CakeHandler)
	srv := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("starting http server on port :8000")
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
	if err := srv.ListenAndServe(); err != nil {
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
