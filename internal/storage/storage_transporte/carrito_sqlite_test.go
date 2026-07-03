package storage_test

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSQLite_CrearYListarCarritos(t *testing.T) {
	repo := nuevoRepo(t)

	// Crear un carrito de prueba
	carrito := modelos.Carrito{
		NombreCarrito: "Carrito 1 - Rectorado",
		Capacidad:     5,
		Estado:        "disponible",
	}
	repo.CrearCarrito(carrito)

	// Verificar listado
	lista := repo.ListarCarritos()
	assert.Len(t, lista, 1)
	assert.Equal(t, "Carrito 1 - Rectorado", lista[0].NombreCarrito)
	assert.Equal(t, 5, lista[0].Capacidad)
}

func TestSQLite_CrearYBuscarCarritoPorID(t *testing.T) {
	repo := nuevoRepo(t)

	// Crear carrito pequeño
	carrito := repo.CrearCarrito(modelos.Carrito{
		NombreCarrito: "Carrito 2 - FCVT",
		Capacidad:     3,
		Estado:        "en_ruta",
	})

	// Buscar
	encontrado, ok := repo.BuscarCarritoPorID(carrito.ID)

	assert.True(t, ok)
	assert.NotZero(t, encontrado.ID)
	assert.Equal(t, "Carrito 2 - FCVT", encontrado.NombreCarrito)
	assert.Equal(t, 3, encontrado.Capacidad)
}

func TestSQLite_ActualizarCarrito(t *testing.T) {
	repo := nuevoRepo(t)

	// Crear inicial
	carrito := repo.CrearCarrito(modelos.Carrito{
		NombreCarrito: "Carrito 3 - Biblioteca",
		Capacidad:     5,
		Estado:        "mantenimiento",
	})

	// Actualizar
	carrito.Estado = "disponible"
	carrito.Capacidad = 3 // Cambio de capacidad
	actualizado, ok := repo.ActualizarCarrito(carrito.ID, carrito)

	assert.True(t, ok)
	assert.Equal(t, "disponible", actualizado.Estado)
	assert.Equal(t, 3, actualizado.Capacidad)
}

func TestSQLite_BorrarCarrito(t *testing.T) {
	repo := nuevoRepo(t)

	// Crear
	carrito := repo.CrearCarrito(modelos.Carrito{
		NombreCarrito: "Carrito para borrar",
		Capacidad:     5,
	})

	// Borrar
	ok := repo.BorrarCarrito(carrito.ID)
	assert.True(t, ok)

	// Verificar que ya no existe
	_, ok = repo.BuscarCarritoPorID(carrito.ID)
	assert.False(t, ok)
}