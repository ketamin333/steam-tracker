package handlers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func JSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(Response{
		Success: status < 400,
		Data:    data,
	})
}

func Error(w http.ResponseWriter, status int, msg string) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(Response{
		Success: false,
		Error:   msg,
	})
}

func NoContent(w http.ResponseWriter, status int) error {
	w.WriteHeader(status)
	return nil
}
