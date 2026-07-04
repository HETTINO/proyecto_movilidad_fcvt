package handler_parqueadero

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"proyecto_movilidad_fcvt/internal/modelos"
)

// ListarEspacios atiende GET /api/v1/espacios
func (s *Server) ListarEspacios(w http.ResponseWriter, _ *http.Request) {
	espacios := s.Espacio.Listar()
	responderJSON(w, http.StatusOK, espacios)
}

// ObtenerEspacio atiende GET /api/v1/espacios/{id}
func (s *Server) ObtenerEspacio(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	espacio, encontrado := s.Espacio.Obtener(id)
	if !encontrado {
		responderError(w, http.StatusNotFound, "espacio no encontrado")
		return
	}

	responderJSON(w, http.StatusOK, espacio)
}

// CrearEspacio atiende POST /api/v1/espacios
func (s *Server) CrearEspacio(w http.ResponseWriter, r *http.Request) {
	var nuevo modelos.Espacio

	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		responderError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	creado, err := s.Espacio.Crear(nuevo)
	if err != nil {
		responderError(w, statusDeError(err), err.Error())
		return
	}

	responderJSON(w, http.StatusCreated, creado)
}

// ActualizarEspacio atiende PUT /api/v1/espacios/{id}
func (s *Server) ActualizarEspacio(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	var datos modelos.Espacio
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		responderError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	actualizado, encontrado, err := s.Espacio.Actualizar(id, datos)
	if err != nil {
		responderError(w, statusDeError(err), err.Error())
		return
	}
	if !encontrado {
		responderError(w, http.StatusNotFound, "espacio no encontrado")
		return
	}

	responderJSON(w, http.StatusOK, actualizado)
}

// BorrarEspacio atiende DELETE /api/v1/espacios/{id}
func (s *Server) BorrarEspacio(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	if err := s.Espacio.Borrar(id); err != nil {
		responderError(w, http.StatusNotFound, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
