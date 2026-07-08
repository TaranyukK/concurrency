package main

import (
	"fmt"
	"go/concurrency/configs"
	"go/concurrency/internals/verify"
	"log"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	verify.NewVerifyHandler(router, verify.VerifyHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server is listening on 8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
