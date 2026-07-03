package handler_parqueadero

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

// Deps agrupa las dependencias OBLIGATORIAS de este servidor.
// Agregar una entidad nueva = un campo aquí, no un parámetro más.
type Deps struct {
	Parqueadero *sp.ParqueaderoService
	Espacio     *sp.EspacioService
	Ocupacion   *sp.OcupacionService
	Auth        *service.AuthService
}

func NewServer(d Deps) *Server {
	return &Server{
		Parqueadero: d.Parqueadero,
		Espacio:     d.Espacio,
		Ocupacion:   d.Ocupacion,
		Auth:        d.Auth,
	}
}
