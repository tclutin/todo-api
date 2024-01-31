package http

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response[T any] struct {
	Status  int
	Message T
}

func SendJSON[T any](writer http.ResponseWriter, statusCode int, data T) {
	resp := Response[T]{
		Status:  statusCode,
		Message: data,
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	if err := json.NewEncoder(writer).Encode(resp); err != nil {
		log.Fatalln("Error send json", err)
	}
}
