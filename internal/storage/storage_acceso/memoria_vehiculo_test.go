package storage_acceso_test

import (
	"testing"

	"proyecto_movilidad_fcvt/internal/modelos"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_acceso"
)

func TestMemoria_CrearYBuscarVehiculo(t *testing.T) {
	m := storage.NuevoMemoriaAcceso()

	m.CrearVehiculo(modelos.Vehiculo{
		Placa:        "ABC-999",
		IDUsuario:    "0102030405",
		TipoVehiculo: "Automóvil",
		Marca:        "Toyota",
		Modelo:       "Corolla",
		Color:        "Blanco",
		Año:          2022,
	})

	encontrado, ok := m.BuscarVehiculoPorPlaca("ABC-999")
	if !ok {
		t.Fatalf("no se encontró el vehículo")
	}

	if encontrado.Marca != "Toyota" {
		t.Errorf("marca = %q; esperaba %q", encontrado.Marca, "Toyota")
	}
}

func TestMemoria_BorrarVehiculo(t *testing.T) {
	m := storage.NuevoMemoriaAcceso()

	m.CrearVehiculo(modelos.Vehiculo{
		Placa:        "MCE-2026",
		IDUsuario:    "0102030405",
		TipoVehiculo: "Automóvil",
	})

	if !m.BorrarVehiculo("MCE-2026") {
		t.Errorf("esperaba poder borrar el vehículo")
	}
}
