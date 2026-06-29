package storage

import (
	"proyecto_movilidad_fcvt/internal/models"
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

func (a *AlmacenSQLite) ListarRutas() []models.Ruta {
	var rutas []models.Ruta
	a.db.Find(&rutas)
	return rutas
}

func (a *AlmacenSQLite) BuscarRutaPorID(id int) (models.Ruta, bool) {
	var r models.Ruta
	if err := a.db.First(&r, id).Error; err != nil {
		return models.Ruta{}, false
	}
	return r, true
}

func (a *AlmacenSQLite) CrearRuta(r models.Ruta) models.Ruta {
	a.db.Create(&r)
	return r
}

func (a *AlmacenSQLite) ActualizarRuta(id int, datos models.Ruta) (models.Ruta, bool) {
	var existente models.Ruta
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Ruta{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarRuta(id int) bool {
	res := a.db.Delete(&models.Ruta{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// CARRITOS
// =========================================================

func (a *AlmacenSQLite) ListarCarritos() []models.Carrito {
	var carritos []models.Carrito
	a.db.Find(&carritos)
	return carritos
}

func (a *AlmacenSQLite) ListarCarritosPorRuta(idRuta int) []models.Carrito {
	var carritos []models.Carrito
	a.db.Where("ruta_id = ?", idRuta).Find(&carritos)
	return carritos
}

func (a *AlmacenSQLite) BuscarCarritoPorID(id int) (models.Carrito, bool) {
	var c models.Carrito
	if err := a.db.First(&c, id).Error; err != nil {
		return models.Carrito{}, false
	}
	return c, true
}

func (a *AlmacenSQLite) CrearCarrito(c models.Carrito) models.Carrito {
	a.db.Create(&c)
	return c
}

func (a *AlmacenSQLite) ActualizarCarrito(id int, datos models.Carrito) (models.Carrito, bool) {
	var existente models.Carrito
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Carrito{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarCarrito(id int) bool {
	res := a.db.Delete(&models.Carrito{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// PARADAS
// =========================================================

func (a *AlmacenSQLite) ListarParadas() []models.Parada {
	var paradas []models.Parada
	a.db.Find(&paradas)
	return paradas
}

func (a *AlmacenSQLite) BuscarParadaPorID(id int) (models.Parada, bool) {
	var p models.Parada
	if err := a.db.First(&p, id).Error; err != nil {
		return models.Parada{}, false
	}
	return p, true
}

func (a *AlmacenSQLite) CrearParada(p models.Parada) models.Parada {
	a.db.Create(&p)
	return p
}

func (a *AlmacenSQLite) ActualizarParada(id int, datos models.Parada) (models.Parada, bool) {
	var existente models.Parada
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Parada{}, false
	}
	datos.IDParada = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarParada(id int) bool {
	res := a.db.Delete(&models.Parada{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// LOCACIONES
// =========================================================

func (a *AlmacenSQLite) ListarLocaciones() []models.Locacion {
	var locaciones []models.Locacion
	a.db.Find(&locaciones)
	return locaciones
}

func (a *AlmacenSQLite) ObtenerUltimaLocacionPorCarrito(carritoID int) (models.Locacion, bool) {
	var l models.Locacion
	if err := a.db.Where("carrito_id = ?", carritoID).Order("time_stamp DESC").First(&l).Error; err != nil {
		return models.Locacion{}, false
	}
	return l, true
}

func (a *AlmacenSQLite) RegistrarLocacion(l models.Locacion) models.Locacion {
	a.db.Create(&l)
	return l
}

// =========================================================
// SOLICITUDES
// =========================================================

func (a *AlmacenSQLite) ListarSolicitudes() []models.Solicitud {
	var solicitudes []models.Solicitud
	a.db.Find(&solicitudes)
	return solicitudes
}

func (a *AlmacenSQLite) BuscarSolicitudPorID(id int) (models.Solicitud, bool) {
	var s models.Solicitud
	if err := a.db.First(&s, id).Error; err != nil {
		return models.Solicitud{}, false
	}
	return s, true
}

func (a *AlmacenSQLite) CrearSolicitud(s models.Solicitud) models.Solicitud {
	a.db.Create(&s)
	return s
}

func (a *AlmacenSQLite) ActualizarSolicitud(id int, datos models.Solicitud) (models.Solicitud, bool) {
	var existente models.Solicitud
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Solicitud{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarSolicitud(id int) bool {
	res := a.db.Delete(&models.Solicitud{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// SEEDS
// =========================================================

func (a *AlmacenSQLite) SembrarSiVacio() {
	// Verificar si ya hay datos
	var n int64
	a.db.Model(&models.Ruta{}).Count(&n)
	if n > 0 {
		return
	}

	// Crear rutas
	rutas := []models.Ruta{
		{ID: 1, Nombre: "Ruta Centro", Descripcion: "Recorrido por el centro de la ciudad"},
		{ID: 2, Nombre: "Ruta Norte", Descripcion: "Recorrido por la zona norte"},
		{ID: 3, Nombre: "Ruta Sur", Descripcion: "Recorrido por la zona sur"},
	}
	a.db.Create(&rutas)

	// Crear carritos
	carritos := []models.Carrito{
		{ID: 1, NombreCarrito: "Carrito A", Capacidad: 20, Estado: "disponible", RutaID: 1},
		{ID: 2, NombreCarrito: "Carrito B", Capacidad: 15, Estado: "disponible", RutaID: 1},
		{ID: 3, NombreCarrito: "Carrito C", Capacidad: 25, Estado: "disponible", RutaID: 2},
	}
	a.db.Create(&carritos)

	// Crear paradas
	paradas := []models.Parada{
		{IDParada: 1, Nombre: "Parada Central", Latitud: -0.9281, Longitud: -78.6245},
		{IDParada: 2, Nombre: "Parada Norte", Latitud: -0.8950, Longitud: -78.6200},
		{IDParada: 3, Nombre: "Parada Sur", Latitud: -0.9500, Longitud: -78.6300},
	}
	a.db.Create(&paradas)

	// Crear solicitudes
	solicitudes := []models.Solicitud{
		{ID: 1, CedulaUsuario: "1234567890", CantPersonas: 3, ParadaOrigen: 1, PuntoDestino: "Biblioteca", Estado: "pendiente"},
		{ID: 2, CedulaUsuario: "0987654321", CantPersonas: 2, ParadaOrigen: 2, PuntoDestino: "Hospital", Estado: "asignada"},
	}
	a.db.Create(&solicitudes)
}

var _ Almacen = (*AlmacenSQLite)(nil)