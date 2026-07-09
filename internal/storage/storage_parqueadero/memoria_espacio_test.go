package storage_parqueadero_test

import (
	"testing"

	"proyecto_movilidad_fcvt/internal/modelos"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_parqueadero"
)

func TestMemoria_CrearYBuscarEspacio(t *testing.T) {
	m := storage.NuevaMemoria()

	creado := m.CrearEspacio(modelos.Espacio{
		IDParqueadero: 1,
		Numero:        1,
		Estado:        "libre",
		TipoEspacio:   "auto",
	})
	if creado.IDEspacio == 0 {
		t.Fatalf("esperaba un ID asignado, obtuve 0")
	}

	encontrado, ok := m.BuscarEspacioPorID(creado.IDEspacio)
	if !ok {
		t.Fatalf("no se encontró el espacio recién creado (id=%d)", creado.IDEspacio)
	}
	if encontrado.Estado != "libre" {
		t.Errorf("estado = %q; esperaba %q", encontrado.Estado, "libre")
	}
}

func TestMemoria_BuscarEspacioInexistente(t *testing.T) {
	m := storage.NuevaMemoria()

	if _, ok := m.BuscarEspacioPorID(999); ok {
		t.Errorf("esperaba ok=false para un id inexistente")
	}
}

func TestMemoria_ActualizarYBorrarEspacio(t *testing.T) {
	m := storage.NuevaMemoria()
	creado := m.CrearEspacio(modelos.Espacio{
		IDParqueadero: 1,
		Numero:        2,
		Estado:        "libre",
		TipoEspacio:   "moto",
	})

	_, ok := m.ActualizarEspacio(creado.IDEspacio, modelos.Espacio{
		IDParqueadero: 1,
		Numero:        2,
		Estado:        "ocupado",
		TipoEspacio:   "moto",
	})
	if !ok {
		t.Fatalf("no se pudo actualizar el espacio id=%d", creado.IDEspacio)
	}

	if !m.BorrarEspacio(creado.IDEspacio) {
		t.Errorf("esperaba poder borrar el espacio id=%d", creado.IDEspacio)
	}
	if _, ok := m.BuscarEspacioPorID(creado.IDEspacio); ok {
		t.Errorf("el espacio id=%d debería haber sido borrado", creado.IDEspacio)
	}
}
