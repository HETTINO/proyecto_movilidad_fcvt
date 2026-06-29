package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	modelos "proyecto_movilidad_fcvt/internal/models"
)

// ListarRutas atiende GET /api/v1/rutas
func (s *Server) ListarRutas(w http.ResponseWriter, _ *http.Request) {
	rutas := s.Ruta.Listar()
	responderJSON(w, http.StatusOK, rutas)
}

// ObtenerRuta atiende GET /api/v1/rutas/{id}
func (s *Server) ObtenerRuta(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	ruta, encontrado := s.Ruta.Obtener(id)
	if !encontrado {
		responderError(w, http.StatusNotFound, "ruta no encontrada")
		return
	}

	responderJSON(w, http.StatusOK, ruta)
}

// CrearRuta atiende POST /api/v1/rutas
func (s *Server) CrearRuta(w http.ResponseWriter, r *http.Request) {
	var nuevo modelos.Ruta

	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		responderError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	creado, err := s.Ruta.Crear(nuevo)
	if err != nil {
		responderError(w, http.StatusBadRequest, err.Error())
		return
	}

	responderJSON(w, http.StatusCreated, creado)
}

// ActualizarRuta atiende PUT /api/v1/rutas/{id}
func (s *Server) ActualizarRuta(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	var datos modelos.Ruta
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		responderError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	actualizado, encontrado, err := s.Ruta.Actualizar(id, datos)
	if err != nil {
		responderError(w, http.StatusBadRequest, err.Error())
		return
	}
	if !encontrado {
		responderError(w, http.StatusNotFound, "ruta no encontrada")
		return
	}

	responderJSON(w, http.StatusOK, actualizado)
}

// BorrarRuta atiende DELETE /api/v1/rutas/{id}
func (s *Server) BorrarRuta(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	if err := s.Ruta.Borrar(id); err != nil {
		responderError(w, http.StatusNotFound, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
