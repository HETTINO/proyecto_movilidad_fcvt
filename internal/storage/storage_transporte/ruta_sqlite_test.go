package storage_test

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSQLite_CrearYListarRutas(t *testing.T) {
	repo := nuevoRepo(t)

	// Crear una ruta de prueba
	ruta := modelos.Ruta{
		Nombre:      "Ruta: Facultad FCVT - Tasty",
		Descripcion: "Recorrido desde la Facultad de Ciencias Informáticas hacia el Tasty",
	}
	repo.CrearRuta(ruta)

	// Verificar listado
	lista := repo.ListarRutas()
	assert.Len(t, lista, 1)
	assert.Equal(t, "Ruta: Facultad FCVT - Tasty", lista[0].Nombre)
}

func TestSQLite_CrearYBuscarRutaPorID(t *testing.T) {
	repo := nuevoRepo(t)

	// Crear
	ruta := repo.CrearRuta(modelos.Ruta{
		Nombre:      "Ruta: Paraninfo - Facultad de Ingeniería",
		Descripcion: "Conexión entre el centro de eventos y el bloque de aulas",
	})

	// Buscar
	encontrado, ok := repo.BuscarRutaPorID(ruta.ID)

	assert.True(t, ok)
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