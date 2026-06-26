package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"

	"proyecto_movilidad_fcvt/internal/handlers"
	"proyecto_movilidad_fcvt/internal/middleware"
	"proyecto_movilidad_fcvt/internal/modelos"
	"proyecto_movilidad_fcvt/internal/service"

	sp "proyecto_movilidad_fcvt/internal/service/service_parqueadero"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_parqueadero"
)

// =====================================================
// Fake de usuarios
// =====================================================

type usuarioRepoFake struct {
	porEmail map[string]modelos.Usuario
	nextID   uint
}

func nuevoUsuarioRepoFake() *usuarioRepoFake {
	return &usuarioRepoFake{
		porEmail: map[string]modelos.Usuario{},
		nextID:   1,
	}
}

func (f *usuarioRepoFake) CrearUsuario(u modelos.Usuario) (modelos.Usuario, error) {
	u.ID = f.nextID
	f.nextID++

	f.porEmail[u.Email] = u

	return u, nil
}

func (f *usuarioRepoFake) BuscarUsuarioPorEmail(email string) (modelos.Usuario, bool) {
	u, ok := f.porEmail[email]
	return u, ok
}

// =====================================================
// Construcción del entorno
// =====================================================

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

	srv := handlers.NewServer(
		parqueaderoSvc,
		espacioSvc,
		ocupacionSvc,
		authSvc,
	)

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

// =====================================================
// Obtener Token
// =====================================================

func registrarYObtenerToken(t *testing.T, h http.Handler) string {

	t.Helper()

	cred := `{"email":"docente@uleam.edu.ec","password":"secreta123"}`

	reqReg := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/auth/register",
		strings.NewReader(cred),
	)

	h.ServeHTTP(httptest.NewRecorder(), reqReg)

	reqLogin := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/auth/login",
		strings.NewReader(cred),
	)

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
