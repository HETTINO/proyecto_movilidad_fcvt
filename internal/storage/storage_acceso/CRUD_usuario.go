package storage_acceso

import "proyecto_movilidad_fcvt/internal/modelos"

func (m *MemoriaAcceso) ListarUsuarios() []modelos.Usuario {
	// m.mu.Lock() // Descomenta estas líneas si usas exclusión mutua 'mu'
	// defer m.mu.Unlock()

	copia := make([]modelos.Usuario, len(m.Usuarios))
	copy(copia, m.Usuarios)
	return copia
}

func (m *MemoriaAcceso) BuscarUsuarioPorCedula(cedula string) (modelos.Usuario, bool) {
	// m.mu.Lock()
	// defer m.mu.Unlock()

	for _, u := range m.Usuarios {
		if u.Cedula == cedula {
			return u, true
		}
	}
	return modelos.Usuario{}, false
}

func (m *MemoriaAcceso) CrearUsuario(u modelos.Usuario) modelos.Usuario {
	// m.mu.Lock()
	// defer m.mu.Unlock()

	m.Usuarios = append(m.Usuarios, u)
	return u
}

func (m *MemoriaAcceso) ActualizarUsuario(cedula string, datos modelos.Usuario) (modelos.Usuario, bool) {
	// m.mu.Lock()
	// defer m.mu.Unlock()

	for i, u := range m.Usuarios {
		if u.Cedula == cedula {
			datos.Cedula = cedula
			m.Usuarios[i] = datos
			return datos, true
		}
	}
	return modelos.Usuario{}, false
}

func (m *MemoriaAcceso) BorrarUsuario(cedula string) bool {
	// m.mu.Lock()
	// defer m.mu.Unlock()

	for i, u := range m.Usuarios {
		if u.Cedula == cedula {
			m.Usuarios = append(m.Usuarios[:i], m.Usuarios[i+1:]...)
			return true
		}
	}
	return false
}
