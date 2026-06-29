package storage

import modelos "proyecto_movilidad_fcvt/internal/modelos"

func (m *Memoria) ListarLocaciones() []modelos.Locacion {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]modelos.Locacion, len(m.locaciones))
	copy(copia, m.locaciones)
	return copia
}

func (m *Memoria) RegistrarLocacion(l modelos.Locacion) modelos.Locacion {
	m.mu.Lock()
	defer m.mu.Unlock()
	l.ID = m.nextLocacionID
	m.nextLocacionID++
	m.locaciones = append(m.locaciones, l)
	return l
}

func (m *Memoria) ObtenerUltimaLocacionPorCarrito(carritoID int) (modelos.Locacion, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	var ultima modelos.Locacion
	encontrado := false
	for _, l := range m.locaciones {
		if l.CarritoID == carritoID {
			ultima = l
			encontrado = true
		}
	}
	return ultima, encontrado
}
