package storage_test

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSQLite_CrearCarrito(t *testing.T) {
	repo := nuevoRepo(t)
	c := modelos.Carrito{
		NombreCarrito: "Carrito 1 - Rectorado",
		Capacidad:     5,
		Estado:        "disponible",
	}

	creado := repo.CrearCarrito(c)

	assert.NotZero(t, creado.ID)
	assert.Equal(t, "Carrito 1 - Rectorado", creado.NombreCarrito)
}

func TestSQLite_ListarCarritos(t *testing.T) {
	repo := nuevoRepo(t)
	repo.CrearCarrito(modelos.Carrito{NombreCarrito: "C1"})
	repo.CrearCarrito(modelos.Carrito{NombreCarrito: "C2"})

	lista := repo.ListarCarritos()

	assert.Len(t, lista, 2)
}

func TestSQLite_BuscarCarritoPorID(t *testing.T) {
	repo := nuevoRepo(t)
	c := repo.CrearCarrito(modelos.Carrito{
		NombreCarrito: "Carrito 2 - FCVT",
		Capacidad:     3,
	})

	encontrado, ok := repo.BuscarCarritoPorID(c.ID)

	assert.True(t, ok)
	assert.Equal(t, c.ID, encontrado.ID)
	assert.Equal(t, "Carrito 2 - FCVT", encontrado.NombreCarrito)
}

func TestSQLite_ActualizarCarrito(t *testing.T) {
	repo := nuevoRepo(t)
	c := repo.CrearCarrito(modelos.Carrito{
		NombreCarrito: "Carrito 3 - Biblioteca",
		Capacidad:     5,
		Estado:        "mantenimiento",
	})

	c.Estado = "disponible"
	c.Capacidad = 3
	actualizado, ok := repo.ActualizarCarrito(c.ID, c)

	assert.True(t, ok)
	assert.Equal(t, "disponible", actualizado.Estado)
	assert.Equal(t, 3, actualizado.Capacidad)
}

func TestSQLite_BorrarCarrito(t *testing.T) {
	repo := nuevoRepo(t)
	c := repo.CrearCarrito(modelos.Carrito{NombreCarrito: "Para borrar"})

	ok := repo.BorrarCarrito(c.ID)
	assert.True(t, ok)

	_, existe := repo.BuscarCarritoPorID(c.ID)
	assert.False(t, existe)
}
