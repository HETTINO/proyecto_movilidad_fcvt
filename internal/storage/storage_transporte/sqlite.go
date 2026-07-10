package storage

import (
	"proyecto_movilidad_fcvt/internal/modelos"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	a.db.Omit(clause.Associations).Create(&r)
	return r
}

func (a *AlmacenSQLite) ActualizarRuta(id int, datos modelos.Ruta) (modelos.Ruta, bool) {
	var existente modelos.Ruta
	if err := a.db.First(&existente, id).Error; err != nil {
		return modelos.Ruta{}, false
	}
	datos.ID = id
	a.db.Omit(clause.Associations).Save(&datos)
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
	a.db.Preload("Ruta").Find(&carritos)
	return carritos
}

func (a *AlmacenSQLite) ListarCarritosPorRuta(idRuta int) []modelos.Carrito {
	var carritos []modelos.Carrito
	a.db.Preload("Ruta").Where("ruta_id = ?", idRuta).Find(&carritos)
	return carritos
}

func (a *AlmacenSQLite) BuscarCarritoPorID(id int) (modelos.Carrito, bool) {
	var c modelos.Carrito
	if err := a.db.Preload("Ruta").First(&c, id).Error; err != nil {
		return modelos.Carrito{}, false
	}
	return c, true
}

func (a *AlmacenSQLite) CrearCarrito(c modelos.Carrito) modelos.Carrito {
	a.db.Omit(clause.Associations).Create(&c)
	return c
}

func (a *AlmacenSQLite) ActualizarCarrito(id int, datos modelos.Carrito) (modelos.Carrito, bool) {
	var existente modelos.Carrito
	if err := a.db.First(&existente, id).Error; err != nil {
		return modelos.Carrito{}, false
	}
	datos.ID = id
	a.db.Omit(clause.Associations).Save(&datos)
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
	a.db.Preload("Ruta").Find(&paradas)
	return paradas
}

func (a *AlmacenSQLite) BuscarParadaPorID(id int) (modelos.Parada, bool) {
	var p modelos.Parada
	if err := a.db.Preload("Ruta").First(&p, id).Error; err != nil {
		return modelos.Parada{}, false
	}
	return p, true
}

func (a *AlmacenSQLite) CrearParada(p modelos.Parada) modelos.Parada {
	a.db.Omit(clause.Associations).Create(&p)
	return p
}

func (a *AlmacenSQLite) ActualizarParada(id int, datos modelos.Parada) (modelos.Parada, bool) {
	var existente modelos.Parada
	if err := a.db.First(&existente, id).Error; err != nil {
		return modelos.Parada{}, false
	}
	datos.IDParada = id
	a.db.Omit(clause.Associations).Save(&datos)
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
	a.db.Preload("Carrito").Find(&locaciones)
	return locaciones
}

func (a *AlmacenSQLite) ObtenerUltimaLocacionPorCarrito(carritoID int) (modelos.Locacion, bool) {
	var l modelos.Locacion
	if err := a.db.Preload("Carrito").Where("carrito_id = ?", carritoID).Order("time_stamp DESC").First(&l).Error; err != nil {
		return modelos.Locacion{}, false
	}
	return l, true
}

func (a *AlmacenSQLite) RegistrarLocacion(l modelos.Locacion) modelos.Locacion {
	a.db.Omit(clause.Associations).Create(&l)
	return l
}

// =========================================================
// SOLICITUDES
// =========================================================

func (a *AlmacenSQLite) ListarSolicitudes() []modelos.Solicitud {
	var solicitudes []modelos.Solicitud
	a.db.Preload("Parada.Ruta").Preload("Carrito.Ruta").Find(&solicitudes)
	return solicitudes
}

func (a *AlmacenSQLite) BuscarSolicitudPorID(id int) (modelos.Solicitud, bool) {
	var s modelos.Solicitud
	if err := a.db.Preload("Parada.Ruta").Preload("Carrito.Ruta").First(&s, id).Error; err != nil {
		return modelos.Solicitud{}, false
	}
	return s, true
}

