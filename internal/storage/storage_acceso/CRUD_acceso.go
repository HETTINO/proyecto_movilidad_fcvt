package storage_acceso

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	"time"
)

func (a *Almacen) RegistrarEntrada(acceso *modelos.Acceso) error {
	acceso.TiempoEntrada = time.Now()
	return a.DB.Create(acceso).Error
}

func (a *Almacen) RegistrarSalida(placa string) error {
	ahora := time.Now()
	return a.DB.Model(&modelos.Acceso{}).
		Where("placa_vehiculo = ? AND tiempo_salida IS NULL", placa).
		Updates(map[string]interface{}{
			"tiempo_salida": &ahora,
			"estado":        "salido",
		}).Error
}
