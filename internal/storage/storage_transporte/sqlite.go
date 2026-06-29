package storage

import (
	"proyecto_movilidad_fcvt/internal/modelos"

	"gorm.io/gorm"
)

type AlmacenSQLite struct {
	db *gorm.DB
}

func NuevoAlmacenSQLite(db *gorm.DB) *AlmacenSQLite {
	return &AlmacenSQLite{db: db}
}

// =========================================================
// RUTAS
// =========================================================

func (a *AlmacenSQLite) ListarRutas() []modelos.Ruta {
	var rutas []modelos.Ruta
	a.db.Find(&rutas)
	return rutas
}

func (a *AlmacenSQLite) BuscarRutaPorID(id int) (modelos.Ruta, bool) {
	var r modelos.Ruta
	if err := a.db.First(&r, id).Error; err != nil {
		return modelos.Ruta{}, false
	}
	return r, true
}

func (a *AlmacenSQLite) CrearRuta(r modelos.Ruta) modelos.Ruta {
	a.db.Create(&r)
	return r
}

func (a *AlmacenSQLite) ActualizarRuta(id int, datos modelos.Ruta) (modelos.Ruta, bool) {
	var existente modelos.Ruta
	if err := a.db.First(&existente, id).Error; err != nil {
		return modelos.Ruta{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarRuta(id int) bool {
	res := a.db.Delete(&modelos.Ruta{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// CARRITOS
// =========================================================

func (a *AlmacenSQLite) ListarCarritos() []modelos.Carrito {
	var carritos []modelos.Carrito
	a.db.Find(&carritos)
	return carritos
}

func (a *AlmacenSQLite) ListarCarritosPorRuta(idRuta int) []modelos.Carrito {
	var carritos []modelos.Carrito
	a.db.Where("ruta_id = ?", idRuta).Find(&carritos)
	return carritos
}

func (a *AlmacenSQLite) BuscarCarritoPorID(id int) (modelos.Carrito, bool) {
	var c modelos.Carrito
	if err := a.db.First(&c, id).Error; err != nil {
		return modelos.Carrito{}, false
	}
	return c, true
}

func (a *AlmacenSQLite) CrearCarrito(c modelos.Carrito) modelos.Carrito {
	a.db.Create(&c)
	return c
}

func (a *AlmacenSQLite) ActualizarCarrito(id int, datos modelos.Carrito) (modelos.Carrito, bool) {
	var existente modelos.Carrito
	if err := a.db.First(&existente, id).Error; err != nil {
		return modelos.Carrito{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarCarrito(id int) bool {
	res := a.db.Delete(&modelos.Carrito{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// PARADAS
// =========================================================

func (a *AlmacenSQLite) ListarParadas() []modelos.Parada {
	var paradas []modelos.Parada
	a.db.Find(&paradas)
	return paradas
}

func (a *AlmacenSQLite) BuscarParadaPorID(id int) (modelos.Parada, bool) {
	var p modelos.Parada
	if err := a.db.First(&p, id).Error; err != nil {
		return modelos.Parada{}, false
	}
	return p, true
}

func (a *AlmacenSQLite) CrearParada(p modelos.Parada) modelos.Parada {
	a.db.Create(&p)
	return p
}

func (a *AlmacenSQLite) ActualizarParada(id int, datos modelos.Parada) (modelos.Parada, bool) {
	var existente modelos.Parada
	if err := a.db.First(&existente, id).Error; err != nil {
		return modelos.Parada{}, false
	}
	datos.IDParada = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarParada(id int) bool {
	res := a.db.Delete(&modelos.Parada{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// LOCACIONES
// =========================================================

func (a *AlmacenSQLite) ListarLocaciones() []modelos.Locacion {
	var locaciones []modelos.Locacion
	a.db.Find(&locaciones)
	return locaciones
}

func (a *AlmacenSQLite) ObtenerUltimaLocacionPorCarrito(carritoID int) (modelos.Locacion, bool) {
	var l modelos.Locacion
	if err := a.db.Where("carrito_id = ?", carritoID).Order("time_stamp DESC").First(&l).Error; err != nil {
		return modelos.Locacion{}, false
	}
	return l, true
}

func (a *AlmacenSQLite) RegistrarLocacion(l modelos.Locacion) modelos.Locacion {
	a.db.Create(&l)
	return l
}

// =========================================================
// SOLICITUDES
// =========================================================

func (a *AlmacenSQLite) ListarSolicitudes() []modelos.Solicitud {
	var solicitudes []modelos.Solicitud
	a.db.Find(&solicitudes)
	return solicitudes
}

func (a *AlmacenSQLite) BuscarSolicitudPorID(id int) (modelos.Solicitud, bool) {
	var s modelos.Solicitud
	if err := a.db.First(&s, id).Error; err != nil {
		return modelos.Solicitud{}, false
	}
	return s, true
}

func (a *AlmacenSQLite) CrearSolicitud(s modelos.Solicitud) modelos.Solicitud {
	a.db.Create(&s)
	return s
}

func (a *AlmacenSQLite) ActualizarSolicitud(id int, datos modelos.Solicitud) (modelos.Solicitud, bool) {
	var existente modelos.Solicitud
	if err := a.db.First(&existente, id).Error; err != nil {
		return modelos.Solicitud{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarSolicitud(id int) bool {
	res := a.db.Delete(&modelos.Solicitud{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// SEEDS
// =========================================================

func (a *AlmacenSQLite) SembrarSiVacio() {
	// Verificar si ya hay datos
	var n int64
	a.db.Model(&modelos.Ruta{}).Count(&n)
	if n > 0 {
		return
	}

	// Crear rutas
	rutas := []modelos.Ruta{
		{ID: 1, Nombre: "Ruta Centro", Descripcion: "Recorrido por el centro de la ciudad"},
		{ID: 2, Nombre: "Ruta Norte", Descripcion: "Recorrido por la zona norte"},
		{ID: 3, Nombre: "Ruta Sur", Descripcion: "Recorrido por la zona sur"},
	}
	a.db.Create(&rutas)

	// Crear carritos
	carritos := []modelos.Carrito{
		{ID: 1, NombreCarrito: "Carrito A", Capacidad: 20, Estado: "disponible", RutaID: 1},
		{ID: 2, NombreCarrito: "Carrito B", Capacidad: 15, Estado: "disponible", RutaID: 1},
		{ID: 3, NombreCarrito: "Carrito C", Capacidad: 25, Estado: "disponible", RutaID: 2},
	}
	a.db.Create(&carritos)

	// Crear paradas
	paradas := []modelos.Parada{
		{IDParada: 1, Nombre: "Parada Central", Latitud: -0.9281, Longitud: -78.6245},
		{IDParada: 2, Nombre: "Parada Norte", Latitud: -0.8950, Longitud: -78.6200},
		{IDParada: 3, Nombre: "Parada Sur", Latitud: -0.9500, Longitud: -78.6300},
	}
	a.db.Create(&paradas)

	// Crear solicitudes
	solicitudes := []modelos.Solicitud{
		{ID: 1, CedulaUsuario: "1234567890", CantPersonas: 3, ParadaOrigen: 1, PuntoDestino: "Biblioteca", Estado: "pendiente"},
		{ID: 2, CedulaUsuario: "0987654321", CantPersonas: 2, ParadaOrigen: 2, PuntoDestino: "Hospital", Estado: "asignada"},
	}
	a.db.Create(&solicitudes)
}

var _ Almacen = (*AlmacenSQLite)(nil)
