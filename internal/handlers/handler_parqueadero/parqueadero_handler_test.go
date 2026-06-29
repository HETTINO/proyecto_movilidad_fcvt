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

func TestCrearParqueadero(t *testing.T) {
	hp, token := construirEntorno(t)

	body := `{
		"nombre":"Parqueadero Central",
		"capacidad":100,
		"tipo":"cubierto"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/parqueaderos",
		strings.NewReader(body),
	)

	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()

	hp.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)
}

func TestCrearParqueadero_Exitoso(t *testing.T) {
	hp, token := construirEntorno(t)

	body := `{
		"nombre":"Parqueadero Central",
		"capacidad":100,
		"tipo":"cubierto"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/parqueaderos",
		strings.NewReader(body),
	)

	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()

	hp.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)

	var creado modelos.Parqueadero

	require.NoError(t, json.NewDecoder(rec.Body).Decode(&creado))

	assert.NotZero(t, creado.IDParqueadero)
	assert.Equal(t, "Parqueadero Central", creado.Nombre)
}

func TestObtenerParqueadero_NoEncontrado(t *testing.T) {

	hp, token := construirEntorno(t)

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/v1/parqueaderos/9999",
		nil,
	)

	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()

	hp.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}
