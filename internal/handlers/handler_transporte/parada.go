package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	modelos "proyecto_movilidad_fcvt/internal/modelos"
)

// ListarParadas atiende GET /api/v1/paradas
func (s *Server) ListarParadas(w http.ResponseWriter, _ *http.Request) {
	paradas := s.Parada.Listar()
	responderJSON(w, http.StatusOK, paradas)
}

// ObtenerParada atiende GET /api/v1/paradas/{id}
func (s *Server) ObtenerParada(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	parada, encontrado := s.Parada.Obtener(id)
	if !encontrado {
		responderError(w, http.StatusNotFound, "parada no encontrada")
		return
	}

	responderJSON(w, http.StatusOK, parada)
}

// CrearParada atiende POST /api/v1/paradas
func (s *Server) CrearParada(w http.ResponseWriter, r *http.Request) {
	var nuevo modelos.Parada

	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		responderError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	creado, err := s.Parada.Crear(nuevo)
	if err != nil {
		responderError(w, http.StatusBadRequest, err.Error())
		return
	}

	responderJSON(w, http.StatusCreated, creado)
}

// ActualizarParada atiende PUT /api/v1/paradas/{id}
func (s *Server) ActualizarParada(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	var datos modelos.Parada
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		responderError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	actualizado, encontrado, err := s.Parada.Actualizar(id, datos)
	if err != nil {
		responderError(w, statusDeError(err), err.Error())
		return
	}
	if !encontrado {
		responderError(w, http.StatusNotFound, "parada no encontrada")
		return
	}

	responderJSON(w, http.StatusOK, actualizado)
}

// BorrarParada atiende DELETE /api/v1/paradas/{id}
func (s *Server) BorrarParada(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	if err := s.Parada.Borrar(id); err != nil {
		responderError(w, http.StatusNotFound, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
