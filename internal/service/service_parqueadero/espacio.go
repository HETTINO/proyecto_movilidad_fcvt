package serviceparqueadero

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	service "proyecto_movilidad_fcvt/internal/service"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_parqueadero"
)

type EspacioService struct {
	repo          storage.EspacioRepository
	ocupacionRepo storage.OcupacionesRepository
}

// NewEspacioService recibe además el repositorio de Ocupaciones, necesario
// para validar integridad referencial antes de borrar un espacio.
func NewEspacioService(repo storage.EspacioRepository, ocupacionRepo storage.OcupacionesRepository) *EspacioService {
	return &EspacioService{repo: repo, ocupacionRepo: ocupacionRepo}
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
	if _, encontrado := s.repo.BuscarEspacioPorID(id); !encontrado {
		return service.ErrNoEncontrado
	}

	if activas := s.ocupacionRepo.ListarOcupacionesActivas(id); len(activas) > 0 {
		return service.ErrEspacioConOcupacionesActivas
	}

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
