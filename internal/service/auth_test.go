package service

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"

	"proyecto_movilidad_fcvt/internal/modelos"
)

// =========================================================
// MOCK — UsuarioRepository (para probar AuthService aislado
// de la base de datos real, tal como pide la rúbrica del H3)
// =========================================================

type mockUsuarioRepo struct {
	mock.Mock
}

func (m *mockUsuarioRepo) ListarUsuarios() []modelos.Usuario {
	return m.Called().Get(0).([]modelos.Usuario)
}

func (m *mockUsuarioRepo) BuscarUsuarioPorCedula(cedula string) (modelos.Usuario, bool) {
	args := m.Called(cedula)
	return args.Get(0).(modelos.Usuario), args.Bool(1)
}

func (m *mockUsuarioRepo) CrearUsuario(u modelos.Usuario) modelos.Usuario {
	return m.Called(u).Get(0).(modelos.Usuario)
}

func (m *mockUsuarioRepo) ActualizarUsuario(cedula string, u modelos.Usuario) (modelos.Usuario, bool) {
	args := m.Called(cedula, u)
	return args.Get(0).(modelos.Usuario), args.Bool(1)
}

func (m *mockUsuarioRepo) BorrarUsuario(cedula string) bool {
	return m.Called(cedula).Bool(0)
}

// =========================================================
// TESTS — Registrar
// =========================================================

func TestAuthService_Registrar_OK(t *testing.T) {
	repo := new(mockUsuarioRepo)
	svc := NewAuthService(repo)

	// El usuario todavía no existe
	repo.On("BuscarUsuarioPorCedula", "1300000000").
		Return(modelos.Usuario{}, false)

	// CrearUsuario devuelve el usuario ya "guardado" (comportamiento típico del repo)
	repo.On("CrearUsuario", mock.AnythingOfType("modelos.Usuario")).
		Return(modelos.Usuario{Cedula: "1300000000", Nombre: "Ana", Email: "ana@test.com", Rol: "usuario"})

	u, err := svc.Registrar("1300000000", "Ana", "ana@test.com", "clave123", "usuario")

	assert.NoError(t, err)
	assert.Equal(t, "1300000000", u.Cedula)
	repo.AssertExpectations(t)
}

func TestAuthService_Registrar_CamposVacios(t *testing.T) {
	repo := new(mockUsuarioRepo)
	svc := NewAuthService(repo)

	casos := []struct {
		nombre                       string
		cedula, email, password, rol string
	}{
		{"cedula vacía", "", "a@test.com", "1234", "usuario"},
		{"email vacío", "1300000000", "", "1234", "usuario"},
		{"password vacío", "1300000000", "a@test.com", "", "usuario"},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			_, err := svc.Registrar(c.cedula, "Ana", c.email, c.password, c.rol)
			assert.ErrorIs(t, err, ErrCredencialesInvalidas)
		})
	}
	// El repo no debería haber sido llamado ni una vez: validación falla antes
	repo.AssertNotCalled(t, "BuscarUsuarioPorCedula", mock.Anything)
}

func TestAuthService_Registrar_CedulaYaExiste(t *testing.T) {
	repo := new(mockUsuarioRepo)
	svc := NewAuthService(repo)

	existente := modelos.Usuario{Cedula: "1300000000", Email: "ana@test.com"}
	repo.On("BuscarUsuarioPorCedula", "1300000000").Return(existente, true)

	_, err := svc.Registrar("1300000000", "Ana", "ana@test.com", "clave123", "usuario")

	assert.ErrorIs(t, err, ErrEmailenUso)
	repo.AssertNotCalled(t, "CrearUsuario", mock.Anything)
}

// =========================================================
// TESTS — Login
// =========================================================

func TestAuthService_Login_OK(t *testing.T) {
	repo := new(mockUsuarioRepo)
	svc := NewAuthService(repo, WithSecreto([]byte("secreto-test")), WithDuracion(time.Hour))

	hash, _ := bcrypt.GenerateFromPassword([]byte("clave123"), bcrypt.DefaultCost)
	usuario := modelos.Usuario{Cedula: "1300000000", Contrasena: string(hash)}

	repo.On("BuscarUsuarioPorCedula", "1300000000").Return(usuario, true)

	token, err := svc.Login("1300000000", "clave123")

	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// El token generado debe ser válido con el mismo servicio
	cedula, err := svc.ValidarToken(token)
	assert.NoError(t, err)
	assert.Equal(t, "1300000000", cedula)
}

func TestAuthService_Login_UsuarioNoExiste(t *testing.T) {
	repo := new(mockUsuarioRepo)
	svc := NewAuthService(repo)

	repo.On("BuscarUsuarioPorCedula", "9999999999").Return(modelos.Usuario{}, false)

	_, err := svc.Login("9999999999", "cualquiera")

	assert.ErrorIs(t, err, ErrCredencialesInvalidas)
}

func TestAuthService_Login_PasswordIncorrecta(t *testing.T) {
	repo := new(mockUsuarioRepo)
	svc := NewAuthService(repo)

	hash, _ := bcrypt.GenerateFromPassword([]byte("claveCorrecta"), bcrypt.DefaultCost)
	usuario := modelos.Usuario{Cedula: "1300000000", Contrasena: string(hash)}
	repo.On("BuscarUsuarioPorCedula", "1300000000").Return(usuario, true)

	_, err := svc.Login("1300000000", "claveIncorrecta")

	assert.ErrorIs(t, err, ErrCredencialesInvalidas)
}

func TestAuthService_Login_CamposVacios(t *testing.T) {
	repo := new(mockUsuarioRepo)
	svc := NewAuthService(repo)

	_, err := svc.Login("", "")

	assert.ErrorIs(t, err, ErrCredencialesInvalidas)
	repo.AssertNotCalled(t, "BuscarUsuarioPorCedula", mock.Anything)
}

// =========================================================
// TESTS — ValidarToken
// =========================================================

func TestAuthService_ValidarToken_TokenInvalido(t *testing.T) {
	svc := NewAuthService(nil)

	_, err := svc.ValidarToken("esto-no-es-un-jwt-valido")

	assert.ErrorIs(t, err, ErrCredencialesInvalidas)
}

func TestAuthService_ValidarToken_TokenExpirado(t *testing.T) {
	secreto := []byte("secreto-test")
	svc := NewAuthService(nil, WithSecreto(secreto))

	// Fabricamos a mano un token ya vencido para probar el caso edge
	claims := &Claims{
		Cedula: "1300000000",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Hour)), // expiró hace 1h
			IssuedAt:  jwt.NewNumericDate(time.Now().Add(-2 * time.Hour)),
		},
	}
	tokenExpirado, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secreto)
	assert.NoError(t, err)

	_, err = svc.ValidarToken(tokenExpirado)
	assert.ErrorIs(t, err, ErrCredencialesInvalidas)
}

func TestAuthService_ValidarToken_FirmaConSecretoDistinto(t *testing.T) {
	svc := NewAuthService(nil, WithSecreto([]byte("secreto-correcto")))

	claims := &Claims{
		Cedula: "1300000000",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}
	// Firmado con un secreto DISTINTO al que usa el servicio -> debe rechazarse
	tokenFalso, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("secreto-atacante"))

	_, err := svc.ValidarToken(tokenFalso)
	assert.ErrorIs(t, err, ErrCredencialesInvalidas)
}
