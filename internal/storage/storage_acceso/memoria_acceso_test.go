package storage_acceso_test

import (
	"testing"

	"proyecto_movilidad_fcvt/internal/modelos"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_acceso"
)

func TestMemoria_CrearYBuscarAcceso(t *testing.T) {
	m := storage.NuevoMemoriaAcceso()

	creado := m.CrearAcceso(modelos.Acceso{
		PlacaVehiculo: "ABC-1234",
		Estado:        "activo",
		Observaciones: "Prueba en memoria nativa",
	})

	if creado.ID == 0 {
		t.Fatalf("esperaba un ID asignado, obtuve 0")
	}

	encontrado, ok := m.BuscarAccesoPorID(creado.ID)
	if !ok {
		t.Fatalf("no se encontró el acceso recién creado (id=%d)", creado.ID)
	}

	if encontrado.Estado != "activo" {
		t.Errorf("estado = %q; esperaba %q", encontrado.Estado, "activo")
	}
}

func TestMemoria_BuscarAccesoInexistente(t *testing.T) {
	m := storage.NuevoMemoriaAcceso()

	if _, ok := m.BuscarAccesoPorID(999); ok {
		t.Errorf("esperaba ok=false para un id inexistente")
	}
}

func TestMemoria_ActualizarYBorrarAcceso(t *testing.T) {
	m := storage.NuevoMemoriaAcceso()

	creado := m.CrearAcceso(modelos.Acceso{
		PlacaVehiculo: "XYZ-9876",
		Estado:        "activo",
	})

	_, ok := m.ActualizarAcceso(creado.ID, modelos.Acceso{
		PlacaVehiculo: "XYZ-9876",
		Estado:        "finalizado",
	})

	if !ok {
		t.Fatalf("no se pudo actualizar el acceso id=%d", creado.ID)
	}

	if !m.BorrarAcceso(creado.ID) {
		t.Errorf("esperaba poder borrar el acceso id=%d", creado.ID)
	}

	if _, ok := m.BuscarAccesoPorID(creado.ID); ok {
		t.Errorf("el acceso id=%d debería haber sido borrado", creado.ID)
	}
}
