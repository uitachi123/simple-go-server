package main

import (
	"flag"
	"io"
	"net/http"
	"time"

	"go-server/pkg/echo"
	"go.uber.org/zap"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome to new server!")
}

func main() {

	loggingLevel := flag.String("logging", "INFO", "logging level")
	flag.Parse()

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("Starting web server...",
		// Structured context as strongly typed Field values.
		zap.String("time", time.Now().String()),
		zap.String("logging level", *loggingLevel),
	)

	mux := http.NewServeMux()
	mux.HandleFunc("/", welcome)
	mux.HandleFunc("/echo/", echo.Echo)
	// listen to port
	http.ListenAndServe(":8080", mux)
}
