package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	modelos "proyecto_movilidad_fcvt/internal/modelos"
)

// ListarLocaciones atiende GET /api/v1/locaciones
func (s *Server) ListarLocaciones(w http.ResponseWriter, _ *http.Request) {
	locaciones := s.Locacion.Listar()
	responderJSON(w, http.StatusOK, locaciones)
}

// ObtenerUbicacionCarrito atiende GET /api/v1/locaciones/carrito/{id}
func (s *Server) ObtenerUbicacionCarrito(w http.ResponseWriter, r *http.Request) {
	carritoID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	locacion, encontrado := s.Locacion.ObtenerUltimaDelCarrito(carritoID)
	if !encontrado {
		responderError(w, http.StatusNotFound, "ubicación no encontrada")
		return
	}

	responderJSON(w, http.StatusOK, locacion)
}

// RegistrarLocacion atiende POST /api/v1/locaciones
func (s *Server) RegistrarLocacion(w http.ResponseWriter, r *http.Request) {
	var nueva modelos.Locacion

	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		responderError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	registrada, err := s.Locacion.Registrar(nueva)
	if err != nil {
		responderError(w, http.StatusBadRequest, err.Error())
		return
	}

	responderJSON(w, http.StatusCreated, registrada)
}

// GetTiempoEstimado atiende GET /api/v1/tiempo-estimado
func (s *Server) GetTiempoEstimado(w http.ResponseWriter, r *http.Request) {
	carritoIDStr := r.URL.Query().Get("carrito_id")
	destino := r.URL.Query().Get("destino")

	if carritoIDStr == "" || destino == "" {
		responderError(w, http.StatusBadRequest, "carrito_id y destino son requeridos")
		return
	}

	carritoID, err := strconv.Atoi(carritoIDStr)
	if err != nil {
		responderError(w, http.StatusBadRequest, "carrito_id inválido")
		return
	}

	// TODO: Implementar lógica de cálculo de tiempo
	tiempoEstimado := 15 // Valor dummy por ahora

	responderJSON(w, http.StatusOK, map[string]interface{}{
		"carrito_id":     carritoID,
		"destino":        destino,
		"tiempo_minutos": tiempoEstimado,
	})
}
