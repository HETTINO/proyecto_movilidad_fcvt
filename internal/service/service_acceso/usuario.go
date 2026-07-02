package service_acceso

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	service "proyecto_movilidad_fcvt/internal/service"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_acceso"
)

type UsuarioService struct {
	repo storage.UsuarioRepository
}

func NewUsuarioService(repo storage.UsuarioRepository) *UsuarioService {
	return &UsuarioService{repo: repo}
}

func (s *UsuarioService) Listar() []modelos.Usuario {
	return s.repo.ListarUsuarios()
}

func (s *UsuarioService) Obtener(cedula string) (modelos.Usuario, bool) {
	return s.repo.BuscarUsuarioPorCedula(cedula)
}

func (s *UsuarioService) Crear(u modelos.Usuario) (modelos.Usuario, error) {
	if err := validarUsuario(u); err != nil {
		return modelos.Usuario{}, err
	}
	return s.repo.CrearUsuario(u), nil
}

func (s *UsuarioService) Actualizar(cedula string, datos modelos.Usuario) (modelos.Usuario, bool, error) {
	if err := validarUsuario(datos); err != nil {
		return modelos.Usuario{}, false, err
	}

	actualizado, encontrado := s.repo.ActualizarUsuario(cedula, datos)
	if !encontrado {
		return modelos.Usuario{}, false, service.ErrNoEncontrado
	}

	return actualizado, true, nil
}

func (s *UsuarioService) Borrar(cedula string) error {
	if !s.repo.BorrarUsuario(cedula) {
		return service.ErrNoEncontrado
	}
	return nil
}

func validarUsuario(u modelos.Usuario) error {
	if u.Cedula == "" || u.Nombre == "" || u.Email == "" {
		return service.ErrCampoRequerido
	}
	return nil
}
