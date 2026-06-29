package service_acceso_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"proyecto_movilidad_fcvt/internal/modelos"
	sa "proyecto_movilidad_fcvt/internal/service/service_acceso"
)

func TestVehiculo_Crear(t *testing.T) {

	repo := new(vehiculoRepoMock)

	input := modelos.Vehiculo{
		Placa:        "ABC123",
		IDUsuario:    "12345678",
		TipoVehiculo: "Carro",
		Marca:        "Toyota",
	}

	repo.On("CrearVehiculo", input).Return(input)

	svc := sa.NewVehiculoService(repo)

	res, err := svc.Crear(input)

	assert.NoError(t, err)
	assert.Equal(t, input.Placa, res.Placa)

	repo.AssertExpectations(t)
}
