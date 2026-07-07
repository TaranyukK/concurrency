package main

import (
	"fmt"
	randomapi "go/concurrency/2-random-api"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	randomapi.NewRandomNumHandler(router)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server running on port :8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
