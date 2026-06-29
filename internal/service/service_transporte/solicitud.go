package servicetransporte

import (
	modelos "proyecto_movilidad_fcvt/internal/modelos"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_transporte"
)

type SolicitudService struct {
	repo storage.Almacen
}

func NewSolicitudService(repo storage.Almacen) *SolicitudService {
	return &SolicitudService{repo: repo}
}

func (s *SolicitudService) Listar() []modelos.Solicitud {
	return s.repo.ListarSolicitudes()
}

func (s *SolicitudService) Obtener(id int) (modelos.Solicitud, bool) {
	return s.repo.BuscarSolicitudPorID(id)
}

func (s *SolicitudService) Crear(sol modelos.Solicitud) (modelos.Solicitud, error) {
	if err := validarSolicitud(sol); err != nil {
		return modelos.Solicitud{}, err
	}
	sol.Estado = "pendiente"
	return s.repo.CrearSolicitud(sol), nil
}

func (s *SolicitudService) Actualizar(id int, datos modelos.Solicitud) (modelos.Solicitud, bool, error) {
	if err := validarSolicitud(datos); err != nil {
		return modelos.Solicitud{}, false, err
	}
	actualizado, encontrado := s.repo.ActualizarSolicitud(id, datos)
	if !encontrado {
		return modelos.Solicitud{}, false, ErrNoEncontrado
	}
	return actualizado, true, nil
}

func (s *SolicitudService) Borrar(id int) error {
	if !s.repo.BorrarSolicitud(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validarSolicitud(sol modelos.Solicitud) error {
	if sol.CedulaUsuario == "" {
		return ErrCampoRequerido
	}
	if sol.CantPersonas <= 0 {
		return ErrDatosInvalidos
	}
	if sol.PuntoDestino == "" {
		return ErrCampoRequerido
	}
	return nil
}
