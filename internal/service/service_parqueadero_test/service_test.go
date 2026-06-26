package service_parqueadero_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"proyecto_movilidad_fcvt/internal/modelos"
	sp "proyecto_movilidad_fcvt/internal/service/service_parqueadero"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_parqueadero"
)

// =========================================================
// MOCK — ParqueaderoRepository
// =========================================================

type parqueaderoRepoMock struct{ mock.Mock }

func (m *parqueaderoRepoMock) ListarParqueaderos() []modelos.Parqueadero {
	return m.Called().Get(0).([]modelos.Parqueadero)
}
func (m *parqueaderoRepoMock) BuscarParqueaderoPorID(id int) (modelos.Parqueadero, bool) {
	args := m.Called(id)
	return args.Get(0).(modelos.Parqueadero), args.Bool(1)
}
func (m *parqueaderoRepoMock) CrearParqueadero(p modelos.Parqueadero) modelos.Parqueadero {
	return m.Called(p).Get(0).(modelos.Parqueadero)
}
func (m *parqueaderoRepoMock) ActualizarParqueadero(id int, datos modelos.Parqueadero) (modelos.Parqueadero, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(modelos.Parqueadero), args.Bool(1)
}
func (m *parqueaderoRepoMock) BorrarParqueadero(id int) bool {
	return m.Called(id).Bool(0)
}

var _ storage.ParqueaderoRepository = (*parqueaderoRepoMock)(nil)

// =========================================================
// MOCK — OcupacionesRepository
// =========================================================

type ocupacionRepoMock struct{ mock.Mock }

func (m *ocupacionRepoMock) ListarOcupaciones() []modelos.Ocupacion {
	return m.Called().Get(0).([]modelos.Ocupacion)
}
func (m *ocupacionRepoMock) BuscarOcupacionPorID(id int) (modelos.Ocupacion, bool) {
	args := m.Called(id)
	return args.Get(0).(modelos.Ocupacion), args.Bool(1)
}
func (m *ocupacionRepoMock) CrearOcupacion(o modelos.Ocupacion) modelos.Ocupacion {
	return m.Called(o).Get(0).(modelos.Ocupacion)
}
func (m *ocupacionRepoMock) ActualizarOcupacion(id int, datos modelos.Ocupacion) (modelos.Ocupacion, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(modelos.Ocupacion), args.Bool(1)
}
func (m *ocupacionRepoMock) BorrarOcupacion(id int) bool {
	return m.Called(id).Bool(0)
}
func (m *ocupacionRepoMock) LiberarOcupacion(id int) (modelos.Ocupacion, bool) {
	args := m.Called(id)
	return args.Get(0).(modelos.Ocupacion), args.Bool(1)
}

var _ storage.OcupacionesRepository = (*ocupacionRepoMock)(nil)

// =========================================================
// TESTS — ParqueaderoService
// =========================================================

func TestParqueaderoService_Crear(t *testing.T) {
	casos := []struct {
		nombre        string
		entrada       modelos.Parqueadero
		debeFallar    bool
		debePersistir bool
	}{
		{
			nombre:        "nombre vacío -> no persiste",
			entrada:       modelos.Parqueadero{Nombre: "   "},
			debeFallar:    true,
			debePersistir: false,
		},
		{
			nombre:        "parqueadero válido -> se persiste",
			entrada:       modelos.Parqueadero{Nombre: "Norte"},
			debeFallar:    false,
			debePersistir: true,
		},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {

			repo := new(parqueaderoRepoMock)

			if c.debePersistir {
				guardado := c.entrada
				guardado.IDParqueadero = 1

				repo.
					On("CrearParqueadero", c.entrada).
					Return(guardado)
			}

			svc := sp.NewParqueaderoService(repo)

			creado, err := svc.Crear(c.entrada)

			if c.debeFallar {
				assert.Error(t, err)
				repo.AssertNotCalled(t, "CrearParqueadero")
			} else {
				assert.NoError(t, err)
				assert.Equal(t, 1, creado.IDParqueadero)
				repo.AssertCalled(t, "CrearParqueadero", c.entrada)
			}
		})
	}
}

func TestParqueaderoService_Obtener_NoEncontrado(t *testing.T) {
	repo := new(parqueaderoRepoMock)
	repo.On("BuscarParqueaderoPorID", 999).Return(modelos.Parqueadero{}, false)
	svc := sp.NewParqueaderoService(repo)

	_, ok := svc.Obtener(999)

	assert.False(t, ok)
	repo.AssertExpectations(t)
}

func TestParqueaderoService_Obtener_Exitoso(t *testing.T) {
	repo := new(parqueaderoRepoMock)
	esperado := modelos.Parqueadero{IDParqueadero: 1, Nombre: "Norte", Capacidad: 50, Tipo: "cubierto"}
	repo.On("BuscarParqueaderoPorID", 1).Return(esperado, true)
	svc := sp.NewParqueaderoService(repo)

	resultado, ok := svc.Obtener(1)

	assert.True(t, ok)
	assert.Equal(t, "Norte", resultado.Nombre)
	repo.AssertExpectations(t)
}

// =========================================================
// TESTS — OcupacionService
// =========================================================

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
