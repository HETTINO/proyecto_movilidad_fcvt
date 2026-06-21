package storage

import modelos "proyecto_movilidad_fcvt/internal/models"

func (m *Memoria) ListarParadas() []modelos.Parada {
	m.mu.Lock()
	defer m.mu.Unlock()



	return m.paradas
}

func (m *Memoria) BuscarParadaPorID(id int) (modelos.Parada, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, p := range m.paradas {
		if p.IDParada == id {
			return p, true
		}
	}

	return modelos.Parada{}, false
}

func (m *Memoria) CrearParada(p modelos.Parada) modelos.Parada {
	m.mu.Lock()
	defer m.mu.Unlock()

	p.IDParada = m.nextParadaID
	m.nextParadaID++

	m.paradas = append(m.paradas, p)

	return p
}

func (m *Memoria) ActualizarParada(id int, datos modelos.Parada) (modelos.Parada, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, p := range m.paradas {
		if p.IDParada == id {
			datos.IDParada = id
			m.paradas[i] = datos
			return datos, true
		}
	}

	return modelos.Parada{}, false
}

func (m *Memoria) BorrarParada(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, p := range m.paradas {
		if p.IDParada == id {
			m.paradas = append(
				m.paradas[:i],
				m.paradas[i+1:]...,
			)
			return true
		}
	}

	return false
}
