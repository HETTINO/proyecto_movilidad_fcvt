package handlers_test

import (
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"

	mw "proyecto_movilidad_fcvt/internal/middleware"
	st "proyecto_movilidad_fcvt/internal/service/service_transporte"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_transporte"
	ht "proyecto_movilidad_fcvt/internal/handlers/handler_transporte"


)

// =====================================================
// Construcción del entorno
// =====================================================

func construirEntorno(t *testing.T) http.Handler {
	t.Helper()

	mem := storage.NuevaMemoria()

	mem.SeedRutas()
	mem.SeedParadas()
	mem.SeedCarritos()
	mem.SeedLocaciones()
	mem.SeedSolicitudes()

	rutaSvc := st.NewRutaService(mem)
	carritoSvc := st.NewCarritoService(mem)
	paradaSvc := st.NewParadaService(mem)
	locacionSvc := st.NewLocacionService(mem)
	solicitudSvc := st.NewSolicitudService(mem)

  	srv := ht.NewServer(ht.Deps{
		Ruta:      rutaSvc,
		Carrito:   carritoSvc,
		Parada:    paradaSvc,
		Locacion:  locacionSvc,
		Solicitud: solicitudSvc,
	})

	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {

		r.Get("/rutas", srv.ListarRutas)
		r.Post("/rutas", srv.CrearRuta)
		r.Get("/rutas/{id}", srv.ObtenerRuta)
		r.Put("/rutas/{id}", srv.ActualizarRuta)
		r.Delete("/rutas/{id}", srv.BorrarRuta)

		r.Get("/paradas", srv.ListarParadas)
		r.Post("/paradas", srv.CrearParada)
		r.Get("/paradas/{id}", srv.ObtenerParada)
		r.Put("/paradas/{id}", srv.ActualizarParada)
		r.Delete("/paradas/{id}", srv.BorrarParada)

		r.Get("/carritos", srv.ListarCarritos)
		r.Post("/carritos", srv.CrearCarrito)
		r.Get("/carritos/{id}", srv.ObtenerCarrito)
		r.Put("/carritos/{id}", srv.ActualizarCarrito)
		r.Delete("/carritos/{id}", srv.BorrarCarrito)

		r.Get("/locaciones", srv.ListarLocaciones)
		r.Post("/locaciones", srv.RegistrarLocacion)
		r.Get("/locaciones/carrito/{id}", srv.ObtenerUbicacionCarrito)
		r.Get("/tiempo-estimado", srv.GetTiempoEstimado)

		r.Get("/solicitudes", srv.ListarSolicitudes)
		r.With(mw.RequireAuth).Post("/solicitudes", srv.CrearSolicitud)
		r.Get("/solicitudes/{id}", srv.ObtenerSolicitud)
		r.With(mw.RequireAuth).Put("/solicitudes/{id}", srv.ActualizarSolicitud)
		r.With(mw.RequireAuth).Delete("/solicitudes/{id}", srv.BorrarSolicitud)

	})

	return r
}
