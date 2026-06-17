package handlers

import (
	"encoding/json"
	"net/http"
)

// responderJSON escribe una respuesta JSON con el status code dado
func responderJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "error al serializar respuesta", http.StatusInternalServerError)
	}
}

// responderError escribe un mensaje de error en formato JSON
func responderError(w http.ResponseWriter, status int, mensaje string) {
	responderJSON(w, status, map[string]string{"error": mensaje})
}
