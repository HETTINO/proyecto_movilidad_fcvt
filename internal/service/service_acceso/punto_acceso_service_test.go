package service_acceso_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"proyecto_movilidad_fcvt/internal/modelos"
	service "proyecto_movilidad_fcvt/internal/service"
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

func TestPuntoAccesoService_Listar(t *testing.T) {
	repo := new(puntoAccesoRepoMock)

	esperados := []modelos.PuntoDeAcceso{
		{ID: 1, Ubicacion: "Bloque A"},
		{ID: 2, Ubicacion: "Bloque B"},
	}

	repo.On("ListarPuntosAcceso").Return(esperados)

	svc := sa.NewPuntoAccesoService(repo)

	res := svc.Listar()

	assert.Len(t, res, 2)
	assert.Equal(t, esperados, res)
	repo.AssertExpectations(t)
}

func TestPuntoAccesoService_Obtener_Encontrado(t *testing.T) {
	repo := new(puntoAccesoRepoMock)

	esperado := modelos.PuntoDeAcceso{ID: 1, Ubicacion: "Bloque A"}
	repo.On("BuscarPuntoAccesoPorID", 1).Return(esperado, true)

	svc := sa.NewPuntoAccesoService(repo)

	res, ok := svc.Obtener(1)

	assert.True(t, ok)
	assert.Equal(t, esperado, res)
	repo.AssertExpectations(t)
}

func TestPuntoAccesoService_Obtener_NoEncontrado(t *testing.T) {
	repo := new(puntoAccesoRepoMock)

	repo.On("BuscarPuntoAccesoPorID", 999).Return(modelos.PuntoDeAcceso{}, false)

	svc := sa.NewPuntoAccesoService(repo)

	res, ok := svc.Obtener(999)

	assert.False(t, ok)
	assert.Equal(t, modelos.PuntoDeAcceso{}, res)
	repo.AssertExpectations(t)
}

func TestPuntoAccesoService_Crear_CampoRequeridoFaltante(t *testing.T) {
	repo := new(puntoAccesoRepoMock)

	svc := sa.NewPuntoAccesoService(repo)

	entrada := modelos.PuntoDeAcceso{Frecuencia: "Alta"} // sin Ubicacion

	res, err := svc.Crear(entrada)

	assert.ErrorIs(t, err, service.ErrCampoRequerido)
	assert.Equal(t, modelos.PuntoDeAcceso{}, res)
	repo.AssertNotCalled(t, "CrearPuntoAcceso", entrada)
}

func TestPuntoAccesoService_Actualizar_Exitoso(t *testing.T) {
	repo := new(puntoAccesoRepoMock)

	datos := modelos.PuntoDeAcceso{Frecuencia: "Alta", Ubicacion: "Bloque A"}
	actualizado := datos
	actualizado.ID = 1

	repo.On("ActualizarPuntoAcceso", 1, datos).Return(actualizado, true)

	svc := sa.NewPuntoAccesoService(repo)

	res, ok, err := svc.Actualizar(1, datos)

	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, actualizado, res)
	repo.AssertExpectations(t)
}

func TestPuntoAccesoService_Actualizar_NoEncontrado(t *testing.T) {
	repo := new(puntoAccesoRepoMock)

	datos := modelos.PuntoDeAcceso{Ubicacion: "Bloque Z"}
	repo.On("ActualizarPuntoAcceso", 999, datos).Return(modelos.PuntoDeAcceso{}, false)

	svc := sa.NewPuntoAccesoService(repo)

	res, ok, err := svc.Actualizar(999, datos)

	assert.False(t, ok)
	assert.ErrorIs(t, err, service.ErrNoEncontrado)
	assert.Equal(t, modelos.PuntoDeAcceso{}, res)
	repo.AssertExpectations(t)
}

func TestPuntoAccesoService_Actualizar_CampoRequeridoFaltante(t *testing.T) {
	repo := new(puntoAccesoRepoMock)

	svc := sa.NewPuntoAccesoService(repo)

	datos := modelos.PuntoDeAcceso{Frecuencia: "Alta"} // sin Ubicacion

	res, ok, err := svc.Actualizar(1, datos)

	assert.False(t, ok)
	assert.ErrorIs(t, err, service.ErrCampoRequerido)
	assert.Equal(t, modelos.PuntoDeAcceso{}, res)
	repo.AssertNotCalled(t, "ActualizarPuntoAcceso", 1, datos)
}

func TestPuntoAccesoService_Borrar_Exitoso(t *testing.T) {
	repo := new(puntoAccesoRepoMock)

	repo.On("BorrarPuntoAcceso", 1).Return(true)

	svc := sa.NewPuntoAccesoService(repo)

	err := svc.Borrar(1)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestPuntoAccesoService_Borrar_NoEncontrado(t *testing.T) {
	repo := new(puntoAccesoRepoMock)

	repo.On("BorrarPuntoAcceso", 999).Return(false)

	svc := sa.NewPuntoAccesoService(repo)

	err := svc.Borrar(999)

	assert.ErrorIs(t, err, service.ErrNoEncontrado)
	repo.AssertExpectations(t)
}
