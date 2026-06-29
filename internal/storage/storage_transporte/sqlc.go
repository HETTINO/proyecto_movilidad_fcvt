package storage

import (
	"context"
	"database/sql"

	"proyecto_movilidad_fcvt/internal/models"
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

func aRutaDominio(r sqlcdb.Ruta) models.Ruta {
	return models.Ruta{
		ID:          int(r.ID),
		Nombre:      r.Nombre,
		Descripcion: r.Descripcion.String,
	}
}

func aCarritoDominio(c sqlcdb.Carrito) models.Carrito {
	return models.Carrito{
		ID:            int(c.ID),
		NombreCarrito: c.NombreCarrito,
		Capacidad:     int(c.Capacidad),
		Estado:        c.Estado.String,
		RutaID:        int(c.RutaID.Int64),
	}
}

func aParadaDominio(p sqlcdb.Parada) models.Parada {
	return models.Parada{
		IDParada: int(p.IDParada),
		Nombre:   p.Nombre,
		Latitud:  p.Latitud.Float64,
		Longitud: p.Longitud.Float64,
	}
}

func aLocacionDominio(l sqlcdb.Locacione) models.Locacion {
	return models.Locacion{
		ID:        int(l.ID),
		Latitud:   l.Latitud,
		Longitud:  l.Longitud,
		TimeStamp: l.TimeStamp.Time,
		CarritoID: int(l.CarritoID),
	}
}

func aSolicitudDominio(s sqlcdb.Solicitude) models.Solicitud {
	var idCarrito *int
	if s.IDCarrito.Valid {
		id := int(s.IDCarrito.Int64)
		idCarrito = &id
	}
	return models.Solicitud{
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

func (a *AlmacenSQLC) ListarRutas() []models.Ruta {
	filas, err := a.q.ListarRutas(context.Background())
	if err != nil {
		return nil
	}
	out := make([]models.Ruta, 0, len(filas))
	for _, f := range filas {
		out = append(out, aRutaDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarRutaPorID(id int) (models.Ruta, bool) {
	f, err := a.q.BuscarRutaPorID(context.Background(), int64(id))
	if err != nil {
		return models.Ruta{}, false
	}
	return aRutaDominio(f), true
}

func (a *AlmacenSQLC) CrearRuta(r models.Ruta) models.Ruta {
	f, err := a.q.CrearRuta(context.Background(), sqlcdb.CrearRutaParams{
		Nombre:      r.Nombre,
		Descripcion: sql.NullString{String: r.Descripcion, Valid: r.Descripcion != ""},
	})
	if err != nil {
		return models.Ruta{}
	}
	return aRutaDominio(f)
}

func (a *AlmacenSQLC) ActualizarRuta(id int, datos models.Ruta) (models.Ruta, bool) {
	f, err := a.q.ActualizarRuta(context.Background(), sqlcdb.ActualizarRutaParams{
		Nombre:      datos.Nombre,
		Descripcion: sql.NullString{String: datos.Descripcion, Valid: datos.Descripcion != ""},
		ID:          int64(id),
	})
	if err != nil {
		return models.Ruta{}, false
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

func (a *AlmacenSQLC) ListarCarritos() []models.Carrito {
	filas, err := a.q.ListarCarritos(context.Background())
	if err != nil {
		return nil
	}
	out := make([]models.Carrito, 0, len(filas))
	for _, f := range filas {
		out = append(out, aCarritoDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) ListarCarritosPorRuta(idRuta int) []models.Carrito {
	filas, err := a.q.ListarCarritosPorRuta(context.Background(), sql.NullInt64{Int64: int64(idRuta), Valid: true})
	if err != nil {
		return nil
	}
	out := make([]models.Carrito, 0, len(filas))
	for _, f := range filas {
		out = append(out, aCarritoDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarCarritoPorID(id int) (models.Carrito, bool) {
	f, err := a.q.BuscarCarritoPorID(context.Background(), int64(id))
	if err != nil {
		return models.Carrito{}, false
	}
	return aCarritoDominio(f), true
}

func (a *AlmacenSQLC) CrearCarrito(c models.Carrito) models.Carrito {
	f, err := a.q.CrearCarrito(context.Background(), sqlcdb.CrearCarritoParams{
		NombreCarrito: c.NombreCarrito,
		Capacidad:     int64(c.Capacidad),
		Estado:        sql.NullString{String: c.Estado, Valid: c.Estado != ""},
		RutaID:        sql.NullInt64{Int64: int64(c.RutaID), Valid: true},
	})
	if err != nil {
		return models.Carrito{}
	}
	return aCarritoDominio(f)
}

func (a *AlmacenSQLC) ActualizarCarrito(id int, datos models.Carrito) (models.Carrito, bool) {
	f, err := a.q.ActualizarCarrito(context.Background(), sqlcdb.ActualizarCarritoParams{
		NombreCarrito: datos.NombreCarrito,
		Capacidad:     int64(datos.Capacidad),
		Estado:        sql.NullString{String: datos.Estado, Valid: datos.Estado != ""},
		RutaID:        sql.NullInt64{Int64: int64(datos.RutaID), Valid: true},
		ID:            int64(id),
	})
	if err != nil {
		return models.Carrito{}, false
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

func (a *AlmacenSQLC) ListarParadas() []models.Parada {
	filas, err := a.q.ListarParadas(context.Background())
	if err != nil {
		return nil
	}
	out := make([]models.Parada, 0, len(filas))
	for _, f := range filas {
		out = append(out, aParadaDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarParadaPorID(id int) (models.Parada, bool) {
	f, err := a.q.BuscarParadaPorID(context.Background(), int64(id))
	if err != nil {
		return models.Parada{}, false
	}
	return aParadaDominio(f), true
}

func (a *AlmacenSQLC) CrearParada(p models.Parada) models.Parada {
	f, err := a.q.CrearParada(context.Background(), sqlcdb.CrearParadaParams{
		Nombre:   p.Nombre,
		Latitud:  sql.NullFloat64{Float64: p.Latitud, Valid: true},
		Longitud: sql.NullFloat64{Float64: p.Longitud, Valid: true},
	})
	if err != nil {
		return models.Parada{}
	}
	return aParadaDominio(f)
}

func (a *AlmacenSQLC) ActualizarParada(id int, datos models.Parada) (models.Parada, bool) {
	f, err := a.q.ActualizarParada(context.Background(), sqlcdb.ActualizarParadaParams{
		Nombre:   datos.Nombre,
		Latitud:  sql.NullFloat64{Float64: datos.Latitud, Valid: true},
		Longitud: sql.NullFloat64{Float64: datos.Longitud, Valid: true},
		IDParada: int64(id),
	})
	if err != nil {
		return models.Parada{}, false
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

func (a *AlmacenSQLC) ListarLocaciones() []models.Locacion {
	filas, err := a.q.ListarLocaciones(context.Background())
	if err != nil {
		return nil
	}
	out := make([]models.Locacion, 0, len(filas))
	for _, f := range filas {
		out = append(out, aLocacionDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) ObtenerUltimaLocacionPorCarrito(carritoID int) (models.Locacion, bool) {
	f, err := a.q.ObtenerUltimaLocacionPorCarrito(context.Background(), int64(carritoID))
	if err != nil {
		return models.Locacion{}, false
	}
	return aLocacionDominio(f), true
}

func (a *AlmacenSQLC) RegistrarLocacion(l models.Locacion) models.Locacion {
	f, err := a.q.RegistrarLocacion(context.Background(), sqlcdb.RegistrarLocacionParams{
		Latitud:   l.Latitud,
		Longitud:  l.Longitud,
		TimeStamp: sql.NullTime{Time: l.TimeStamp, Valid: true},
		CarritoID: int64(l.CarritoID),
	})
	if err != nil {
		return models.Locacion{}
	}
	return aLocacionDominio(f)
}

// =========================================================
// SOLICITUDES
// =========================================================

func (a *AlmacenSQLC) ListarSolicitudes() []models.Solicitud {
	filas, err := a.q.ListarSolicitudes(context.Background())
	if err != nil {
		return nil
	}
	out := make([]models.Solicitud, 0, len(filas))
	for _, f := range filas {
		out = append(out, aSolicitudDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarSolicitudPorID(id int) (models.Solicitud, bool) {
	f, err := a.q.BuscarSolicitudPorID(context.Background(), int64(id))
	if err != nil {
		return models.Solicitud{}, false
	}
	return aSolicitudDominio(f), true
}

func (a *AlmacenSQLC) CrearSolicitud(s models.Solicitud) models.Solicitud {
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
		return models.Solicitud{}
	}
	return aSolicitudDominio(f)
}

func (a *AlmacenSQLC) ActualizarSolicitud(id int, datos models.Solicitud) (models.Solicitud, bool) {
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
		return models.Solicitud{}, false
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