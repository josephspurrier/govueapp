package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"app/api/config"
	"app/api/internal/requestcontext"
	"app/api/pkg/logger"
)

func main() {
	// Create the logger.
	l := logger.New(log.New(os.Stderr, "", log.Lshortfile))

	// Load the environment variables.
	settings := config.LoadEnv(l, "")

	// Setup the services.
	core := config.Services(l, settings, config.Database(l), requestcontext.New(), nil)
	config.LoadRoutes(core)

	// Start the web server.
	l.Printf("Server started.")
	err := http.ListenAndServe(fmt.Sprintf(":%v", settings.Port), config.Middleware(core))
	if err != nil {
		l.Printf(err.Error())
	}
}
