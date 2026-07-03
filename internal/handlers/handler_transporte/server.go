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

// Deps agrupa las dependencias OBLIGATORIAS de este servidor.
type Deps struct {
	Ruta      *sp.RutaService
	Carrito   *sp.CarritoService
	Parada    *sp.ParadaService
	Locacion  *sp.LocacionService
	Solicitud *sp.SolicitudService
}

func NewServer(d Deps) *Server {
	return &Server{
		Ruta:      d.Ruta,
		Carrito:   d.Carrito,
		Parada:    d.Parada,
		Locacion:  d.Locacion,
		Solicitud: d.Solicitud,
	}
}
