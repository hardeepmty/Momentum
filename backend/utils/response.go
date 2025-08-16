package utils

import (
	"encoding/json"
	"net/http"
)

func SendResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func SendError(w http.ResponseWriter, message string, statusCode int) {
	SendResponse(w, map[string]string{"error": message}, statusCode)
}

func ParseBody(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}