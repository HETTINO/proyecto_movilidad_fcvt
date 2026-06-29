package service_acceso_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"proyecto_movilidad_fcvt/internal/modelos"
	sa "proyecto_movilidad_fcvt/internal/service/service_acceso"
)

func TestAccesoService_Crear(t *testing.T) {

	repo := new(accesoRepoMock)

	repo.
		On("CrearAcceso", mock.Anything).
		Return(func(a modelos.Acceso) modelos.Acceso {
			a.ID = 1
			return a
		})

	svc := sa.NewAccesoService(repo)

	input := modelos.Acceso{
		PlacaVehiculo: "ABC123",
		Estado:        "activo",
	}

	res := svc.Crear(input) // 👈 IMPORTANTE: SOLO 1 valor

	assert.Equal(t, 1, res.ID)
	assert.Equal(t, "ABC123", res.PlacaVehiculo)

	repo.AssertExpectations(t)
}
