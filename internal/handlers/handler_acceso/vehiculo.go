package handler_acceso

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"proyecto_movilidad_fcvt/internal/modelos"
)

// ListarVehiculos maneja GET /api/v1/vehiculos
func (s *Server) ListarVehiculos(w http.ResponseWriter, r *http.Request) {
	vehiculos := s.Vehiculo.Listar()
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(vehiculos)
}

// CrearVehiculo maneja POST /api/v1/vehiculos
func (s *Server) CrearVehiculo(w http.ResponseWriter, r *http.Request) {
	var entrada modelos.Vehiculo
	if err := json.NewDecoder(r.Body).Decode(&entrada); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	creado, err := s.Vehiculo.Crear(entrada)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(creado)
}

// ObtenerVehiculo maneja GET /api/v1/vehiculos/{placa}
func (s *Server) ObtenerVehiculo(w http.ResponseWriter, r *http.Request) {
	placa := chi.URLParam(r, "placa")

	vehiculo, ok := s.Vehiculo.Obtener(placa)
	if !ok {
		http.Error(w, "Vehículo no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(vehiculo)
}

// ActualizarVehiculo maneja PUT /api/v1/vehiculos/{placa}
func (s *Server) ActualizarVehiculo(w http.ResponseWriter, r *http.Request) {
	placa := chi.URLParam(r, "placa")

	var datos modelos.Vehiculo
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	actualizado, ok, err := s.Vehiculo.Actualizar(placa, datos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !ok {
		http.Error(w, "Vehículo no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(actualizado)
}

// BorrarVehiculo maneja DELETE /api/v1/vehiculos/{placa}
func (s *Server) BorrarVehiculo(w http.ResponseWriter, r *http.Request) {
	placa := chi.URLParam(r, "placa")

	if err := s.Vehiculo.Borrar(placa); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
