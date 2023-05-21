package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ValeriiAnisimov/load_observer/internal/observer"
)

func Handler() http.Handler {
	h := http.NewServeMux()
	h.HandleFunc("/api", handleConnection)
	return h
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fullJson := observer.GetResultT()
	result, err := json.Marshal(fullJson)
	if err != nil {
		log.Println(err)
	}
	w.Write(result)
}
