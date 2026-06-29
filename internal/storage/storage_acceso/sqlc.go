package storage_acceso

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"proyecto_movilidad_fcvt/internal/modelos"
	sqlcdb "proyecto_movilidad_fcvt/internal/storage/sqlcdb_acceso"
)

// =========================================================
// ESTRUCTURA
// =========================================================

type AlmacenSQLC struct {
	q *sqlcdb.Queries
}

func NuevoAlmacenSQLC(db *sql.DB) *AlmacenSQLC {
	return &AlmacenSQLC{q: sqlcdb.New(db)}
}

// =========================================================
// HELPERS
// =========================================================

func nullString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

func aNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

// =========================================================
// USUARIO MAP
// =========================================================

func aUsuario(u sqlcdb.Usuario) modelos.Usuario {
	return modelos.Usuario{
		Cedula:     fmt.Sprintf("%d", u.CedulaInt),
		Nombre:     u.NombreUsuario,
		Contrasena: u.Contrasena,
		Email:      u.Email,
		Rol:        u.Rol,
	}
}

// =========================================================
// VEHICULO MAP
// =========================================================

func aVehiculo(v sqlcdb.Vehiculo) modelos.Vehiculo {
	return modelos.Vehiculo{
		Placa:        v.PlacaVehiculo,
		IDUsuario:    fmt.Sprintf("%d", v.IDUsuario), // int64 -> string
		TipoVehiculo: v.TipoUsuario,                  // OJO: columna en BD se llama "Tipo_usuario", revisar si debería ser "Tipo_vehiculo"
		Marca:        v.Marca,
		Modelo:       v.Modelo,
		Color:        v.Color,
		Año:          int(v.Anio),
	}
}

// =========================================================
// PUNTO DE ACCESO MAP
// =========================================================

func aPuntoAcceso(p sqlcdb.PuntoDeAcceso) modelos.PuntoDeAcceso {
	return modelos.PuntoDeAcceso{
		ID:         int(p.IDPuntoacceso),
		Frecuencia: p.TipoAcceso, // OJO: la columna real es "Tipo_acceso", no existe "Frecuencia" en la BD
		Ubicacion:  p.Ubicacion,
	}
}

// =========================================================
// ACCESO MAP
// =========================================================

func aAcceso(a sqlcdb.Acceso) modelos.Acceso {

	var salida *time.Time
	if a.TiempoSalida.Valid {
		t := a.TiempoSalida.Time
		salida = &t
	}

	return modelos.Acceso{
		ID:            int(a.IDAcceso),
		PlacaVehiculo: a.PlacaVehiculo,
		PuntoAccesoID: int(a.IDPuntoacceso),
		TiempoEntrada: a.TiempoEntrada,
		TiempoSalida:  salida,
		Estado:        a.Estado,
		Observaciones: nullString(a.Observaciones),
	}
}

// =========================================================
// USUARIOS
// =========================================================

