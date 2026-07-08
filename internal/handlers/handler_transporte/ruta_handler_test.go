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

func TestCrearRuta_Exitoso(t *testing.T) {

	h := construirEntorno(t)

	body := `{
		"nombre":"Ruta Centro",
		"descripcion":"Recorrido por el centro"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/rutas",
		strings.NewReader(body),
	)
	req.Header.Set("Authorization", "Bearer test_token")

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)

	var creada modelos.Ruta

	require.NoError(t, json.NewDecoder(rec.Body).Decode(&creada))

	assert.NotZero(t, creada.ID)
	assert.Equal(t, "Ruta Centro", creada.Nombre)
}

func TestObtenerRuta_NoEncontrado(t *testing.T) {

	h := construirEntorno(t)

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/v1/rutas/9999",
		nil,
	)

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}
