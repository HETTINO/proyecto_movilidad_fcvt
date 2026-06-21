package storage_acceso

import "proyecto_movilidad_fcvt/internal/modelos"

func (a *Almacen) CrearPunto(p *modelos.PuntoDeAcceso) error {
	return a.DB.Create(p).Error
}

func (a *Almacen) ObtenerPuntos() ([]modelos.PuntoDeAcceso, error) {
	var puntos []modelos.PuntoDeAcceso
	err := a.DB.Find(&puntos).Error
	return puntos, err
}
