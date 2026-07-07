package storage_test

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSQLite_CrearRuta(t *testing.T) {
	repo := nuevoRepo(t)
	ruta := modelos.Ruta{
		Nombre:      "Ruta: Facultad FCVT - Tasty",
		Descripcion: "Recorrido desde la Facultad de Ciencias Informáticas hacia el Tasty",
	}

	creada := repo.CrearRuta(ruta)

	assert.NotZero(t, creada.ID)
	assert.Equal(t, "Ruta: Facultad FCVT - Tasty", creada.Nombre)
}

func TestSQLite_ListarRutas(t *testing.T) {
	repo := nuevoRepo(t)
	repo.CrearRuta(modelos.Ruta{Nombre: "Ruta 1"})
	repo.CrearRuta(modelos.Ruta{Nombre: "Ruta 2"})

	lista := repo.ListarRutas()

	assert.Len(t, lista, 2)
}

func TestSQLite_BuscarRutaPorID(t *testing.T) {
	repo := nuevoRepo(t)
	ruta := repo.CrearRuta(modelos.Ruta{
		Nombre:      "Ruta: Paraninfo - Facultad de Ingeniería",
		Descripcion: "Conexión entre el centro de eventos y el bloque de aulas",
	})

	encontrado, ok := repo.BuscarRutaPorID(ruta.ID)

	assert.True(t, ok)
	assert.Equal(t, ruta.ID, encontrado.ID)
	assert.Equal(t, "Ruta: Paraninfo - Facultad de Ingeniería", encontrado.Nombre)
}

func TestSQLite_ActualizarRuta(t *testing.T) {
	repo := nuevoRepo(t)

	// Crear inicial
	ruta := repo.CrearRuta(modelos.Ruta{
		Nombre:      "Ruta Antigua",
		Descripcion: "Desc vieja",
	})

	// Actualizar
	ruta.Nombre = "Ruta Nueva"
	ruta.Descripcion = "Desc actualizada"
	actualizado, ok := repo.ActualizarRuta(ruta.ID, ruta)

	assert.True(t, ok)
	assert.Equal(t, "Ruta Nueva", actualizado.Nombre)
	assert.Equal(t, "Desc actualizada", actualizado.Descripcion)
}

func TestSQLite_BorrarRuta(t *testing.T) {
	repo := nuevoRepo(t)

	// Crear
	ruta := repo.CrearRuta(modelos.Ruta{Nombre: "Para borrar"})

	// Borrar
	ok := repo.BorrarRuta(ruta.ID)
	assert.True(t, ok)

	// Verificar que ya no existe
	_, ok = repo.BuscarRutaPorID(ruta.ID)
	assert.False(t, ok)
}