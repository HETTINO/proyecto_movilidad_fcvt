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

func TestCrearUsuario_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)

	body := `{
		"cedula": "13165432",
		"nombre": "Shirley Juleidy",
		"email": "shirley.j@example.com",
		"rol": "admin"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/usuarios",
		strings.NewReader(body),
	)

	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)

	var creado modelos.Usuario
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&creado))

	assert.Equal(t, "13165432", creado.Cedula)
	assert.Equal(t, "Shirley Juleidy", creado.Nombre)
}
