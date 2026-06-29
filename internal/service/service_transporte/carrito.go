package servicetransporte

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_transporte"
)

type CarritoService struct {
	repo storage.Almacen
}

func NewCarritoService(repo storage.Almacen) *CarritoService {
	return &CarritoService{repo: repo}
}

func (s *CarritoService) Listar() []modelos.Carrito {
	return s.repo.ListarCarritos()
}

func (s *CarritoService) Obtener(id int) (modelos.Carrito, bool) {
	return s.repo.BuscarCarritoPorID(id)
}

func (s *CarritoService) Crear(c modelos.Carrito) (modelos.Carrito, error) {
	if err := validarCarrito(c); err != nil {
		return modelos.Carrito{}, err
	}
	return s.repo.CrearCarrito(c), nil
}

func (s *CarritoService) Actualizar(id int, datos modelos.Carrito) (modelos.Carrito, bool, error) {
	if err := validarCarrito(datos); err != nil {
		return modelos.Carrito{}, false, err
	}
	actualizado, encontrado := s.repo.ActualizarCarrito(id, datos)
	if !encontrado {
		return modelos.Carrito{}, false, ErrNoEncontrado
	}
	return actualizado, true, nil
}

func (s *CarritoService) Borrar(id int) error {
	if !s.repo.BorrarCarrito(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validarCarrito(c modelos.Carrito) error {
	if c.NombreCarrito == "" {
		return ErrCampoRequerido
	}
	if c.Capacidad <= 0 {
		return ErrDatosInvalidos
	}
	return nil
}
