package storage_acceso

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
// USUARIOS
// =========================================================

func (a *AlmacenSQLite) ListarUsuarios() []modelos.Usuario {
	var usuarios []modelos.Usuario
	a.db.Find(&usuarios)
	return usuarios
}

func (a *AlmacenSQLite) BuscarUsuarioPorCedula(cedula string) (modelos.Usuario, bool) {
	var u modelos.Usuario
	if err := a.db.First(&u, "cedula = ?", cedula).Error; err != nil {
		return modelos.Usuario{}, false
	}
	return u, true
}

func (a *AlmacenSQLite) CrearUsuario(u modelos.Usuario) modelos.Usuario {
	a.db.Create(&u)
	return u
}

func (a *AlmacenSQLite) ActualizarUsuario(cedula string, datos modelos.Usuario) (modelos.Usuario, bool) {
	var existente modelos.Usuario
	if err := a.db.First(&existente, "cedula = ?", cedula).Error; err != nil {
		return modelos.Usuario{}, false
	}

	datos.Cedula = cedula
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarUsuario(cedula string) bool {
	res := a.db.Delete(&modelos.Usuario{}, "cedula = ?", cedula)
	return res.RowsAffected > 0
}

// =========================================================
// VEHICULOS
// =========================================================

func (a *AlmacenSQLite) ListarVehiculos() []modelos.Vehiculo {
	var vehiculos []modelos.Vehiculo
	a.db.Find(&vehiculos)
	return vehiculos
}

func (a *AlmacenSQLite) BuscarVehiculoPorPlaca(placa string) (modelos.Vehiculo, bool) {
	var v modelos.Vehiculo
	if err := a.db.First(&v, "placa = ?", placa).Error; err != nil {
		return modelos.Vehiculo{}, false
	}
	return v, true
}

func (a *AlmacenSQLite) CrearVehiculo(v modelos.Vehiculo) modelos.Vehiculo {
	a.db.Create(&v)
	return v
}

func (a *AlmacenSQLite) ActualizarVehiculo(placa string, datos modelos.Vehiculo) (modelos.Vehiculo, bool) {
	var existente modelos.Vehiculo
	if err := a.db.First(&existente, "placa = ?", placa).Error; err != nil {
		return modelos.Vehiculo{}, false
	}

	datos.Placa = placa
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarVehiculo(placa string) bool {
	res := a.db.Delete(&modelos.Vehiculo{}, "placa = ?", placa)
	return res.RowsAffected > 0
}

// =========================================================
// PUNTOS DE ACCESO
// =========================================================

func (a *AlmacenSQLite) ListarPuntosAcceso() []modelos.PuntoDeAcceso {
	var puntos []modelos.PuntoDeAcceso
	a.db.Find(&puntos)
	return puntos
}

func (a *AlmacenSQLite) BuscarPuntoAccesoPorID(id int) (modelos.PuntoDeAcceso, bool) {
	var p modelos.PuntoDeAcceso
	if err := a.db.First(&p, id).Error; err != nil {
		return modelos.PuntoDeAcceso{}, false
	}
	return p, true
}

func (a *AlmacenSQLite) CrearPuntoAcceso(p modelos.PuntoDeAcceso) modelos.PuntoDeAcceso {
	a.db.Create(&p)
	return p
}

func (a *AlmacenSQLite) ActualizarPuntoAcceso(id int, datos modelos.PuntoDeAcceso) (modelos.PuntoDeAcceso, bool) {
	var existente modelos.PuntoDeAcceso
	if err := a.db.First(&existente, id).Error; err != nil {
		return modelos.PuntoDeAcceso{}, false
	}

	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarPuntoAcceso(id int) bool {
	res := a.db.Delete(&modelos.PuntoDeAcceso{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// ACCESOS
// =========================================================

func (a *AlmacenSQLite) ListarAccesos() []modelos.Acceso {
	var accesos []modelos.Acceso
	a.db.Find(&accesos)
	return accesos
}

func (a *AlmacenSQLite) BuscarAccesoPorID(id int) (modelos.Acceso, bool) {
	var ac modelos.Acceso
	if err := a.db.First(&ac, id).Error; err != nil {
		return modelos.Acceso{}, false
	}
	return ac, true
}

func (a *AlmacenSQLite) CrearAcceso(ac modelos.Acceso) modelos.Acceso {
	a.db.Create(&ac)
	return ac
}

func (a *AlmacenSQLite) ActualizarAcceso(id int, datos modelos.Acceso) (modelos.Acceso, bool) {
	var existente modelos.Acceso
	if err := a.db.First(&existente, id).Error; err != nil {
		return modelos.Acceso{}, false
	}

	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarAcceso(id int) bool {
	res := a.db.Delete(&modelos.Acceso{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// SEEDS
// =========================================================

func (a *AlmacenSQLite) SembrarSiVacio() {
	var n int64
	a.db.Model(&modelos.Usuario{}).Count(&n)
	if n > 0 {
		return
	}

	usuarios := []modelos.Usuario{
		{
			Cedula: "131555",
			Nombre: "Shirley Juleidy",
			Email:  "shirley@example.com",
			Rol:    "admin",
		},
	}

	a.db.Create(&usuarios)
}

// interfaces
var _ AccesoRepository = (*AlmacenSQLite)(nil)
var _ UsuarioRepository = (*AlmacenSQLite)(nil)
var _ VehiculoRepository = (*AlmacenSQLite)(nil)
var _ PuntoAccesoRepository = (*AlmacenSQLite)(nil)
