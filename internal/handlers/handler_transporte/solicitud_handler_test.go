package handlers_test

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

	require.Equal(t, http.StatusTeapot, rec.Code)

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
