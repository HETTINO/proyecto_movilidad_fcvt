package storage

import modelos "proyecto_movilidad_fcvt/internal/models"

func (m *Memoria) ListarCarritos() []modelos.Carrito {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]modelos.Carrito, len(m.carritos))
	copy(copia, m.carritos)
	return copia
}

func (m *Memoria) BuscarCarritoPorID(id int) (modelos.Carrito, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, c := range m.carritos {
		if c.ID == id {
			return c, true
		}
	}
	return modelos.Carrito{}, false
}

func (m *Memoria) CrearCarrito(c modelos.Carrito) modelos.Carrito {
	m.mu.Lock()
	defer m.mu.Unlock()
	c.ID = m.nextCarritoID
	m.nextCarritoID++
	m.carritos = append(m.carritos, c)
	return c
}

func (m *Memoria) ActualizarCarrito(id int, datos modelos.Carrito) (modelos.Carrito, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, c := range m.carritos {
		if c.ID == id {
			datos.ID = id
			m.carritos[i] = datos
			return datos, true
		}
	}
	return modelos.Carrito{}, false
}

func (m *Memoria) BorrarCarrito(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, c := range m.carritos {
		if c.ID == id {
			m.carritos = append(m.carritos[:i], m.carritos[i+1:]...)
			return true
		}
	}
	return false
}
//var _ Almacen = (*Memoria)(nil)
