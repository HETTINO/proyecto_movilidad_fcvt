package service_acceso_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"proyecto_movilidad_fcvt/internal/modelos"
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
