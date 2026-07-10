package storage_test

import (
	"testing"

	"proyecto_movilidad_fcvt/internal/modelos"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_transporte"
)

func TestMemoria_CrearYBuscarRuta(t *testing.T) {
	m := storage.NuevaMemoria()

	creada := m.CrearRuta(modelos.Ruta{Nombre: "Ruta Test", Descripcion: "desc"})
	if creada.ID == 0 {
		t.Fatalf("esperaba un ID asignado, obtuve 0")
	}

	encontrada, ok := m.BuscarRutaPorID(creada.ID)
	if !ok {
		t.Fatalf("no se encontró la ruta recién creada (id=%d)", creada.ID)
	}
	if encontrada.Nombre != "Ruta Test" {
		t.Errorf("nombre = %q; esperaba %q", encontrada.Nombre, "Ruta Test")
	}
}

func TestMemoria_BuscarRutaInexistente(t *testing.T) {
	m := storage.NuevaMemoria()

	if _, ok := m.BuscarRutaPorID(999); ok {
		t.Errorf("esperaba ok=false para un id inexistente")
	}
}

func TestMemoria_ActualizarYBorrarRuta(t *testing.T) {
	m := storage.NuevaMemoria()
	creada := m.CrearRuta(modelos.Ruta{Nombre: "Ruta Original", Descripcion: "desc"})

	_, ok := m.ActualizarRuta(creada.ID, modelos.Ruta{Nombre: "Ruta Editada", Descripcion: "otra desc"})
	if !ok {
		t.Fatalf("no se pudo actualizar la ruta id=%d", creada.ID)
	}

	if !m.BorrarRuta(creada.ID) {
		t.Errorf("esperaba poder borrar la ruta id=%d", creada.ID)
	}
	if _, ok := m.BuscarRutaPorID(creada.ID); ok {
		t.Errorf("la ruta id=%d debería haber sido borrada", creada.ID)
	}
}
