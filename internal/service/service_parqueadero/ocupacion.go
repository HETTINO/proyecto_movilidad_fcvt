package serviceparqueadero

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_parqueadero"
)

type OcupacionService struct {
	repo storage.OcupacionesRepository
}

func NewOcupacionService(repo storage.OcupacionesRepository) *OcupacionService {
	return &OcupacionService{repo: repo}
}

func (s *OcupacionService) Listar() []modelos.Ocupacion {
	return s.repo.ListarOcupaciones()
}

func (s *OcupacionService) Obtener(id int) (modelos.Ocupacion, bool) {
	return s.repo.BuscarOcupacionPorID(id)
}

func (s *OcupacionService) Crear(o modelos.Ocupacion) (modelos.Ocupacion, error) {
	if err := validarOcupacion(o); err != nil {
		return modelos.Ocupacion{}, err
	}
	return s.repo.CrearOcupacion(o), nil
}

func (s *OcupacionService) Actualizar(id int, datos modelos.Ocupacion) (modelos.Ocupacion, bool, error) {
	if err := validarOcupacion(datos); err != nil {
		return modelos.Ocupacion{}, false, err
	}
	actualizado, encontrado := s.repo.ActualizarOcupacion(id, datos)
	if !encontrado {
		return modelos.Ocupacion{}, false, ErrNoEncontrado
	}
	return actualizado, true, nil
}

func (s *OcupacionService) Borrar(id int) error {
	if !s.repo.BorrarOcupacion(id) {
		return ErrNoEncontrado
	}
	return nil
}

func (s *OcupacionService) Liberar(id int) (modelos.Ocupacion, bool) {
	return s.repo.LiberarOcupacion(id)
}

func validarOcupacion(o modelos.Ocupacion) error {
	if o.IDEspacio == 0 {
		return ErrCampoRequerido
	}
	return nil
}
