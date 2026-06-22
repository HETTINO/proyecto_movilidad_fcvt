package service_acceso

import (
	"proyecto_movilidad_fcvt/internal/modelos"
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
		return modelos.PuntoDeAcceso{}, false, ErrNoEncontrado
	}
	return actualizado, true, nil
}

func (s *PuntoAccesoService) Borrar(id int) error {
	if !s.repo.BorrarPuntoAcceso(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validarPuntoAcceso(p modelos.PuntoDeAcceso) error {
	if p.ID == 0 && p.Ubicacion == "" { // Usa campos estándar como ID o Ubicación si existen
		return ErrCampoRequerido
	}
	return nil
}
