package storage_acceso

import "proyecto_movilidad_fcvt/internal/modelos"

func (m *MemoriaAcceso) ListarVehiculos() []modelos.Vehiculo {
	// m.mu.Lock()
	// defer m.mu.Unlock()

	copia := make([]modelos.Vehiculo, len(m.Vehiculos))
	copy(copia, m.Vehiculos)
	return copia
}

func (m *MemoriaAcceso) BuscarVehiculoPorPlaca(placa string) (modelos.Vehiculo, bool) {
	// m.mu.Lock()
	// defer m.mu.Unlock()

	for _, v := range m.Vehiculos {
		if v.Placa == placa {
			return v, true
		}
	}
	return modelos.Vehiculo{}, false
}

func (m *MemoriaAcceso) CrearVehiculo(v modelos.Vehiculo) modelos.Vehiculo {
	// m.mu.Lock()
	// defer m.mu.Unlock()

	m.Vehiculos = append(m.Vehiculos, v)
	return v
}

func (m *MemoriaAcceso) ActualizarVehiculo(placa string, datos modelos.Vehiculo) (modelos.Vehiculo, bool) {
	// m.mu.Lock()
	// defer m.mu.Unlock()

	for i, v := range m.Vehiculos {
		if v.Placa == placa {
			datos.Placa = placa
			m.Vehiculos[i] = datos
			return datos, true
		}
	}
	return modelos.Vehiculo{}, false
}

func (m *MemoriaAcceso) BorrarVehiculo(placa string) bool {
	// m.mu.Lock()
	// defer m.mu.Unlock()

	for i, v := range m.Vehiculos {
		if v.Placa == placa {
			m.Vehiculos = append(m.Vehiculos[:i], m.Vehiculos[i+1:]...)
			return true
		}
	}
	return false
}
