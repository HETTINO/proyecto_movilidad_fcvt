package handlers

import (
	"encoding/json"
	"net/http"
)

type credenciales struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *Server) Registrar(w http.ResponseWriter, r *http.Request) {
	var creds credenciales
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		responderJSON(w, http.StatusBadRequest, map[string]string{"error": "cuerpo inválido"})
		return
	}

	usuario, err := s.Auth.Registrar(creds.Email, creds.Password)
	if err != nil {
		responderJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	responderJSON(w, http.StatusCreated, usuario)
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	var creds credenciales
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		responderJSON(w, http.StatusBadRequest, map[string]string{"error": "cuerpo inválido"})
		return
	}

	token, err := s.Auth.Login(creds.Email, creds.Password)
	if err != nil {
		responderJSON(w, http.StatusUnauthorized, map[string]string{"error": err.Error()})
		return
	}

	responderJSON(w, http.StatusOK, map[string]string{"token": token})
}
