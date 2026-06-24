package storage_parqueadero

import "proyecto_movilidad_fcvt/internal/modelos"

func (m *Memoria) ListarEspacios() []modelos.Espacio {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]modelos.Espacio, len(m.espacios))
	copy(copia, m.espacios)

	return copia
}

func (m *Memoria) BuscarEspacioPorID(id int) (modelos.Espacio, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, e := range m.espacios {
		if e.IDEspacio == id {
			return e, true
		}
	}

	return modelos.Espacio{}, false
}

func (m *Memoria) CrearEspacio(e modelos.Espacio) modelos.Espacio {
	m.mu.Lock()
	defer m.mu.Unlock()

	e.IDEspacio = m.nextEspacioID
	m.nextEspacioID++

	m.espacios = append(m.espacios, e)

	return e
}

func (m *Memoria) ActualizarEspacio(id int, datos modelos.Espacio) (modelos.Espacio, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, e := range m.espacios {
		if e.IDEspacio == id {

			datos.IDEspacio = id
			m.espacios[i] = datos

			return datos, true
		}
	}

	return modelos.Espacio{}, false
}

func (m *Memoria) BorrarEspacio(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, e := range m.espacios {
		if e.IDEspacio == id {

			m.espacios = append(
				m.espacios[:i],
				m.espacios[i+1:]...,
			)

			return true
		}
	}

	return false
}
