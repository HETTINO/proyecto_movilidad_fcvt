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

func crearVehiculoAuxiliar(t *testing.T, h http.Handler, token string) modelos.Vehiculo {
	t.Helper()

	body := `{
		"placa": "AUX-9999",
		"tipo_vehiculo": "Auxiliar",
		"marca": "Chevrolet"
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/vehiculos", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)

	var creado modelos.Vehiculo
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&creado))
	return creado
}

func TestCrearVehiculo_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)

	body := `{
		"placa": "MCE-2026",
		"tipo_vehiculo": "Docente",
		"marca": "Toyota"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/vehiculos",
		strings.NewReader(body),
	)

	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)

	var creado modelos.Vehiculo
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&creado))

	assert.Equal(t, "MCE-2026", creado.Placa)
}

func TestCrearVehiculo_JSONInvalido(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/vehiculos", strings.NewReader(`{invalido`))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCrearVehiculo_CampoRequeridoFaltante(t *testing.T) {
	h, token := construirEntorno(t)

	body := `{"placa": "", "marca": "Toyota"}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/vehiculos", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestListarVehiculos_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)
	crearVehiculoAuxiliar(t, h, token)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/vehiculos", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	var lista []modelos.Vehiculo
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&lista))
	assert.NotEmpty(t, lista)
}

func TestObtenerVehiculo_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)
	crearVehiculoAuxiliar(t, h, token)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/vehiculos/AUX-9999", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	var obtenido modelos.Vehiculo
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&obtenido))
	assert.Equal(t, "Chevrolet", obtenido.Marca)
}

func TestObtenerVehiculo_NoEncontrado(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/vehiculos/ZZZ-0000", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestActualizarVehiculo_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)
	crearVehiculoAuxiliar(t, h, token)

	body := `{
		"placa": "AUX-9999",
		"tipo_vehiculo": "Auxiliar",
		"marca": "Mazda"
	}`

	req := httptest.NewRequest(http.MethodPut, "/api/v1/vehiculos/AUX-9999", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	var actualizado modelos.Vehiculo
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&actualizado))
	assert.Equal(t, "Mazda", actualizado.Marca)
}

func TestActualizarVehiculo_JSONInvalido(t *testing.T) {
	h, token := construirEntorno(t)
	crearVehiculoAuxiliar(t, h, token)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/vehiculos/AUX-9999", strings.NewReader(`{invalido`))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestActualizarVehiculo_NoEncontrado(t *testing.T) {
	h, token := construirEntorno(t)

	body := `{"placa": "ZZZ-0000", "marca": "Nada"}`

	req := httptest.NewRequest(http.MethodPut, "/api/v1/vehiculos/ZZZ-0000", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestBorrarVehiculo_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)
	crearVehiculoAuxiliar(t, h, token)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/vehiculos/AUX-9999", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNoContent, rec.Code)
}

func TestBorrarVehiculo_NoEncontrado(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/vehiculos/ZZZ-0000", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}
