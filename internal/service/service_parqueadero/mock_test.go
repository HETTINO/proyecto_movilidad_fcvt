package serviceparqueadero_test
import (
	"github.com/stretchr/testify/mock"

	"proyecto_movilidad_fcvt/internal/modelos"
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

type espacioRepoMock struct {
	mock.Mock
}

func (m *espacioRepoMock) ListarEspacios() []modelos.Espacio {
	return m.Called().Get(0).([]modelos.Espacio)
}

func (m *espacioRepoMock) BuscarEspacioPorID(id int) (modelos.Espacio, bool) {
	args := m.Called(id)
	return args.Get(0).(modelos.Espacio), args.Bool(1)
}

func (m *espacioRepoMock) CrearEspacio(e modelos.Espacio) modelos.Espacio {
	return m.Called(e).Get(0).(modelos.Espacio)
}

func (m *espacioRepoMock) ActualizarEspacio(id int, datos modelos.Espacio) (modelos.Espacio, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(modelos.Espacio), args.Bool(1)
}

func (m *espacioRepoMock) BorrarEspacio(id int) bool {
	return m.Called(id).Bool(0)
}

var _ storage.ParqueaderoRepository = (*parqueaderoRepoMock)(nil)
var _ storage.EspacioRepository = (*espacioRepoMock)(nil)
var _ storage.OcupacionesRepository = (*ocupacionRepoMock)(nil)
