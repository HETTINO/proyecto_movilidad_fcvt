package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"

	"proyecto_movilidad_fcvt/internal/middleware"
	"proyecto_movilidad_fcvt/internal/service"
)

const secretoDePrueba = "secreto-de-pruebas"

// handlerDePrueba es el "next" que el middleware debe (o no debe) alcanzar.
// Si se ejecuta, escribe 200 y, si hay un usuarioID en el contexto, lo refleja
// en el body — así comprobamos que el middleware sí inyectó el dato esperado.
func handlerDePrueba() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if v := r.Context().Value(middleware.ClaveUsuarioID); v != nil {
			w.Write([]byte(v.(string)))
		}
	})
}

// generarToken firma un token válido usando el mismo secreto que usará el middleware,
// para no depender de una base de datos real.
func generarToken(t *testing.T, cedula string, vencidoHace time.Duration) string {
	t.Helper()

	claims := &service.Claims{
		Cedula: cedula,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(vencidoHace)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	firmado, err := token.SignedString([]byte(secretoDePrueba))
	assert.NoError(t, err)
	return firmado
}

func construirAuthService() *service.AuthService {
	// repo=nil porque ValidarToken no necesita tocar el repositorio
	return service.NewAuthService(nil, service.WithSecreto([]byte(secretoDePrueba)))
}

func TestAuth_TokenValido_PermiteElPaso(t *testing.T) {
	auth := construirAuthService()
	tokenValido := generarToken(t, "1300000000", time.Hour)

	mw := middleware.Auth(auth)(handlerDePrueba())

	req := httptest.NewRequest(http.MethodGet, "/protegido", nil)
	req.Header.Set("Authorization", "Bearer "+tokenValido)
	rec := httptest.NewRecorder()

	mw.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "1300000000", rec.Body.String()) // confirma que inyectó la cédula en el contexto
}

func TestAuth_SinHeaderAuthorization_Rechaza(t *testing.T) {
	auth := construirAuthService()
	mw := middleware.Auth(auth)(handlerDePrueba())

	req := httptest.NewRequest(http.MethodGet, "/protegido", nil)
	rec := httptest.NewRecorder()

	mw.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestAuth_HeaderSinBearer_Rechaza(t *testing.T) {
	auth := construirAuthService()
	tokenValido := generarToken(t, "1300000000", time.Hour)
	mw := middleware.Auth(auth)(handlerDePrueba())

	req := httptest.NewRequest(http.MethodGet, "/protegido", nil)
	req.Header.Set("Authorization", tokenValido) // falta el prefijo "Bearer "
	rec := httptest.NewRecorder()

	mw.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestAuth_TokenMalFormado_Rechaza(t *testing.T) {
	auth := construirAuthService()
	mw := middleware.Auth(auth)(handlerDePrueba())

	req := httptest.NewRequest(http.MethodGet, "/protegido", nil)
	req.Header.Set("Authorization", "Bearer esto-no-es-un-jwt")
	rec := httptest.NewRecorder()

	mw.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestAuth_TokenExpirado_Rechaza(t *testing.T) {
	auth := construirAuthService()
	tokenExpirado := generarToken(t, "1300000000", -time.Hour) // venció hace 1h

	mw := middleware.Auth(auth)(handlerDePrueba())

	req := httptest.NewRequest(http.MethodGet, "/protegido", nil)
	req.Header.Set("Authorization", "Bearer "+tokenExpirado)
	rec := httptest.NewRecorder()

	mw.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestAuth_TokenFirmadoConOtroSecreto_Rechaza(t *testing.T) {
	auth := construirAuthService() // espera "secreto-de-pruebas"

	claims := &service.Claims{
		Cedula: "1300000000",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}
	tokenFalso, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte("secreto-de-un-atacante"))

	mw := middleware.Auth(auth)(handlerDePrueba())

	req := httptest.NewRequest(http.MethodGet, "/protegido", nil)
	req.Header.Set("Authorization", "Bearer "+tokenFalso)
	rec := httptest.NewRecorder()

	mw.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

// =========================================================
// TESTS — RequireAuth (el middleware simple que solo exige
// que exista el header, sin validar el JWT)
// =========================================================

func TestRequireAuth_SinToken_Rechaza(t *testing.T) {
	mw := middleware.RequireAuth(handlerDePrueba())

	req := httptest.NewRequest(http.MethodGet, "/protegido", nil)
	rec := httptest.NewRecorder()

	mw.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestRequireAuth_ConToken_PermiteElPaso(t *testing.T) {
	mw := middleware.RequireAuth(handlerDePrueba())

	req := httptest.NewRequest(http.MethodGet, "/protegido", nil)
	req.Header.Set("Authorization", "Bearer cualquier-cosa")
	rec := httptest.NewRecorder()

	mw.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
