package storage_test

import (
	"testing"

	"proyecto_movilidad_fcvt/internal/modelos"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_transporte"
)

func TestMemoria_RegistrarYObtenerUltimaLocacion(t *testing.T) {
	m := storage.NuevaMemoria()

	registrada := m.RegistrarLocacion(modelos.Locacion{Latitud: 1, Longitud: 2, CarritoID: 5})
	if registrada.ID == 0 {
		t.Fatalf("esperaba un ID asignado, obtuve 0")
	}

	ultima, ok := m.ObtenerUltimaLocacionPorCarrito(5)
	if !ok {
		t.Fatalf("no se encontró la última locación del carrito 5")
	}
	if ultima.ID != registrada.ID {
		t.Errorf("id = %d; esperaba %d", ultima.ID, registrada.ID)
	}
}

func TestMemoria_ObtenerUltimaLocacionInexistente(t *testing.T) {
	m := storage.NuevaMemoria()

	if _, ok := m.ObtenerUltimaLocacionPorCarrito(999); ok {
		t.Errorf("esperaba ok=false para un carrito sin locaciones")
	}
}