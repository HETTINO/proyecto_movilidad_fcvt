package servicetransporte

import "errors"

var (
	ErrNoEncontrado   = errors.New("recurso no encontrado")
	ErrCampoRequerido = errors.New("campo requerido faltando")
	ErrDatosInvalidos = errors.New("datos inválidos")
)