func (a *AlmacenSQLite) CrearSolicitud(s modelos.Solicitud) modelos.Solicitud {
	a.db.Omit(clause.Associations).Create(&s)
	return s
}

func (a *AlmacenSQLite) ActualizarSolicitud(id int, datos modelos.Solicitud) (modelos.Solicitud, bool) {
	var existente modelos.Solicitud
	if err := a.db.First(&existente, id).Error; err != nil {
		return modelos.Solicitud{}, false
	}
	datos.ID = id
	a.db.Omit(clause.Associations).Save(&datos)
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

	// Crear rutas (sin ID manual: dejamos que Postgres asigne el autoincremento)
	rutas := []modelos.Ruta{
		{Nombre: "Ruta Campus Norte", Descripcion: "Recorrido por la zona norte de la ULEAM: acceso principal y facultades"},
		{Nombre: "Ruta Biblioteca - Bienestar", Descripcion: "Conecta la biblioteca con el departamento de bienestar estudiantil"},
		{Nombre: "Ruta Comedor - Parqueaderos", Descripcion: "Recorrido entre el comedor universitario y las zonas de parqueo"},
	}
	a.db.Omit(clause.Associations).Create(&rutas)
	// Tras el Create, rutas[0].ID, rutas[1].ID, rutas[2].ID ya tienen los IDs reales asignados por Postgres.

	// Crear carritos, referenciando las rutas recién creadas
	carritos := []modelos.Carrito{
		{NombreCarrito: "Carrito 1 - Rectorado", Capacidad: 20, Estado: "disponible", RutaID: rutas[0].ID},
		{NombreCarrito: "Carrito 2 - FCVT", Capacidad: 15, Estado: "disponible", RutaID: rutas[0].ID},
		{NombreCarrito: "Carrito 3 - Biblioteca", Capacidad: 12, Estado: "disponible", RutaID: rutas[1].ID},
		{NombreCarrito: "Carrito 4 - Comedor", Capacidad: 25, Estado: "mantenimiento", RutaID: rutas[2].ID},
	}
	a.db.Omit(clause.Associations).Create(&carritos)

	// Crear paradas
	paradas := []modelos.Parada{
		{Nombre: "Parada Facultad de Ciencias Informáticas", Latitud: -0.9521, Longitud: -80.7485, RutaID: rutas[0].ID},
		{Nombre: "Parada Paraninfo Universitario", Latitud: -0.9505, Longitud: -80.7490, RutaID: rutas[0].ID},
		{Nombre: "Parada Biblioteca Central", Latitud: -0.9515, Longitud: -80.7495, RutaID: rutas[1].ID},
		{Nombre: "Parada Departamento de Bienestar", Latitud: -0.9525, Longitud: -80.7480, RutaID: rutas[1].ID},
		{Nombre: "Parada Plaza Centenaria", Latitud: -0.9535, Longitud: -80.7502, RutaID: rutas[2].ID},
		{Nombre: "Parada Comedor Tasty", Latitud: -0.9542, Longitud: -80.7515, RutaID: rutas[2].ID},
	}
	a.db.Omit(clause.Associations).Create(&paradas)

	// Crear solicitudes, referenciando las paradas recién creadas
	solicitudes := []modelos.Solicitud{
		{CedulaUsuario: "1234567890", CantPersonas: 3, ParadaOrigen: paradas[0].IDParada, PuntoDestino: "Biblioteca Central", Estado: "pendiente"},
		{CedulaUsuario: "0987654321", CantPersonas: 2, ParadaOrigen: paradas[2].IDParada, PuntoDestino: "Departamento de Bienestar", Estado: "asignada"},
		{CedulaUsuario: "1122334455", CantPersonas: 1, ParadaOrigen: paradas[4].IDParada, PuntoDestino: "Comedor Tasty", Estado: "completada"},
	}
	a.db.Omit(clause.Associations).Create(&solicitudes)
}

//var _ Almacen = (*AlmacenSQLite)(nil)