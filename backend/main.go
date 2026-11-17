package main

import (
	"backend/routes"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	ADDR string = ":5461"
)

func main() {
	r := gin.Default()

	r.GET("/api/color-graphs", routes.ColorGraph)
	r.GET("/api/alg/2", routes.Algorithm_2)
	r.GET("/api/alg/3", routes.Algorithm_3)

	go func() {
		log.Printf("Server running at: http://localhost%s\n", ADDR)
		if err := r.Run(ADDR); err != nil {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	_, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()
}
