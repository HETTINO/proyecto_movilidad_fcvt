package storage_test_parqueadero

import (
	"testing"

	"proyecto_movilidad_fcvt/internal/modelos"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_parqueadero"
)

func TestMemoria_CrearYBuscarOcupacion(t *testing.T) {
	m := storage.NuevaMemoria()

	creada := m.CrearOcupacion(modelos.Ocupacion{
		PlacaVehiculo: "ABC-1234",
		IDEspacio:     1,
		IDAcceso:      1,
	})
	if creada.IDOcupacion == 0 {
		t.Fatalf("esperaba un ID asignado, obtuve 0")
	}
	if creada.HoraInicio.IsZero() {
		t.Errorf("esperaba HoraInicio asignada automáticamente")
	}

	encontrada, ok := m.BuscarOcupacionPorID(creada.IDOcupacion)
	if !ok {
		t.Fatalf("no se encontró la ocupación recién creada (id=%d)", creada.IDOcupacion)
	}
	if encontrada.PlacaVehiculo != "ABC-1234" {
		t.Errorf("placa = %q; esperaba %q", encontrada.PlacaVehiculo, "ABC-1234")
	}
}

func TestMemoria_LiberarOcupacion(t *testing.T) {
	m := storage.NuevaMemoria()
	creada := m.CrearOcupacion(modelos.Ocupacion{
		PlacaVehiculo: "XYZ-9999",
		IDEspacio:     1,
		IDAcceso:      1,
	})

	// HoraFin debe ser nil antes de liberar
	if creada.HoraFin != nil {
		t.Errorf("esperaba HoraFin nil antes de liberar")
	}

	liberada, ok := m.LiberarOcupacion(creada.IDOcupacion)
	if !ok {
		t.Fatalf("no se pudo liberar la ocupación id=%d", creada.IDOcupacion)
	}
	if liberada.HoraFin == nil {
		t.Errorf("esperaba HoraFin asignada después de liberar")
	}
}

func TestMemoria_BorrarOcupacion(t *testing.T) {
	m := storage.NuevaMemoria()
	creada := m.CrearOcupacion(modelos.Ocupacion{
		PlacaVehiculo: "DEL-0001",
		IDEspacio:     1,
		IDAcceso:      1,
	})

	if !m.BorrarOcupacion(creada.IDOcupacion) {
		t.Errorf("esperaba poder borrar la ocupación id=%d", creada.IDOcupacion)
	}
	if _, ok := m.BuscarOcupacionPorID(creada.IDOcupacion); ok {
		t.Errorf("la ocupación id=%d debería haber sido borrada", creada.IDOcupacion)
	}
}
