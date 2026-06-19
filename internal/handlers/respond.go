package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"proyecto_movilidad_fcvt/internal/service"
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
func statusDeError(err error) int {
	switch {
	case errors.Is(err, service.ErrNombreVacio):
		return http.StatusBadRequest

	case errors.Is(err, service.ErrPrecioNegativo):
		return http.StatusBadRequest

	case errors.Is(err, service.ErrEmailVacio):
		return http.StatusBadRequest

	case errors.Is(err, service.ErrPasswordVacio):
		return http.StatusBadRequest

	case errors.Is(err, service.ErrEmailenUso):
		return http.StatusConflict

	case errors.Is(err, service.ErrCredencialesInvalidas):
		return http.StatusUnauthorized

	case errors.Is(err, service.ErrTokenInvalido):
		return http.StatusUnauthorized

	case errors.Is(err, service.ErrNoEncontrado):
		return http.StatusNotFound

	default:
		return http.StatusInternalServerError
	}
}
