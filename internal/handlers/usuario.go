package handlers

import (
	"encoding/json"
	"net/http"
	"proyecto_movilidad_fcvt/internal/modelos"
	"proyecto_movilidad_fcvt/internal/service/service_acceso" // <-- Subcarpeta de tu módulo
)

type UsuarioHandler struct {
	servicio *service_acceso.UsuarioService // <-- Cambiado al nuevo servicio específico
}

func NuevoUsuarioHandler(serv *service_acceso.UsuarioService) *UsuarioHandler { // <-- Cambiado aquí también
	return &UsuarioHandler{servicio: serv}
}

// RegistrarUsuarioHandler maneja POST /api/usuarios
func (h *UsuarioHandler) RegistrarUsuarioHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var nuevoUsuario modelos.Usuario
	if err := json.NewDecoder(r.Body).Decode(&nuevoUsuario); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Consumimos el método .Crear() estructurado de tu servicio de acceso
	if _, err := h.servicio.Crear(nuevoUsuario); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]string{"mensaje": "Usuario registrado exitosamente"})
}
