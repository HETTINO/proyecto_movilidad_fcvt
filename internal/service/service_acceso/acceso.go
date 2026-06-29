package service_acceso

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	service "proyecto_movilidad_fcvt/internal/service"
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

func (s *AccesoService) Crear(a modelos.Acceso) modelos.Acceso {
	return s.repo.CrearAcceso(a)
}

func (s *AccesoService) Actualizar(id int, datos modelos.Acceso) (modelos.Acceso, bool, error) {
	actualizado, ok := s.repo.ActualizarAcceso(id, datos)
	if !ok {
		return modelos.Acceso{}, false, service.ErrNoEncontrado
	}
	return actualizado, true, nil
}

func (s *AccesoService) Borrar(id int) error {
	if !s.repo.BorrarAcceso(id) {
		return service.ErrNoEncontrado
	}
	return nil
}
