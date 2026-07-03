package handler_acceso

import (
	"proyecto_movilidad_fcvt/internal/service"
	"proyecto_movilidad_fcvt/internal/service/service_acceso" // <-- Tu subcarpeta de accesos
)

type Server struct {
	Auth        *service.AuthService
	Acceso      *service_acceso.AccesoService
	Usuario     *service_acceso.UsuarioService
	Vehiculo    *service_acceso.VehiculoService
	PuntoAcceso *service_acceso.PuntoAccesoService
}

func NewServer(
	auth *service.AuthService,
	acceso *service_acceso.AccesoService,
	usuario *service_acceso.UsuarioService,
	vehiculo *service_acceso.VehiculoService,
	puntoAcceso *service_acceso.PuntoAccesoService,
) *Server {
	return &Server{
		Auth:        auth,
		Acceso:      acceso,
		Usuario:     usuario,
		Vehiculo:    vehiculo,
		PuntoAcceso: puntoAcceso,
	}
}
