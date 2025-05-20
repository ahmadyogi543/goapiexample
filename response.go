package main

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if v != nil {
		return json.NewEncoder(w).Encode(v)
	}
	return nil
}

func WriteHttpOKJSON(w http.ResponseWriter, v any) error {
	return WriteJSON(w, http.StatusOK, v)
}

func WriteHttpBadRequestJSON(w http.ResponseWriter, v any) error {
	return WriteJSON(w, http.StatusBadRequest, v)
}

func WriteHttpNoContentJSON(w http.ResponseWriter) error {
	return WriteJSON(w, http.StatusNoContent, nil)
}

func WriteHttpInternalServerErrorJSON(w http.ResponseWriter) error {
	return WriteJSON(w, http.StatusInternalServerError, nil)
}
