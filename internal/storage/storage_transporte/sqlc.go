package storage

import (
	"context"
	"database/sql"

	modelos "proyecto_movilidad_fcvt/internal/modelos"
	sqlcdb "proyecto_movilidad_fcvt/internal/storage/sqlcdb_transporte"
)

type AlmacenSQLC struct {
	q *sqlcdb.Queries
}

func NuevoAlmacenSQLC(db *sql.DB) *AlmacenSQLC {
	return &AlmacenSQLC{q: sqlcdb.New(db)}
}

// =========================================================
// MAPEO sqlc <-> dominio
// =========================================================

func aRutaDominio(r sqlcdb.Ruta) modelos.Ruta {
	return modelos.Ruta{
		ID:          int(r.ID),
		Nombre:      r.Nombre,
		Descripcion: r.Descripcion.String,
	}
}

func aCarritoDominio(c sqlcdb.Carrito) modelos.Carrito {
	return modelos.Carrito{
		ID:            int(c.ID),
		NombreCarrito: c.NombreCarrito,
		Capacidad:     int(c.Capacidad),
		Estado:        c.Estado.String,
		RutaID:        int(c.RutaID.Int64),
	}
}

func aParadaDominio(p sqlcdb.Parada) modelos.Parada {
	return modelos.Parada{
		IDParada: int(p.IDParada),
		Nombre:   p.Nombre,
		Latitud:  p.Latitud.Float64,
		Longitud: p.Longitud.Float64,
	}
}

func aLocacionDominio(l sqlcdb.Locacione) modelos.Locacion {
	return modelos.Locacion{
		ID:        int(l.ID),
		Latitud:   l.Latitud,
		Longitud:  l.Longitud,
		TimeStamp: l.TimeStamp.Time,
		CarritoID: int(l.CarritoID),
	}
}

func aSolicitudDominio(s sqlcdb.Solicitude) modelos.Solicitud {
	var idCarrito *int
	if s.IDCarrito.Valid {
		id := int(s.IDCarrito.Int64)
		idCarrito = &id
	}
	return modelos.Solicitud{
		ID:            int(s.ID),
		CedulaUsuario: s.CedulaUsuario,
		CantPersonas:  int(s.CantPersonas.Int64),
		ParadaOrigen:  int(s.ParadaOrigen.Int64),
		PuntoDestino:  s.PuntoDestino.String,
		Estado:        s.Estado.String,
		IDCarrito:     idCarrito,
	}
}

// =========================================================
// RUTAS
// =========================================================

