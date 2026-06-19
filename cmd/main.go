package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"proyecto_movilidad_fcvt/internal/handlers"
	"proyecto_movilidad_fcvt/internal/middleware"
	"proyecto_movilidad_fcvt/internal/modelos"
	"proyecto_movilidad_fcvt/internal/service"
	sp "proyecto_movilidad_fcvt/internal/service/service_parqueadero"
	us "proyecto_movilidad_fcvt/internal/storage"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_parqueadero"
)

func main() {
	// 1. GORM abre la DB y migra el esquema.
	gdb, err := gorm.Open(sqlite.Open("parqueadero.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("no se pudo abrir la base de datos: ", err)
	}
	if err := gdb.AutoMigrate(
		&modelos.Parqueadero{},
		&modelos.Espacio{},
		&modelos.Ocupacion{},
		&modelos.Usuario{},
	); err != nil {
		log.Fatal("falló AutoMigrate: ", err)
	}

	// 2. Almacén en memoria + seed de datos de prueba.
	memoria := storage.NuevaMemoria()
	memoria.SeedParqueaderos()
	memoria.SeedEspacios()
	memoria.SeedOcupaciones()
	log.Println("Datos de prueba cargados")

	// 3. Servicios con inyección de dependencias.
	usuarioRepo := us.NewUsuarioGORM(gdb)
	authService := service.NewAuthService(usuarioRepo)
	parqueaderoService := sp.NewParqueaderoService(memoria)
	espacioService := sp.NewEspacioService(memoria)
	ocupacionService := sp.NewOcupacionService(memoria)

	// 4. Server central con todos los servicios.
	servidor := handlers.NewServer(parqueaderoService, espacioService, ocupacionService, authService)

	// 5. Router + middleware.
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	// 6. Rutas versionadas /api/v1/.
	r.Route("/api/v1", func(r chi.Router) {

		// Públicas
		r.Post("/auth/register", servidor.Registrar)
		r.Post("/auth/login", servidor.Login)

		// Protegidas
		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authService))

			r.Get("/parqueaderos", servidor.ListarParqueaderos)
			r.Post("/parqueaderos", servidor.CrearParqueadero)
			r.Get("/parqueaderos/{id}", servidor.ObtenerParqueadero)
			r.Put("/parqueaderos/{id}", servidor.ActualizarParqueadero)
			r.Delete("/parqueaderos/{id}", servidor.BorrarParqueadero)

			r.Get("/espacios", servidor.ListarEspacios)
			r.Post("/espacios", servidor.CrearEspacio)
			r.Get("/espacios/{id}", servidor.ObtenerEspacio)
			r.Put("/espacios/{id}", servidor.ActualizarEspacio)
			r.Delete("/espacios/{id}", servidor.BorrarEspacio)

			r.Get("/ocupaciones", servidor.ListarOcupaciones)
			r.Post("/ocupaciones", servidor.CrearOcupacion)
			r.Get("/ocupaciones/{id}", servidor.ObtenerOcupacion)
			r.Put("/ocupaciones/{id}", servidor.ActualizarOcupacion)
			r.Delete("/ocupaciones/{id}", servidor.BorrarOcupacion)
			r.Patch("/ocupaciones/{id}/liberar", servidor.LiberarOcupacion)
		})
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
