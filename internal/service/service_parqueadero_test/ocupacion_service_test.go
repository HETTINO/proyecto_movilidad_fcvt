package service_parqueadero_test

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	sp "proyecto_movilidad_fcvt/internal/service/service_parqueadero"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOcupacionService_Liberar_NoEncontrado(t *testing.T) {
	repo := new(ocupacionRepoMock)
	repo.On("LiberarOcupacion", 999).Return(modelos.Ocupacion{}, false)
	svc := sp.NewOcupacionService(repo)

	_, ok := svc.Liberar(999)

	assert.False(t, ok)
	repo.AssertExpectations(t)
}

func TestOcupacionService_Liberar_Exitoso(t *testing.T) {
	repo := new(ocupacionRepoMock)
	liberada := modelos.Ocupacion{IDOcupacion: 1, PlacaVehiculo: "ABC-1234"}
	repo.On("LiberarOcupacion", 1).Return(liberada, true)
	svc := sp.NewOcupacionService(repo)

	resultado, ok := svc.Liberar(1)

	assert.True(t, ok)
	assert.Equal(t, "ABC-1234", resultado.PlacaVehiculo)
	repo.AssertExpectations(t)
}
