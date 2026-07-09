package service_acceso_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"proyecto_movilidad_fcvt/internal/modelos"
	service "proyecto_movilidad_fcvt/internal/service"
	sa "proyecto_movilidad_fcvt/internal/service/service_acceso"
)

// =========================================================
// TESTS — UsuarioService
// =========================================================

func TestUsuarioService_Crear(t *testing.T) {

	repo := new(usuarioRepoMock)

	input := modelos.Usuario{
		Cedula:     "12345678",
		Nombre:     "Shirley",
		Contrasena: "1234",
		Email:      "shirley@test.com",
		Rol:        "admin",
	}

	repo.On("CrearUsuario", input).Return(input, nil)

	svc := sa.NewUsuarioService(repo)

	res, err := svc.Crear(input)

	assert.NoError(t, err)
	assert.Equal(t, input.Email, res.Email)

	repo.AssertExpectations(t)
}

func TestUsuarioService_Listar(t *testing.T) {
	repo := new(usuarioRepoMock)

	esperados := []modelos.Usuario{
		{Cedula: "111", Nombre: "Ana", Email: "ana@test.com"},
		{Cedula: "222", Nombre: "Luis", Email: "luis@test.com"},
	}

	repo.On("ListarUsuarios").Return(esperados)

	svc := sa.NewUsuarioService(repo)

	res := svc.Listar()

	assert.Len(t, res, 2)
	assert.Equal(t, esperados, res)
	repo.AssertExpectations(t)
}

func TestUsuarioService_Obtener_Encontrado(t *testing.T) {
	repo := new(usuarioRepoMock)

	esperado := modelos.Usuario{Cedula: "111", Nombre: "Ana", Email: "ana@test.com"}
	repo.On("BuscarUsuarioPorCedula", "111").Return(esperado, true)

	svc := sa.NewUsuarioService(repo)

	res, ok := svc.Obtener("111")

	assert.True(t, ok)
	assert.Equal(t, esperado, res)
	repo.AssertExpectations(t)
}

func TestUsuarioService_Obtener_NoEncontrado(t *testing.T) {
	repo := new(usuarioRepoMock)

	repo.On("BuscarUsuarioPorCedula", "999").Return(modelos.Usuario{}, false)

	svc := sa.NewUsuarioService(repo)

	res, ok := svc.Obtener("999")

	assert.False(t, ok)
	assert.Equal(t, modelos.Usuario{}, res)
	repo.AssertExpectations(t)
}

func TestUsuarioService_Crear_CampoRequeridoFaltante(t *testing.T) {
	repo := new(usuarioRepoMock)

	svc := sa.NewUsuarioService(repo)

	// Sin Nombre -> debe fallar la validación antes de tocar el repo
	entrada := modelos.Usuario{Cedula: "111", Email: "ana@test.com"}

	res, err := svc.Crear(entrada)

	assert.ErrorIs(t, err, service.ErrCampoRequerido)
	assert.Equal(t, modelos.Usuario{}, res)
	repo.AssertNotCalled(t, "CrearUsuario", entrada)
}

func TestUsuarioService_Actualizar_Exitoso(t *testing.T) {
	repo := new(usuarioRepoMock)

	datos := modelos.Usuario{Cedula: "111", Nombre: "Ana Actualizada", Email: "ana@test.com"}
	repo.On("ActualizarUsuario", "111", datos).Return(datos, true)

	svc := sa.NewUsuarioService(repo)

	res, ok, err := svc.Actualizar("111", datos)

	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, datos, res)
	repo.AssertExpectations(t)
}

func TestUsuarioService_Actualizar_NoEncontrado(t *testing.T) {
	repo := new(usuarioRepoMock)

	datos := modelos.Usuario{Cedula: "999", Nombre: "Nadie", Email: "nadie@test.com"}
	repo.On("ActualizarUsuario", "999", datos).Return(modelos.Usuario{}, false)

	svc := sa.NewUsuarioService(repo)

	res, ok, err := svc.Actualizar("999", datos)

	assert.False(t, ok)
	assert.ErrorIs(t, err, service.ErrNoEncontrado)
	assert.Equal(t, modelos.Usuario{}, res)
	repo.AssertExpectations(t)
}

func TestUsuarioService_Actualizar_CampoRequeridoFaltante(t *testing.T) {
	repo := new(usuarioRepoMock)

	svc := sa.NewUsuarioService(repo)

	// Sin Email -> debe fallar la validación antes de tocar el repo
	datos := modelos.Usuario{Cedula: "111", Nombre: "Ana"}

	res, ok, err := svc.Actualizar("111", datos)

	assert.False(t, ok)
	assert.ErrorIs(t, err, service.ErrCampoRequerido)
	assert.Equal(t, modelos.Usuario{}, res)
	repo.AssertNotCalled(t, "ActualizarUsuario", "111", datos)
}

func TestUsuarioService_Borrar_Exitoso(t *testing.T) {
	repo := new(usuarioRepoMock)

	repo.On("BorrarUsuario", "111").Return(true)

	svc := sa.NewUsuarioService(repo)

	err := svc.Borrar("111")

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestUsuarioService_Borrar_NoEncontrado(t *testing.T) {
	repo := new(usuarioRepoMock)

	repo.On("BorrarUsuario", "999").Return(false)

	svc := sa.NewUsuarioService(repo)

	err := svc.Borrar("999")

	assert.ErrorIs(t, err, service.ErrNoEncontrado)
	repo.AssertExpectations(t)
}
