package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /order", Handler())

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
