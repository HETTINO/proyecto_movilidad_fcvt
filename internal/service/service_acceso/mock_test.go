package service_acceso_test

import (
	"github.com/stretchr/testify/mock"

	"proyecto_movilidad_fcvt/internal/modelos"
	storage_acceso "proyecto_movilidad_fcvt/internal/storage/storage_acceso"
)

// =========================================================
// MOCK — USUARIO
// =========================================================

type usuarioRepoMock struct {
	mock.Mock
}

func (m *usuarioRepoMock) ListarUsuarios() []modelos.Usuario {
	return m.Called().Get(0).([]modelos.Usuario)
}

func (m *usuarioRepoMock) BuscarUsuarioPorCedula(cedula string) (modelos.Usuario, bool) {
	args := m.Called(cedula)
	return args.Get(0).(modelos.Usuario), args.Bool(1)
}

func (m *usuarioRepoMock) CrearUsuario(u modelos.Usuario) modelos.Usuario {
	args := m.Called(u)
	return args.Get(0).(modelos.Usuario)
}

func (m *usuarioRepoMock) ActualizarUsuario(cedula string, datos modelos.Usuario) (modelos.Usuario, bool) {
	args := m.Called(cedula, datos)
	return args.Get(0).(modelos.Usuario), args.Bool(1)
}

func (m *usuarioRepoMock) BorrarUsuario(cedula string) bool {
	return m.Called(cedula).Bool(0)
}

// validación de interfaz (IMPORTANTE)
var _ storage_acceso.UsuarioRepository = (*usuarioRepoMock)(nil)

// =====================================================
// MOCK VEHICULO
// =====================================================

type vehiculoRepoMock struct {
	mock.Mock
}

func (m *vehiculoRepoMock) ListarVehiculos() []modelos.Vehiculo {
	return m.Called().Get(0).([]modelos.Vehiculo)
}

func (m *vehiculoRepoMock) BuscarVehiculoPorPlaca(placa string) (modelos.Vehiculo, bool) {
	args := m.Called(placa)
	return args.Get(0).(modelos.Vehiculo), args.Bool(1)
}

func (m *vehiculoRepoMock) CrearVehiculo(v modelos.Vehiculo) modelos.Vehiculo {
	return m.Called(v).Get(0).(modelos.Vehiculo)
}

func (m *vehiculoRepoMock) ActualizarVehiculo(placa string, v modelos.Vehiculo) (modelos.Vehiculo, bool) {
	args := m.Called(placa, v)
	return args.Get(0).(modelos.Vehiculo), args.Bool(1)
}

func (m *vehiculoRepoMock) BorrarVehiculo(placa string) bool {
	return m.Called(placa).Bool(0)
}

// validación de interfaz
var _ storage_acceso.VehiculoRepository = (*vehiculoRepoMock)(nil)

// =========================================================
// MOCK — PUNTO DE ACCESO
// =========================================================

type puntoAccesoRepoMock struct {
	mock.Mock
}

func (m *puntoAccesoRepoMock) ListarPuntosAcceso() []modelos.PuntoDeAcceso {
	return m.Called().Get(0).([]modelos.PuntoDeAcceso)
}

func (m *puntoAccesoRepoMock) BuscarPuntoAccesoPorID(id int) (modelos.PuntoDeAcceso, bool) {
	args := m.Called(id)
	return args.Get(0).(modelos.PuntoDeAcceso), args.Bool(1)
}

func (m *puntoAccesoRepoMock) CrearPuntoAcceso(p modelos.PuntoDeAcceso) modelos.PuntoDeAcceso {
	return m.Called(p).Get(0).(modelos.PuntoDeAcceso)
}

// ❗ ESTE ERA EL QUE TE FALTABA
func (m *puntoAccesoRepoMock) ActualizarPuntoAcceso(id int, datos modelos.PuntoDeAcceso) (modelos.PuntoDeAcceso, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(modelos.PuntoDeAcceso), args.Bool(1)
}

func (m *puntoAccesoRepoMock) BorrarPuntoAcceso(id int) bool {
	return m.Called(id).Bool(0)
}

// validación de interfaz
var _ storage_acceso.PuntoAccesoRepository = (*puntoAccesoRepoMock)(nil)

// =========================================================
// MOCK — ACCESO
// =========================================================

type accesoRepoMock struct {
	mock.Mock
}

func (m *accesoRepoMock) ListarAccesos() []modelos.Acceso {
	return m.Called().Get(0).([]modelos.Acceso)
}

func (m *accesoRepoMock) BuscarAccesoPorID(id int) (modelos.Acceso, bool) {
	args := m.Called(id)
	return args.Get(0).(modelos.Acceso), args.Bool(1)
}

func (m *accesoRepoMock) CrearAcceso(a modelos.Acceso) modelos.Acceso {
	args := m.Called(a)
	if fn, ok := args.Get(0).(func(modelos.Acceso) modelos.Acceso); ok {
		return fn(a)
	}
	return args.Get(0).(modelos.Acceso)
}

func (m *accesoRepoMock) ActualizarAcceso(id int, datos modelos.Acceso) (modelos.Acceso, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(modelos.Acceso), args.Bool(1)
}

func (m *accesoRepoMock) BorrarAcceso(id int) bool {
	return m.Called(id).Bool(0)
}

var _ storage_acceso.AccesoRepository = (*accesoRepoMock)(nil)
