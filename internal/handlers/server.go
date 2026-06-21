package handlers

import (
	"proyecto_movilidad_fcvt/internal/service"
	// sp "proyecto_movilidad_fcvt/internal/service/service_parqueadero" // <-- COMENTADO: Tu rama no tiene esta carpeta
)

type Server struct {
	// Parqueadero *sp.ParqueaderoService // <-- COMENTADO
	// Espacio     *sp.EspacioService     // <-- COMENTADO
	// Ocupacion   *sp.OcupacionService   // <-- COMENTADO
	Auth   *service.AuthService
	Acceso *service.AccesoServicio
}

func NewServer(
	// parqueadero *sp.ParqueaderoService, // <-- COMENTADO
	// espacio *sp.EspacioService,         // <-- COMENTADO
	// ocupacion *sp.OcupacionService,     // <-- COMENTADO
	auth *service.AuthService,
	acceso *service.AccesoServicio,
) *Server {
	return &Server{
		// Parqueadero: parqueadero,
		// Espacio:     espacio,
		// Ocupacion:   ocupacion,
		Auth:   auth,
		Acceso: acceso,
	}
}
