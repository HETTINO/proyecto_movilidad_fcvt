package handlers

import (
	"encoding/json"
	"net/http"
	"proyecto_movilidad_fcvt/internal/modelos"
	"proyecto_movilidad_fcvt/internal/service/service_acceso" // <-- Tu subcarpeta de módulo
)

type VehiculoHandler struct {
	servicio *service_acceso.VehiculoService // <-- Cambiado al servicio específico de Vehículos
}

func NuevoVehiculoHandler(serv *service_acceso.VehiculoService) *VehiculoHandler { // <-- Cambiado aquí también
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

	// Consumimos el método .Crear() estructurado de tu servicio de vehículos
	if _, err := h.servicio.Crear(nuevoVehiculo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]string{"mensaje": "Vehículo registrado exitosamente"})
}
