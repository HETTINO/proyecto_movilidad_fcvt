package storage_acceso

import "proyecto_movilidad_fcvt/internal/modelos"

// InsertarSemillasAcceso inyecta datos iniciales para pruebas del equipo
func (a *Almacen) InsertarSemillasAcceso() error {
	var conteo int64
	a.DB.Model(&modelos.PuntoDeAcceso{}).Count(&conteo)

	if conteo == 0 {
		puntoInicial := modelos.PuntoDeAcceso{
			ID:         1,
			Frecuencia: "Alta", // Usa Frecuencia, tal como está en tu modelo
			Ubicacion:  "Garita Principal",
		}
		return a.DB.Create(&puntoInicial).Error
	}
	return nil
}
