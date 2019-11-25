package main

import (
	"context"
	"flag"
	"fmt"
	"go-starter/ch15OHIS/ohisrouter"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	port := flag.Int64("p", 8080, "port")
	flag.Parse()

	engine := gin.Default()
	ohisrouter.InitRouter(engine)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen and Serve: %v\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server Shutdown: %v\n", err)
	}
	log.Println("Server Existing In 5 Seconds.")
}
