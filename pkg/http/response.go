package model

import "encoding/json"

type Response[T any] struct {
	Status  int
	Message T
}

func SendJSON[T any]() {
	bytes, _ := json.Marshal(response)

}
