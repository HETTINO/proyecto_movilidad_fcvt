package service_acceso_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"proyecto_movilidad_fcvt/internal/modelos"
	sa "proyecto_movilidad_fcvt/internal/service/service_acceso"
)

func TestPuntoAcceso_Crear(t *testing.T) {

	repo := new(puntoAccesoRepoMock)

	input := modelos.PuntoDeAcceso{
		Frecuencia: "Alta",
		Ubicacion:  "Bloque A",
	}

	expected := input
	expected.ID = 1

	repo.On("CrearPuntoAcceso", input).Return(expected)

	svc := sa.NewPuntoAccesoService(repo)

	res, err := svc.Crear(input)

	assert.NoError(t, err)
	assert.Equal(t, 1, res.ID)

	repo.AssertExpectations(t)
}
