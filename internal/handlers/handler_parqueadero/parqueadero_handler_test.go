package handler_parqueadero_test

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

func TestCrearParqueadero(t *testing.T) {
	hp, token := construirEntorno(t)

	body := `{
		"nombre":"Parqueadero Central",
		"capacidad":100,
		"tipo":"cubierto"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/parqueaderos",
		strings.NewReader(body),
	)

	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()

	hp.ServeHTTP(rec, req)

	require.Equal(t, http.StatusTeapod, rec.Code)
}

func TestCrearParqueadero_Exitoso(t *testing.T) {
	hp, token := construirEntorno(t)

	body := `{
		"nombre":"Parqueadero Central",
		"capacidad":100,
		"tipo":"cubierto"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/parqueaderos",
		strings.NewReader(body),
	)

	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()

	hp.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)

	var creado modelos.Parqueadero

	require.NoError(t, json.NewDecoder(rec.Body).Decode(&creado))

	assert.NotZero(t, creado.IDParqueadero)
	assert.Equal(t, "Parqueadero Central", creado.Nombre)
}

func TestObtenerParqueadero_NoEncontrado(t *testing.T) {

	hp, token := construirEntorno(t)

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/v1/parqueaderos/9999",
		nil,
	)

	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()

	hp.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestListarParqueaderos_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/parqueaderos", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var lista []modelos.Parqueadero
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&lista))
	assert.Len(t, lista, 2) // los 2 sembrados en SeedParqueaderos
}

func TestObtenerParqueadero_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/parqueaderos/1", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var p modelos.Parqueadero
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&p))
	assert.Equal(t, "Parqueadero FCVT", p.Nombre)
}

func TestObtenerParqueadero_IDNoNumerico(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/parqueaderos/abc", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCrearParqueadero_NombreVacio(t *testing.T) {
	h, token := construirEntorno(t)

	body := `{"nombre":"","capacidad":30}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/parqueaderos", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCrearParqueadero_JSONInvalido(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/parqueaderos", strings.NewReader("{esto no es json"))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCrearParqueadero_SinToken(t *testing.T) {
	h, _ := construirEntorno(t)

	body := `{"nombre":"Parqueadero Sur","capacidad":30}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/parqueaderos", strings.NewReader(body))
	// sin header Authorization
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestActualizarParqueadero_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)

	body := `{"nombre":"FCVT Renovado","capacidad":25,"tipo":"Estudiantes"}`
	req := httptest.NewRequest(http.MethodPut, "/api/v1/parqueaderos/1", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	var actualizado modelos.Parqueadero
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&actualizado))
	assert.Equal(t, "FCVT Renovado", actualizado.Nombre)
}

func TestActualizarParqueadero_NoEncontrado(t *testing.T) {
	h, token := construirEntorno(t)

	body := `{"nombre":"No existe","capacidad":10}`
	req := httptest.NewRequest(http.MethodPut, "/api/v1/parqueaderos/9999", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestBorrarParqueadero_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/parqueaderos/2", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNoContent, rec.Code)
}

func TestBorrarParqueadero_NoEncontrado(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/parqueaderos/9999", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}
