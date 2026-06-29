package storage

import modelos "proyecto_movilidad_fcvt/internal/modelos"

func (m *Memoria) ListarRutas() []modelos.Ruta {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.rutas

}

func (m *Memoria) BuscarRutaPorID(id int) (modelos.Ruta, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, r := range m.rutas {
		if r.ID == id {
			return r, true
		}
	}

	return modelos.Ruta{}, false
}

func (m *Memoria) CrearRuta(r modelos.Ruta) modelos.Ruta {
	m.mu.Lock()
	defer m.mu.Unlock()

	r.ID = m.nextRutaID
	m.nextRutaID++

	m.rutas = append(m.rutas, r)

	return r
}

func (m *Memoria) ActualizarRuta(id int, datos modelos.Ruta) (modelos.Ruta, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, r := range m.rutas {
		if r.ID == id {

			datos.ID = id
			m.rutas[i] = datos

			return datos, true
		}
	}

	return modelos.Ruta{}, false
}

func (m *Memoria) BorrarRuta(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, r := range m.rutas {
		if r.ID == id {

			m.rutas = append(
				m.rutas[:i],
				m.rutas[i+1:]...,
			)

			return true
		}
	}

	return false
}
