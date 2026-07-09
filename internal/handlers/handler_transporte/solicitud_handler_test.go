package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	modelos "proyecto_movilidad_fcvt/internal/modelos"
)

func TestCrearSolicitud_Exitoso(t *testing.T) {

	h := construirEntorno(t)

	body := `{
		"cedula_usuario":"0102030405",
		"cant_personas":2,
		"parada_origen":1,
		"punto_destino":"Biblioteca"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/solicitudes",
		strings.NewReader(body),
	)
	req.Header.Set("Authorization", "Bearer test_token")

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)

	var creada modelos.Solicitud

	require.NoError(t, json.NewDecoder(rec.Body).Decode(&creada))

	assert.NotZero(t, creada.ID)
	assert.Equal(t, "pendiente", creada.Estado)
}

func TestObtenerSolicitud_NoEncontrado(t *testing.T) {

	h := construirEntorno(t)

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/v1/solicitudes/9999",
		nil,
	)

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestObtenerSolicitud_Exitoso(t *testing.T) {

	h := construirEntorno(t)

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/v1/solicitudes/1",
		nil,
	)

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	var sol modelos.Solicitud
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&sol))

	assert.Equal(t, 1, sol.ID)
	assert.Equal(t, "Biblioteca", sol.PuntoDestino)
}

func TestObtenerSolicitud_IDInvalido(t *testing.T) {

	h := construirEntorno(t)

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/v1/solicitudes/abc",
		nil,
	)

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
func TestListarSolicitudes(t *testing.T) {

	h := construirEntorno(t)

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/v1/solicitudes",
		nil,
	)

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	var lista []modelos.Solicitud
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&lista))

	assert.Len(t, lista, 3) // seed carga 3 solicitudes
}

func TestCrearSolicitud_DatosInvalidos(t *testing.T) {

	h := construirEntorno(t)

	body := `{
		"cedula_usuario":"",
		"cant_personas":2,
		"parada_origen":1,
		"punto_destino":"Biblioteca"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/solicitudes",
		strings.NewReader(body),
	)
	req.Header.Set("Authorization", "Bearer test_token")

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
func TestActualizarSolicitud_Exitoso(t *testing.T) {

	h := construirEntorno(t)

	body := `{
		"cedula_usuario":"1234567890",
		"cant_personas":3,
		"parada_origen":1,
		"punto_destino":"Biblioteca",
		"estado":"completada"
	}`

	req := httptest.NewRequest(
		http.MethodPut,
		"/api/v1/solicitudes/1",
		strings.NewReader(body),
	)
	req.Header.Set("Authorization", "Bearer test_token")

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	var actualizada modelos.Solicitud
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&actualizada))
	assert.Equal(t, "completada", actualizada.Estado)
}

func TestActualizarSolicitud_NoEncontrado(t *testing.T) {

	h := construirEntorno(t)

	body := `{
		"cedula_usuario":"1234567890",
		"cant_personas":3,
		"parada_origen":1,
		"punto_destino":"Biblioteca"
	}`

	req := httptest.NewRequest(
		http.MethodPut,
		"/api/v1/solicitudes/9999",
		strings.NewReader(body),
	)
	req.Header.Set("Authorization", "Bearer test_token")

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestBorrarSolicitud_Exitoso(t *testing.T) {

	h := construirEntorno(t)

	req := httptest.NewRequest(
		http.MethodDelete,
		"/api/v1/solicitudes/1",
		nil,
	)
	req.Header.Set("Authorization", "Bearer test_token")

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNoContent, rec.Code)
}
func TestBorrarSolicitud_NoEncontrado(t *testing.T) {

	h := construirEntorno(t)

	req := httptest.NewRequest(
		http.MethodDelete,
		"/api/v1/solicitudes/9999",
		nil,
	)
	req.Header.Set("Authorization", "Bearer test_token")

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}