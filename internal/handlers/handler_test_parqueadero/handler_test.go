package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"proyecto_movilidad_fcvt/internal/handlers"
	"proyecto_movilidad_fcvt/internal/middleware"
	"proyecto_movilidad_fcvt/internal/modelos"
	"proyecto_movilidad_fcvt/internal/service"
	sp "proyecto_movilidad_fcvt/internal/service/service_parqueadero"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_parqueadero"
)

// usuarioRepoFake: repositorio de usuarios en memoria para los tests.
type usuarioRepoFake struct {
	porEmail map[string]modelos.Usuario
	nextID   uint // <- corregido de int a uint
}

func nuevoUsuarioRepoFake() *usuarioRepoFake {
	return &usuarioRepoFake{porEmail: map[string]modelos.Usuario{}, nextID: 1}
}

func (f *usuarioRepoFake) CrearUsuario(u modelos.Usuario) (modelos.Usuario, error) {
	u.ID = f.nextID // ahora ambos son uint
	f.nextID++
	f.porEmail[u.Email] = u
	return u, nil
}

func (f *usuarioRepoFake) BuscarUsuarioPorEmail(email string) (modelos.Usuario, bool) {
	u, ok := f.porEmail[email]
	return u, ok
}

func construirEntorno(t *testing.T) (http.Handler, string) {
	t.Helper()

	mem := storage.NuevaMemoria()
	mem.SeedParqueaderos()
	mem.SeedEspacios()
	mem.SeedOcupaciones()
	usuarios := nuevoUsuarioRepoFake()

	parqueaderoSvc := sp.NewParqueaderoService(mem)
	espacioSvc := sp.NewEspacioService(mem)
	ocupacionSvc := sp.NewOcupacionService(mem)
	authSvc := service.NewAuthService(usuarios)
	srv := handlers.NewServer(parqueaderoSvc, espacioSvc, ocupacionSvc, authSvc)

	r := chi.NewRouter()
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/register", srv.Registrar)
		r.Post("/auth/login", srv.Login)

		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authSvc))

			r.Get("/parqueaderos", srv.ListarParqueaderos)
			r.Post("/parqueaderos", srv.CrearParqueadero)
			r.Get("/parqueaderos/{id}", srv.ObtenerParqueadero)
			r.Put("/parqueaderos/{id}", srv.ActualizarParqueadero)
			r.Delete("/parqueaderos/{id}", srv.BorrarParqueadero)

			r.Get("/espacios", srv.ListarEspacios)
			r.Post("/espacios", srv.CrearEspacio)
			r.Get("/espacios/{id}", srv.ObtenerEspacio)
			r.Put("/espacios/{id}", srv.ActualizarEspacio)
			r.Delete("/espacios/{id}", srv.BorrarEspacio)

			r.Get("/ocupaciones", srv.ListarOcupaciones)
			r.Post("/ocupaciones", srv.CrearOcupacion)
			r.Get("/ocupaciones/{id}", srv.ObtenerOcupacion)
			r.Put("/ocupaciones/{id}", srv.ActualizarOcupacion)
			r.Delete("/ocupaciones/{id}", srv.BorrarOcupacion)
			r.Patch("/ocupaciones/{id}/liberar", srv.LiberarOcupacion)
		})
	})

	token := registrarYObtenerToken(t, r)
	return r, token
}

func registrarYObtenerToken(t *testing.T, h http.Handler) string {
	t.Helper()
	cred := `{"email":"docente@uleam.edu.ec","password":"secreta123"}`

	reqReg := httptest.NewRequest(http.MethodPost, "/api/v1/auth/register", strings.NewReader(cred))
	h.ServeHTTP(httptest.NewRecorder(), reqReg)

	reqLogin := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", strings.NewReader(cred))
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, reqLogin)
	require.Equal(t, http.StatusOK, rec.Code)

	var resp struct {
		Token string `json:"token"`
	}
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
	require.NotEmpty(t, resp.Token)
	return resp.Token
}

// =========================================================
// TESTS — Parqueaderos
// =========================================================

func TestCrearParqueadero_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)
	body := `{"nombre":"Parqueadero Central","capacidad":100,"tipo":"cubierto"}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/parqueaderos", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)
	var creado modelos.Parqueadero
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&creado))
	assert.NotZero(t, creado.IDParqueadero)
	assert.Equal(t, "Parqueadero Central", creado.Nombre)
}

func TestObtenerParqueadero_NoEncontrado(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/parqueaderos/9999", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

// =========================================================
// TESTS — Espacios
// =========================================================

func TestCrearEspacio_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)
	body := `{"id_parqueadero":1,"numero":99,"estado":"libre","tipo_espacio":"auto"}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/espacios", strings.NewReader(body))
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

	req := httptest.NewRequest(http.MethodGet, "/api/v1/espacios/9999", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

// =========================================================
// TESTS — Ocupaciones
// =========================================================

func TestCrearOcupacion_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)
	body := `{"placa_vehiculo":"ABC-1234","id_espacio":1,"id_acceso":1}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/ocupaciones", strings.NewReader(body))
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

	req := httptest.NewRequest(http.MethodPatch, "/api/v1/ocupaciones/9999/liberar", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

// =========================================================
// TESTS — Middleware
// =========================================================

func TestRutaProtegida_SinToken(t *testing.T) {
	h, _ := construirEntorno(t)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/parqueaderos",
		strings.NewReader(`{"nombre":"Test","capacidad":10,"tipo":"abierto"}`))
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}
