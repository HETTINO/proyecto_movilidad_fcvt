package servicetransporte_test

import (
	"github.com/stretchr/testify/mock"

	modelos "proyecto_movilidad_fcvt/internal/modelos"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_transporte"
)

// =========================================================
// MOCK — Almacen
// =========================================================
type almacenMock struct{ mock.Mock }

// Rutas

func (m *almacenMock) ListarRutas() []modelos.Ruta {
	return m.Called().Get(0).([]modelos.Ruta)
}

func (m *almacenMock) BuscarRutaPorID(id int) (modelos.Ruta, bool) {
	args := m.Called(id)
	return args.Get(0).(modelos.Ruta), args.Bool(1)
}

func (m *almacenMock) CrearRuta(r modelos.Ruta) modelos.Ruta {
	return m.Called(r).Get(0).(modelos.Ruta)
}

func (m *almacenMock) ActualizarRuta(id int, datos modelos.Ruta) (modelos.Ruta, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(modelos.Ruta), args.Bool(1)
}

func (m *almacenMock) BorrarRuta(id int) bool {
	return m.Called(id).Bool(0)
}

// Paradas

func (m *almacenMock) ListarParadas() []modelos.Parada {
	return m.Called().Get(0).([]modelos.Parada)
}

func (m *almacenMock) BuscarParadaPorID(id int) (modelos.Parada, bool) {
	args := m.Called(id)
	return args.Get(0).(modelos.Parada), args.Bool(1)
}

func (m *almacenMock) CrearParada(p modelos.Parada) modelos.Parada {
	return m.Called(p).Get(0).(modelos.Parada)
}

func (m *almacenMock) ActualizarParada(id int, datos modelos.Parada) (modelos.Parada, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(modelos.Parada), args.Bool(1)
}

func (m *almacenMock) BorrarParada(id int) bool {
	return m.Called(id).Bool(0)
}

// Carritos

func (m *almacenMock) ListarCarritos() []modelos.Carrito {
	return m.Called().Get(0).([]modelos.Carrito)
}

func (m *almacenMock) BuscarCarritoPorID(id int) (modelos.Carrito, bool) {
	args := m.Called(id)
	return args.Get(0).(modelos.Carrito), args.Bool(1)
}

func (m *almacenMock) CrearCarrito(c modelos.Carrito) modelos.Carrito {
	return m.Called(c).Get(0).(modelos.Carrito)
}

func (m *almacenMock) ActualizarCarrito(id int, datos modelos.Carrito) (modelos.Carrito, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(modelos.Carrito), args.Bool(1)
}

func (m *almacenMock) BorrarCarrito(id int) bool {
	return m.Called(id).Bool(0)
}

// Locaciones

func (m *almacenMock) ListarLocaciones() []modelos.Locacion {
	return m.Called().Get(0).([]modelos.Locacion)
}

func (m *almacenMock) RegistrarLocacion(l modelos.Locacion) modelos.Locacion {
	return m.Called(l).Get(0).(modelos.Locacion)
}

func (m *almacenMock) ObtenerUltimaLocacionPorCarrito(carritoID int) (modelos.Locacion, bool) {
	args := m.Called(carritoID)
	return args.Get(0).(modelos.Locacion), args.Bool(1)
}

// Solicitudes

func (m *almacenMock) ListarSolicitudes() []modelos.Solicitud {
	return m.Called().Get(0).([]modelos.Solicitud)
}

func (m *almacenMock) BuscarSolicitudPorID(id int) (modelos.Solicitud, bool) {
	args := m.Called(id)
	return args.Get(0).(modelos.Solicitud), args.Bool(1)
}

func (m *almacenMock) CrearSolicitud(s modelos.Solicitud) modelos.Solicitud {
	return m.Called(s).Get(0).(modelos.Solicitud)
}

func (m *almacenMock) ActualizarSolicitud(id int, datos modelos.Solicitud) (modelos.Solicitud, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(modelos.Solicitud), args.Bool(1)
}

func (m *almacenMock) BorrarSolicitud(id int) bool {
	return m.Called(id).Bool(0)
}

var _ storage.Almacen = (*almacenMock)(nil)
