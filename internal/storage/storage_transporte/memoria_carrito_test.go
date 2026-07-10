package storage_test

import (
	"testing"

	"proyecto_movilidad_fcvt/internal/modelos"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_transporte"
)

func TestMemoria_CrearYBuscarCarrito(t *testing.T) {
	m := storage.NuevaMemoria()

	creado := m.CrearCarrito(modelos.Carrito{NombreCarrito: "Carrito Test", Capacidad: 4})
	if creado.ID == 0 {
		t.Fatalf("esperaba un ID asignado, obtuve 0")
	}

	encontrado, ok := m.BuscarCarritoPorID(creado.ID)
	if !ok {
		t.Fatalf("no se encontró el carrito recién creado (id=%d)", creado.ID)
	}
	if encontrado.NombreCarrito != "Carrito Test" {
		t.Errorf("nombre = %q; esperaba %q", encontrado.NombreCarrito, "Carrito Test")
	}
}

func TestMemoria_BuscarCarritoInexistente(t *testing.T) {
	m := storage.NuevaMemoria()

	if _, ok := m.BuscarCarritoPorID(999); ok {
		t.Errorf("esperaba ok=false para un id inexistente")
	}
}

func TestMemoria_ActualizarYBorrarCarrito(t *testing.T) {
	m := storage.NuevaMemoria()
	creado := m.CrearCarrito(modelos.Carrito{NombreCarrito: "Carrito Original", Capacidad: 4})

	_, ok := m.ActualizarCarrito(creado.ID, modelos.Carrito{NombreCarrito: "Carrito Editado", Capacidad: 6})
	if !ok {
		t.Fatalf("no se pudo actualizar el carrito id=%d", creado.ID)
	}

	if !m.BorrarCarrito(creado.ID) {
		t.Errorf("esperaba poder borrar el carrito id=%d", creado.ID)
	}
}