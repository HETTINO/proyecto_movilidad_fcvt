package storage_test

import (
	"testing"

	"proyecto_movilidad_fcvt/internal/modelos"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_transporte"
)

func TestMemoria_CrearYBuscarSolicitud(t *testing.T) {
	m := storage.NuevaMemoria()

	creada := m.CrearSolicitud(modelos.Solicitud{CedulaUsuario: "1234567890", CantPersonas: 2, ParadaOrigen: 1})
	if creada.ID == 0 {
		t.Fatalf("esperaba un ID asignado, obtuve 0")
	}

	encontrada, ok := m.BuscarSolicitudPorID(creada.ID)
	if !ok {
		t.Fatalf("no se encontró la solicitud recién creada (id=%d)", creada.ID)
	}
	if encontrada.CedulaUsuario != "1234567890" {
		t.Errorf("cedula = %q; esperaba %q", encontrada.CedulaUsuario, "1234567890")
	}
}

func TestMemoria_BuscarSolicitudInexistente(t *testing.T) {
	m := storage.NuevaMemoria()

	if _, ok := m.BuscarSolicitudPorID(999); ok {
		t.Errorf("esperaba ok=false para un id inexistente")
	}
}

func TestMemoria_ActualizarYBorrarSolicitud(t *testing.T) {
	m := storage.NuevaMemoria()
	creada := m.CrearSolicitud(modelos.Solicitud{CedulaUsuario: "1234567890", CantPersonas: 2, ParadaOrigen: 1})

	_, ok := m.ActualizarSolicitud(creada.ID, modelos.Solicitud{Estado: "asignada"})
	if !ok {
		t.Fatalf("no se pudo actualizar la solicitud id=%d", creada.ID)
	}

	if !m.BorrarSolicitud(creada.ID) {
		t.Errorf("esperaba poder borrar la solicitud id=%d", creada.ID)
	}
}
