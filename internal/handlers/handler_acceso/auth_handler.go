package handler_acceso

import (
	"encoding/json"
	"net/http"
)

// =========================================================
// DTOs de entrada/salida
// =========================================================

type registrarRequest struct {
	Cedula     string `json:"cedula"`
	Nombre     string `json:"nombre"`
	Email      string `json:"email"`
	Contrasena string `json:"contrasena"`
	Rol        string `json:"rol"`
}

type loginRequest struct {
	Cedula     string `json:"cedula"`
	Contrasena string `json:"contrasena"`
}

type loginResponse struct {
	Token string `json:"token"`
}

// =========================================================
// HANDLERS
// =========================================================

// Registrar maneja POST /api/v1/auth/register
func (s *Server) Registrar(w http.ResponseWriter, r *http.Request) {
	var req registrarRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	usuario, err := s.Auth.Registrar(req.Cedula, req.Nombre, req.Email, req.Contrasena, req.Rol)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(usuario)
}

// Login maneja POST /api/v1/auth/login
func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	token, err := s.Auth.Login(req.Cedula, req.Contrasena)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(loginResponse{Token: token})
}
