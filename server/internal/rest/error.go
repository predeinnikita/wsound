package rest

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

func ResponseError(w http.ResponseWriter, error Error) {
	response, _ := json.Marshal(error)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(error.Status)
	_, _ = w.Write(response)
}
