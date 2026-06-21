package handlers

import (
	"encoding/json"
	"net/http"
	"proyecto_movilidad_fcvt/internal/modelos"
	"proyecto_movilidad_fcvt/internal/service"
)

type AccesoHandler struct {
	servicio *service.AccesoServicio
}

func NuevoAccesoHandler(serv *service.AccesoServicio) *AccesoHandler {
	return &AccesoHandler{servicio: serv}
}

// RegistrarEntradaHandler maneja POST /api/acceso/entrada
func (h *AccesoHandler) RegistrarEntradaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var nuevoAcceso modelos.Acceso
	if err := json.NewDecoder(r.Body).Decode(&nuevoAcceso); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if err := h.servicio.RegistrarIngreso(&nuevoAcceso); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Entrada registrada exitosamente"})
}

// RegistrarSalidaHandler maneja POST /api/acceso/salida
func (h *AccesoHandler) RegistrarSalidaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var data struct {
		Placa string `json:"placa_vehiculo"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if err := h.servicio.RegistrarEgresoVehiculo(data.Placa); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Salida registrada exitosamente"})
}
