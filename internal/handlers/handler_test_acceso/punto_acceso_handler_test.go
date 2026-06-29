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

func TestCrearPuntoAcceso_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)

	body := `{
		"frecuencia": "Vehicular",
		"ubicacion": "Bloque Sur"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/puntos-acceso",
		strings.NewReader(body),
	)

	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)

	var creado modelos.PuntoDeAcceso
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&creado))

	assert.NotZero(t, creado.ID)
	assert.Equal(t, "Bloque Sur", creado.Ubicacion)
}
