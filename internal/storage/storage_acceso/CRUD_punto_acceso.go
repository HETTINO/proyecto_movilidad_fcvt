package storage_acceso

import "proyecto_movilidad_fcvt/internal/modelos"

func (m *MemoriaAcceso) ListarPuntosDeAcceso() []modelos.PuntoDeAcceso {
	// m.mu.Lock()
	// defer m.mu.Unlock()

	copia := make([]modelos.PuntoDeAcceso, len(m.PuntosDeAcceso))
	copy(copia, m.PuntosDeAcceso)
	return copia
}

func (m *MemoriaAcceso) BuscarPuntoPorID(id int) (modelos.PuntoDeAcceso, bool) {
	// m.mu.Lock()
	// defer m.mu.Unlock()

	for _, p := range m.PuntosDeAcceso {
		if p.ID == id {
			return p, true
		}
	}
	return modelos.PuntoDeAcceso{}, false
}

func (m *MemoriaAcceso) CrearPuntoDeAcceso(p modelos.PuntoDeAcceso) modelos.PuntoDeAcceso {
	// m.mu.Lock()
	// defer m.mu.Unlock()

	// p.ID = m.nextPuntoID // Descomenta si usas un contador secuencial secuencial en tu struct
	// m.nextPuntoID++

	m.PuntosDeAcceso = append(m.PuntosDeAcceso, p)
	return p
}

func (m *MemoriaAcceso) ActualizarPuntoDeAcceso(id int, datos modelos.PuntoDeAcceso) (modelos.PuntoDeAcceso, bool) {
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

func (m *MemoriaAcceso) BorrarPuntoDeAcceso(id int) bool {
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
