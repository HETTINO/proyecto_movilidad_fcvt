package storage_test

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSQLite_CrearSolicitud(t *testing.T) {
	repo := nuevoRepo(t)
	sol := modelos.Solicitud{
		CedulaUsuario: "1234567890",
		CantPersonas:  2,
		ParadaOrigen:  1,
		PuntoDestino:  "Tasty Food - Comedor",
		Estado:        "pendiente",
	}

	creada := repo.CrearSolicitud(sol)

	assert.NotZero(t, creada.ID)
	assert.Equal(t, "1234567890", creada.CedulaUsuario)
	assert.Equal(t, "Tasty Food - Comedor", creada.PuntoDestino)
}

func TestSQLite_ListarSolicitudes(t *testing.T) {
	repo := nuevoRepo(t)
	repo.CrearSolicitud(modelos.Solicitud{CedulaUsuario: "User1"})
	repo.CrearSolicitud(modelos.Solicitud{CedulaUsuario: "User2"})

	lista := repo.ListarSolicitudes()

	assert.Len(t, lista, 2)
}

func TestSQLite_BuscarSolicitudPorID(t *testing.T) {
	repo := nuevoRepo(t)
	sol := repo.CrearSolicitud(modelos.Solicitud{
		CedulaUsuario: "0987654321",
		CantPersonas:  1,
		ParadaOrigen:  2,
		PuntoDestino:  "Facultad de Ingeniería - Bloque A",
	})

	encontrado, ok := repo.BuscarSolicitudPorID(sol.ID)

	assert.True(t, ok)
	assert.Equal(t, sol.ID, encontrado.ID)
	assert.Equal(t, "0987654321", encontrado.CedulaUsuario)
}

func TestSQLite_BuscarSolicitudInexistente(t *testing.T) {
	repo := nuevoRepo(t)
	_, ok := repo.BuscarSolicitudPorID(999)
	assert.False(t, ok)
}

func TestSQLite_ActualizarSolicitud(t *testing.T) {
	repo := nuevoRepo(t)

	// Crear inicial
	sol := repo.CrearSolicitud(modelos.Solicitud{
		CedulaUsuario: "111",
		CantPersonas:  1,
		PuntoDestino:  "Facultas de Salud - Bloque B",
	})

	// Actualizar
	sol.Estado = "completado"
	actualizado, ok := repo.ActualizarSolicitud(sol.ID, sol)

	assert.True(t, ok)
	assert.Equal(t, "completado", actualizado.Estado)
}

func TestSQLite_BorrarSolicitud(t *testing.T) {
	repo := nuevoRepo(t)

	sol := repo.CrearSolicitud(modelos.Solicitud{
		CedulaUsuario: "222",
		CantPersonas:  1,
		PuntoDestino:  "Casa",
	})

	eliminado := repo.BorrarSolicitud(sol.ID)
	assert.True(t, eliminado)

	_, ok := repo.BuscarSolicitudPorID(sol.ID)
	assert.False(t, ok)
}

func TestSQLite_BorrarSolicitudInexistente(t *testing.T) {
	repo := nuevoRepo(t)

	eliminado := repo.BorrarSolicitud(999)
	assert.False(t, eliminado)
}
