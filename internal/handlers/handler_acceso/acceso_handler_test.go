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

func crearAccesoAuxiliar(t *testing.T, h http.Handler, token string) modelos.Acceso {
	t.Helper()

	body := `{
		"placa_vehiculo": "AUX-0001",
		"estado": "activo",
		"observaciones": "Ingreso auxiliar para pruebas"
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/accesos", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)

	var creado modelos.Acceso
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&creado))
	return creado
}

func TestCrearAcceso_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)

	body := `{
		"placa_vehiculo": "ABC-1234",
		"estado": "activo",
		"observaciones": "Ingreso regular"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/accesos",
		strings.NewReader(body),
	)

	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)

	var creado modelos.Acceso
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&creado))

	assert.NotZero(t, creado.ID)
	assert.Equal(t, "ABC-1234", creado.PlacaVehiculo)
}

func TestCrearAcceso_JSONInvalido(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/accesos", strings.NewReader(`{invalido`))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestListarAccesos_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)
	crearAccesoAuxiliar(t, h, token)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/accesos", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	var lista []modelos.Acceso
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&lista))
	assert.NotEmpty(t, lista)
}

func TestObtenerAcceso_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)
	crearAccesoAuxiliar(t, h, token)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/accesos/1", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	var obtenido modelos.Acceso
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&obtenido))
	assert.Equal(t, "AUX-0001", obtenido.PlacaVehiculo)
}

func TestObtenerAcceso_NoEncontrado(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/v1/accesos/9999",
		nil,
	)

	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestObtenerAcceso_IDInvalido(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/accesos/no-es-numero", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestActualizarAcceso_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)
	crearAccesoAuxiliar(t, h, token)

	body := `{
		"placa_vehiculo": "AUX-0001",
		"estado": "cerrado",
		"observaciones": "Salida registrada"
	}`

	req := httptest.NewRequest(http.MethodPut, "/api/v1/accesos/1", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	var actualizado modelos.Acceso
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&actualizado))
	assert.Equal(t, "cerrado", actualizado.Estado)
}

func TestActualizarAcceso_IDInvalido(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/accesos/no-es-numero", strings.NewReader(`{}`))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestActualizarAcceso_JSONInvalido(t *testing.T) {
	h, token := construirEntorno(t)
	crearAccesoAuxiliar(t, h, token)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/accesos/1", strings.NewReader(`{invalido`))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestActualizarAcceso_NoEncontrado(t *testing.T) {
	h, token := construirEntorno(t)

	body := `{"placa_vehiculo": "ZZZ-9999", "estado": "activo"}`

	req := httptest.NewRequest(http.MethodPut, "/api/v1/accesos/9999", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestBorrarAcceso_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)
	crearAccesoAuxiliar(t, h, token)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/accesos/1", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNoContent, rec.Code)
}

func TestBorrarAcceso_IDInvalido(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/accesos/no-es-numero", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestBorrarAcceso_NoEncontrado(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/accesos/9999", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

// =====================================================
// Test de Handler 401 (sin token / token inválido)
// =====================================================

func TestListarAccesos_SinToken_Devuelve401(t *testing.T) {
	h, _ := construirEntorno(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/accesos", nil)
	// Sin cabecera Authorization

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestListarAccesos_TokenInvalido_Devuelve401(t *testing.T) {
	h, _ := construirEntorno(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/accesos", nil)
	req.Header.Set("Authorization", "Bearer token-falso-o-vencido")

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestListarAccesos_FormatoAuthorizationMalo_Devuelve401(t *testing.T) {
	h, _ := construirEntorno(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/accesos", nil)
	// Sin el prefijo "Bearer "
	req.Header.Set("Authorization", "token-sin-prefijo")

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}
