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

func crearPuntoAccesoAuxiliar(t *testing.T, h http.Handler, token string) modelos.PuntoDeAcceso {
	t.Helper()

	body := `{
		"frecuencia": "Peatonal",
		"ubicacion": "Bloque Auxiliar"
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/puntos-acceso", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)

	var creado modelos.PuntoDeAcceso
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&creado))
	return creado
}

func TestCrearPuntoAcceso_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)

	body := `{
		"frecuencia": "Vehicular",
		"ubicacion": "Bloque Sur"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/puntos-acceso",
		strings.NewReader(body),
	)

	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)

	var creado modelos.PuntoDeAcceso
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&creado))

	assert.NotZero(t, creado.ID)
	assert.Equal(t, "Bloque Sur", creado.Ubicacion)
}

func TestCrearPuntoAcceso_JSONInvalido(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/puntos-acceso", strings.NewReader(`{invalido`))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCrearPuntoAcceso_CampoRequeridoFaltante(t *testing.T) {
	h, token := construirEntorno(t)

	body := `{"frecuencia": "Alta", "ubicacion": ""}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/puntos-acceso", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestListarPuntosAcceso_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)
	crearPuntoAccesoAuxiliar(t, h, token)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/puntos-acceso", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	var lista []modelos.PuntoDeAcceso
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&lista))
	assert.NotEmpty(t, lista)
}

func TestObtenerPuntoAcceso_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)
	crearPuntoAccesoAuxiliar(t, h, token)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/puntos-acceso/1", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	var obtenido modelos.PuntoDeAcceso
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&obtenido))
	assert.Equal(t, "Bloque Auxiliar", obtenido.Ubicacion)
}

func TestObtenerPuntoAcceso_IDInvalido(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/puntos-acceso/no-es-numero", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestObtenerPuntoAcceso_NoEncontrado(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/puntos-acceso/9999", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestActualizarPuntoAcceso_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)
	crearPuntoAccesoAuxiliar(t, h, token)

	body := `{
		"frecuencia": "Baja",
		"ubicacion": "Bloque Modificado"
	}`

	req := httptest.NewRequest(http.MethodPut, "/api/v1/puntos-acceso/1", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	var actualizado modelos.PuntoDeAcceso
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&actualizado))
	assert.Equal(t, "Bloque Modificado", actualizado.Ubicacion)
}

func TestActualizarPuntoAcceso_IDInvalido(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/puntos-acceso/no-es-numero", strings.NewReader(`{}`))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestActualizarPuntoAcceso_JSONInvalido(t *testing.T) {
	h, token := construirEntorno(t)
	crearPuntoAccesoAuxiliar(t, h, token)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/puntos-acceso/1", strings.NewReader(`{invalido`))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestActualizarPuntoAcceso_NoEncontrado(t *testing.T) {
	h, token := construirEntorno(t)

	body := `{"frecuencia": "Alta", "ubicacion": "Bloque Z"}`

	req := httptest.NewRequest(http.MethodPut, "/api/v1/puntos-acceso/9999", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestBorrarPuntoAcceso_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)
	crearPuntoAccesoAuxiliar(t, h, token)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/puntos-acceso/1", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNoContent, rec.Code)
}

func TestBorrarPuntoAcceso_IDInvalido(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/puntos-acceso/no-es-numero", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestBorrarPuntoAcceso_NoEncontrado(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/puntos-acceso/9999", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}
