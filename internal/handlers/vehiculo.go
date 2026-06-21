package handlers

import (
	"encoding/json"
	"net/http"
	"proyecto_movilidad_fcvt/internal/modelos"
	"proyecto_movilidad_fcvt/internal/service"
)

type VehiculoHandler struct {
	servicio *service.AccesoServicio
}

func NuevoVehiculoHandler(serv *service.AccesoServicio) *VehiculoHandler {
	return &VehiculoHandler{servicio: serv}
}

// RegistrarVehiculoHandler maneja POST /api/vehiculos
func (h *VehiculoHandler) RegistrarVehiculoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var nuevoVehiculo modelos.Vehiculo
	if err := json.NewDecoder(r.Body).Decode(&nuevoVehiculo); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if err := h.servicio.GuardarVehiculo(&nuevoVehiculo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Vehículo registrado exitosamente"})
}
