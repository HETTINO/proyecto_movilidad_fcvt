package storage_parqueadero_test

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSQLite_CrearYListarParqueadero(t *testing.T) {
	repo := nuevoRepo(t) // ← directo, sin prefijo sqtp.
	repo.CrearParqueadero(modelos.Parqueadero{
		Nombre: "Parqueadero Norte", Capacidad: 50, Tipo: "cubierto",
	})
	lista := repo.ListarParqueaderos()
	assert.Len(t, lista, 1)
	assert.Equal(t, "Parqueadero Norte", lista[0].Nombre)
}

func TestSQLite_CrearYBuscarParqueaderoPorID(t *testing.T) {
	repo := nuevoRepo(t)
	creado, _ := repo.CrearParqueadero(modelos.Parqueadero{
		Nombre: "Parqueadero Sur", Capacidad: 30, Tipo: "abierto",
	})
	assert.NotZero(t, creado.IDParqueadero)
	encontrado, ok := repo.BuscarParqueaderoPorID(creado.IDParqueadero)
	assert.True(t, ok)
	assert.Equal(t, "Parqueadero Sur", encontrado.Nombre)
	assert.Equal(t, 30, encontrado.Capacidad)
}

func TestSQLite_BuscarParqueaderoInexistente(t *testing.T) {
	repo := nuevoRepo(t)
	_, ok := repo.BuscarParqueaderoPorID(999)
	assert.False(t, ok)
}
