package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/jubelio/go-logging/logging"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logging.Info("Test request from dev")
		w.Write([]byte("Hello, world!"))
	})

	logging.Info("Running server")

	http.ListenAndServe(":8080", nil)
}
