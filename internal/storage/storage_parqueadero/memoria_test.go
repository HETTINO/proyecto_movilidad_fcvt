package storage_parqueadero

import (
	"testing"

	"proyecto_movilidad_fcvt/internal/modelos"
)

// =========================================================
// PARQUEADEROS
// =========================================================

func TestMemoria_CrearYBuscarParqueadero(t *testing.T) {
	m := NuevaMemoria()

	creado := m.CrearParqueadero(modelos.Parqueadero{
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
	m := NuevaMemoria()

	if _, ok := m.BuscarParqueaderoPorID(999); ok {
		t.Errorf("esperaba ok=false para un id inexistente")
	}
}

func TestMemoria_ActualizarYBorrarParqueadero(t *testing.T) {
	m := NuevaMemoria()
	creado := m.CrearParqueadero(modelos.Parqueadero{
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

// =========================================================
// ESPACIOS
// =========================================================

func TestMemoria_CrearYBuscarEspacio(t *testing.T) {
	m := NuevaMemoria()

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
	m := NuevaMemoria()

	if _, ok := m.BuscarEspacioPorID(999); ok {
		t.Errorf("esperaba ok=false para un id inexistente")
	}
}

func TestMemoria_ActualizarYBorrarEspacio(t *testing.T) {
	m := NuevaMemoria()
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

// =========================================================
// OCUPACIONES
// =========================================================

func TestMemoria_CrearYBuscarOcupacion(t *testing.T) {
	m := NuevaMemoria()

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
	m := NuevaMemoria()
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
	m := NuevaMemoria()
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
