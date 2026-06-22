package service_acceso

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_acceso"
)

type AccesoService struct {
	repo storage.AccesoRepository
}

func NewAccesoService(repo storage.AccesoRepository) *AccesoService {
	return &AccesoService{repo: repo}
}

func (s *AccesoService) Listar() []modelos.Acceso {
	return s.repo.ListarAccesos()
}

func (s *AccesoService) Obtener(id int) (modelos.Acceso, bool) {
	return s.repo.BuscarAccesoPorID(id)
}

func (s *AccesoService) Crear(a modelos.Acceso) (modelos.Acceso, error) {
	if err := validarAcceso(a); err != nil {
		return modelos.Acceso{}, err
	}
	return s.repo.CrearAcceso(a), nil
}

func (s *AccesoService) Actualizar(id int, datos modelos.Acceso) (modelos.Acceso, bool, error) {
	if err := validarAcceso(datos); err != nil {
		return modelos.Acceso{}, false, err
	}
	actualizado, encontrado := s.repo.ActualizarAcceso(id, datos)
	if !encontrado {
		return modelos.Acceso{}, false, ErrNoEncontrado
	}
	return actualizado, true, nil
}

func (s *AccesoService) Borrar(id int) error {
	if !s.repo.BorrarAcceso(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validarAcceso(a modelos.Acceso) error {
	if a.PlacaVehiculo == "" {
		return ErrCampoRequerido
	}
	return nil
}
