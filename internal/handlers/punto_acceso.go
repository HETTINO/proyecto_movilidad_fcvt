package handlers

import (
	"encoding/json"
	"net/http"
	"proyecto_movilidad_fcvt/internal/modelos"
	"proyecto_movilidad_fcvt/internal/service/service_acceso" // <-- Importa tu subcarpeta corregida
)

type PuntoAccesoHandler struct {
	servicio *service_acceso.PuntoAccesoService // <-- Cambiado al servicio específico del Punto de Acceso
}

func NuevoPuntoAccesoHandler(serv *service_acceso.PuntoAccesoService) *PuntoAccesoHandler { // <-- Cambiado aquí también
	return &PuntoAccesoHandler{servicio: serv}
}

// CrearPuntoHandler maneja POST /api/puntos-acceso
func (h *PuntoAccesoHandler) CrearPuntoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var nuevoPunto modelos.PuntoDeAcceso
	if err := json.NewDecoder(r.Body).Decode(&nuevoPunto); err != nil {
		http.Error(w, "Cuerpo de petición inválido", http.StatusBadRequest)
		return
	}

	// Conectamos la lógica real llamando al método Crear del servicio estructurado
	if _, err := h.servicio.Crear(nuevoPunto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]string{"mensaje": "Punto de acceso creado exitosamente"})
}
