package sqlite_test_acceso

import (
	"testing"
	"time"

	"proyecto_movilidad_fcvt/internal/modelos"

	"github.com/stretchr/testify/assert"
)

func TestSQLite_CrearYListarAcceso(t *testing.T) {
	repo := nuevoRepo(t)

	acceso := modelos.Acceso{
		PlacaVehiculo: "ABC-1234",
		PuntoAccesoID: 1,
		TiempoEntrada: time.Now(),
		Estado:        "activo",
		Observaciones: "Ingreso inicial",
	}

	creado := repo.CrearAcceso(acceso)

	assert.Equal(t, "ABC-1234", creado.PlacaVehiculo)
	assert.Equal(t, "activo", creado.Estado)

	lista := repo.ListarAccesos()

	assert.Len(t, lista, 1)
	assert.Equal(t, "ABC-1234", lista[0].PlacaVehiculo)
}

func TestSQLite_BuscarAccesoPorID(t *testing.T) {
	repo := nuevoRepo(t)

	acceso := modelos.Acceso{
		PlacaVehiculo: "XYZ-999",
		PuntoAccesoID: 1,
		TiempoEntrada: time.Now(),
		Estado:        "activo",
		Observaciones: "Prueba búsqueda",
	}

	creado := repo.CrearAcceso(acceso)

	assert.NotZero(t, creado.ID)

	encontrado, ok := repo.BuscarAccesoPorID(creado.ID)

	assert.True(t, ok)
	assert.Equal(t, "XYZ-999", encontrado.PlacaVehiculo)
}