func (a *AlmacenSQLC) ListarRutas() []modelos.Ruta {
	filas, err := a.q.ListarRutas(context.Background())
	if err != nil {
		return nil
	}
	out := make([]modelos.Ruta, 0, len(filas))
	for _, f := range filas {
		out = append(out, aRutaDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarRutaPorID(id int) (modelos.Ruta, bool) {
	f, err := a.q.BuscarRutaPorID(context.Background(), int64(id))
	if err != nil {
		return modelos.Ruta{}, false
	}
	return aRutaDominio(f), true
}

func (a *AlmacenSQLC) CrearRuta(r modelos.Ruta) modelos.Ruta {
	f, err := a.q.CrearRuta(context.Background(), sqlcdb.CrearRutaParams{
		Nombre:      r.Nombre,
		Descripcion: sql.NullString{String: r.Descripcion, Valid: r.Descripcion != ""},
	})
	if err != nil {
		return modelos.Ruta{}
	}
	return aRutaDominio(f)
}

func (a *AlmacenSQLC) ActualizarRuta(id int, datos modelos.Ruta) (modelos.Ruta, bool) {
	f, err := a.q.ActualizarRuta(context.Background(), sqlcdb.ActualizarRutaParams{
		Nombre:      datos.Nombre,
		Descripcion: sql.NullString{String: datos.Descripcion, Valid: datos.Descripcion != ""},
		ID:          int64(id),
	})
	if err != nil {
		return modelos.Ruta{}, false
	}
	return aRutaDominio(f), true
}

func (a *AlmacenSQLC) BorrarRuta(id int) bool {
	filas, err := a.q.BorrarRuta(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

// =========================================================
// CARRITOS
// =========================================================

func (a *AlmacenSQLC) ListarCarritos() []modelos.Carrito {
	filas, err := a.q.ListarCarritos(context.Background())
	if err != nil {
		return nil
	}
	out := make([]modelos.Carrito, 0, len(filas))
	for _, f := range filas {
		out = append(out, aCarritoDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) ListarCarritosPorRuta(idRuta int) []modelos.Carrito {
	filas, err := a.q.ListarCarritosPorRuta(context.Background(), sql.NullInt64{Int64: int64(idRuta), Valid: true})
	if err != nil {
		return nil
	}
	out := make([]modelos.Carrito, 0, len(filas))
	for _, f := range filas {
		out = append(out, aCarritoDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarCarritoPorID(id int) (modelos.Carrito, bool) {
	f, err := a.q.BuscarCarritoPorID(context.Background(), int64(id))
	if err != nil {
		return modelos.Carrito{}, false
	}
	return aCarritoDominio(f), true
}

func (a *AlmacenSQLC) CrearCarrito(c modelos.Carrito) modelos.Carrito {
	f, err := a.q.CrearCarrito(context.Background(), sqlcdb.CrearCarritoParams{
		NombreCarrito: c.NombreCarrito,
		Capacidad:     int64(c.Capacidad),
		Estado:        sql.NullString{String: c.Estado, Valid: c.Estado != ""},
		RutaID:        sql.NullInt64{Int64: int64(c.RutaID), Valid: true},
	})
	if err != nil {
		return modelos.Carrito{}
	}
	return aCarritoDominio(f)
}

func (a *AlmacenSQLC) ActualizarCarrito(id int, datos modelos.Carrito) (modelos.Carrito, bool) {
	f, err := a.q.ActualizarCarrito(context.Background(), sqlcdb.ActualizarCarritoParams{
		NombreCarrito: datos.NombreCarrito,
		Capacidad:     int64(datos.Capacidad),
		Estado:        sql.NullString{String: datos.Estado, Valid: datos.Estado != ""},
		RutaID:        sql.NullInt64{Int64: int64(datos.RutaID), Valid: true},
		ID:            int64(id),
	})
	if err != nil {
		return modelos.Carrito{}, false
	}
	return aCarritoDominio(f), true
}

func (a *AlmacenSQLC) BorrarCarrito(id int) bool {
	filas, err := a.q.BorrarCarrito(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

// =========================================================
// PARADAS
// =========================================================

func (a *AlmacenSQLC) ListarParadas() []modelos.Parada {
	filas, err := a.q.ListarParadas(context.Background())
	if err != nil {
		return nil
	}
	out := make([]modelos.Parada, 0, len(filas))
	for _, f := range filas {
		out = append(out, aParadaDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarParadaPorID(id int) (modelos.Parada, bool) {
	f, err := a.q.BuscarParadaPorID(context.Background(), int64(id))
	if err != nil {
		return modelos.Parada{}, false
	}
	return aParadaDominio(f), true
}

func (a *AlmacenSQLC) CrearParada(p modelos.Parada) modelos.Parada {
	f, err := a.q.CrearParada(context.Background(), sqlcdb.CrearParadaParams{
		Nombre:   p.Nombre,
		Latitud:  sql.NullFloat64{Float64: p.Latitud, Valid: true},
		Longitud: sql.NullFloat64{Float64: p.Longitud, Valid: true},
	})
	if err != nil {
		return modelos.Parada{}
	}
	return aParadaDominio(f)
}

func (a *AlmacenSQLC) ActualizarParada(id int, datos modelos.Parada) (modelos.Parada, bool) {
	f, err := a.q.ActualizarParada(context.Background(), sqlcdb.ActualizarParadaParams{
		Nombre:   datos.Nombre,
		Latitud:  sql.NullFloat64{Float64: datos.Latitud, Valid: true},
		Longitud: sql.NullFloat64{Float64: datos.Longitud, Valid: true},
		IDParada: int64(id),
	})
	if err != nil {
		return modelos.Parada{}, false
	}
	return aParadaDominio(f), true
}

func (a *AlmacenSQLC) BorrarParada(id int) bool {
	filas, err := a.q.BorrarParada(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

// =========================================================
// LOCACIONES
// =========================================================

func (a *AlmacenSQLC) ListarLocaciones() []modelos.Locacion {
	filas, err := a.q.ListarLocaciones(context.Background())
	if err != nil {
		return nil
	}
	out := make([]modelos.Locacion, 0, len(filas))
	for _, f := range filas {
		out = append(out, aLocacionDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) ObtenerUltimaLocacionPorCarrito(carritoID int) (modelos.Locacion, bool) {
	f, err := a.q.ObtenerUltimaLocacionPorCarrito(context.Background(), int64(carritoID))
	if err != nil {
		return modelos.Locacion{}, false
	}
	return aLocacionDominio(f), true
}

func (a *AlmacenSQLC) RegistrarLocacion(l modelos.Locacion) modelos.Locacion {
	f, err := a.q.RegistrarLocacion(context.Background(), sqlcdb.RegistrarLocacionParams{
		Latitud:   l.Latitud,
		Longitud:  l.Longitud,
		TimeStamp: sql.NullTime{Time: l.TimeStamp, Valid: true},
		CarritoID: int64(l.CarritoID),
	})
	if err != nil {
		return modelos.Locacion{}
	}
	return aLocacionDominio(f)
}

// =========================================================
// SOLICITUDES
// =========================================================

func (a *AlmacenSQLC) ListarSolicitudes() []modelos.Solicitud {
	filas, err := a.q.ListarSolicitudes(context.Background())
	if err != nil {
		return nil
	}
	out := make([]modelos.Solicitud, 0, len(filas))
	for _, f := range filas {
		out = append(out, aSolicitudDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarSolicitudPorID(id int) (modelos.Solicitud, bool) {
	f, err := a.q.BuscarSolicitudPorID(context.Background(), int64(id))
	if err != nil {
		return modelos.Solicitud{}, false
	}
	return aSolicitudDominio(f), true
}

func (a *AlmacenSQLC) CrearSolicitud(s modelos.Solicitud) modelos.Solicitud {
	var idCarrito sql.NullInt64
	if s.IDCarrito != nil {
		idCarrito = sql.NullInt64{Int64: int64(*s.IDCarrito), Valid: true}
	}

	f, err := a.q.CrearSolicitud(context.Background(), sqlcdb.CrearSolicitudParams{
		CedulaUsuario: s.CedulaUsuario,
		CantPersonas:  sql.NullInt64{Int64: int64(s.CantPersonas), Valid: true},
		ParadaOrigen:  sql.NullInt64{Int64: int64(s.ParadaOrigen), Valid: true},
		PuntoDestino:  sql.NullString{String: s.PuntoDestino, Valid: s.PuntoDestino != ""},
		Estado:        sql.NullString{String: s.Estado, Valid: s.Estado != ""},
		IDCarrito:     idCarrito,
	})
	if err != nil {
		return modelos.Solicitud{}
	}
	return aSolicitudDominio(f)
}

func (a *AlmacenSQLC) ActualizarSolicitud(id int, datos modelos.Solicitud) (modelos.Solicitud, bool) {
	var idCarrito sql.NullInt64
	if datos.IDCarrito != nil {
		idCarrito = sql.NullInt64{Int64: int64(*datos.IDCarrito), Valid: true}
	}

	f, err := a.q.ActualizarSolicitud(context.Background(), sqlcdb.ActualizarSolicitudParams{
		CedulaUsuario: datos.CedulaUsuario,
		CantPersonas:  sql.NullInt64{Int64: int64(datos.CantPersonas), Valid: true},
		ParadaOrigen:  sql.NullInt64{Int64: int64(datos.ParadaOrigen), Valid: true},
		PuntoDestino:  sql.NullString{String: datos.PuntoDestino, Valid: datos.PuntoDestino != ""},
		Estado:        sql.NullString{String: datos.Estado, Valid: datos.Estado != ""},
		IDCarrito:     idCarrito,
		ID:            int64(id),
	})
	if err != nil {
		return modelos.Solicitud{}, false
	}
	return aSolicitudDominio(f), true
}

func (a *AlmacenSQLC) BorrarSolicitud(id int) bool {
	filas, err := a.q.BorrarSolicitud(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

var _ Almacen = (*AlmacenSQLC)(nil)
