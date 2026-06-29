package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/glebarez/go-sqlite" // driver database/sql "sqlite" (pure-Go) para el backend sqlc
	"github.com/glebarez/sqlite"      // driver GORM (pure-Go)
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"

	handlers "proyecto_movilidad_fcvt/internal/handlers/handler_parqueadero"
	"proyecto_movilidad_fcvt/internal/middleware"
	"proyecto_movilidad_fcvt/internal/modelos"
	"proyecto_movilidad_fcvt/internal/service"
	sp "proyecto_movilidad_fcvt/internal/service/service_parqueadero"
	us "proyecto_movilidad_fcvt/internal/storage"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_parqueadero"
)

func main() {
	// 1. GORM es el DUEÑO DEL ESQUEMA: abre la DB, migra y siembra.
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
	almacenGorm := storage.NuevoAlmacenSQLite(gdb)
	almacenGorm.SembrarSiVacio()

	// 2. Elegir el backend según STORAGE.
	var almacen storage.Almacen
	switch os.Getenv("STORAGE") {
	case "sqlc":
		sdb, err := sql.Open("sqlite", "parqueadero.db")
		if err != nil {
			log.Fatal("no se pudo abrir sql.DB para sqlc: ", err)
		}
		almacen = storage.NuevoAlmacenSQLC(sdb)
		log.Println("Backend: sqlc (database/sql)")
	case "memoria":
		mem := storage.NuevaMemoria()
		mem.SeedParqueaderos()
		mem.SeedEspacios()
		mem.SeedOcupaciones()
		almacen = mem
		log.Println("Backend: MEMORIA")
	default:
		almacen = almacenGorm
		log.Println("Backend: GORM")
	}

	// 3. Los usuarios viven SIEMPRE en GORM.
	usuarioRepo := us.NewUsuarioGORM(gdb)

	// 4. Capa de servicio.
	authService := service.NewAuthService(usuarioRepo)
	parqueaderoService := sp.NewParqueaderoService(almacen)
	espacioService := sp.NewEspacioService(almacen)
	ocupacionService := sp.NewOcupacionService(almacen)

	// 5. Server con los servicios inyectados.
	servidor := handlers.NewServer(parqueaderoService, espacioService, ocupacionService, authService)

	// 6. Router + middleware global.
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	// 7. Rutas versionadas /api/v1/.
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
