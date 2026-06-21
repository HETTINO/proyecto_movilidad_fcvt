package service

import (
	"errors"
	"proyecto_movilidad_fcvt/internal/modelos"
	"proyecto_movilidad_fcvt/internal/storage/storage_acceso"
)

type AccesoServicio struct {
	storage *storage_acceso.MemoriaAcceso // <-- Cambiado de *Almacen a *MemoriaAcceso
}

// NuevoAccesoServicio inicializa el servicio inyectando el almacenamiento en memoria
func NuevoAccesoServicio(str *storage_acceso.MemoriaAcceso) *AccesoServicio { // <-- Cambiado aquí también
	return &AccesoServicio{storage: str}
}

// ==========================================
// MÉTODOS DEL TRANSACCIONAL (Entrada y Salida)
// ==========================================

// RegistrarIngreso verifica si el auto puede entrar y crea el registro
func (s *AccesoServicio) RegistrarIngreso(acceso *modelos.Acceso) error {
	if acceso.PlacaVehiculo == "" {
		return errors.New("la placa del vehículo es obligatoria")
	}
	acceso.Estado = "ENTRADA"
	s.storage.CrearAcceso(*acceso) // <-- Llama al método correcto del CRUD en memoria
	return nil
}

// RegistrarEgresoVehiculo procesa la salida
func (s *AccesoServicio) RegistrarEgresoVehiculo(placa string) error {
	if placa == "" {
		return errors.New("la placa del vehículo es requerida")
	}

	// Busca el acceso activo en memoria y lo actualiza
	accesos := s.storage.ListarAccesos()
	for _, a := range accesos {
		if a.PlacaVehiculo == placa && a.TiempoSalida == nil {
			a.Estado = "salido"
			s.storage.ActualizarAcceso(a.ID, a)
			return nil
		}
	}
	return errors.New("no se encontró un acceso activo para esa placa")
}

// ==========================================
// MÉTODOS DE PUENTE PARA ENTIDADES ADICIONALES
// ==========================================

// GuardarUsuario delega la persistencia del usuario al Almacen en memoria
func (s *AccesoServicio) GuardarUsuario(u *modelos.Usuario) error {
	if u.Cedula == "" {
		return errors.New("la cédula del usuario es obligatoria")
	}
	s.storage.CrearUsuario(*u)
	return nil
}

// GuardarVehiculo delega la persistencia del vehículo al Almacen en memoria
func (s *AccesoServicio) GuardarVehiculo(v *modelos.Vehiculo) error {
	if v.Placa == "" {
		return errors.New("la placa del vehículo es obligatoria")
	}
	s.storage.CrearVehiculo(*v)
	return nil
}
