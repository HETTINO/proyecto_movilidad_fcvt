package storage_acceso

import "proyecto_movilidad_fcvt/internal/modelos"

func (m *MemoriaAcceso) ListarPuntosAcceso() []modelos.PuntoDeAcceso {
	// m.mu.Lock()
	// defer m.mu.Unlock()

	copia := make([]modelos.PuntoDeAcceso, len(m.PuntosDeAcceso))
	copy(copia, m.PuntosDeAcceso)
	return copia
}

func (m *MemoriaAcceso) BuscarPuntoAccesoPorID(id int) (modelos.PuntoDeAcceso, bool) {
	// m.mu.Lock()
	// defer m.mu.Unlock()

	for _, p := range m.PuntosDeAcceso {
		if p.ID == id {
			return p, true
		}
	}
	return modelos.PuntoDeAcceso{}, false
}

func (m *MemoriaAcceso) CrearPuntoAcceso(p modelos.PuntoDeAcceso) modelos.PuntoDeAcceso {
	// m.mu.Lock()
	// defer m.mu.Unlock()

	p.ID = m.nextIDPuntoAcceso
	m.nextIDPuntoAcceso++

	m.PuntosDeAcceso = append(m.PuntosDeAcceso, p)
	return p
}

func (m *MemoriaAcceso) ActualizarPuntoAcceso(id int, datos modelos.PuntoDeAcceso) (modelos.PuntoDeAcceso, bool) {
	// m.mu.Lock()
	// defer m.mu.Unlock()

	for i, p := range m.PuntosDeAcceso {
		if p.ID == id {
			datos.ID = id
			m.PuntosDeAcceso[i] = datos
			return datos, true
		}
	}
	return modelos.PuntoDeAcceso{}, false
}

func (m *MemoriaAcceso) BorrarPuntoAcceso(id int) bool {
	// m.mu.Lock()
	// defer m.mu.Unlock()

	for i, p := range m.PuntosDeAcceso {
		if p.ID == id {
			m.PuntosDeAcceso = append(m.PuntosDeAcceso[:i], m.PuntosDeAcceso[i+1:]...)
			return true
		}
	}
	return false
}
