package handlers

import (
	sp "proyecto_movilidad_fcvt/internal/service/service_transporte"
)

type Server struct {
	Ruta      *sp.RutaService
	Carrito   *sp.CarritoService
	Parada    *sp.ParadaService
	Locacion  *sp.LocacionService
	Solicitud *sp.SolicitudService
}

func NewServer(
	ruta *sp.RutaService,
	carrito *sp.CarritoService,
	parada *sp.ParadaService,
	locacion *sp.LocacionService,
	solicitud *sp.SolicitudService,

) *Server {
	return &Server{
		Ruta:      ruta,
		Carrito:   carrito,
		Parada:    parada,
		Locacion:  locacion,
		Solicitud: solicitud,
	}
}
