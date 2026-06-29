package servicetransporte

import (
	modelos "proyecto_movilidad_fcvt/internal/models"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_transporte"
)

type RutaService struct {
	repo storage.Almacen
}

func NewRutaService(repo storage.Almacen) *RutaService {
	return &RutaService{repo: repo}
}

func (s *RutaService) Listar() []modelos.Ruta {
	return s.repo.ListarRutas()
}

func (s *RutaService) Obtener(id int) (modelos.Ruta, bool) {
	return s.repo.BuscarRutaPorID(id)
}

func (s *RutaService) Crear(r modelos.Ruta) (modelos.Ruta, error) {
	if err := validarRuta(r); err != nil {
		return modelos.Ruta{}, err
	}
	return s.repo.CrearRuta(r), nil
}

func (s *RutaService) Actualizar(id int, datos modelos.Ruta) (modelos.Ruta, bool, error) {
	if err := validarRuta(datos); err != nil {
		return modelos.Ruta{}, false, err
	}
	actualizado, encontrado := s.repo.ActualizarRuta(id, datos)
	if !encontrado {
		return modelos.Ruta{}, false, ErrNoEncontrado
	}
	return actualizado, true, nil
}

func (s *RutaService) Borrar(id int) error {
	if !s.repo.BorrarRuta(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validarRuta(r modelos.Ruta) error {
	if r.Nombre == "" {
		return ErrCampoRequerido
	}
	return nil
}
