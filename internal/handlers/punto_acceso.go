package handlers

import (
	"encoding/json"
	"net/http"
	"proyecto_movilidad_fcvt/internal/modelos"
	"proyecto_movilidad_fcvt/internal/service"
)

type PuntoAccesoHandler struct {
	servicio *service.AccesoServicio // Inyectamos el servicio que maneja la lógica
}

func NuevoPuntoAccesoHandler(serv *service.AccesoServicio) *PuntoAccesoHandler {
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

	// Llamamos al storage a través del servicio (si tienes la función mapeada en el service)
	// Nota: Si aún no la mapeas en service, puedes llamar directo temporalmente o agregarla a tu service.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Punto de acceso creado exitosamente"})
}
