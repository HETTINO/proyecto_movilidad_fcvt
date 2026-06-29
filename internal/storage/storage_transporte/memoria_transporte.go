package storage

import (
	modelos "proyecto_movilidad_fcvt/internal/modelos"
	"sync"
)

type Memoria struct {
	rutas       []modelos.Ruta
	paradas     []modelos.Parada
	carritos    []modelos.Carrito
	locaciones  []modelos.Locacion
	solicitudes []modelos.Solicitud

	nextRutaID      int
	nextParadaID    int
	nextCarritoID   int
	nextLocacionID  int
	nextSolicitudID int

	mu sync.Mutex
}

func NuevaMemoria() *Memoria {
	return &Memoria{
		rutas:       []modelos.Ruta{},
		paradas:     []modelos.Parada{},
		carritos:    []modelos.Carrito{},
		locaciones:  []modelos.Locacion{},
		solicitudes: []modelos.Solicitud{},

		nextRutaID:      1,
		nextParadaID:    1,
		nextCarritoID:   1,
		nextLocacionID:  1,
		nextSolicitudID: 1,
	}
}

var _ Almacen = (*Memoria)(nil)
