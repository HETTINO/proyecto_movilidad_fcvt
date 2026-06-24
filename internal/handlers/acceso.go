package handlers

import (
	"encoding/json"
	"net/http"
	"proyecto_movilidad_fcvt/internal/modelos"
	"proyecto_movilidad_fcvt/internal/service/service_acceso" // <-- Importa tu nueva subcarpeta
)

type AccesoHandler struct {
	servicio *service_acceso.AccesoService // <-- Cambiado al nuevo servicio estructurado
}

func NuevoAccesoHandler(serv *service_acceso.AccesoService) *AccesoHandler { // <-- Cambiado aquí también
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

	// Usamos .Crear() que es el método unificado en tu nuevo AccesoService
	if _, err := h.servicio.Crear(nuevoAcceso); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]string{"mensaje": "Entrada registrada exitosamente"})
}

// RegistrarSalidaHandler maneja POST /api/acceso/salida
func (h *AccesoHandler) RegistrarSalidaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var data struct {
		ID int `json:"id"` // Ajustado para buscar por ID o el campo de datos que use tu tabla
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Primero obtenemos el registro de acceso para poder actualizarlo
	accesoExistente, encontrado := h.servicio.Obtener(data.ID)
	if !encontrado {
		http.Error(w, "Registro de acceso no encontrado", http.StatusNotFound)
		return
	}

	// Modificamos el estado o datos necesarios usando tu nuevo método unificado .Actualizar()
	if _, _, err := h.servicio.Actualizar(data.ID, accesoExistente); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"mensaje": "Salida registrada exitosamente"})
}
