package handler_test_acceso

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
