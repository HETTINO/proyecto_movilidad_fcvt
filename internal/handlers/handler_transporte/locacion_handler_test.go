package handlers

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

func TestRegistrarLocacion_Exitoso(t *testing.T) {
	h := construirEntorno(t)

	// Simulación de un punto en el campus ULEAM
	body := `{
		"latitud": -0.950,
		"longitud": -80.750,
		"carrito_id": 1,
		"time_stamp": "2026-07-03T10:00:00Z"
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/locaciones", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)

	var registrada modelos.Locacion
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&registrada))

	assert.Equal(t, 1, registrada.CarritoID)
	assert.Equal(t, -0.950, registrada.Latitud)
}

func TestObtenerUbicacionCarrito_Exitoso(t *testing.T) {
	h := construirEntorno(t)

	// Consultar ubicación del carrito 1 (ya existe por el SeedLocaciones en setup_test.go)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/locaciones/carrito/1", nil)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var loc modelos.Locacion
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&loc))
	assert.Equal(t, 1, loc.CarritoID)
}

func TestGetTiempoEstimado_Exitoso(t *testing.T) {
	h := construirEntorno(t)

	// Probar cálculo estimado hacia el Paraninfo
	req := httptest.NewRequest(http.MethodGet, "/api/v1/tiempo-estimado?carrito_id=1&destino=Paraninfo", nil)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	
	// Validar que recibimos el tiempo dummy
	var res map[string]interface{}
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&res))
	assert.Equal(t, float64(15), res["tiempo_minutos"])
}

func TestListarLocaciones_Exitoso(t *testing.T) {
	h := construirEntorno(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/locaciones", nil)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	
	var lista []modelos.Locacion
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&lista))
	assert.GreaterOrEqual(t, len(lista), 0)
}