package storage_acceso

import "proyecto_movilidad_fcvt/internal/modelos"

type MemoriaAcceso struct {
	Usuarios       []modelos.Usuario
	Vehiculos      []modelos.Vehiculo
	PuntosDeAcceso []modelos.PuntoDeAcceso
	Accesos        []modelos.Acceso

	nextIDPuntoAcceso int
}

func NuevoMemoriaAcceso() *MemoriaAcceso {
	return &MemoriaAcceso{
		Usuarios:          make([]modelos.Usuario, 0),
		Vehiculos:         make([]modelos.Vehiculo, 0),
		PuntosDeAcceso:    make([]modelos.PuntoDeAcceso, 0),
		Accesos:           make([]modelos.Acceso, 0),
		nextIDPuntoAcceso: 1,
	}
}

// ============================================================================
// CONTRATOS (INTERFACES) QUE SE ENLAZAN CON TUS SERVICIOS
//
// Implementaciones:
//   - AccesoRepository      -> CRUD_acceso.go
//   - PuntoAccesoRepository -> CRUD_punto_acceso.go
//   - UsuarioRepository     -> CRUD_usuario.go
//   - VehiculoRepository    -> CRUD_vehiculo.go
// ============================================================================

type AccesoRepository interface {
	ListarAccesos() []modelos.Acceso
	BuscarAccesoPorID(id int) (modelos.Acceso, bool)
	CrearAcceso(a modelos.Acceso) modelos.Acceso
	ActualizarAcceso(id int, a modelos.Acceso) (modelos.Acceso, bool)
	BorrarAcceso(id int) bool
}

type PuntoAccesoRepository interface {
	ListarPuntosAcceso() []modelos.PuntoDeAcceso
	BuscarPuntoAccesoPorID(id int) (modelos.PuntoDeAcceso, bool)
	CrearPuntoAcceso(p modelos.PuntoDeAcceso) modelos.PuntoDeAcceso
	ActualizarPuntoAcceso(id int, p modelos.PuntoDeAcceso) (modelos.PuntoDeAcceso, bool)
	BorrarPuntoAcceso(id int) bool
}

type UsuarioRepository interface {
	ListarUsuarios() []modelos.Usuario
	BuscarUsuarioPorCedula(cedula string) (modelos.Usuario, bool)
	CrearUsuario(u modelos.Usuario) modelos.Usuario
	ActualizarUsuario(cedula string, u modelos.Usuario) (modelos.Usuario, bool)
	BorrarUsuario(cedula string) bool
}

type VehiculoRepository interface {
	ListarVehiculos() []modelos.Vehiculo
	BuscarVehiculoPorPlaca(placa string) (modelos.Vehiculo, bool)
	CrearVehiculo(v modelos.Vehiculo) modelos.Vehiculo
	ActualizarVehiculo(placa string, v modelos.Vehiculo) (modelos.Vehiculo, bool)
	BorrarVehiculo(placa string) bool
}
