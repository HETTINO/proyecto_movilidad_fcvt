package handler_parqueadero_test
import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRutaProtegida_SinToken(t *testing.T) {

	h, _ := construirEntorno(t)

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/parqueaderos",
		strings.NewReader(`{
			"nombre":"Test",
			"capacidad":10,
			"tipo":"abierto"
		}`),
	)

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}
