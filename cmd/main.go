package main

import (
	"log"
	"net/http"

	_ "github.com/glebarez/go-sqlite"
	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"

	// Parqueadero
	handlerParqueadero "proyecto_movilidad_fcvt/internal/handlers/handler_parqueadero"
	modelosParqueadero "proyecto_movilidad_fcvt/internal/modelos"
	serviceParqueadero "proyecto_movilidad_fcvt/internal/service/service_parqueadero"
	storageParqueadero "proyecto_movilidad_fcvt/internal/storage/storage_parqueadero"

	// Transporte
	handlerTransporte "proyecto_movilidad_fcvt/internal/handlers/handler_transporte"
	modelosTransporte "proyecto_movilidad_fcvt/internal/modelos"
	serviceTransporte "proyecto_movilidad_fcvt/internal/service/service_transporte"
	storageTransporte "proyecto_movilidad_fcvt/internal/storage/storage_transporte"

	// Compartidos
	"proyecto_movilidad_fcvt/internal/middleware"
	"proyecto_movilidad_fcvt/internal/service"
	storageUser "proyecto_movilidad_fcvt/internal/storage"
)

func main() {

	// DB GORM
	gdb, err := gorm.Open(sqlite.Open("parqueadero.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("no se pudo abrir la base de datos: ", err)
	}

	// MIGRACIONES (TODO JUNTO)
	if err := gdb.AutoMigrate(
		&modelosParqueadero.Parqueadero{},
		&modelosParqueadero.Espacio{},
		&modelosParqueadero.Ocupacion{},

		&modelosTransporte.Ruta{},
		&modelosTransporte.Carrito{},
		&modelosTransporte.Parada{},
		&modelosTransporte.Locacion{},
		&modelosTransporte.Solicitud{},

		&modelosParqueadero.Usuario{},
	); err != nil {
		log.Fatal("falló AutoMigrate: ", err)
	}

	// =========================
	// STORAGE PARQUEADERO
	// =========================
	memParqueadero := storageParqueadero.NuevoAlmacenSQLite(gdb)
	memParqueadero.SembrarSiVacio()

	// =========================
	// STORAGE TRANSPORTE
	// =========================
	memTransporte := storageTransporte.NuevaMemoria()
	memTransporte.SeedRutas()
	memTransporte.SeedCarritos()
	memTransporte.SeedLocaciones()
	memTransporte.SeedParadas()
	memTransporte.SeedSolicitudes()

	// =========================
	// SERVICIOS PARQUEADERO
	// =========================
	usuarioRepo := storageUser.NewUsuarioGORM(gdb)

	authService := service.NewAuthService(usuarioRepo)

	parqueaderoService := serviceParqueadero.NewParqueaderoService(memParqueadero)
	espacioService := serviceParqueadero.NewEspacioService(memParqueadero)
	ocupacionService := serviceParqueadero.NewOcupacionService(memParqueadero)

	// =========================
	// SERVICIOS TRANSPORTE
	// =========================
	rutaService := serviceTransporte.NewRutaService(memTransporte)
	carritoService := serviceTransporte.NewCarritoService(memTransporte)
	paradaService := serviceTransporte.NewParadaService(memTransporte)
	locacionService := serviceTransporte.NewLocacionService(memTransporte)
	solicitudService := serviceTransporte.NewSolicitudService(memTransporte)

	// =========================
	// SERVERS
	// =========================
	parqueaderoServer := handlerParqueadero.NewServer(
		parqueaderoService,
		espacioService,
		ocupacionService,
		authService,
	)

	transporteServer := handlerTransporte.NewServer(
		rutaService,
		carritoService,
		paradaService,
		locacionService,
		solicitudService,
	)

	// =========================
	// ROUTER
	// =========================
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	r.Route("/api/v1", func(r chi.Router) {

		// AUTH (usa parqueadero)
		r.Post("/auth/register", parqueaderoServer.Registrar)
		r.Post("/auth/login", parqueaderoServer.Login)

		// PROTEGIDAS PARQUEADERO
		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authService))

			r.Get("/parqueaderos", parqueaderoServer.ListarParqueaderos)
			r.Post("/parqueaderos", parqueaderoServer.CrearParqueadero)
			r.Get("/parqueaderos/{id}", parqueaderoServer.ObtenerParqueadero)
			r.Put("/parqueaderos/{id}", parqueaderoServer.ActualizarParqueadero)
			r.Delete("/parqueaderos/{id}", parqueaderoServer.BorrarParqueadero)

			r.Get("/espacios", parqueaderoServer.ListarEspacios)
			r.Post("/espacios", parqueaderoServer.CrearEspacio)
			r.Get("/espacios/{id}", parqueaderoServer.ObtenerEspacio)
			r.Put("/espacios/{id}", parqueaderoServer.ActualizarEspacio)
			r.Delete("/espacios/{id}", parqueaderoServer.BorrarEspacio)

			r.Get("/ocupaciones", parqueaderoServer.ListarOcupaciones)
			r.Post("/ocupaciones", parqueaderoServer.CrearOcupacion)
			r.Get("/ocupaciones/{id}", parqueaderoServer.ObtenerOcupacion)
			r.Put("/ocupaciones/{id}", parqueaderoServer.ActualizarOcupacion)
			r.Delete("/ocupaciones/{id}", parqueaderoServer.BorrarOcupacion)
			r.Patch("/ocupaciones/{id}/liberar", parqueaderoServer.LiberarOcupacion)
		})

		// TRANSPORTE (sin auth o con auth si quieres)
		r.Group(func(r chi.Router) {

			r.Get("/rutas", transporteServer.ListarRutas)
			r.Post("/rutas", transporteServer.CrearRuta)
			r.Get("/rutas/{id}", transporteServer.ObtenerRuta)
			r.Put("/rutas/{id}", transporteServer.ActualizarRuta)
			r.Delete("/rutas/{id}", transporteServer.BorrarRuta)

			r.Get("/carritos", transporteServer.ListarCarritos)
			r.Post("/carritos", transporteServer.CrearCarrito)
			r.Get("/carritos/{id}", transporteServer.ObtenerCarrito)
			r.Put("/carritos/{id}", transporteServer.ActualizarCarrito)
			r.Delete("/carritos/{id}", transporteServer.BorrarCarrito)

			r.Get("/paradas", transporteServer.ListarParadas)
			r.Post("/paradas", transporteServer.CrearParada)
			r.Get("/paradas/{id}", transporteServer.ObtenerParada)
			r.Put("/paradas/{id}", transporteServer.ActualizarParada)
			r.Delete("/paradas/{id}", transporteServer.BorrarParada)

			r.Get("/locaciones", transporteServer.ListarLocaciones)
			r.Post("/locaciones", transporteServer.RegistrarLocacion)
			r.Get("/tiempo-estimado", transporteServer.GetTiempoEstimado)

			r.Get("/solicitudes", transporteServer.ListarSolicitudes)
			r.Post("/solicitudes", transporteServer.CrearSolicitud)
		})
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
