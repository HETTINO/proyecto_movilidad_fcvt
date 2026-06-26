package sqlite_test_parqueadero

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSQLite_CrearYListarEspacio(t *testing.T) {
	repo := nuevoRepo(t)
	park := repo.CrearParqueadero(modelos.Parqueadero{
		Nombre: "Parqueadero Central", Capacidad: 20, Tipo: "cubierto",
	})
	repo.CrearEspacio(modelos.Espacio{
		IDParqueadero: park.IDParqueadero, Numero: 1,
		Estado: "libre", TipoEspacio: "auto",
	})
	lista := repo.ListarEspacios()
	assert.Len(t, lista, 1)
	assert.Equal(t, "libre", lista[0].Estado)
}

func TestSQLite_CrearYBuscarEspacioPorID(t *testing.T) {
	repo := nuevoRepo(t)
	park := repo.CrearParqueadero(modelos.Parqueadero{
		Nombre: "Parqueadero Este", Capacidad: 10, Tipo: "abierto",
	})
	creado := repo.CrearEspacio(modelos.Espacio{
		IDParqueadero: park.IDParqueadero, Numero: 5,
		Estado: "ocupado", TipoEspacio: "moto",
	})
	assert.NotZero(t, creado.IDEspacio)
	encontrado, ok := repo.BuscarEspacioPorID(creado.IDEspacio)
	assert.True(t, ok)
	assert.Equal(t, "ocupado", encontrado.Estado)
	assert.Equal(t, "moto", encontrado.TipoEspacio)
}

func TestSQLite_BuscarEspacioInexistente(t *testing.T) {
	repo := nuevoRepo(t)
	_, ok := repo.BuscarEspacioPorID(999)
	assert.False(t, ok)
}

func TestSQLite_ListarEspaciosPorParqueadero(t *testing.T) {
	repo := nuevoRepo(t)
	park1 := repo.CrearParqueadero(modelos.Parqueadero{Nombre: "Norte", Capacidad: 10, Tipo: "cubierto"})
	park2 := repo.CrearParqueadero(modelos.Parqueadero{Nombre: "Sur", Capacidad: 10, Tipo: "abierto"})
	repo.CrearEspacio(modelos.Espacio{IDParqueadero: park1.IDParqueadero, Numero: 1, Estado: "libre", TipoEspacio: "auto"})
	repo.CrearEspacio(modelos.Espacio{IDParqueadero: park1.IDParqueadero, Numero: 2, Estado: "libre", TipoEspacio: "moto"})
	repo.CrearEspacio(modelos.Espacio{IDParqueadero: park2.IDParqueadero, Numero: 1, Estado: "libre", TipoEspacio: "auto"})
	espaciosPark1 := repo.ListarEspaciosPorParqueadero(park1.IDParqueadero)
	assert.Len(t, espaciosPark1, 2)
}
