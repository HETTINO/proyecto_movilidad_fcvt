package storage_acceso

import "proyecto_movilidad_fcvt/internal/modelos"

func (a *Almacen) GuardarVehiculo(v *modelos.Vehiculo) error {
	return a.DB.Create(v).Error
}

func (a *Almacen) BuscarVehiculo(placa string) (*modelos.Vehiculo, error) {
	var v modelos.Vehiculo
	err := a.DB.Where("placa = ?", placa).First(&v).Error
	if err != nil {
		return nil, err
	}
	return &v, nil
}