func (a *AlmacenSQLC) ListarUsuarios() []modelos.Usuario {
	filas, err := a.q.ListarUsuarios(context.Background())
	if err != nil {
		return nil
	}

	out := make([]modelos.Usuario, 0, len(filas))
	for _, f := range filas {
		out = append(out, aUsuario(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarUsuarioPorCedula(cedula string) (modelos.Usuario, bool) {
	c := 0
	fmt.Sscanf(cedula, "%d", &c)

	f, err := a.q.ObtenerUsuarioPorCedula(context.Background(), int64(c))
	if err != nil {
		return modelos.Usuario{}, false
	}

	return aUsuario(f), true
}

func (a *AlmacenSQLC) CrearUsuario(u modelos.Usuario) modelos.Usuario {
	c := 0
	fmt.Sscanf(u.Cedula, "%d", &c)

	f, err := a.q.RegistrarUsuario(context.Background(), sqlcdb.RegistrarUsuarioParams{
		CedulaInt:     int64(c),
		NombreUsuario: u.Nombre,
		Contrasena:    u.Contrasena,
		Email:         u.Email,
		Rol:           u.Rol,
	})
	if err != nil {
		return modelos.Usuario{}
	}
	return aUsuario(f)
}

// =========================================================
// VEHICULOS
// =========================================================

func (a *AlmacenSQLC) ListarVehiculos() []modelos.Vehiculo {
	filas, err := a.q.ListarVehiculos(context.Background())
	if err != nil {
		return nil
	}

	out := make([]modelos.Vehiculo, 0, len(filas))
	for _, f := range filas {
		out = append(out, aVehiculo(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarVehiculoPorPlaca(placa string) (modelos.Vehiculo, bool) {
	f, err := a.q.ObtenerVehiculoPorPlaca(context.Background(), placa)
	if err != nil {
		return modelos.Vehiculo{}, false
	}
	return aVehiculo(f), true
}

func (a *AlmacenSQLC) CrearVehiculo(v modelos.Vehiculo) modelos.Vehiculo {
	idUsuario := 0
	fmt.Sscanf(v.IDUsuario, "%d", &idUsuario)

	f, err := a.q.RegistrarVehiculo(context.Background(), sqlcdb.RegistrarVehiculoParams{
		PlacaVehiculo: v.Placa,
		IDUsuario:     int64(idUsuario),
		TipoUsuario:   v.TipoVehiculo,
		Marca:         v.Marca,
		Modelo:        v.Modelo,
		Color:         v.Color,
		Anio:          int64(v.Año),
	})
	if err != nil {
		return modelos.Vehiculo{}
	}
	return aVehiculo(f)
}

// =========================================================
// PUNTOS DE ACCESO
// =========================================================

func (a *AlmacenSQLC) ListarPuntosAcceso() []modelos.PuntoDeAcceso {
	filas, err := a.q.ListarPuntosAcceso(context.Background())
	if err != nil {
		return nil
	}

	out := make([]modelos.PuntoDeAcceso, 0, len(filas))
	for _, f := range filas {
		out = append(out, aPuntoAcceso(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarPuntoAccesoPorID(id int) (modelos.PuntoDeAcceso, bool) {
	f, err := a.q.ObtenerPuntoAccesoPorId(context.Background(), int64(id))
	if err != nil {
		return modelos.PuntoDeAcceso{}, false
	}
	return aPuntoAcceso(f), true
}

func (a *AlmacenSQLC) CrearPuntoAcceso(p modelos.PuntoDeAcceso) modelos.PuntoDeAcceso {
	f, err := a.q.CrearPuntoAcceso(context.Background(), sqlcdb.CrearPuntoAccesoParams{
		TipoAcceso: p.Frecuencia,
		Ubicacion:  p.Ubicacion,
	})
	if err != nil {
		return modelos.PuntoDeAcceso{}
	}
	return aPuntoAcceso(f)
}

// =========================================================
// ACCESOS
// =========================================================

func (a *AlmacenSQLC) ListarAccesos() []modelos.Acceso {
	filas, err := a.q.ListarAccesos(context.Background())
	if err != nil {
		return nil
	}

	out := make([]modelos.Acceso, 0, len(filas))
	for _, f := range filas {
		out = append(out, aAcceso(f))
	}
	return out
}

func (a *AlmacenSQLC) RegistrarEntrada(ac modelos.Acceso) modelos.Acceso {
	f, err := a.q.RegistrarAccesoEntrada(context.Background(), sqlcdb.RegistrarAccesoEntradaParams{
		PlacaVehiculo: ac.PlacaVehiculo,
		IDPuntoacceso: int64(ac.PuntoAccesoID),
		TiempoEntrada: ac.TiempoEntrada,
		Estado:        ac.Estado,
		Observaciones: aNullString(ac.Observaciones),
	})
	if err != nil {
		return modelos.Acceso{}
	}
	return aAcceso(f)
}

func (a *AlmacenSQLC) RegistrarSalida(id int, ac modelos.Acceso) (modelos.Acceso, bool) {
	var tiempoSalida time.Time
	if ac.TiempoSalida != nil {
		tiempoSalida = *ac.TiempoSalida
	}

	f, err := a.q.RegistrarAccesoSalida(context.Background(), sqlcdb.RegistrarAccesoSalidaParams{
		TiempoSalida:  sql.NullTime{Time: tiempoSalida, Valid: ac.TiempoSalida != nil},
		Estado:        ac.Estado,
		Observaciones: aNullString(ac.Observaciones),
		IDAcceso:      int64(id),
	})
	if err != nil {
		return modelos.Acceso{}, false
	}
	return aAcceso(f), true
}

func (a *AlmacenSQLC) LiberarAcceso(id int) (modelos.Acceso, bool) {
	ahora := time.Now()
	f, err := a.q.RegistrarAccesoSalida(context.Background(), sqlcdb.RegistrarAccesoSalidaParams{
		TiempoSalida: sql.NullTime{Time: ahora, Valid: true},
		Estado:       "finalizado",
		IDAcceso:     int64(id),
	})
	if err != nil {
		return modelos.Acceso{}, false
	}
	return aAcceso(f), true
}
