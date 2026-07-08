package storage_parqueadero_test

import (
	"testing"

	"proyecto_movilidad_fcvt/internal/modelos"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_parqueadero"
)

func TestMemoria_CrearYBuscarParqueadero(t *testing.T) {
	m := storage.NuevaMemoria()

	creado, _ := m.CrearParqueadero(modelos.Parqueadero{
		Nombre:    "Parqueadero Norte",
		Capacidad: 50,
		Tipo:      "cubierto",
	})
	if creado.IDParqueadero == 0 {
		t.Fatalf("esperaba un ID asignado, obtuve 0")
	}

	encontrado, ok := m.BuscarParqueaderoPorID(creado.IDParqueadero)
	if !ok {
		t.Fatalf("no se encontró el parqueadero recién creado (id=%d)", creado.IDParqueadero)
	}
	if encontrado.Nombre != "Parqueadero Norte" {
		t.Errorf("nombre = %q; esperaba %q", encontrado.Nombre, "Parqueadero Norte")
	}
}

func TestMemoria_BuscarParqueaderoInexistente(t *testing.T) {
	m := storage.NuevaMemoria()

	if _, ok := m.BuscarParqueaderoPorID(999); ok {
		t.Errorf("esperaba ok=false para un id inexistente")
	}
}

func TestMemoria_ActualizarYBorrarParqueadero(t *testing.T) {
	m := storage.NuevaMemoria()
	creado, _ := m.CrearParqueadero(modelos.Parqueadero{
		Nombre:    "Parqueadero Sur",
		Capacidad: 30,
		Tipo:      "abierto",
	})

	_, ok := m.ActualizarParqueadero(creado.IDParqueadero, modelos.Parqueadero{
		Nombre:    "Parqueadero Sur Ampliado",
		Capacidad: 60,
		Tipo:      "cubierto",
	})
	if !ok {
		t.Fatalf("no se pudo actualizar el parqueadero id=%d", creado.IDParqueadero)
	}

	if !m.BorrarParqueadero(creado.IDParqueadero) {
		t.Errorf("esperaba poder borrar el parqueadero id=%d", creado.IDParqueadero)
	}
	if _, ok := m.BuscarParqueaderoPorID(creado.IDParqueadero); ok {
		t.Errorf("el parqueadero id=%d debería haber sido borrado", creado.IDParqueadero)
	}
}
