package storage_parqueadero_test

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSQLite_CrearYListarOcupacion(t *testing.T) {
	repo := nuevoRepo(t)
	park := repo.CrearParqueadero(modelos.Parqueadero{
		Nombre: "Parqueadero Oeste", Capacidad: 15, Tipo: "cubierto",
	})
	espacio := repo.CrearEspacio(modelos.Espacio{
		IDParqueadero: park.IDParqueadero, Numero: 1,
		Estado: "libre", TipoEspacio: "auto",
	})
	repo.CrearOcupacion(modelos.Ocupacion{
		PlacaVehiculo: "ABC-1234", IDEspacio: espacio.IDEspacio, IDAcceso: 1,
	})
	lista := repo.ListarOcupaciones()
	assert.Len(t, lista, 1)
	assert.Equal(t, "ABC-1234", lista[0].PlacaVehiculo)
}

func TestSQLite_CrearYBuscarOcupacionPorID(t *testing.T) {
	repo := nuevoRepo(t)
	park := repo.CrearParqueadero(modelos.Parqueadero{Nombre: "P1", Capacidad: 5, Tipo: "abierto"})
	espacio := repo.CrearEspacio(modelos.Espacio{
		IDParqueadero: park.IDParqueadero, Numero: 2,
		Estado: "libre", TipoEspacio: "moto",
	})
	creada := repo.CrearOcupacion(modelos.Ocupacion{
		PlacaVehiculo: "XYZ-9999", IDEspacio: espacio.IDEspacio, IDAcceso: 2,
	})
	assert.NotZero(t, creada.IDOcupacion)
	encontrada, ok := repo.BuscarOcupacionPorID(creada.IDOcupacion)
	assert.True(t, ok)
	assert.Equal(t, "XYZ-9999", encontrada.PlacaVehiculo)
}

func TestSQLite_BuscarOcupacionInexistente(t *testing.T) {
	repo := nuevoRepo(t)
	_, ok := repo.BuscarOcupacionPorID(999)
	assert.False(t, ok)
}

func TestSQLite_LiberarOcupacion(t *testing.T) {
	repo := nuevoRepo(t)
	park := repo.CrearParqueadero(modelos.Parqueadero{Nombre: "P2", Capacidad: 5, Tipo: "cubierto"})
	espacio := repo.CrearEspacio(modelos.Espacio{
		IDParqueadero: park.IDParqueadero, Numero: 3,
		Estado: "ocupado", TipoEspacio: "auto",
	})
	creada := repo.CrearOcupacion(modelos.Ocupacion{
		PlacaVehiculo: "LMN-5555", IDEspacio: espacio.IDEspacio, IDAcceso: 3,
	})
	assert.Nil(t, creada.HoraFin)
	liberada, ok := repo.LiberarOcupacion(creada.IDOcupacion)
	assert.True(t, ok)
	assert.NotNil(t, liberada.HoraFin)
}
