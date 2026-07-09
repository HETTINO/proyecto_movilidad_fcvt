package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	modelos "proyecto_movilidad_fcvt/internal/modelos"
)

func TestCrearParada_Exitoso(t *testing.T) {
	h := construirEntorno(t)

	body := `{
		"nombre": "Parada: Facultad FCVT",
		"latitud": -0.950,
		"longitud": -80.750,
		"ruta_id": 1
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/paradas", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)

	var creada modelos.Parada
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&creada))

	assert.NotZero(t, creada.IDParada)
	assert.Equal(t, "Parada: Facultad FCVT", creada.Nombre)
}

func TestObtenerParada_Exitoso(t *testing.T) {
	h := construirEntorno(t)

	// Consultar parada (asumiendo que SeedParadas cargó al menos una)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/paradas/1", nil)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestActualizarParada_Exitoso(t *testing.T) {
	h := construirEntorno(t)

	body := `{
		"nombre": "Parada: FCVT Actualizada",
		"latitud": -0.960,
		"longitud": -80.760,
		"ruta_id": 1
	}`

	req := httptest.NewRequest(http.MethodPut, "/api/v1/paradas/1", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestBorrarParada_Exitoso(t *testing.T) {
	h := construirEntorno(t)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/paradas/1", nil)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNoContent, rec.Code)
}

func TestListarParadas_Exitoso(t *testing.T) {
	h := construirEntorno(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/paradas", nil)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var lista []modelos.Parada
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&lista))
	assert.GreaterOrEqual(t, len(lista), 0)
}
