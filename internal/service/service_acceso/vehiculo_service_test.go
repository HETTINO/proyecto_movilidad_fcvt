package service_acceso_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"proyecto_movilidad_fcvt/internal/modelos"
	service "proyecto_movilidad_fcvt/internal/service"
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

func TestVehiculoService_Listar(t *testing.T) {
	repo := new(vehiculoRepoMock)

	esperados := []modelos.Vehiculo{
		{Placa: "ABC123", Marca: "Toyota"},
		{Placa: "XYZ789", Marca: "Chevrolet"},
	}

	repo.On("ListarVehiculos").Return(esperados)

	svc := sa.NewVehiculoService(repo)

	res := svc.Listar()

	assert.Len(t, res, 2)
	assert.Equal(t, esperados, res)
	repo.AssertExpectations(t)
}

func TestVehiculoService_Obtener_Encontrado(t *testing.T) {
	repo := new(vehiculoRepoMock)

	esperado := modelos.Vehiculo{Placa: "ABC123", Marca: "Toyota"}
	repo.On("BuscarVehiculoPorPlaca", "ABC123").Return(esperado, true)

	svc := sa.NewVehiculoService(repo)

	res, ok := svc.Obtener("ABC123")

	assert.True(t, ok)
	assert.Equal(t, esperado, res)
	repo.AssertExpectations(t)
}

func TestVehiculoService_Obtener_NoEncontrado(t *testing.T) {
	repo := new(vehiculoRepoMock)

	repo.On("BuscarVehiculoPorPlaca", "ZZZ999").Return(modelos.Vehiculo{}, false)

	svc := sa.NewVehiculoService(repo)

	res, ok := svc.Obtener("ZZZ999")

	assert.False(t, ok)
	assert.Equal(t, modelos.Vehiculo{}, res)
	repo.AssertExpectations(t)
}

func TestVehiculoService_Crear_CampoRequeridoFaltante(t *testing.T) {
	repo := new(vehiculoRepoMock)

	svc := sa.NewVehiculoService(repo)

	entrada := modelos.Vehiculo{Marca: "Toyota"} // sin Placa

	res, err := svc.Crear(entrada)

	assert.ErrorIs(t, err, service.ErrCampoRequerido)
	assert.Equal(t, modelos.Vehiculo{}, res)
	repo.AssertNotCalled(t, "CrearVehiculo", entrada)
}

func TestVehiculoService_Actualizar_Exitoso(t *testing.T) {
	repo := new(vehiculoRepoMock)

	datos := modelos.Vehiculo{Placa: "ABC123", Marca: "Toyota", Color: "Rojo"}
	repo.On("ActualizarVehiculo", "ABC123", datos).Return(datos, true)

	svc := sa.NewVehiculoService(repo)

	res, ok, err := svc.Actualizar("ABC123", datos)

	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, datos, res)
	repo.AssertExpectations(t)
}

func TestVehiculoService_Actualizar_NoEncontrado(t *testing.T) {
	repo := new(vehiculoRepoMock)

	datos := modelos.Vehiculo{Placa: "ZZZ999", Marca: "Nada"}
	repo.On("ActualizarVehiculo", "ZZZ999", datos).Return(modelos.Vehiculo{}, false)

	svc := sa.NewVehiculoService(repo)

	res, ok, err := svc.Actualizar("ZZZ999", datos)

	assert.False(t, ok)
	assert.ErrorIs(t, err, service.ErrNoEncontrado)
	assert.Equal(t, modelos.Vehiculo{}, res)
	repo.AssertExpectations(t)
}

func TestVehiculoService_Actualizar_CampoRequeridoFaltante(t *testing.T) {
	repo := new(vehiculoRepoMock)

	svc := sa.NewVehiculoService(repo)

	datos := modelos.Vehiculo{Marca: "Toyota"} // sin Placa

	res, ok, err := svc.Actualizar("ABC123", datos)

	assert.False(t, ok)
	assert.ErrorIs(t, err, service.ErrCampoRequerido)
	assert.Equal(t, modelos.Vehiculo{}, res)
	repo.AssertNotCalled(t, "ActualizarVehiculo", "ABC123", datos)
}

func TestVehiculoService_Borrar_Exitoso(t *testing.T) {
	repo := new(vehiculoRepoMock)

	repo.On("BorrarVehiculo", "ABC123").Return(true)

	svc := sa.NewVehiculoService(repo)

	err := svc.Borrar("ABC123")

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestVehiculoService_Borrar_NoEncontrado(t *testing.T) {
	repo := new(vehiculoRepoMock)

	repo.On("BorrarVehiculo", "ZZZ999").Return(false)

	svc := sa.NewVehiculoService(repo)

	err := svc.Borrar("ZZZ999")

	assert.ErrorIs(t, err, service.ErrNoEncontrado)
	repo.AssertExpectations(t)
}
