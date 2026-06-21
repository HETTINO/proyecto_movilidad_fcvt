package handlers

import (
	"encoding/json"
	"net/http"
	"proyecto_movilidad_fcvt/internal/modelos"
	"proyecto_movilidad_fcvt/internal/service"
)

type UsuarioHandler struct {
	servicio *service.AccesoServicio
}

func NuevoUsuarioHandler(serv *service.AccesoServicio) *UsuarioHandler {
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

	// Llama a la persistencia a través del servicio
	if err := h.servicio.GuardarUsuario(&nuevoUsuario); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Usuario registrado exitosamente"})
}
