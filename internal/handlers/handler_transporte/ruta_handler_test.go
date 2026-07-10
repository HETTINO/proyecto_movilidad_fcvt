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

func TestActualizarRuta_Exitoso(t *testing.T) {

	h := construirEntorno(t)

	body := `{
		"nombre":"Ruta Editada",
		"descripcion":"Descripción editada"
	}`

	req := httptest.NewRequest(
		http.MethodPut,
		"/api/v1/rutas/1",
		strings.NewReader(body),
	)
	req.Header.Set("Authorization", "Bearer test_token")

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	var actualizada modelos.Ruta
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&actualizada))
	assert.Equal(t, "Ruta Editada", actualizada.Nombre)
}

func TestActualizarRuta_NoEncontrado(t *testing.T) {

	h := construirEntorno(t)

	body := `{
		"nombre":"No existe",
		"descripcion":"desc"
	}`

	req := httptest.NewRequest(
		http.MethodPut,
		"/api/v1/rutas/9999",
		strings.NewReader(body),
	)
	req.Header.Set("Authorization", "Bearer test_token")

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestBorrarRuta_Exitoso(t *testing.T) {

	h := construirEntorno(t)

	req := httptest.NewRequest(
		http.MethodDelete,
		"/api/v1/rutas/1",
		nil,
	)
	req.Header.Set("Authorization", "Bearer test_token")

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNoContent, rec.Code)
}

func TestObtenerRuta_IDInvalido(t *testing.T) {

	h := construirEntorno(t)

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/v1/rutas/abc",
		nil,
	)

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
