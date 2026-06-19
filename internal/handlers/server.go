package handlers

import (
	"proyecto_movilidad_fcvt/internal/service"
	sp "proyecto_movilidad_fcvt/internal/service/service_parqueadero"
)

type Server struct {
	Parqueadero *sp.ParqueaderoService
	Espacio     *sp.EspacioService
	Ocupacion   *sp.OcupacionService
	Auth        *service.AuthService
}

func NewServer(
	parqueadero *sp.ParqueaderoService,
	espacio *sp.EspacioService,
	ocupacion *sp.OcupacionService,
	auth *service.AuthService,
) *Server {
	return &Server{
		Parqueadero: parqueadero,
		Espacio:     espacio,
		Ocupacion:   ocupacion,
		Auth:        auth,
	}
}
