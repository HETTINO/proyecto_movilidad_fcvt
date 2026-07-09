package handler_parqueadero

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"proyecto_movilidad_fcvt/internal/modelos"
)

// ListarParqueaderos atiende GET /api/v1/parqueaderos
func (s *Server) ListarParqueaderos(w http.ResponseWriter, _ *http.Request) {
	parqueaderos := s.Parqueadero.Listar()
	responderJSON(w, http.StatusOK, parqueaderos)
}

// ObtenerParqueadero atiende GET /api/v1/parqueaderos/{id}
func (s *Server) ObtenerParqueadero(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	parqueadero, encontrado := s.Parqueadero.Obtener(id)
	if !encontrado {
		responderError(w, http.StatusNotFound, "parqueadero no encontrado")
		return
	}

	responderJSON(w, http.StatusOK, parqueadero)
}

// CrearParqueadero atiende POST /api/v1/parqueaderos
func (s *Server) CrearParqueadero(w http.ResponseWriter, r *http.Request) {
	var nuevo modelos.Parqueadero

	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		responderError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	creado, err := s.Parqueadero.Crear(nuevo)
	if err != nil {
		responderError(w, statusDeError(err), err.Error())
		return
	}

	responderJSON(w, http.StatusCreated, creado)
}

// ActualizarParqueadero atiende PUT /api/v1/parqueaderos/{id}
func (s *Server) ActualizarParqueadero(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	var datos modelos.Parqueadero
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		responderError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	actualizado, encontrado, err := s.Parqueadero.Actualizar(id, datos)
	if err != nil {
		responderError(w, statusDeError(err), err.Error())
		return
	}
	if !encontrado {
		responderError(w, http.StatusNotFound, "parqueadero no encontrado")
		return
	}

	responderJSON(w, http.StatusOK, actualizado)
}

// BorrarParqueadero atiende DELETE /api/v1/parqueaderos/{id}
func (s *Server) BorrarParqueadero(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	if err := s.Parqueadero.Borrar(id); err != nil {
		responderError(w, http.StatusNotFound, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
