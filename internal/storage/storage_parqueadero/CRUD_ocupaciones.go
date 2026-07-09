package storage_parqueadero

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	"time"
)

// CRUD Ocupaciones
func (m *Memoria) ListarOcupaciones() []modelos.Ocupacion {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]modelos.Ocupacion, len(m.ocupaciones))
	copy(copia, m.ocupaciones)

	return copia
}

func (m *Memoria) BuscarOcupacionPorID(id int) (modelos.Ocupacion, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, o := range m.ocupaciones {
		if o.IDOcupacion == id {
			return o, true
		}
	}

	return modelos.Ocupacion{}, false
}

func (m *Memoria) CrearOcupacion(o modelos.Ocupacion) modelos.Ocupacion {
	m.mu.Lock()
	defer m.mu.Unlock()

	o.IDOcupacion = m.nextOcupacionID
	o.HoraInicio = time.Now()

	m.nextOcupacionID++

	m.ocupaciones = append(m.ocupaciones, o)

	return o
}

func (m *Memoria) ActualizarOcupacion(id int, datos modelos.Ocupacion) (modelos.Ocupacion, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, o := range m.ocupaciones {
		if o.IDOcupacion == id {

			datos.IDOcupacion = id
			m.ocupaciones[i] = datos

			return datos, true
		}
	}

	return modelos.Ocupacion{}, false
}

func (m *Memoria) BorrarOcupacion(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, o := range m.ocupaciones {
		if o.IDOcupacion == id {

			m.ocupaciones = append(
				m.ocupaciones[:i],
				m.ocupaciones[i+1:]...,
			)

			return true
		}
	}

	return false
}

// ListarOcupacionesActivas devuelve las ocupaciones del espacio indicado
// que todavía no tienen HoraFin (siguen en curso).
func (m *Memoria) ListarOcupacionesActivas(idEspacio int) []modelos.Ocupacion {
	m.mu.Lock()
	defer m.mu.Unlock()

	var activas []modelos.Ocupacion
	for _, o := range m.ocupaciones {
		if o.IDEspacio == idEspacio && o.HoraFin == nil {
			activas = append(activas, o)
		}
	}

	return activas
}

func (m *Memoria) LiberarOcupacion(id int) (modelos.Ocupacion, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, o := range m.ocupaciones {

		if o.IDOcupacion == id {

			ahora := time.Now()

			m.ocupaciones[i].HoraFin = &ahora

			return m.ocupaciones[i], true
		}
	}

	return modelos.Ocupacion{}, false
}
