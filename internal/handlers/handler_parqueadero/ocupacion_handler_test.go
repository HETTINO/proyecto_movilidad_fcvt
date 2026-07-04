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

func TestCrearOcupacion_Exitoso(t *testing.T) {

	h, token := construirEntorno(t)

	body := `{
		"placa_vehiculo":"ABC-1234",
		"id_espacio":1,
		"id_acceso":1
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/ocupaciones",
		strings.NewReader(body),
	)

	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)

	var creada modelos.Ocupacion

	require.NoError(t, json.NewDecoder(rec.Body).Decode(&creada))

	assert.NotZero(t, creada.IDOcupacion)
	assert.Equal(t, "ABC-1234", creada.PlacaVehiculo)
	assert.Nil(t, creada.HoraFin)
}

func TestLiberarOcupacion_NoEncontrado(t *testing.T) {

	h, token := construirEntorno(t)

	req := httptest.NewRequest(
		http.MethodPatch,
		"/api/v1/ocupaciones/9999/liberar",
		nil,
	)

	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestListarOcupaciones_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/ocupaciones", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var lista []modelos.Ocupacion
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&lista))
	assert.Len(t, lista, 2) // las 2 sembradas en SeedOcupaciones
}

func TestObtenerOcupacion_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/ocupaciones/1", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var o modelos.Ocupacion
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&o))
	assert.Equal(t, "ABC1234", o.PlacaVehiculo)
}

func TestObtenerOcupacion_NoEncontrado(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/ocupaciones/9999", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestCrearOcupacion_JSONInvalido(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/ocupaciones", strings.NewReader("no-es-json"))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestActualizarOcupacion_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)

	body := `{"id_espacio":2,"placa_vehiculo":"UPD-0001"}`
	req := httptest.NewRequest(http.MethodPut, "/api/v1/ocupaciones/1", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	var actualizada modelos.Ocupacion
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&actualizada))
	assert.Equal(t, "UPD-0001", actualizada.PlacaVehiculo)
}

func TestActualizarOcupacion_NoEncontrado(t *testing.T) {
	h, token := construirEntorno(t)

	body := `{"id_espacio":1,"placa_vehiculo":"XXX-0000"}`
	req := httptest.NewRequest(http.MethodPut, "/api/v1/ocupaciones/9999", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestBorrarOcupacion_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/ocupaciones/2", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNoContent, rec.Code)
}

func TestBorrarOcupacion_NoEncontrado(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/ocupaciones/9999", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestLiberarOcupacion_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodPatch, "/api/v1/ocupaciones/1/liberar", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
