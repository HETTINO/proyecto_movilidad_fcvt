package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	modelos "proyecto_movilidad_fcvt/internal/modelos"
)

// ListarSolicitudes atiende GET /api/v1/solicitudes
func (s *Server) ListarSolicitudes(w http.ResponseWriter, _ *http.Request) {
	solicitudes := s.Solicitud.Listar()
	responderJSON(w, http.StatusOK, solicitudes)
}

// ObtenerSolicitud atiende GET /api/v1/solicitudes/{id}
func (s *Server) ObtenerSolicitud(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	solicitud, encontrado := s.Solicitud.Obtener(id)
	if !encontrado {
		responderError(w, http.StatusNotFound, "solicitud no encontrada")
		return
	}

	responderJSON(w, http.StatusOK, solicitud)
}

// CrearSolicitud atiende POST /api/v1/solicitudes
func (s *Server) CrearSolicitud(w http.ResponseWriter, r *http.Request) {
	var nueva modelos.Solicitud

	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		responderError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	creada, err := s.Solicitud.Crear(nueva)
	if err != nil {
		responderError(w, http.StatusBadRequest, err.Error())
		return
	}

	responderJSON(w, http.StatusCreated, creada)
}

// ActualizarSolicitud atiende PUT /api/v1/solicitudes/{id}
func (s *Server) ActualizarSolicitud(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	var datos modelos.Solicitud
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		responderError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	actualizada, encontrado, err := s.Solicitud.Actualizar(id, datos)
	if err != nil {
		responderError(w, statusDeError(err), err.Error())
		return
	}
	if !encontrado {
		responderError(w, http.StatusNotFound, "solicitud no encontrada")
		return
	}

	responderJSON(w, http.StatusOK, actualizada)
}

// BorrarSolicitud atiende DELETE /api/v1/solicitudes/{id}
func (s *Server) BorrarSolicitud(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	if err := s.Solicitud.Borrar(id); err != nil {
		responderError(w, http.StatusNotFound, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
