package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"errors"
	st "proyecto_movilidad_fcvt/internal/service/service_transporte"
)

// RespondJSON escribe data como JSON con el status HTTP indicado.
//
// Centraliza tres cosas que antes repetíamos en CADA handler:
//   - poner el header Content-Type
//   - escribir el status code
//   - codificar el cuerpo y registrar el error si la codificación falla
//
// Si data es nil (por ejemplo en un 204 No Content) no escribe cuerpo.
// responderJSON escribe data como JSON con el status HTTP indicado.
func responderJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// RespondError escribe un error en un formato JSON consistente: {"error": "..."}.
//
// Así el cliente siempre recibe los errores con la misma forma, en lugar de
// texto plano unas veces y JSON otras.
// responderError escribe un error en formato JSON consistente.
func responderError(w http.ResponseWriter, status int, mensaje string) {
	responderJSON(w, status, map[string]string{"error": mensaje})
}
// statusDeError traduce un error del service al código HTTP correcto.
// Centraliza esta decisión en un solo lugar en vez de repetirla en cada handler.
func statusDeError(err error) int {
	switch {
	case errors.Is(err, st.ErrNoEncontrado):
		return http.StatusNotFound

	case errors.Is(err, st.ErrCampoRequerido):
		return http.StatusBadRequest

	case errors.Is(err, st.ErrDatosInvalidos):
		return http.StatusBadRequest

	default:
		return http.StatusInternalServerError
	}
}
