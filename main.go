package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "strive_go/config"
	"strive_go/db"
)

func main() {
	mainLogger := GetLoggerWithName("main")
	mainLogger.Info("Starting server...")

	dbLogger := GetLoggerWithName("database")

	// Initialize Database
	db.Connect(dbLogger)
	db.Migrate(dbLogger)

	// Initialize Router
	router := initRouter()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	srvErrs := make(chan error, 1)
	go func() {
		srvErrs <- srv.ListenAndServe()
	}()

	gracefulShutdown := shutdown(srv, mainLogger)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-srvErrs:
		gracefulShutdown(err)
	case sig := <-sigs:
		gracefulShutdown(sig)
	}

	mainLogger.Info("Server stopped")

}

func init() {
	createLogger()
}
