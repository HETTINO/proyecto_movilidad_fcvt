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

func TestCrearEspacio_Exitoso(t *testing.T) {

	h, token := construirEntorno(t)

	body := `{
		"id_parqueadero":1,
		"numero":99,
		"estado":"libre",
		"tipo_espacio":"auto"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/espacios",
		strings.NewReader(body),
	)

	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)

	var creado modelos.Espacio

	require.NoError(t, json.NewDecoder(rec.Body).Decode(&creado))

	assert.NotZero(t, creado.IDEspacio)
	assert.Equal(t, "libre", creado.Estado)
}

func TestObtenerEspacio_NoEncontrado(t *testing.T) {

	h, token := construirEntorno(t)

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/v1/espacios/9999",
		nil,
	)

	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}
