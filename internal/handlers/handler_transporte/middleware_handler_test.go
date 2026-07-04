package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRutaProtegida_SinToken(t *testing.T) {

	h := construirEntorno(t) // ← Solo 1 valor

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/solicitudes",
		strings.NewReader(`{
			"cedula_usuario":"0102030405",
			"cant_personas":2,
			"parada_origen":1,
			"punto_destino":"Biblioteca"
		}`),
	)

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}
