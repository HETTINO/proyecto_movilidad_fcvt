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
