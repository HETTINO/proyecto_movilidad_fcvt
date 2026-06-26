package serviceparqueadero

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	service "proyecto_movilidad_fcvt/internal/service"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_parqueadero"
)

type EspacioService struct {
	repo storage.EspacioRepository
}

func NewEspacioService(repo storage.EspacioRepository) *EspacioService {
	return &EspacioService{repo: repo}
}

func (s *EspacioService) Listar() []modelos.Espacio {
	return s.repo.ListarEspacios()
}

func (s *EspacioService) Obtener(id int) (modelos.Espacio, bool) {
	return s.repo.BuscarEspacioPorID(id)
}

func (s *EspacioService) Crear(e modelos.Espacio) (modelos.Espacio, error) {
	if err := validarEspacio(e); err != nil {
		return modelos.Espacio{}, err
	}
	return s.repo.CrearEspacio(e), nil
}

func (s *EspacioService) Actualizar(id int, datos modelos.Espacio) (modelos.Espacio, bool, error) {
	if err := validarEspacio(datos); err != nil {
		return modelos.Espacio{}, false, err
	}
	actualizado, encontrado := s.repo.ActualizarEspacio(id, datos)
	if !encontrado {
		return modelos.Espacio{}, false, service.ErrNoEncontrado
	}
	return actualizado, true, nil
}

func (s *EspacioService) Borrar(id int) error {
	if !s.repo.BorrarEspacio(id) {
		return service.ErrNoEncontrado
	}
	return nil
}

func validarEspacio(e modelos.Espacio) error {
	if e.IDParqueadero == 0 {
		return service.ErrCampoRequerido
	}
	return nil
}
