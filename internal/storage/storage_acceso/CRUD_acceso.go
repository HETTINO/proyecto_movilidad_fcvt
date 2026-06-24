package storage_acceso

import "proyecto_movilidad_fcvt/internal/modelos"

func (m *MemoriaAcceso) ListarAccesos() []modelos.Acceso {
	// m.mu.Lock()
	// defer m.mu.Unlock()

	copia := make([]modelos.Acceso, len(m.Accesos))
	copy(copia, m.Accesos)
	return copia
}

func (m *MemoriaAcceso) BuscarAccesoPorID(id int) (modelos.Acceso, bool) {
	// m.mu.Lock()
	// defer m.mu.Unlock()

	for _, a := range m.Accesos {
		if a.ID == id {
			return a, true
		}
	}
	return modelos.Acceso{}, false
}

func (m *MemoriaAcceso) CrearAcceso(a modelos.Acceso) modelos.Acceso {
	// m.mu.Lock()
	// defer m.mu.Unlock()

	// a.ID = m.nextAccesoID
	// m.nextAccesoID++

	m.Accesos = append(m.Accesos, a)
	return a
}

func (m *MemoriaAcceso) ActualizarAcceso(id int, datos modelos.Acceso) (modelos.Acceso, bool) {
	// m.mu.Lock()
	// defer m.mu.Unlock()

	for i, a := range m.Accesos {
		if a.ID == id {
			datos.ID = id
			m.Accesos[i] = datos
			return datos, true
		}
	}
	return modelos.Acceso{}, false
}

func (m *MemoriaAcceso) BorrarAcceso(id int) bool {
	// m.mu.Lock()
	// defer m.mu.Unlock()

	for i, a := range m.Accesos {
		if a.ID == id {
			m.Accesos = append(m.Accesos[:i], m.Accesos[i+1:]...)
			return true
		}
	}
	return false
}
