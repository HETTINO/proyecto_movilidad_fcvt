package storage_test
import (
	"proyecto_movilidad_fcvt/internal/modelos"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSQLite_CrearYListarSolicitudes(t *testing.T) {
	repo := nuevoRepo(t)

	// Crear una solicitud de prueba
	sol := modelos.Solicitud{
		CedulaUsuario: "1234567890",
		CantPersonas:  2,
		ParadaOrigen:  1,
		PuntoDestino:  "Terminal Norte",
		Estado:        "pendiente",
	}
	repo.CrearSolicitud(sol)

	// Verificar listado
	lista := repo.ListarSolicitudes()
	assert.Len(t, lista, 1)
	assert.Equal(t, "1234567890", lista[0].CedulaUsuario)
	assert.Equal(t, "Terminal Norte", lista[0].PuntoDestino)
}

func TestSQLite_CrearYBuscarSolicitudPorID(t *testing.T) {
	repo := nuevoRepo(t)

	// Crear
	sol := repo.CrearSolicitud(modelos.Solicitud{
		CedulaUsuario: "0987654321",
		CantPersonas:  1,
		ParadaOrigen:  2,
		PuntoDestino:  "Centro",
	})

	// Buscar
	encontrado, ok := repo.BuscarSolicitudPorID(sol.ID)

	assert.True(t, ok)
	assert.NotZero(t, encontrado.ID)
	assert.Equal(t, "0987654321", encontrado.CedulaUsuario)
	assert.Equal(t, "Centro", encontrado.PuntoDestino)
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
		PuntoDestino:  "Casa",
	})

	// Actualizar
	sol.Estado = "completado"
	actualizado, ok := repo.ActualizarSolicitud(sol.ID, sol)

	assert.True(t, ok)
	assert.Equal(t, "completado", actualizado.Estado)
}
