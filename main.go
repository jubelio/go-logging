package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/jubelio/go-logging/logging"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	type info struct {
		Test  string
		Test2 string
	}

	var x info
	x.Test = "TESTTTTTTTT"
	x.Test2 = "TESSSTTTTTTTT2222"

	logging.Infof("Testwith extra", x)

}
