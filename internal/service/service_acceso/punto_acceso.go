package service_acceso

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	service "proyecto_movilidad_fcvt/internal/service"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_acceso"
)

type PuntoAccesoService struct {
	repo storage.PuntoAccesoRepository
}

func NewPuntoAccesoService(repo storage.PuntoAccesoRepository) *PuntoAccesoService {
	return &PuntoAccesoService{repo: repo}
}

func (s *PuntoAccesoService) Listar() []modelos.PuntoDeAcceso {
	return s.repo.ListarPuntosAcceso()
}

func (s *PuntoAccesoService) Obtener(id int) (modelos.PuntoDeAcceso, bool) {
	return s.repo.BuscarPuntoAccesoPorID(id)
}

func (s *PuntoAccesoService) Crear(p modelos.PuntoDeAcceso) (modelos.PuntoDeAcceso, error) {
	if err := validarPuntoAcceso(p); err != nil {
		return modelos.PuntoDeAcceso{}, err
	}
	return s.repo.CrearPuntoAcceso(p), nil
}

func (s *PuntoAccesoService) Actualizar(id int, datos modelos.PuntoDeAcceso) (modelos.PuntoDeAcceso, bool, error) {
	if err := validarPuntoAcceso(datos); err != nil {
		return modelos.PuntoDeAcceso{}, false, err
	}

	actualizado, encontrado := s.repo.ActualizarPuntoAcceso(id, datos)
	if !encontrado {
		return modelos.PuntoDeAcceso{}, false, service.ErrNoEncontrado
	}

	return actualizado, true, nil
}

func (s *PuntoAccesoService) Borrar(id int) error {
	if !s.repo.BorrarPuntoAcceso(id) {
		return service.ErrNoEncontrado
	}
	return nil
}

// VALIDACIÓN CORRECTA SEGÚN TU MODELO
func validarPuntoAcceso(p modelos.PuntoDeAcceso) error {
	if p.Ubicacion == "" {
		return service.ErrCampoRequerido
	}
	return nil
}
