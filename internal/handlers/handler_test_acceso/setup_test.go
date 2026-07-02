package handler_test_acceso

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"

	ha "proyecto_movilidad_fcvt/internal/handlers/handler_acceso"
	"proyecto_movilidad_fcvt/internal/middleware"
	"proyecto_movilidad_fcvt/internal/service"

	sa "proyecto_movilidad_fcvt/internal/service/service_acceso"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_acceso"
)

// =====================================================
// Construcción del entorno
// =====================================================

func construirEntorno(t *testing.T) (http.Handler, string) {
	t.Helper()

	// Un único almacén en memoria para todos los servicios,
	// incluido AuthService (que requiere *storage_acceso.MemoriaAcceso).
	mem := storage.NuevoMemoriaAcceso()

	// Servicios
	accesoSvc := sa.NewAccesoService(mem)
	usuarioSvc := sa.NewUsuarioService(mem)
	vehiculoSvc := sa.NewVehiculoService(mem)
	puntoAccesoSvc := sa.NewPuntoAccesoService(mem)
	authSvc := service.NewAuthService(mem)

	// El orden debe coincidir exactamente con la firma de ha.NewServer:
	// (auth, acceso, usuario, vehiculo, puntoAcceso)
	srv := ha.NewServer(
		authSvc,
		accesoSvc,
		usuarioSvc,
		vehiculoSvc,
		puntoAccesoSvc,
	)

	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {

		r.Post("/auth/register", srv.Registrar)
		r.Post("/auth/login", srv.Login)

		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authSvc))

			// Accesos
			r.Get("/accesos", srv.ListarAccesos)
			r.Post("/accesos", srv.CrearAcceso)
			r.Get("/accesos/{id}", srv.ObtenerAcceso)
			r.Put("/accesos/{id}", srv.ActualizarAcceso)
			r.Delete("/accesos/{id}", srv.BorrarAcceso)

			// Usuarios
			r.Get("/usuarios", srv.ListarUsuarios)
			r.Post("/usuarios", srv.CrearUsuario)
			r.Get("/usuarios/{id}", srv.ObtenerUsuario)
			r.Put("/usuarios/{id}", srv.ActualizarUsuario)
			r.Delete("/usuarios/{id}", srv.BorrarUsuario)

			// Vehículos
			r.Get("/vehiculos", srv.ListarVehiculos)
			r.Post("/vehiculos", srv.CrearVehiculo)
			r.Get("/vehiculos/{placa}", srv.ObtenerVehiculo)
			r.Put("/vehiculos/{placa}", srv.ActualizarVehiculo)
			r.Delete("/vehiculos/{placa}", srv.BorrarVehiculo)

			// Puntos de acceso
			r.Get("/puntos-acceso", srv.ListarPuntosAcceso)
			r.Post("/puntos-acceso", srv.CrearPuntoAcceso)
			r.Get("/puntos-acceso/{id}", srv.ObtenerPuntoAcceso)
			r.Put("/puntos-acceso/{id}", srv.ActualizarPuntoAcceso)
			r.Delete("/puntos-acceso/{id}", srv.BorrarPuntoAcceso)
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

	// Registrar requiere: cedula, nombre, email, contrasena, rol
	registro := `{
		"cedula": "1234567890",
		"nombre": "Docente Uleam",
		"email": "docente@uleam.edu.ec",
		"contrasena": "secreta123",
		"rol": "docente"
	}`

	reqReg := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/auth/register",
		strings.NewReader(registro),
	)

	recReg := httptest.NewRecorder()
	h.ServeHTTP(recReg, reqReg)
	require.Equal(t, http.StatusCreated, recReg.Code)

	// Login requiere: cedula, contrasena
	credenciales := `{
		"cedula": "1234567890",
		"contrasena": "secreta123"
	}`

	reqLogin := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/auth/login",
		strings.NewReader(credenciales),
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
