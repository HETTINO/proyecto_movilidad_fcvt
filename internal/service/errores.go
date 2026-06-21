package service

import "errors"

var (
	// Errores de Autenticación (Simétricos a los de parqueadero)
	ErrCredencialesInvalidas = errors.New("credenciales inválidas")
	ErrEmailenUso            = errors.New("el email ya está en uso")

	// Errores específicos para tu módulo de Accesos
	ErrUsuarioNoEncontrado  = errors.New("el usuario solicitado no existe")
	ErrVehiculoNoRegistrado = errors.New("el vehículo no está registrado en el sistema")
	ErrPuntoNoValido        = errors.New("el punto de acceso especificado no está activo o no existe")
)
