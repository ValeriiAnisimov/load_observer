package controller

import (
	"log"
	"net/http"
)

func StartServer() {
	log.Println("Starting server")

	srv := &http.Server{
		Addr:    "localhost:8282",
		Handler: Handler(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Println("Starting server failed")
	}
}
