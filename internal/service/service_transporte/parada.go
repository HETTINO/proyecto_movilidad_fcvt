package servicetransporte

import (
	modelos "proyecto_movilidad_fcvt/internal/models"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_transporte"
)

type ParadaService struct {
	repo storage.Almacen
}

func NewParadaService(repo storage.Almacen) *ParadaService {
	return &ParadaService{repo: repo}
}

func (s *ParadaService) Listar() []modelos.Parada {
	return s.repo.ListarParadas()
}

func (s *ParadaService) Obtener(id int) (modelos.Parada, bool) {
	return s.repo.BuscarParadaPorID(id)
}

func (s *ParadaService) Crear(p modelos.Parada) (modelos.Parada, error) {
	if err := validarParada(p); err != nil {
		return modelos.Parada{}, err
	}
	return s.repo.CrearParada(p), nil
}

func (s *ParadaService) Actualizar(id int, datos modelos.Parada) (modelos.Parada, bool, error) {
	if err := validarParada(datos); err != nil {
		return modelos.Parada{}, false, err
	}
	actualizado, encontrado := s.repo.ActualizarParada(id, datos)
	if !encontrado {
		return modelos.Parada{}, false, ErrNoEncontrado
	}
	return actualizado, true, nil
}

func (s *ParadaService) Borrar(id int) error {
	if !s.repo.BorrarParada(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validarParada(p modelos.Parada) error {
	if p.Nombre == "" {
		return ErrCampoRequerido
	}
	if p.Latitud == 0 || p.Longitud == 0 {
		return ErrDatosInvalidos
	}
	return nil
}
