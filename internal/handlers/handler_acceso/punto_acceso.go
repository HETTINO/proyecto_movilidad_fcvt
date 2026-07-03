package handler_acceso

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"proyecto_movilidad_fcvt/internal/modelos"
)

// ListarPuntosAcceso maneja GET /api/v1/puntos-acceso
func (s *Server) ListarPuntosAcceso(w http.ResponseWriter, r *http.Request) {
	puntos := s.PuntoAcceso.Listar()
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(puntos)
}

// CrearPuntoAcceso maneja POST /api/v1/puntos-acceso
func (s *Server) CrearPuntoAcceso(w http.ResponseWriter, r *http.Request) {
	var entrada modelos.PuntoDeAcceso
	if err := json.NewDecoder(r.Body).Decode(&entrada); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	creado, err := s.PuntoAcceso.Crear(entrada)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(creado)
}

// ObtenerPuntoAcceso maneja GET /api/v1/puntos-acceso/{id}
func (s *Server) ObtenerPuntoAcceso(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	punto, ok := s.PuntoAcceso.Obtener(id)
	if !ok {
		http.Error(w, "Punto de acceso no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(punto)
}

// ActualizarPuntoAcceso maneja PUT /api/v1/puntos-acceso/{id}
func (s *Server) ActualizarPuntoAcceso(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var datos modelos.PuntoDeAcceso
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	actualizado, ok, err := s.PuntoAcceso.Actualizar(id, datos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !ok {
		http.Error(w, "Punto de acceso no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(actualizado)
}

// BorrarPuntoAcceso maneja DELETE /api/v1/puntos-acceso/{id}
func (s *Server) BorrarPuntoAcceso(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := s.PuntoAcceso.Borrar(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
