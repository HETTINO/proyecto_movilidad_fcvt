package storage_acceso

import "proyecto_movilidad_fcvt/internal/modelos"

type MemoriaAcceso struct {
	Usuarios       []modelos.Usuario
	Vehiculos      []modelos.Vehiculo
	PuntosDeAcceso []modelos.PuntoDeAcceso
	Accesos        []modelos.Acceso
}

// NuevoMemoriaAcceso inicializa los vectores en memoria RAM (para simulación)
func NuevoMemoriaAcceso() *MemoriaAcceso {
	return &MemoriaAcceso{
		Usuarios:       make([]modelos.Usuario, 0),
		Vehiculos:      make([]modelos.Vehiculo, 0),
		PuntosDeAcceso: make([]modelos.PuntoDeAcceso, 0),
		Accesos:        make([]modelos.Acceso, 0),
	}
}
