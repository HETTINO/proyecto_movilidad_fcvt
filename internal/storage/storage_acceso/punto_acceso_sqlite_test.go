package storage_acceso_test

import (
	"testing"

	"proyecto_movilidad_fcvt/internal/modelos"

	"github.com/stretchr/testify/assert"
)

func TestSQLite_CrearYListarPuntoAcceso(t *testing.T) {
	repo := nuevoRepo(t)

	repo.CrearPuntoAcceso(modelos.PuntoDeAcceso{
		Frecuencia: "Vehicular",
		Ubicacion:  "Bloque Sur",
	})

	lista := repo.ListarPuntosAcceso()

	assert.Len(t, lista, 1)
	assert.Equal(t, "Bloque Sur", lista[0].Ubicacion)
}

func TestSQLite_CrearYBuscarPuntoAccesoPorID(t *testing.T) {
	repo := nuevoRepo(t)

	creado := repo.CrearPuntoAcceso(modelos.PuntoDeAcceso{
		Frecuencia: "Peatonal",
		Ubicacion:  "Bloque Norte",
	})

	assert.NotZero(t, creado.ID)

	encontrado, ok := repo.BuscarPuntoAccesoPorID(creado.ID)

	assert.True(t, ok)
	assert.Equal(t, "Peatonal", encontrado.Frecuencia)
}
