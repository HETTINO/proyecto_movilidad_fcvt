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

func TestActualizarEspacio_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)

	body := `{"id_parqueadero":1,"numero":1,"estado":"ocupado","tipo_espacio":"auto"}`
	req := httptest.NewRequest(http.MethodPut, "/api/v1/espacios/1", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	var actualizado modelos.Espacio
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&actualizado))
	assert.Equal(t, "ocupado", actualizado.Estado)
}

func TestActualizarEspacio_NoEncontrado(t *testing.T) {
	h, token := construirEntorno(t)

	body := `{"id_parqueadero":1,"numero":1,"estado":"libre"}`
	req := httptest.NewRequest(http.MethodPut, "/api/v1/espacios/9999", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

// TestBorrarEspacio_Exitoso borra el espacio 1, que en la semilla NO tiene
// ninguna ocupación activa apuntándole (las ocupaciones sembradas están en
// los espacios 2 y 3), por eso aquí sí se espera 204.
func TestBorrarEspacio_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/espacios/1", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNoContent, rec.Code)
}

// TestBorrarEspacio_ConOcupacionesActivas verifica que el sistema NO permita
// borrar un espacio que todavía tiene una ocupación activa (sin HoraFin)
// apuntándole: el espacio 3 tiene la ocupación sembrada con IDOcupacion=2.
func TestBorrarEspacio_ConOcupacionesActivas(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/espacios/3", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusConflict, rec.Code)
}

func TestBorrarEspacio_NoEncontrado(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/espacios/9999", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}
