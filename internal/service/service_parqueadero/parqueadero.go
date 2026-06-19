package serviceparqueadero

import (
	"strings"

	"proyecto_movilidad_fcvt/internal/modelos"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_parqueadero"
)

type ParqueaderoService struct {
	repo storage.ParqueaderoRepository
}

func NewParqueaderoService(repo storage.ParqueaderoRepository) *ParqueaderoService {
	return &ParqueaderoService{repo: repo}
}

func (s *ParqueaderoService) Listar() []modelos.Parqueadero {
	return s.repo.ListarParqueaderos()
}

func (s *ParqueaderoService) Obtener(id int) (modelos.Parqueadero, bool) {
	return s.repo.BuscarParqueaderoPorID(id)
}

func (s *ParqueaderoService) Crear(p modelos.Parqueadero) (modelos.Parqueadero, error) {
	if err := validarParqueadero(p); err != nil {
		return modelos.Parqueadero{}, err
	}
	return s.repo.CrearParqueadero(p), nil
}

func (s *ParqueaderoService) Actualizar(id int, datos modelos.Parqueadero) (modelos.Parqueadero, bool, error) {
	if err := validarParqueadero(datos); err != nil {
		return modelos.Parqueadero{}, false, err
	}
	actualizado, encontrado := s.repo.ActualizarParqueadero(id, datos)
	if !encontrado {
		return modelos.Parqueadero{}, false, ErrNoEncontrado
	}
	return actualizado, true, nil
}

func (s *ParqueaderoService) Borrar(id int) error {
	if !s.repo.BorrarParqueadero(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validarParqueadero(p modelos.Parqueadero) error {
	if strings.TrimSpace(p.Nombre) == "" {
		return ErrNombreVacio
	}
	return nil
}
