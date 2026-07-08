package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/glebarez/go-sqlite"
	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/postgres"
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

	// Acceso
	handlerAcceso "proyecto_movilidad_fcvt/internal/handlers/handler_acceso"
	serviceAcceso "proyecto_movilidad_fcvt/internal/service/service_acceso"
	storageAcceso "proyecto_movilidad_fcvt/internal/storage/storage_acceso"

	// Compartidos
	"proyecto_movilidad_fcvt/internal/config"
	"proyecto_movilidad_fcvt/internal/httpserver"
	"proyecto_movilidad_fcvt/internal/middleware"
	"proyecto_movilidad_fcvt/internal/service"
)

func main() {
	cfg := config.Cargar()

	// DB GORM: elegimos el dialector según DB_DRIVER (sqlite local o postgres en docker)
	var gdb *gorm.DB
	var err error

	switch cfg.DBDriver {
	case "postgres":
		gdb, err = gorm.Open(postgres.Open(cfg.PostgresDSN), &gorm.Config{})
	default: // "sqlite" o cualquier valor no reconocido
		gdb, err = gorm.Open(sqlite.Open(cfg.RutaDB), &gorm.Config{})
	}
	if err != nil {
		log.Fatal("no se pudo abrir la base de datos: ", err)
	}

	// MIGRACIONES (TODO JUNTO)
	if err := gdb.AutoMigrate(
		&modelosParqueadero.Parqueadero{},
		&modelosParqueadero.Espacio{},
		&modelosParqueadero.Ocupacion{},

		// Transporte: el orden importa por las FKs (Ruta antes de Parada/Carrito,
		// Carrito antes de Locacion, Parada antes de Solicitud)
		&modelosTransporte.Ruta{},
		&modelosTransporte.Parada{},
		&modelosTransporte.Carrito{},
		&modelosTransporte.Locacion{},
		&modelosTransporte.Solicitud{},

		&modelosParqueadero.Usuario{},
		&modelosParqueadero.Vehiculo{},
		&modelosParqueadero.PuntoDeAcceso{},
		&modelosParqueadero.Acceso{},
	); err != nil {
		log.Fatal("falló AutoMigrate: ", err)
	}

	// =========================
	// STORAGE PARQUEADERO
	// =========================
	memParqueadero := storageParqueadero.NuevoAlmacenSQLite(gdb)
	memParqueadero.SembrarSiVacio()

	// =========================
	// STORAGE TRANSPORTE (GORM: SQLite en local, Postgres en docker)
	// =========================
	memTransporte := storageTransporte.NuevoAlmacenSQLite(gdb)
	memTransporte.SembrarSiVacio()

	// =========================
	// STORAGE ACCESO (sqlite + gorm)
	// =========================
	memAcceso := storageAcceso.NuevoAlmacenSQLite(gdb)
	memAcceso.SembrarSiVacio()

	// =========================
	// SERVICIOS PARQUEADERO
	// =========================
	authService := service.NewAuthService(memAcceso,
		service.WithSecreto(cfg.JWTSecreto),
		service.WithDuracion(cfg.JWTDuracion),
	)

	parqueaderoService := serviceParqueadero.NewParqueaderoService(memParqueadero)
	espacioService := serviceParqueadero.NewEspacioService(memParqueadero, memParqueadero)
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
	// SERVICIOS ACCESO
	// =========================
	accesoService := serviceAcceso.NewAccesoService(memAcceso)
	usuarioService := serviceAcceso.NewUsuarioService(memAcceso)
	vehiculoService := serviceAcceso.NewVehiculoService(memAcceso)
	puntoAccesoService := serviceAcceso.NewPuntoAccesoService(memAcceso)

	// =========================
	// SERVERS
	// =========================
	parqueaderoServer := handlerParqueadero.NewServer(handlerParqueadero.Deps{
		Parqueadero: parqueaderoService,
		Espacio:     espacioService,
		Ocupacion:   ocupacionService,
		Auth:        authService,
	})

	transporteServer := handlerTransporte.NewServer(handlerTransporte.Deps{
		Ruta:      rutaService,
		Carrito:   carritoService,
		Parada:    paradaService,
		Locacion:  locacionService,
		Solicitud: solicitudService,
	})

	accesoServer := handlerAcceso.NewServer(handlerAcceso.Deps{
		Auth:        authService,
		Acceso:      accesoService,
		Usuario:     usuarioService,
		Vehiculo:    vehiculoService,
		PuntoAcceso: puntoAccesoService,
	})

	// =========================
	// ROUTER
	// =========================
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	r.Route("/api/v1", func(r chi.Router) {

		// AUTH (usa el módulo de acceso)
		r.Post("/auth/register", accesoServer.Registrar)
		r.Post("/auth/login", accesoServer.Login)

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

		// TRANSPORTE
		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authService))

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
			r.Get("/locaciones/carrito/{id}", transporteServer.ObtenerUbicacionCarrito)
			r.Get("/tiempo-estimado", transporteServer.GetTiempoEstimado)

			r.Get("/solicitudes", transporteServer.ListarSolicitudes)
			r.Post("/solicitudes", transporteServer.CrearSolicitud)
			r.Get("/solicitudes/{id}", transporteServer.ObtenerSolicitud)
			r.Put("/solicitudes/{id}", transporteServer.ActualizarSolicitud)
			r.Delete("/solicitudes/{id}", transporteServer.BorrarSolicitud)
		})

		// PROTEGIDAS ACCESO
		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authService))

			r.Get("/usuarios", accesoServer.ListarUsuarios)
			r.Post("/usuarios", accesoServer.CrearUsuario)
			r.Get("/usuarios/{id}", accesoServer.ObtenerUsuario)
			r.Put("/usuarios/{id}", accesoServer.ActualizarUsuario)
			// Solo un admin puede borrar usuarios
			r.With(middleware.RequireRol("admin")).Delete("/usuarios/{id}", accesoServer.BorrarUsuario)

			r.Get("/vehiculos", accesoServer.ListarVehiculos)
			r.Post("/vehiculos", accesoServer.CrearVehiculo)
			r.Get("/vehiculos/{placa}", accesoServer.ObtenerVehiculo)
			r.Put("/vehiculos/{placa}", accesoServer.ActualizarVehiculo)
			r.Delete("/vehiculos/{placa}", accesoServer.BorrarVehiculo)

			// Puntos de acceso: solo un admin los crea, edita o borra
			r.Get("/puntos-acceso", accesoServer.ListarPuntosAcceso)
			r.Get("/puntos-acceso/{id}", accesoServer.ObtenerPuntoAcceso)
			r.With(middleware.RequireRol("admin")).Post("/puntos-acceso", accesoServer.CrearPuntoAcceso)
			r.With(middleware.RequireRol("admin")).Put("/puntos-acceso/{id}", accesoServer.ActualizarPuntoAcceso)
			r.With(middleware.RequireRol("admin")).Delete("/puntos-acceso/{id}", accesoServer.BorrarPuntoAcceso)

			r.Get("/accesos", accesoServer.ListarAccesos)
			r.Post("/accesos", accesoServer.CrearAcceso)
			r.Get("/accesos/{id}", accesoServer.ObtenerAcceso)
			r.Put("/accesos/{id}", accesoServer.ActualizarAcceso)
			r.Delete("/accesos/{id}", accesoServer.BorrarAcceso)
		})
	})

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	servidor := httpserver.Nuevo(cfg.Puerto, r)
	servidor.IniciarConGracefulShutdown(ctx, 10*time.Second)
}
