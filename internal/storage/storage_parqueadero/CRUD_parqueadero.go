package storage_parqueadero

import (
	"proyecto_movilidad_fcvt/internal/modelos"
)

func (m *Memoria) ListarParqueaderos() []modelos.Parqueadero {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]modelos.Parqueadero, len(m.parqueaderos))
	copy(copia, m.parqueaderos)

	return copia
}

func (m *Memoria) BuscarParqueaderoPorID(id int) (modelos.Parqueadero, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, p := range m.parqueaderos {
		if p.IDParqueadero == id {
			return p, true
		}
	}

	return modelos.Parqueadero{}, false
}

func (m *Memoria) CrearParqueadero(p modelos.Parqueadero) (modelos.Parqueadero, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	p.IDParqueadero = m.nextParqueaderoID
	m.nextParqueaderoID++

	m.parqueaderos = append(m.parqueaderos, p)

	return p, nil
}

func (m *Memoria) ActualizarParqueadero(id int, datos modelos.Parqueadero) (modelos.Parqueadero, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, p := range m.parqueaderos {
		if p.IDParqueadero == id {

			datos.IDParqueadero = id
			m.parqueaderos[i] = datos

			return datos, true
		}
	}

	return modelos.Parqueadero{}, false
}

func (m *Memoria) BorrarParqueadero(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, p := range m.parqueaderos {
		if p.IDParqueadero == id {

			m.parqueaderos = append(
				m.parqueaderos[:i],
				m.parqueaderos[i+1:]...,
			)

			return true
		}
	}

	return false
}
