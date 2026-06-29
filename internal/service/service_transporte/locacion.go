package servicetransporte

import (
	"proyecto_movilidad_fcvt/internal/models"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_transporte"
)

type LocacionService struct {
	repo storage.Almacen
}

func NewLocacionService(repo storage.Almacen) *LocacionService {
	return &LocacionService{repo: repo}
}

func (s *LocacionService) Listar() []models.Locacion {
	return s.repo.ListarLocaciones()
}

func (s *LocacionService) ObtenerUltimaDelCarrito(carritoID int) (models.Locacion, bool) {
	return s.repo.ObtenerUltimaLocacionPorCarrito(carritoID)
}

func (s *LocacionService) Registrar(l models.Locacion) (models.Locacion, error) {
	if err := validarLocacion(l); err != nil {
		return models.Locacion{}, err
	}
	return s.repo.RegistrarLocacion(l), nil
}

func validarLocacion(l models.Locacion) error {
	if l.CarritoID == 0 {
		return ErrCampoRequerido
	}
	if l.Latitud == 0 || l.Longitud == 0 {
		return ErrDatosInvalidos
	}
	return nil
}