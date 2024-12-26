package main

import (
	"encoding/json"
	"log"
	"net/http"
	"order-service/producer"
)

type RequestPayload struct {
	OrderId string `json:"order_id"`
}

func Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var jsonRequest RequestPayload

		err := json.NewDecoder(r.Body).Decode(&jsonRequest)
		if err != nil {
			log.Fatal("error while parsing request json")
			return
		}

		log.Printf("get order_id = %s", jsonRequest.OrderId)

		producer.PublishOrderMessage(jsonRequest.OrderId)

		w.WriteHeader(200)
	}
}
