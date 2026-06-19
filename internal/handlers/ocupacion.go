package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"proyecto_movilidad_fcvt/internal/modelos"
)

// ListarOcupaciones atiende GET /api/v1/ocupaciones
func (s *Server) ListarOcupaciones(w http.ResponseWriter, _ *http.Request) {
	ocupaciones := s.Ocupacion.Listar()
	responderJSON(w, http.StatusOK, ocupaciones)
}

// ObtenerOcupacion atiende GET /api/v1/ocupaciones/{id}
func (s *Server) ObtenerOcupacion(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	ocupacion, encontrado := s.Ocupacion.Obtener(id)
	if !encontrado {
		responderError(w, http.StatusNotFound, "ocupación no encontrada")
		return
	}

	responderJSON(w, http.StatusOK, ocupacion)
}

// CrearOcupacion atiende POST /api/v1/ocupaciones
func (s *Server) CrearOcupacion(w http.ResponseWriter, r *http.Request) {
	var nueva modelos.Ocupacion

	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		responderError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	creada, err := s.Ocupacion.Crear(nueva)
	if err != nil {
		responderError(w, statusDeError(err), err.Error())
		return
	}

	responderJSON(w, http.StatusCreated, creada)
}

// ActualizarOcupacion atiende PUT /api/v1/ocupaciones/{id}
func (s *Server) ActualizarOcupacion(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	var datos modelos.Ocupacion
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		responderError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	actualizada, encontrada, err := s.Ocupacion.Actualizar(id, datos)
	if err != nil {
		responderError(w, http.StatusBadRequest, err.Error())
		return
	}
	if !encontrada {
		responderError(w, http.StatusNotFound, "ocupación no encontrada")
		return
	}

	responderJSON(w, http.StatusOK, actualizada)
}

// BorrarOcupacion atiende DELETE /api/v1/ocupaciones/{id}
func (s *Server) BorrarOcupacion(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	if err := s.Ocupacion.Borrar(id); err != nil {
		responderError(w, http.StatusNotFound, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// LiberarOcupacion atiende PATCH /api/v1/ocupaciones/{id}/liberar
func (s *Server) LiberarOcupacion(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	liberada, encontrada := s.Ocupacion.Liberar(id)
	if !encontrada {
		responderError(w, http.StatusNotFound, "ocupación no encontrada")
		return
	}

	responderJSON(w, http.StatusOK, liberada)
}
