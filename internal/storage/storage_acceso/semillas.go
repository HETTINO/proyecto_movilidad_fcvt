package storage_acceso

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	"time"
)

func (m *MemoriaAcceso) SeedUsuarios() {
	// m.mu.Lock() // Descomenta esta línea si tu estructura MemoriaAcceso tiene un mutex 'mu' como la de tu compañero
	// defer m.mu.Unlock()

	m.Usuarios = []modelos.Usuario{
		{
			Cedula:     "1312345678",
			Nombre:     "Juan Pérez",
			Contrasena: "hash_password_123",
			Email:      "juan.perez@live.uleam.edu.ec",
			Rol:        "Estudiante",
		},
		{
			Cedula:     "1309876543",
			Nombre:     "Ing. María Loor",
			Contrasena: "hash_docente_2026",
			Email:      "maria.loor@uleam.edu.ec",
			Rol:        "Docente",
		},
	}
}

func (m *MemoriaAcceso) SeedVehiculos() {
	// m.mu.Lock()
	// defer m.mu.Unlock()

	m.Vehiculos = []modelos.Vehiculo{
		{
			Placa:        "MAN-1234",
			IDUsuario:    "1312345678",
			TipoVehiculo: "Automovil",
			Marca:        "Chevrolet",
			Modelo:       "Sail",
			Color:        "Negro",
			Año:          2022,
		},
		{
			Placa:        "M-XYZ789",
			IDUsuario:    "1309876543",
			TipoVehiculo: "Motocicleta",
			Marca:        "Honda",
			Modelo:       "CB190R",
			Color:        "Rojo",
			Año:          2025,
		},
	}
}

func (m *MemoriaAcceso) SeedPuntosDeAcceso() {
	// m.mu.Lock()
	// defer m.mu.Unlock()

	m.PuntosDeAcceso = []modelos.PuntoDeAcceso{
		{
			ID:         1,
			Frecuencia: "Alta",
			Ubicacion:  "Garita Principal FCVT",
		},
		{
			ID:         2,
			Frecuencia: "Media",
			Ubicacion:  "Acceso Posterior Canchas",
		},
	}
}

func (m *MemoriaAcceso) SeedAccesos() {
	// m.mu.Lock()
	// defer m.mu.Unlock()

	tiempoEntrada1 := time.Now().Add(-3 * time.Hour)
	tiempoSalida1 := time.Now().Add(-1 * time.Hour)

	tiempoEntrada2 := time.Now().Add(-30 * time.Minute)

	m.Accesos = []modelos.Acceso{
		{
			ID:            1,
			PlacaVehiculo: "MAN-1234",
			PuntoAccesoID: 1,
			TiempoEntrada: tiempoEntrada1,
			TiempoSalida:  &tiempoSalida1,
			Estado:        "salido",
			Observaciones: "Ingreso matutino sin novedades",
		},
		{
			ID:            2,
			PlacaVehiculo: "M-XYZ789",
			PuntoAccesoID: 1,
			TiempoEntrada: tiempoEntrada2,
			TiempoSalida:  nil, // Sigue dentro del campus
			Estado:        "ENTRADA",
			Observaciones: "Vehículo permanece en instalaciones",
		},
	}
}
