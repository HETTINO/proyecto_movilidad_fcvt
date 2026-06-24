package storage_parqueadero

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	"sync"
)

type Memoria struct {
	parqueaderos []modelos.Parqueadero
	espacios     []modelos.Espacio
	ocupaciones  []modelos.Ocupacion

	nextParqueaderoID int
	nextEspacioID     int
	nextOcupacionID   int

	mu sync.Mutex
}

func NuevaMemoria() *Memoria {
	return &Memoria{
		parqueaderos:      []modelos.Parqueadero{},
		espacios:          []modelos.Espacio{},
		ocupaciones:       []modelos.Ocupacion{},
		nextParqueaderoID: 1,
		nextEspacioID:     1,
		nextOcupacionID:   1,
	}
}

var _ Almacen = (*Memoria)(nil)
