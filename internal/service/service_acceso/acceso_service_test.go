package service_acceso_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"proyecto_movilidad_fcvt/internal/modelos"
	service "proyecto_movilidad_fcvt/internal/service"
	sa "proyecto_movilidad_fcvt/internal/service/service_acceso"
)

func TestAccesoService_Crear(t *testing.T) {

	repo := new(accesoRepoMock)

	repo.
		On("CrearAcceso", mock.Anything).
		Return(func(a modelos.Acceso) modelos.Acceso {
			a.ID = 1
			return a
		})

	svc := sa.NewAccesoService(repo)

	input := modelos.Acceso{
		PlacaVehiculo: "ABC123",
		Estado:        "activo",
	}

	res := svc.Crear(input) // 👈 IMPORTANTE: SOLO 1 valor

	assert.Equal(t, 1, res.ID)
	assert.Equal(t, "ABC123", res.PlacaVehiculo)

	repo.AssertExpectations(t)
}

func TestAccesoService_Listar(t *testing.T) {
	repo := new(accesoRepoMock)

	esperados := []modelos.Acceso{
		{ID: 1, PlacaVehiculo: "ABC123"},
		{ID: 2, PlacaVehiculo: "XYZ789"},
	}

	repo.On("ListarAccesos").Return(esperados)

	svc := sa.NewAccesoService(repo)

	res := svc.Listar()

	assert.Len(t, res, 2)
	assert.Equal(t, esperados, res)
	repo.AssertExpectations(t)
}

func TestAccesoService_Obtener_Encontrado(t *testing.T) {
	repo := new(accesoRepoMock)

	esperado := modelos.Acceso{ID: 1, PlacaVehiculo: "ABC123"}
	repo.On("BuscarAccesoPorID", 1).Return(esperado, true)

	svc := sa.NewAccesoService(repo)

	res, ok := svc.Obtener(1)

	assert.True(t, ok)
	assert.Equal(t, esperado, res)
	repo.AssertExpectations(t)
}

func TestAccesoService_Obtener_NoEncontrado(t *testing.T) {
	repo := new(accesoRepoMock)

	repo.On("BuscarAccesoPorID", 999).Return(modelos.Acceso{}, false)

	svc := sa.NewAccesoService(repo)

	res, ok := svc.Obtener(999)

	assert.False(t, ok)
	assert.Equal(t, modelos.Acceso{}, res)
	repo.AssertExpectations(t)
}

func TestAccesoService_Actualizar_Exitoso(t *testing.T) {
	repo := new(accesoRepoMock)

	datos := modelos.Acceso{PlacaVehiculo: "ABC123", Estado: "cerrado"}
	actualizado := datos
	actualizado.ID = 1

	repo.On("ActualizarAcceso", 1, datos).Return(actualizado, true)

	svc := sa.NewAccesoService(repo)

	res, ok, err := svc.Actualizar(1, datos)

	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, actualizado, res)
	repo.AssertExpectations(t)
}

func TestAccesoService_Actualizar_NoEncontrado(t *testing.T) {
	repo := new(accesoRepoMock)

	datos := modelos.Acceso{PlacaVehiculo: "ABC123"}
	repo.On("ActualizarAcceso", 999, datos).Return(modelos.Acceso{}, false)

	svc := sa.NewAccesoService(repo)

	res, ok, err := svc.Actualizar(999, datos)

	assert.False(t, ok)
	assert.ErrorIs(t, err, service.ErrNoEncontrado)
	assert.Equal(t, modelos.Acceso{}, res)
	repo.AssertExpectations(t)
}

func TestAccesoService_Borrar_Exitoso(t *testing.T) {
	repo := new(accesoRepoMock)

	repo.On("BorrarAcceso", 1).Return(true)

	svc := sa.NewAccesoService(repo)

	err := svc.Borrar(1)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestAccesoService_Borrar_NoEncontrado(t *testing.T) {
	repo := new(accesoRepoMock)

	repo.On("BorrarAcceso", 999).Return(false)

	svc := sa.NewAccesoService(repo)

	err := svc.Borrar(999)

	assert.ErrorIs(t, err, service.ErrNoEncontrado)
	repo.AssertExpectations(t)
}
