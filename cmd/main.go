package main

import (
	"log"
	"net/http"

	"proyecto_movilidad_fcvt/internal/handlers"
	"proyecto_movilidad_fcvt/internal/storage"

	"github.com/go-chi/chi/v5"
)

func main() {
	log.Println("Inicializando servidor...")
	memoria := storage.NuevaMemoria()

	memoria.SeedParqueaderos()
	memoria.SeedEspacios()
	memoria.SeedOcupaciones()

	log.Println("Datos de prueba cargados")

	r := chi.NewRouter()

	parqueaderoHandler := handlers.NewParqueaderoHandler(memoria)
	espacioHandler := handlers.NewEspacioHandler(memoria)
	ocupacionHandler := handlers.NewOcupacionHandler(memoria)

	r.Route("/api/parqueaderos", func(r chi.Router) {
		r.Get("/", parqueaderoHandler.Listar)
		r.Get("/{id}", parqueaderoHandler.Obtener)
		r.Post("/", parqueaderoHandler.Crear)
		r.Put("/{id}", parqueaderoHandler.Actualizar)
		r.Delete("/{id}", parqueaderoHandler.Eliminar)
	})

	r.Route("/api/espacios", func(r chi.Router) {
		r.Get("/", espacioHandler.Listar)
		r.Get("/{id}", espacioHandler.Obtener)
		r.Post("/", espacioHandler.Crear)
		r.Put("/{id}", espacioHandler.Actualizar)
		r.Delete("/{id}", espacioHandler.Eliminar)
	})

	r.Route("/api/ocupaciones", func(r chi.Router) {
		r.Get("/", ocupacionHandler.Listar)
		r.Get("/{id}", ocupacionHandler.Obtener)
		r.Post("/", ocupacionHandler.Crear)
		r.Put("/{id}", ocupacionHandler.Actualizar)
		r.Delete("/{id}", ocupacionHandler.Eliminar)

		r.Patch("/{id}/liberar", ocupacionHandler.Liberar)
	})
	log.Println("Servidor escuchando en http://localhost:8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf(" Error al iniciar servidor: %v", err)
	}
}
