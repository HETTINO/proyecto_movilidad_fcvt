package storage

import (
	"proyecto_movilidad_fcvt/internal/models"
	"time"
)

// SeedRutas carga rutas de ejemplo en memoria
func (m *Memoria) SeedRutas() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.rutas = []models.Ruta{
		{
			ID:          1,
			Nombre:      "Ruta A - Campus",
			Descripcion: "Recorre toda la facultad de norte a sur",
		},
		{
			ID:          2,
			Nombre:      "Ruta B - Biblioteca",
			Descripcion: "Del acceso principal a la biblioteca",
		},
		{
			ID:          3,
			Nombre:      "Ruta C - Comedor",
			Descripcion: "Acceso comedor y cafetería",
		},
	}
	m.nextRutaID = 4
}

// SeedParadas carga paradas de ejemplo
func (m *Memoria) SeedParadas() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.paradas = []models.Parada{
		{
			IDParada: 1,
			Nombre:   "Entrada Principal",
			Latitud:  -0.9203,
			Longitud: -80.7346,
		},
		{
			IDParada: 2,
			Nombre:   "Biblioteca",
			Latitud:  -0.9210,
			Longitud: -80.7350,
		},
		{
			IDParada: 3,
			Nombre:   "Comedor",
			Latitud:  -0.9215,
			Longitud: -80.7340,
		},
		{
			IDParada: 4,
			Nombre:   "Parqueadero",
			Latitud:  -0.9200,
			Longitud: -80.7330,
		},
	}
	m.nextParadaID = 5
}

// SeedCarritos carga carritos de ejemplo
func (m *Memoria) SeedCarritos() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.carritos = []models.Carrito{
		{
			ID:            1,
			NombreCarrito: "Carrito 1",
			Capacidad:     8,
			Estado:        "disponible",
			RutaID:        1,
		},
		{
			ID:            2,
			NombreCarrito: "Carrito 2",
			Capacidad:     8,
			Estado:        "disponible",
			RutaID:        2,
		},
		{
			ID:            3,
			NombreCarrito: "Carrito 3",
			Capacidad:     6,
			Estado:        "mantenimiento",
			RutaID:        3,
		},
		{
			ID:            4,
			NombreCarrito: "Carrito 4",
			Capacidad:     8,
			Estado:        "disponible",
			RutaID:        1,
		},
	}
	m.nextCarritoID = 5
}

// SeedLocaciones carga ubicaciones actuales de carritos
func (m *Memoria) SeedLocaciones() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.locaciones = []models.Locacion{
		{
			ID:        1,
			Latitud:   -0.9203,
			Longitud:  -80.7346,
			TimeStamp: time.Now().Add(-5 * time.Minute),
			CarritoID: 1,
		},
		{
			ID:        2,
			Latitud:   -0.9210,
			Longitud:  -80.7350,
			TimeStamp: time.Now().Add(-3 * time.Minute),
			CarritoID: 2,
		},
		{
			ID:        3,
			Latitud:   -0.9215,
			Longitud:  -80.7340,
			TimeStamp: time.Now().Add(-1 * time.Minute),
			CarritoID: 4,
		},
	}
	m.nextLocacionID = 4
}

// SeedSolicitudes carga solicitudes de ejemplo
func (m *Memoria) SeedSolicitudes() {
	m.mu.Lock()
	defer m.mu.Unlock()

	cedulaEst1 := "1234567890"
	carritoID1 := 1
	cedulaEst2 := "0987654321"

	m.solicitudes = []models.Solicitud{
		{
			ID:            1,
			CedulaUsuario: cedulaEst1,
			CantPersonas:  3,
			ParadaOrigen:  1,
			PuntoDestino:  "Biblioteca",
			Estado:        "completada",
			IDCarrito:     &carritoID1,
		},
		{
			ID:            2,
			CedulaUsuario: cedulaEst2,
			CantPersonas:  2,
			ParadaOrigen:  2,
			PuntoDestino:  "Comedor",
			Estado:        "pendiente",
			IDCarrito:     nil,
		},
		{
			ID:            3,
			CedulaUsuario: cedulaEst1,
			CantPersonas:  1,
			ParadaOrigen:  4,
			PuntoDestino:  "Entrada",
			Estado:        "asignada",
			IDCarrito:     &carritoID1,
		},
	}
	m.nextSolicitudID = 4
}