package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type config struct {
	port    int
	env     string
	version string
}
type Application struct {
	config config
	logger *slog.Logger
}

func main() {

	var cfg config
	cfg.version = "v1"
	//flag package enables us to associate type,default values ,helper message to the arguments
	flag.IntVar(&cfg.port, "port", 4000, "enter the port number of webserver")
	flag.StringVar(&cfg.env, "env", "dev", "options:dev,prod")
	flag.Parse()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &Application{
		config: cfg,
		logger: logger,
	}
	router := app.routes()

	//configuring the server manually
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	err := server.ListenAndServe()
	app.logger.Error(err.Error())

}
