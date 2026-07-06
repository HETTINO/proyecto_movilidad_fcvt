package storage_test

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSQLite_RegistrarLocacion(t *testing.T) {
	repo := nuevoRepo(t)
	loc := modelos.Locacion{
		Latitud:   -0.950,
		Longitud:  -80.750,
		TimeStamp: time.Now(),
		CarritoID: 1,
	}

	registrada := repo.RegistrarLocacion(loc)

	assert.NotZero(t, registrada.ID)
	assert.Equal(t, loc.Latitud, registrada.Latitud)
}

func TestSQLite_ListarLocaciones(t *testing.T) {
	repo := nuevoRepo(t)
	repo.RegistrarLocacion(modelos.Locacion{CarritoID: 1})
	repo.RegistrarLocacion(modelos.Locacion{CarritoID: 2})

	lista := repo.ListarLocaciones()

	assert.Len(t, lista, 2)
}

func TestSQLite_ObtenerUltimaLocacionPorCarrito(t *testing.T) {
	repo := nuevoRepo(t)

	// Registrar dos locaciones para el mismo carrito
	carritoID := 1
	loc1 := modelos.Locacion{
		Latitud:   1.0,
		Longitud:  1.0,
		TimeStamp: time.Now().Add(-1 * time.Hour), // Más antigua
		CarritoID: carritoID,
	}
	loc2 := modelos.Locacion{
		Latitud:   2.0, // Más reciente
		Longitud:  2.0,
		TimeStamp: time.Now(),
		CarritoID: carritoID,
	}

	repo.RegistrarLocacion(loc1)
	repo.RegistrarLocacion(loc2)

	// Buscar la última
	encontrada, ok := repo.ObtenerUltimaLocacionPorCarrito(carritoID)

	assert.True(t, ok)
	assert.Equal(t, 2.0, encontrada.Latitud) // Debe ser la loc2
	assert.Equal(t, carritoID, encontrada.CarritoID)
}

func TestSQLite_ObtenerUltimaLocacion_NoEncontrada(t *testing.T) {
	repo := nuevoRepo(t)

	// Buscar un carrito que no tiene locaciones registradas
	_, ok := repo.ObtenerUltimaLocacionPorCarrito(999)

	assert.False(t, ok)
}