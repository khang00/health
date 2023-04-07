package main

import (
	"context"
	"github.com/khang00/health/internal/route"
	"github.com/khang00/health/internal/store"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// Setup
	e := echo.New()

	// Set log level
	e.Logger.SetLevel(log.INFO)

	// Set up storage
	storeClient := store.NewStoreClient()
	defer storeClient.Closed()

	// Setup Route
	route.SetUpRoutes(e, storeClient)

	// Start server
	go func() {
		if err := e.Start(":3000"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
