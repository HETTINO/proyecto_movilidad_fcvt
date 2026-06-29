package storage

import modelos "proyecto_movilidad_fcvt/internal/modelos"

// Almacen define QUÉ sabe hacer un almacén de transporte, sin decir CÓMO.
//
// Memoria (slices) ya cumple esta interfaz sin cambios — por el duck typing
// que vimos en S3 — y AlmacenSQLite (GORM) la cumple igual. El Server depende
// de esta interfaz, no de una implementación concreta: por eso podemos cambiar
// el backend de almacenamiento sin tocar un solo handler.
type Almacen interface {

	// Rutas
	ListarRutas() []modelos.Ruta
	BuscarRutaPorID(id int) (modelos.Ruta, bool)
	CrearRuta(r modelos.Ruta) modelos.Ruta
	ActualizarRuta(id int, datos modelos.Ruta) (modelos.Ruta, bool)
	BorrarRuta(id int) bool

	// Paradas
	ListarParadas() []modelos.Parada
	BuscarParadaPorID(id int) (modelos.Parada, bool)
	CrearParada(p modelos.Parada) modelos.Parada
	ActualizarParada(id int, datos modelos.Parada) (modelos.Parada, bool)
	BorrarParada(id int) bool

	// Carritos

	ListarCarritos() []modelos.Carrito
	BuscarCarritoPorID(id int) (modelos.Carrito, bool)
	CrearCarrito(c modelos.Carrito) modelos.Carrito
	ActualizarCarrito(id int, datos modelos.Carrito) (modelos.Carrito, bool)
	BorrarCarrito(id int) bool

	// Locaciones
	ListarLocaciones() []modelos.Locacion
	RegistrarLocacion(l modelos.Locacion) modelos.Locacion
	ObtenerUltimaLocacionPorCarrito(carritoID int) (modelos.Locacion, bool)

	// Solicitudes
	ListarSolicitudes() []modelos.Solicitud
	BuscarSolicitudPorID(id int) (modelos.Solicitud, bool)
	CrearSolicitud(s modelos.Solicitud) modelos.Solicitud
	ActualizarSolicitud(id int, datos modelos.Solicitud) (modelos.Solicitud, bool)
	BorrarSolicitud(id int) bool
}

// Chequeo en tiempo de compilación: si Memoria dejara de cumplir Almacen,
// el proyecto NO compila. Red de seguridad opcional.
//var _ Almacen = (*Memoria)(nil)
