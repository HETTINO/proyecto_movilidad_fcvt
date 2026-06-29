package main

import (
    "log"
    "net/http"
	

    "github.com/go-chi/chi/v5"
    chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
    
	"proyecto_movilidad_fcvt/internal/handlers/handler_transporte"
	"proyecto_movilidad_fcvt/internal/middleware"
	modelos "proyecto_movilidad_fcvt/internal/models"
	sp "proyecto_movilidad_fcvt/internal/service/service_transporte"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_transporte"
)

func main() {
	// 1. GORM abre la DB y migra el esquema.
	gdb, err := gorm.Open(sqlite.Open("parqueadero.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("no se pudo abrir la base de datos: ", err)
	}
	if err := gdb.AutoMigrate(
		&modelos.Ruta{},
		&modelos.Carrito{},
		&modelos.Locacion{},
		&modelos.Parada{},
		&modelos.Solicitud{},
	); err != nil {
		log.Fatal("falló AutoMigrate: ", err)
	}

	// 2. Almacén en memoria + seed de datos de prueba.
	memoria := storage.NuevaMemoria()
	memoria.SeedRutas()
	memoria.SeedCarritos()
	memoria.SeedLocaciones()
	memoria.SeedParadas()
	memoria.SeedSolicitudes()
	log.Println("Datos de prueba cargados")

	// 3. Servicios con inyección de dependencias.
	rutaService := sp.NewRutaService(memoria)
	carritoService := sp.NewCarritoService(memoria)
	paradaService := sp.NewParadaService(memoria)
	locacionService := sp.NewLocacionService(memoria)
	solicitudService := sp.NewSolicitudService(memoria)

	// 4. Server central con todos los servicios.
	servidor := handlers.NewServer(rutaService, carritoService, paradaService, locacionService, solicitudService)

	// 5. Router + middleware.
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	// 6. Rutas versionadas /api/v1/.
	r.Route("/api/v1", func(r chi.Router) {
    r.Group(func(r chi.Router) {

        // Rutas - RUTA
        r.Get("/rutas", servidor.ListarRutas)
        r.Get("/rutas/{id}", servidor.ObtenerRuta)
        r.Post("/rutas", servidor.CrearRuta)
        r.Put("/rutas/{id}", servidor.ActualizarRuta)
        r.Delete("/rutas/{id}", servidor.BorrarRuta)

        // Rutas - CARRITO
        r.Get("/carritos", servidor.ListarCarritos)
        r.Get("/carritos/{id}", servidor.ObtenerCarrito)
        r.Post("/carritos", servidor.CrearCarrito)
        r.Put("/carritos/{id}", servidor.ActualizarCarrito)
        r.Delete("/carritos/{id}", servidor.BorrarCarrito)

        // Rutas - PARADA
        r.Get("/paradas", servidor.ListarParadas)
        r.Get("/paradas/{id}", servidor.ObtenerParada)
        r.Post("/paradas", servidor.CrearParada)
        r.Put("/paradas/{id}", servidor.ActualizarParada)
        r.Delete("/paradas/{id}", servidor.BorrarParada)

        // Rutas - LOCACION
        r.Get("/locaciones", servidor.ListarLocaciones)
        r.Get("/locaciones/carrito/{id}", servidor.ObtenerUbicacionCarrito)
        r.Post("/locaciones", servidor.RegistrarLocacion)
        r.Get("/tiempo-estimado", servidor.GetTiempoEstimado)

        // Rutas - SOLICITUD
        r.Get("/solicitudes", servidor.ListarSolicitudes)
        r.Get("/solicitudes/{id}", servidor.ObtenerSolicitud)
        r.Post("/solicitudes", servidor.CrearSolicitud)
        r.Put("/solicitudes/{id}", servidor.ActualizarSolicitud)
        r.Delete("/solicitudes/{id}", servidor.BorrarSolicitud)
    })
})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
