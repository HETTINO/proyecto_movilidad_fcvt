package handler_acceso_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"proyecto_movilidad_fcvt/internal/modelos"
)

func crearUsuarioAuxiliar(t *testing.T, h http.Handler, token string) modelos.Usuario {
	t.Helper()

	body := `{
		"cedula": "99999999",
		"nombre": "Usuario Auxiliar",
		"email": "aux@example.com",
		"rol": "usuario"
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/usuarios", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)

	var creado modelos.Usuario
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&creado))
	return creado
}

func TestCrearUsuario_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)

	body := `{
		"cedula": "13165432",
		"nombre": "Shirley Juleidy",
		"email": "shirley.j@example.com",
		"rol": "admin"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/usuarios",
		strings.NewReader(body),
	)

	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)

	var creado modelos.Usuario
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&creado))

	assert.Equal(t, "13165432", creado.Cedula)
	assert.Equal(t, "Shirley Juleidy", creado.Nombre)
}

func TestCrearUsuario_JSONInvalido(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/usuarios", strings.NewReader(`{invalido`))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCrearUsuario_CampoRequeridoFaltante(t *testing.T) {
	h, token := construirEntorno(t)

	body := `{"cedula": "", "nombre": "", "email": ""}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/usuarios", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestListarUsuarios_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)
	crearUsuarioAuxiliar(t, h, token)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/usuarios", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	var lista []modelos.Usuario
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&lista))
	assert.NotEmpty(t, lista)
}

func TestObtenerUsuario_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)
	crearUsuarioAuxiliar(t, h, token)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/usuarios/99999999", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	var obtenido modelos.Usuario
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&obtenido))
	assert.Equal(t, "Usuario Auxiliar", obtenido.Nombre)
}

func TestObtenerUsuario_NoEncontrado(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/usuarios/00000000", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestActualizarUsuario_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)
	crearUsuarioAuxiliar(t, h, token)

	body := `{
		"cedula": "99999999",
		"nombre": "Usuario Modificado",
		"email": "aux@example.com",
		"rol": "usuario"
	}`

	req := httptest.NewRequest(http.MethodPut, "/api/v1/usuarios/99999999", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	var actualizado modelos.Usuario
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&actualizado))
	assert.Equal(t, "Usuario Modificado", actualizado.Nombre)
}

func TestActualizarUsuario_JSONInvalido(t *testing.T) {
	h, token := construirEntorno(t)
	crearUsuarioAuxiliar(t, h, token)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/usuarios/99999999", strings.NewReader(`{invalido`))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestActualizarUsuario_NoEncontrado(t *testing.T) {
	h, token := construirEntorno(t)

	body := `{
		"cedula": "00000000",
		"nombre": "Nadie",
		"email": "nadie@example.com",
		"rol": "usuario"
	}`

	req := httptest.NewRequest(http.MethodPut, "/api/v1/usuarios/00000000", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestBorrarUsuario_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)
	crearUsuarioAuxiliar(t, h, token)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/usuarios/99999999", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNoContent, rec.Code)
}

func TestBorrarUsuario_NoEncontrado(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/usuarios/00000000", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}
