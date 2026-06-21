package storage_acceso

import "proyecto_movilidad_fcvt/internal/modelos"

func (a *Almacen) GuardarUsuario(u *modelos.Usuario) error {
	return a.DB.Create(u).Error
}

func (a *Almacen) BuscarUsuario(cedula string) (*modelos.Usuario, error) {
	var u modelos.Usuario
	err := a.DB.Where("cedula = ?", cedula).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}
