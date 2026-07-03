package handler_acceso

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"proyecto_movilidad_fcvt/internal/modelos"
)

// ListarAccesos maneja GET /api/v1/accesos
func (s *Server) ListarAccesos(w http.ResponseWriter, r *http.Request) {
	accesos := s.Acceso.Listar()
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(accesos)
}

// CrearAcceso maneja POST /api/v1/accesos
func (s *Server) CrearAcceso(w http.ResponseWriter, r *http.Request) {
	var entrada modelos.Acceso
	if err := json.NewDecoder(r.Body).Decode(&entrada); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	creado := s.Acceso.Crear(entrada)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(creado)
}

// ObtenerAcceso maneja GET /api/v1/accesos/{id}
func (s *Server) ObtenerAcceso(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	acceso, ok := s.Acceso.Obtener(id)
	if !ok {
		http.Error(w, "Acceso no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(acceso)
}

// ActualizarAcceso maneja PUT /api/v1/accesos/{id}
func (s *Server) ActualizarAcceso(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var datos modelos.Acceso
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	actualizado, ok, err := s.Acceso.Actualizar(id, datos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if !ok {
		http.Error(w, "Acceso no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(actualizado)
}

// BorrarAcceso maneja DELETE /api/v1/accesos/{id}
func (s *Server) BorrarAcceso(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := s.Acceso.Borrar(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
