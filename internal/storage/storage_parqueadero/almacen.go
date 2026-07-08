package storage_parqueadero

import "proyecto_movilidad_fcvt/internal/modelos"

type ParqueaderoRepository interface {
	// Parqueaderos
	ListarParqueaderos() []modelos.Parqueadero
	BuscarParqueaderoPorID(id int) (modelos.Parqueadero, bool)
	CrearParqueadero(p modelos.Parqueadero) (modelos.Parqueadero, error)
	ActualizarParqueadero(id int, datos modelos.Parqueadero) (modelos.Parqueadero, bool)
	BorrarParqueadero(id int) bool
}

type EspacioRepository interface {
	ListarEspacios() []modelos.Espacio
	BuscarEspacioPorID(id int) (modelos.Espacio, bool)
	CrearEspacio(e modelos.Espacio) modelos.Espacio
	ActualizarEspacio(id int, datos modelos.Espacio) (modelos.Espacio, bool)
	BorrarEspacio(id int) bool
}

type OcupacionesRepository interface {
	ListarOcupaciones() []modelos.Ocupacion
	BuscarOcupacionPorID(id int) (modelos.Ocupacion, bool)
	CrearOcupacion(o modelos.Ocupacion) modelos.Ocupacion
	ActualizarOcupacion(id int, datos modelos.Ocupacion) (modelos.Ocupacion, bool)
	BorrarOcupacion(id int) bool
	LiberarOcupacion(id int) (modelos.Ocupacion, bool)
	// ListarOcupacionesActivas devuelve las ocupaciones de un espacio que
	// aún no tienen HoraFin (es decir, siguen "activas"/en curso).
	ListarOcupacionesActivas(idEspacio int) []modelos.Ocupacion
}
type Almacen interface {
	ParqueaderoRepository
	EspacioRepository
	OcupacionesRepository
}

var _ Almacen = (*Memoria)(nil)
