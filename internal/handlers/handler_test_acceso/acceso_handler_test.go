package handler_test_acceso

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
