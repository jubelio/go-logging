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

	logging.Debug("Debug message")

	logging.Infof("Request payload", logging.Fields{
		"common":  "this is common info",
		"other":   "this is other info",
		"shop_id": 10999,
		"user_id": 10999,
		"details": logging.Fields{
			"name":  "John Doe",
			"email": "john@doe.com",
			"phone": "1234567890",
		},
		"items": []logging.Fields{
			{
				"id":   "123456",
				"name": "Item 1",
				"price": logging.Fields{
					"amount":   100,
					"currency": "USD",
				},
			},
			{
				"id":   "123456",
				"name": "Item 2",
				"price": logging.Fields{
					"amount":   200,
					"currency": "USD",
				},
			},
			{
				"id":   "123456",
				"name": "Item 3",
				"price": logging.Fields{
					"amount":   300,
					"currency": "USD",
				},
			},
		},
	})

}
