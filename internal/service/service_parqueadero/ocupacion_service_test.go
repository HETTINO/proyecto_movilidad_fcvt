package serviceparqueadero_test

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	service "proyecto_movilidad_fcvt/internal/service"
	sp "proyecto_movilidad_fcvt/internal/service/service_parqueadero"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func TestOcupacionService_Crear_OK(t *testing.T) {
	repo := new(ocupacionRepoMock)
	entrada := modelos.Ocupacion{IDEspacio: 1, PlacaVehiculo: "ABC-1234"}
	guardada := entrada
	guardada.IDOcupacion = 1

	repo.On("CrearOcupacion", entrada).Return(guardada)

	svc := sp.NewOcupacionService(repo)
	creada, err := svc.Crear(entrada)

	assert.NoError(t, err)
	assert.Equal(t, 1, creada.IDOcupacion)
}

func TestOcupacionService_Crear_IDEspacioVacio(t *testing.T) {
	repo := new(ocupacionRepoMock)
	svc := sp.NewOcupacionService(repo)

	_, err := svc.Crear(modelos.Ocupacion{IDEspacio: 0})

	assert.ErrorIs(t, err, service.ErrCampoRequerido)
	repo.AssertNotCalled(t, "CrearOcupacion", mock.Anything)
}

func TestOcupacionService_Actualizar_Exitoso(t *testing.T) {
	repo := new(ocupacionRepoMock)
	datos := modelos.Ocupacion{IDEspacio: 1, PlacaVehiculo: "XYZ-9999"}
	actualizado := datos
	actualizado.IDOcupacion = 1

	repo.On("ActualizarOcupacion", 1, datos).Return(actualizado, true)

	svc := sp.NewOcupacionService(repo)
	resultado, ok, err := svc.Actualizar(1, datos)

	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, "XYZ-9999", resultado.PlacaVehiculo)
}

func TestOcupacionService_Actualizar_NoEncontrado(t *testing.T) {
	repo := new(ocupacionRepoMock)
	datos := modelos.Ocupacion{IDEspacio: 1}
	repo.On("ActualizarOcupacion", 999, datos).Return(modelos.Ocupacion{}, false)

	svc := sp.NewOcupacionService(repo)
	_, ok, err := svc.Actualizar(999, datos)

	assert.False(t, ok)
	assert.ErrorIs(t, err, service.ErrNoEncontrado)
}

func TestOcupacionService_Actualizar_IDEspacioVacio(t *testing.T) {
	repo := new(ocupacionRepoMock)
	svc := sp.NewOcupacionService(repo)

	_, ok, err := svc.Actualizar(1, modelos.Ocupacion{IDEspacio: 0})

	assert.False(t, ok)
	assert.ErrorIs(t, err, service.ErrCampoRequerido)
}

func TestOcupacionService_Borrar_Exitoso(t *testing.T) {
	repo := new(ocupacionRepoMock)
	repo.On("BorrarOcupacion", 1).Return(true)

	svc := sp.NewOcupacionService(repo)
	err := svc.Borrar(1)

	assert.NoError(t, err)
}

func TestOcupacionService_Borrar_NoEncontrado(t *testing.T) {
	repo := new(ocupacionRepoMock)
	repo.On("BorrarOcupacion", 999).Return(false)

	svc := sp.NewOcupacionService(repo)
	err := svc.Borrar(999)

	assert.ErrorIs(t, err, service.ErrNoEncontrado)
}

func TestOcupacionService_Listar(t *testing.T) {
	repo := new(ocupacionRepoMock)
	esperado := []modelos.Ocupacion{{IDOcupacion: 1, IDEspacio: 1}}
	repo.On("ListarOcupaciones").Return(esperado)

	svc := sp.NewOcupacionService(repo)
	resultado := svc.Listar()

	assert.Equal(t, esperado, resultado)
}

func TestOcupacionService_Obtener_Exitoso(t *testing.T) {
	repo := new(ocupacionRepoMock)
	esperado := modelos.Ocupacion{IDOcupacion: 1, IDEspacio: 1, PlacaVehiculo: "ABC-1234"}
	repo.On("BuscarOcupacionPorID", 1).Return(esperado, true)

	svc := sp.NewOcupacionService(repo)
	resultado, ok := svc.Obtener(1)

	assert.True(t, ok)
	assert.Equal(t, "ABC-1234", resultado.PlacaVehiculo)
}

func TestOcupacionService_Obtener_NoEncontrado(t *testing.T) {
	repo := new(ocupacionRepoMock)
	repo.On("BuscarOcupacionPorID", 999).Return(modelos.Ocupacion{}, false)

	svc := sp.NewOcupacionService(repo)
	_, ok := svc.Obtener(999)

	assert.False(t, ok)
}
