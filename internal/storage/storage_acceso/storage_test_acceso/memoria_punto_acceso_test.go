package storage_test_acceso

import (
	"testing"

	"proyecto_movilidad_fcvt/internal/modelos"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_acceso"
)

func TestMemoria_CrearYBuscarPuntoAcceso(t *testing.T) {
	m := storage.NuevoMemoriaAcceso()

	m.CrearPuntoAcceso(modelos.PuntoDeAcceso{
		ID:         1,
		Frecuencia: "Vehicular",
		Ubicacion:  "Bloque Sur",
	})

	encontrado, ok := m.BuscarPuntoAccesoPorID(1)
	if !ok {
		t.Fatalf("no se encontró el punto de acceso con ID=1")
	}

	if encontrado.Ubicacion != "Bloque Sur" {
		t.Errorf("ubicacion = %q; esperaba %q", encontrado.Ubicacion, "Bloque Sur")
	}
}

func TestMemoria_BuscarPuntoAccesoInexistente(t *testing.T) {
	m := storage.NuevoMemoriaAcceso()

	if _, ok := m.BuscarPuntoAccesoPorID(999); ok {
		t.Errorf("esperaba ok=false para un ID inexistente")
	}
}
