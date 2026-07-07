package storage_test

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSQLite_CrearParada(t *testing.T) {
	repo := nuevoRepo(t)
	p := modelos.Parada{
		Nombre:   "Parada: Facultad de Ciencias Informáticas",
		Latitud:  -0.9203,
		Longitud: -80.7346,
	}
	
	creada := repo.CrearParada(p)
	
	assert.NotZero(t, creada.IDParada)
	assert.Equal(t, "Parada: Facultad de Ciencias Informáticas", creada.Nombre)
}

func TestSQLite_ListarParadas(t *testing.T) {
	repo := nuevoRepo(t)
	
	// Preparar datos
	repo.CrearParada(modelos.Parada{Nombre: "Parada 1"})
	repo.CrearParada(modelos.Parada{Nombre: "Parada 2"})
	
	lista := repo.ListarParadas()
	
	assert.Len(t, lista, 2)
}

func TestSQLite_BuscarParadaPorID(t *testing.T) {
	repo := nuevoRepo(t)
	p := repo.CrearParada(modelos.Parada{Nombre: "Parada Buscar"})

	encontrado, ok := repo.BuscarParadaPorID(p.IDParada)

	assert.True(t, ok)
	assert.Equal(t, p.IDParada, encontrado.IDParada)
}

func TestSQLite_ActualizarParada(t *testing.T) {
	repo := nuevoRepo(t)
	p := repo.CrearParada(modelos.Parada{Nombre: "Nombre Viejo"})

	p.Nombre = "Nombre Nuevo"
	actualizado, ok := repo.ActualizarParada(p.IDParada, p)

	assert.True(t, ok)
	assert.Equal(t, "Nombre Nuevo", actualizado.Nombre)
}

func TestSQLite_BorrarParada(t *testing.T) {
	repo := nuevoRepo(t)
	p := repo.CrearParada(modelos.Parada{Nombre: "Para borrar"})

	ok := repo.BorrarParada(p.IDParada)
	assert.True(t, ok)

	_, existe := repo.BuscarParadaPorID(p.IDParada)
	assert.False(t, existe)
}