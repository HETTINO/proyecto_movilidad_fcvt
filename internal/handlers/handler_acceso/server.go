package handler_acceso

import (
	"proyecto_movilidad_fcvt/internal/service"
	"proyecto_movilidad_fcvt/internal/service/service_acceso"
)

type Server struct {
	Auth        *service.AuthService
	Acceso      *service_acceso.AccesoService
	Usuario     *service_acceso.UsuarioService
	Vehiculo    *service_acceso.VehiculoService
	PuntoAcceso *service_acceso.PuntoAccesoService
}

// Deps agrupa las dependencias OBLIGATORIAS de este servidor.
type Deps struct {
	Auth        *service.AuthService
	Acceso      *service_acceso.AccesoService
	Usuario     *service_acceso.UsuarioService
	Vehiculo    *service_acceso.VehiculoService
	PuntoAcceso *service_acceso.PuntoAccesoService
}

func NewServer(d Deps) *Server {
	return &Server{
		Auth:        d.Auth,
		Acceso:      d.Acceso,
		Usuario:     d.Usuario,
		Vehiculo:    d.Vehiculo,
		PuntoAcceso: d.PuntoAcceso,
	}
}
