package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"

	handlers "proyecto_movilidad_fcvt/internal/handlers/handler_acceso"
	"proyecto_movilidad_fcvt/internal/middleware"
	"proyecto_movilidad_fcvt/internal/service"
	sa "proyecto_movilidad_fcvt/internal/service/service_acceso"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_acceso"
)

func main() {
	// 1. Almacén en memoria del módulo de acceso
	memoria := storage.NuevoMemoriaAcceso()
	log.Println("Datos de acceso listos")

	// 2. Servicios con inyección de dependencias
	authService := service.NewAuthService(memoria)
	accesoService := sa.NewAccesoService(memoria)
	usuarioService := sa.NewUsuarioService(memoria)
	vehiculoService := sa.NewVehiculoService(memoria)
	puntoAccesoService := sa.NewPuntoAccesoService(memoria)

	// 3. Server central pasándole los servicios en el orden exacto que requiere tu handlers.NewServer
	servidor := handlers.NewServer(authService, accesoService, usuarioService, vehiculoService, puntoAccesoService)

	// 4. Router + middleware
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	// 5. Rutas versionadas /api/v1/
	r.Route("/api/v1", func(r chi.Router) {

		// Públicas
		r.Post("/auth/register", servidor.Registrar)
		r.Post("/auth/login", servidor.Login)

		// Protegidas
		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authService))

			r.Get("/usuarios", servidor.ListarUsuarios)
			r.Post("/usuarios", servidor.CrearUsuario)
			r.Get("/usuarios/{id}", servidor.ObtenerUsuario)
			r.Put("/usuarios/{id}", servidor.ActualizarUsuario)
			r.Delete("/usuarios/{id}", servidor.BorrarUsuario)

			r.Get("/vehiculos", servidor.ListarVehiculos)
			r.Post("/vehiculos", servidor.CrearVehiculo)
			r.Get("/vehiculos/{placa}", servidor.ObtenerVehiculo)
			r.Put("/vehiculos/{placa}", servidor.ActualizarVehiculo)
			r.Delete("/vehiculos/{placa}", servidor.BorrarVehiculo)

			r.Get("/puntos-acceso", servidor.ListarPuntosAcceso)
			r.Post("/puntos-acceso", servidor.CrearPuntoAcceso)
			r.Get("/puntos-acceso/{id}", servidor.ObtenerPuntoAcceso)
			r.Put("/puntos-acceso/{id}", servidor.ActualizarPuntoAcceso)
			r.Delete("/puntos-acceso/{id}", servidor.BorrarPuntoAcceso)

			r.Get("/accesos", servidor.ListarAccesos)
			r.Post("/accesos", servidor.CrearAcceso)
			r.Get("/accesos/{id}", servidor.ObtenerAcceso)
			r.Put("/accesos/{id}", servidor.ActualizarAcceso)
			r.Delete("/accesos/{id}", servidor.BorrarAcceso)
		})
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
