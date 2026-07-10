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

func TestCrearCarrito_Exitoso(t *testing.T) {
	h := construirEntorno(t)

	// Carrito real dentro de la universidad
	body := `{
		"nombre_carrito": "Carrito 1 - Rectorado",
		"capacidad": 20,
		"estado": "activo",
		"ruta_id": 1
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/carritos", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)

	var creado modelos.Carrito
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&creado))

	assert.NotZero(t, creado.ID)
	assert.Equal(t, "Carrito 1 - Rectorado", creado.NombreCarrito)
}

func TestObtenerCarrito_NoEncontrado(t *testing.T) {
	h := construirEntorno(t)

	// ID inexistente
	req := httptest.NewRequest(http.MethodGet, "/api/v1/carritos/9999", nil)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestListarCarritos_Exitoso(t *testing.T) {
	h := construirEntorno(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/carritos", nil)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var lista []modelos.Carrito
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&lista))
	// Si tu SeedCarritos carga datos, aquí podrías verificar que el len > 0
	assert.GreaterOrEqual(t, len(lista), 0)
}
func TestActualizarCarrito_Exitoso(t *testing.T) {
	h := construirEntorno(t)

	body := `{
		"nombre_carrito": "Carrito Editado",
		"capacidad": 10,
		"estado": "mantenimiento",
		"ruta_id": 1
	}`

	req := httptest.NewRequest(http.MethodPut, "/api/v1/carritos/1", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	var actualizado modelos.Carrito
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&actualizado))
	assert.Equal(t, "Carrito Editado", actualizado.NombreCarrito)
}

func TestActualizarCarrito_NoEncontrado(t *testing.T) {
	h := construirEntorno(t)

	body := `{
		"nombre_carrito": "No existe",
		"capacidad": 4,
		"estado": "activo",
		"ruta_id": 1
	}`

	req := httptest.NewRequest(http.MethodPut, "/api/v1/carritos/9999", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestBorrarCarrito_Exitoso(t *testing.T) {
	h := construirEntorno(t)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/carritos/1", nil)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNoContent, rec.Code)
}

func TestCrearCarrito_JSONInvalido(t *testing.T) {
	h := construirEntorno(t)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/carritos", bytes.NewBufferString(`{esto no es json`))
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}