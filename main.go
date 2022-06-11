package main

import (
	"flag"
	"io"
	"net/http"
	"time"

	"github.com/uitachi123/simple-go-server/pkg/api"
	"github.com/uitachi123/simple-go-server/pkg/db"
	"github.com/uitachi123/simple-go-server/pkg/echo"

	"go.uber.org/zap"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome to new server!")
}

func main() {

	loggingLevel := flag.String("logging", "INFO", "logging level")
	port := flag.String("port", "8080", "listening port")
	flag.Parse()

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("Starting web server...",
		// Structured context as strongly typed Field values.
		zap.String("time", time.Now().String()),
		zap.String("logging level", *loggingLevel),
	)

	_, err := db.Init()
	if err != nil {
		logger.Error("Error initializing database", zap.Error(err))
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", welcome)
	mux.HandleFunc("/echo/", echo.Echo)
	mux.HandleFunc("/users", api.Users)
	// listen to port
	http.ListenAndServe(":"+*port, mux)
}
