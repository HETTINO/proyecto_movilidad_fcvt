package storage_parqueadero

import (
	"context"
	"database/sql"
	"time"

	"proyecto_movilidad_fcvt/internal/modelos"
	"proyecto_movilidad_fcvt/internal/storage/sqlcdb"
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

func aParqueaderoDominio(p sqlcdb.Parqueadero) modelos.Parqueadero {
	return modelos.Parqueadero{
		IDParqueadero: int(p.IDParqueadero),
		Nombre:        p.Nombre,
		Capacidad:     int(p.Capacidad),
		Tipo:          p.Tipo,
	}
}

func aEspacioDominio(e sqlcdb.Espacio) modelos.Espacio {
	return modelos.Espacio{
		IDEspacio:     int(e.IDEspacio),
		IDParqueadero: int(e.IDParqueadero),
		Numero:        int(e.Numero),
		Estado:        e.Estado,
		TipoEspacio:   e.TipoEspacio,
	}
}

func aOcupacionDominio(o sqlcdb.Ocupacion) modelos.Ocupacion {
	var horaFin *time.Time
	if o.HoraFin.Valid {
		t := o.HoraFin.Time
		horaFin = &t
	}
	return modelos.Ocupacion{
		IDOcupacion:   int(o.IDOcupacion),
		PlacaVehiculo: o.PlacaVehiculo,
		IDEspacio:     int(o.IDEspacio),
		IDAcceso:      int(o.IDAcceso),
		HoraInicio:    o.HoraInicio,
		HoraFin:       horaFin,
	}
}

// =========================================================
// PARQUEADEROS
// =========================================================

func (a *AlmacenSQLC) ListarParqueaderos() []modelos.Parqueadero {
	filas, err := a.q.ListarParqueaderos(context.Background())
	if err != nil {
		return nil
	}
	out := make([]modelos.Parqueadero, 0, len(filas))
	for _, f := range filas {
		out = append(out, aParqueaderoDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarParqueaderoPorID(id int) (modelos.Parqueadero, bool) {
	f, err := a.q.BuscarParqueaderoPorID(context.Background(), int64(id))
	if err != nil {
		return modelos.Parqueadero{}, false
	}
	return aParqueaderoDominio(f), true
}

func (a *AlmacenSQLC) CrearParqueadero(p modelos.Parqueadero) modelos.Parqueadero {
	f, err := a.q.CrearParqueadero(context.Background(), sqlcdb.CrearParqueaderoParams{
		Nombre:    p.Nombre,
		Capacidad: int64(p.Capacidad),
		Tipo:      p.Tipo,
	})
	if err != nil {
		return modelos.Parqueadero{}
	}
	return aParqueaderoDominio(f)
}

func (a *AlmacenSQLC) ActualizarParqueadero(id int, datos modelos.Parqueadero) (modelos.Parqueadero, bool) {
	f, err := a.q.ActualizarParqueadero(context.Background(), sqlcdb.ActualizarParqueaderoParams{
		Nombre:        datos.Nombre,
		Capacidad:     int64(datos.Capacidad),
		Tipo:          datos.Tipo,
		IDParqueadero: int64(id),
	})
	if err != nil {
		return modelos.Parqueadero{}, false
	}
	return aParqueaderoDominio(f), true
}

func (a *AlmacenSQLC) BorrarParqueadero(id int) bool {
	filas, err := a.q.BorrarParqueadero(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

// =========================================================
// ESPACIOS
// =========================================================

func (a *AlmacenSQLC) ListarEspacios() []modelos.Espacio {
	filas, err := a.q.ListarEspacios(context.Background())
	if err != nil {
		return nil
	}
	out := make([]modelos.Espacio, 0, len(filas))
	for _, f := range filas {
		out = append(out, aEspacioDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) ListarEspaciosPorParqueadero(idParqueadero int) []modelos.Espacio {
	filas, err := a.q.ListarEspaciosPorParqueadero(context.Background(), int64(idParqueadero))
	if err != nil {
		return nil
	}
	out := make([]modelos.Espacio, 0, len(filas))
	for _, f := range filas {
		out = append(out, aEspacioDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarEspacioPorID(id int) (modelos.Espacio, bool) {
	f, err := a.q.BuscarEspacioPorID(context.Background(), int64(id))
	if err != nil {
		return modelos.Espacio{}, false
	}
	return aEspacioDominio(f), true
}

func (a *AlmacenSQLC) CrearEspacio(e modelos.Espacio) modelos.Espacio {
	f, err := a.q.CrearEspacio(context.Background(), sqlcdb.CrearEspacioParams{
		IDParqueadero: int64(e.IDParqueadero),
		Numero:        int64(e.Numero),
		Estado:        e.Estado,
		TipoEspacio:   e.TipoEspacio,
	})
	if err != nil {
		return modelos.Espacio{}
	}
	return aEspacioDominio(f)
}

func (a *AlmacenSQLC) ActualizarEspacio(id int, datos modelos.Espacio) (modelos.Espacio, bool) {
	f, err := a.q.ActualizarEspacio(context.Background(), sqlcdb.ActualizarEspacioParams{
		IDParqueadero: int64(datos.IDParqueadero),
		Numero:        int64(datos.Numero),
		Estado:        datos.Estado,
		TipoEspacio:   datos.TipoEspacio,
		IDEspacio:     int64(id),
	})
	if err != nil {
		return modelos.Espacio{}, false
	}
	return aEspacioDominio(f), true
}

func (a *AlmacenSQLC) BorrarEspacio(id int) bool {
	filas, err := a.q.BorrarEspacio(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

// =========================================================
// OCUPACIONES
// =========================================================

func (a *AlmacenSQLC) ListarOcupaciones() []modelos.Ocupacion {
	filas, err := a.q.ListarOcupaciones(context.Background())
	if err != nil {
		return nil
	}
	out := make([]modelos.Ocupacion, 0, len(filas))
	for _, f := range filas {
		out = append(out, aOcupacionDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarOcupacionPorID(id int) (modelos.Ocupacion, bool) {
	f, err := a.q.BuscarOcupacionPorID(context.Background(), int64(id))
	if err != nil {
		return modelos.Ocupacion{}, false
	}
	return aOcupacionDominio(f), true
}

func (a *AlmacenSQLC) ListarOcupacionesActivas(idEspacio int) []modelos.Ocupacion {
	filas, err := a.q.ListarOcupacionesActivasPorEspacio(context.Background(), int64(idEspacio))
	if err != nil {
		return nil
	}
	out := make([]modelos.Ocupacion, 0, len(filas))
	for _, f := range filas {
		out = append(out, aOcupacionDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) CrearOcupacion(o modelos.Ocupacion) modelos.Ocupacion {
	f, err := a.q.CrearOcupacion(context.Background(), sqlcdb.CrearOcupacionParams{
		PlacaVehiculo: o.PlacaVehiculo,
		IDEspacio:     int64(o.IDEspacio),
		IDAcceso:      int64(o.IDAcceso),
		HoraInicio:    o.HoraInicio,
	})
	if err != nil {
		return modelos.Ocupacion{}
	}
	return aOcupacionDominio(f)
}

func (a *AlmacenSQLC) CerrarOcupacion(id int, datos modelos.Ocupacion) (modelos.Ocupacion, bool) {
	f, err := a.q.CerrarOcupacion(context.Background(), sqlcdb.CerrarOcupacionParams{
		HoraFin:     sql.NullTime{Time: *datos.HoraFin, Valid: datos.HoraFin != nil},
		IDOcupacion: int64(id),
	})
	if err != nil {
		return modelos.Ocupacion{}, false
	}
	return aOcupacionDominio(f), true
}

func (a *AlmacenSQLC) ActualizarOcupacion(id int, datos modelos.Ocupacion) (modelos.Ocupacion, bool) {
	f, err := a.q.ActualizarOcupacion(context.Background(), sqlcdb.ActualizarOcupacionParams{
		PlacaVehiculo: datos.PlacaVehiculo,
		IDEspacio:     int64(datos.IDEspacio),
		IDAcceso:      int64(datos.IDAcceso),
		HoraInicio:    datos.HoraInicio,
		HoraFin: sql.NullTime{Time: func() time.Time {
			if datos.HoraFin != nil {
				return *datos.HoraFin
			}
			return time.Time{}
		}(), Valid: datos.HoraFin != nil},
		IDOcupacion: int64(id),
	})
	if err != nil {
		return modelos.Ocupacion{}, false
	}
	return aOcupacionDominio(f), true
}

func (a *AlmacenSQLC) BorrarOcupacion(id int) bool {
	filas, err := a.q.BorrarOcupacion(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

func (a *AlmacenSQLC) LiberarOcupacion(id int) (modelos.Ocupacion, bool) {
	ahora := time.Now()
	f, err := a.q.CerrarOcupacion(context.Background(), sqlcdb.CerrarOcupacionParams{
		HoraFin:     sql.NullTime{Time: ahora, Valid: true},
		IDOcupacion: int64(id),
	})
	if err != nil {
		return modelos.Ocupacion{}, false
	}
	return aOcupacionDominio(f), true
}

var _ Almacen = (*AlmacenSQLC)(nil)
