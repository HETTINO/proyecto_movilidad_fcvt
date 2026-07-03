package handler_acceso

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"proyecto_movilidad_fcvt/internal/modelos"
)

// ListarUsuarios maneja GET /api/v1/usuarios
func (s *Server) ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	usuarios := s.Usuario.Listar()
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(usuarios)
}

// CrearUsuario maneja POST /api/v1/usuarios
func (s *Server) CrearUsuario(w http.ResponseWriter, r *http.Request) {
	var entrada modelos.Usuario
	if err := json.NewDecoder(r.Body).Decode(&entrada); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	creado, err := s.Usuario.Crear(entrada)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(creado)
}

// ObtenerUsuario maneja GET /api/v1/usuarios/{id}  (id = cedula)
func (s *Server) ObtenerUsuario(w http.ResponseWriter, r *http.Request) {
	cedula := chi.URLParam(r, "id")

	usuario, ok := s.Usuario.Obtener(cedula)
	if !ok {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(usuario)
}

// ActualizarUsuario maneja PUT /api/v1/usuarios/{id}  (id = cedula)
func (s *Server) ActualizarUsuario(w http.ResponseWriter, r *http.Request) {
	cedula := chi.URLParam(r, "id")

	var datos modelos.Usuario
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	actualizado, ok, err := s.Usuario.Actualizar(cedula, datos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !ok {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(actualizado)
}

// BorrarUsuario maneja DELETE /api/v1/usuarios/{id}  (id = cedula)
func (s *Server) BorrarUsuario(w http.ResponseWriter, r *http.Request) {
	cedula := chi.URLParam(r, "id")

	if err := s.Usuario.Borrar(cedula); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
