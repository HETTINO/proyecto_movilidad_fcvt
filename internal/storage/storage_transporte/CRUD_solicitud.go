package storage

import modelos "proyecto_movilidad_fcvt/internal/modelos"

func (m *Memoria) ListarSolicitudes() []modelos.Solicitud {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]modelos.Solicitud, len(m.solicitudes))
	copy(copia, m.solicitudes)
	return copia
}

func (m *Memoria) BuscarSolicitudPorID(id int) (modelos.Solicitud, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, s := range m.solicitudes {
		if s.ID == id {
			return s, true
		}
	}
	return modelos.Solicitud{}, false
}

func (m *Memoria) CrearSolicitud(s modelos.Solicitud) modelos.Solicitud {
	m.mu.Lock()
	defer m.mu.Unlock()
	s.ID = m.nextSolicitudID
	m.nextSolicitudID++
	m.solicitudes = append(m.solicitudes, s)
	return s
}

func (m *Memoria) ActualizarSolicitud(id int, datos modelos.Solicitud) (modelos.Solicitud, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, s := range m.solicitudes {
		if s.ID == id {
			datos.ID = id
			m.solicitudes[i] = datos
			return datos, true
		}
	}
	return modelos.Solicitud{}, false
}

func (m *Memoria) BorrarSolicitud(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, s := range m.solicitudes {
		if s.ID == id {
			m.solicitudes = append(m.solicitudes[:i], m.solicitudes[i+1:]...)
			return true
		}
	}
	return false
}
