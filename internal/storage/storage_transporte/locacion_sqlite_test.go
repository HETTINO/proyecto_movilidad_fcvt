package storage_test

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSQLite_RegistrarYListarLocaciones(t *testing.T) {
	repo := nuevoRepo(t)

	// Registrar una locación (ej. Carrito en Paraninfo)
	loc := modelos.Locacion{
		Latitud:   -0.950,
		Longitud:  -80.750,
		TimeStamp: time.Now(),
		CarritoID: 1,
	}
	repo.RegistrarLocacion(loc)

	// Verificar listado
	lista := repo.ListarLocaciones()
	assert.Len(t, lista, 1)
	assert.Equal(t, -0.950, lista[0].Latitud)
	assert.Equal(t, 1, lista[0].CarritoID)
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