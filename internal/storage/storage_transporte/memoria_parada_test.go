package storage_test

import (
	"testing"

	"proyecto_movilidad_fcvt/internal/modelos"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_transporte"
)

func TestMemoria_CrearYBuscarParada(t *testing.T) {
	m := storage.NuevaMemoria()

	creada := m.CrearParada(modelos.Parada{Nombre: "Parada Test", Latitud: 1, Longitud: 2})
	if creada.IDParada == 0 {
		t.Fatalf("esperaba un ID asignado, obtuve 0")
	}

	encontrada, ok := m.BuscarParadaPorID(creada.IDParada)
	if !ok {
		t.Fatalf("no se encontró la parada recién creada (id=%d)", creada.IDParada)
	}
	if encontrada.Nombre != "Parada Test" {
		t.Errorf("nombre = %q; esperaba %q", encontrada.Nombre, "Parada Test")
	}
}

func TestMemoria_BuscarParadaInexistente(t *testing.T) {
	m := storage.NuevaMemoria()

	if _, ok := m.BuscarParadaPorID(999); ok {
		t.Errorf("esperaba ok=false para un id inexistente")
	}
}

func TestMemoria_ActualizarYBorrarParada(t *testing.T) {
	m := storage.NuevaMemoria()
	creada := m.CrearParada(modelos.Parada{Nombre: "Parada Original", Latitud: 1, Longitud: 2})

	_, ok := m.ActualizarParada(creada.IDParada, modelos.Parada{Nombre: "Parada Editada", Latitud: 3, Longitud: 4})
	if !ok {
		t.Fatalf("no se pudo actualizar la parada id=%d", creada.IDParada)
	}

	if !m.BorrarParada(creada.IDParada) {
		t.Errorf("esperaba poder borrar la parada id=%d", creada.IDParada)
	}
}