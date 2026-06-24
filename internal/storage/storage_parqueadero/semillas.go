package storage_parqueadero

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	"time"
)

func (m *Memoria) SeedParqueaderos() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.parqueaderos = []modelos.Parqueadero{
		{
			IDParqueadero: 1,
			Nombre:        "Parqueadero FCVT",
			Capacidad:     20,
			Tipo:          "Estudiantes",
		},
		{
			IDParqueadero: 2,
			Nombre:        "Parqueadero Docentes",
			Capacidad:     15,
			Tipo:          "Docentes",
		},
	}

	m.nextParqueaderoID = 3
}

func (m *Memoria) SeedEspacios() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.espacios = []modelos.Espacio{
		{
			IDEspacio:     1,
			IDParqueadero: 1,
			Numero:        1,
			Estado:        "Libre",
			TipoEspacio:   "Automovil",
		},
		{
			IDEspacio:     2,
			IDParqueadero: 1,
			Numero:        2,
			Estado:        "Ocupado",
			TipoEspacio:   "Automovil",
		},
		{
			IDEspacio:     3,
			IDParqueadero: 2,
			Numero:        1,
			Estado:        "Libre",
			TipoEspacio:   "Motocicleta",
		},
	}

	m.nextEspacioID = 4
}

func (m *Memoria) SeedOcupaciones() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.ocupaciones = []modelos.Ocupacion{
		{
			IDOcupacion:   1,
			PlacaVehiculo: "ABC1234",
			IDEspacio:     2,
			IDAcceso:      1,
			HoraInicio:    time.Now().Add(-2 * time.Hour),
		},
		{
			IDOcupacion:   2,
			PlacaVehiculo: "XYZ5678",
			IDEspacio:     3,
			IDAcceso:      2,
			HoraInicio:    time.Now().Add(-1 * time.Hour),
		},
	}

	m.nextOcupacionID = 3
}
