package service_acceso

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_acceso"
)

type VehiculoService struct {
	repo storage.VehiculoRepository
}

func NewVehiculoService(repo storage.VehiculoRepository) *VehiculoService {
	return &VehiculoService{repo: repo}
}

func (s *VehiculoService) Listar() []modelos.Vehiculo {
	return s.repo.ListarVehiculos()
}

func (s *VehiculoService) Obtener(placa string) (modelos.Vehiculo, bool) {
	return s.repo.BuscarVehiculoPorPlaca(placa)
}

func (s *VehiculoService) Crear(v modelos.Vehiculo) (modelos.Vehiculo, error) {
	if err := validarVehiculo(v); err != nil {
		return modelos.Vehiculo{}, err
	}
	return s.repo.CrearVehiculo(v), nil
}

func (s *VehiculoService) Actualizar(placa string, datos modelos.Vehiculo) (modelos.Vehiculo, bool, error) {
	if err := validarVehiculo(datos); err != nil {
		return modelos.Vehiculo{}, false, err
	}
	actualizado, encontrado := s.repo.ActualizarVehiculo(placa, datos)
	if !encontrado {
		return modelos.Vehiculo{}, false, ErrNoEncontrado
	}
	return actualizado, true, nil
}

func (s *VehiculoService) Borrar(placa string) error {
	if !s.repo.BorrarVehiculo(placa) {
		return ErrNoEncontrado
	}
	return nil
}

func validarVehiculo(v modelos.Vehiculo) error {
	if v.Placa == "" {
		return ErrCampoRequerido
	}
	return nil
}
