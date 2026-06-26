package service

import "errors"

var (
	ErrNombreVacio           = errors.New("nombre es requerido")
	ErrPrecioNegativo        = errors.New("precio no puede ser negativo")
	ErrNoEncontrado          = errors.New("registro no encontrado")
	ErrEmailenUso            = errors.New("email ya está en uso")
	ErrCredencialesInvalidas = errors.New("credenciales inválidas")
	ErrTokenInvalido         = errors.New("token inválido")
	ErrEmailVacio            = errors.New("email es requerido")
	ErrPasswordVacio         = errors.New("contraseña es requerida")
	ErrCampoRequerido        = errors.New("campo requerido faltante")
	ErrCapacidadInvalida     = errors.New("la capacidad debe ser mayor a cero")
)
