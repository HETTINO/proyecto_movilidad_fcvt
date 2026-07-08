package storage_parqueadero

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	"time"

	"gorm.io/gorm"
)

type AlmacenSQLite struct {
	db *gorm.DB
}

func NuevoAlmacenSQLite(db *gorm.DB) *AlmacenSQLite {
	return &AlmacenSQLite{db: db}
}

// =========================================================
// PARQUEADEROS
// =========================================================

func (a *AlmacenSQLite) ListarParqueaderos() []modelos.Parqueadero {
	var parqueaderos []modelos.Parqueadero
	a.db.Find(&parqueaderos)
	return parqueaderos
}

func (a *AlmacenSQLite) BuscarParqueaderoPorID(id int) (modelos.Parqueadero, bool) {
	var p modelos.Parqueadero
	if err := a.db.First(&p, id).Error; err != nil {
		return modelos.Parqueadero{}, false
	}
	return p, true
}

func (a *AlmacenSQLite) CrearParqueadero(p modelos.Parqueadero) (modelos.Parqueadero, error) {
	if err := a.db.Create(&p).Error; err != nil {
		return modelos.Parqueadero{}, err
	}
	return p, nil
}

func (a *AlmacenSQLite) ActualizarParqueadero(id int, datos modelos.Parqueadero) (modelos.Parqueadero, bool) {
	var existente modelos.Parqueadero
	if err := a.db.First(&existente, id).Error; err != nil {
		return modelos.Parqueadero{}, false
	}
	datos.IDParqueadero = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarParqueadero(id int) bool {
	res := a.db.Delete(&modelos.Parqueadero{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// ESPACIOS
// =========================================================

func (a *AlmacenSQLite) ListarEspacios() []modelos.Espacio {
	var espacios []modelos.Espacio
	a.db.Find(&espacios)
	return espacios
}

func (a *AlmacenSQLite) ListarEspaciosPorParqueadero(idParqueadero int) []modelos.Espacio {
	var espacios []modelos.Espacio
	a.db.Where("id_parqueadero = ?", idParqueadero).Find(&espacios)
	return espacios
}

func (a *AlmacenSQLite) BuscarEspacioPorID(id int) (modelos.Espacio, bool) {
	var e modelos.Espacio
	if err := a.db.First(&e, id).Error; err != nil {
		return modelos.Espacio{}, false
	}
	return e, true
}

func (a *AlmacenSQLite) CrearEspacio(e modelos.Espacio) modelos.Espacio {
	a.db.Create(&e)
	return e
}

func (a *AlmacenSQLite) ActualizarEspacio(id int, datos modelos.Espacio) (modelos.Espacio, bool) {
	var existente modelos.Espacio
	if err := a.db.First(&existente, id).Error; err != nil {
		return modelos.Espacio{}, false
	}
	datos.IDEspacio = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarEspacio(id int) bool {
	res := a.db.Delete(&modelos.Espacio{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// OCUPACIONES
// =========================================================

func (a *AlmacenSQLite) ListarOcupaciones() []modelos.Ocupacion {
	var ocupaciones []modelos.Ocupacion
	a.db.Find(&ocupaciones)
	return ocupaciones
}

func (a *AlmacenSQLite) BuscarOcupacionPorID(id int) (modelos.Ocupacion, bool) {
	var o modelos.Ocupacion
	if err := a.db.First(&o, id).Error; err != nil {
		return modelos.Ocupacion{}, false
	}
	return o, true
}

func (a *AlmacenSQLite) ListarOcupacionesActivas(idEspacio int) []modelos.Ocupacion {
	var ocupaciones []modelos.Ocupacion
	a.db.Where("id_espacio = ? AND hora_fin IS NULL", idEspacio).Find(&ocupaciones)
	return ocupaciones
}

func (a *AlmacenSQLite) CrearOcupacion(o modelos.Ocupacion) modelos.Ocupacion {
	a.db.Create(&o)
	return o
}

func (a *AlmacenSQLite) CerrarOcupacion(id int, datos modelos.Ocupacion) (modelos.Ocupacion, bool) {
	var existente modelos.Ocupacion
	if err := a.db.First(&existente, id).Error; err != nil {
		return modelos.Ocupacion{}, false
	}
	existente.HoraFin = datos.HoraFin
	a.db.Save(&existente)
	return existente, true
}

func (a *AlmacenSQLite) ActualizarOcupacion(id int, datos modelos.Ocupacion) (modelos.Ocupacion, bool) {
	var existente modelos.Ocupacion
	if err := a.db.First(&existente, id).Error; err != nil {
		return modelos.Ocupacion{}, false
	}
	datos.IDOcupacion = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarOcupacion(id int) bool {
	res := a.db.Delete(&modelos.Ocupacion{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// SEEDS
// =========================================================

func (a *AlmacenSQLite) SembrarSiVacio() {
	var n int64
	a.db.Model(&modelos.Parqueadero{}).Count(&n)
	if n > 0 {
		return
	}

	parqueaderos := []modelos.Parqueadero{
		{IDParqueadero: 1, Nombre: "Parqueadero Norte", Capacidad: 50, Tipo: "cubierto"},
		{IDParqueadero: 2, Nombre: "Parqueadero Sur", Capacidad: 30, Tipo: "abierto"},
	}
	a.db.Create(&parqueaderos)

	espacios := []modelos.Espacio{
		{IDEspacio: 1, IDParqueadero: 1, Numero: 1, Estado: "libre", TipoEspacio: "auto"},
		{IDEspacio: 2, IDParqueadero: 1, Numero: 2, Estado: "libre", TipoEspacio: "moto"},
		{IDEspacio: 3, IDParqueadero: 2, Numero: 1, Estado: "libre", TipoEspacio: "auto"},
	}
	a.db.Create(&espacios)
}

func (a *AlmacenSQLite) LiberarOcupacion(id int) (modelos.Ocupacion, bool) {
	var existente modelos.Ocupacion
	if err := a.db.First(&existente, id).Error; err != nil {
		return modelos.Ocupacion{}, false
	}
	ahora := time.Now()
	existente.HoraFin = &ahora
	a.db.Save(&existente)
	return existente, true
}

var _ Almacen = (*AlmacenSQLite)(nil)
