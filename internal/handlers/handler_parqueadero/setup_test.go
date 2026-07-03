package handler_parqueadero_test

import (
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"

	hp "proyecto_movilidad_fcvt/internal/handlers/handler_parqueadero"
	"proyecto_movilidad_fcvt/internal/middleware"
	"proyecto_movilidad_fcvt/internal/service"

	sp "proyecto_movilidad_fcvt/internal/service/service_parqueadero"
	storageAcceso "proyecto_movilidad_fcvt/internal/storage/storage_acceso"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_parqueadero"
)

// =====================================================
// Construcción del entorno
// =====================================================

func construirEntorno(t *testing.T) (http.Handler, string) {
	t.Helper()

	mem := storage.NuevaMemoria()

	mem.SeedParqueaderos()
	mem.SeedEspacios()
	mem.SeedOcupaciones()

	// Repositorio de usuarios del módulo de acceso (en memoria) para autenticación
	usuarios := storageAcceso.NuevoMemoriaAcceso()

	parqueaderoSvc := sp.NewParqueaderoService(mem)
	espacioSvc := sp.NewEspacioService(mem)
	ocupacionSvc := sp.NewOcupacionService(mem)

	authSvc := service.NewAuthService(usuarios)

	srv := hp.NewServer(hp.Deps{
		Parqueadero: parqueaderoSvc,
		Espacio:     espacioSvc,
		Ocupacion:   ocupacionSvc,
		Auth:        authSvc,
	})

	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {

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

	token := registrarYObtenerToken(t, authSvc)

	return r, token
}

// =====================================================
// Obtener Token
// =====================================================

// registrarYObtenerToken crea un usuario directamente contra el AuthService
// (sin pasar por HTTP, ya que el registro/login vive ahora en el módulo de acceso)
// y devuelve un token válido para usarlo en los tests protegidos de parqueadero.
func registrarYObtenerToken(t *testing.T, authSvc *service.AuthService) string {
	t.Helper()

	const cedula = "0102030405"
	const password = "secreta123"

	_, err := authSvc.Registrar(cedula, "Docente Prueba", "docente@uleam.edu.ec", password, "docente")
	require.NoError(t, err)

	token, err := authSvc.Login(cedula, password)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	return token
}
