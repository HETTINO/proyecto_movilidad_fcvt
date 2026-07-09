package servicetransporte_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	modelos "proyecto_movilidad_fcvt/internal/modelos"
	st "proyecto_movilidad_fcvt/internal/service/service_transporte"
)

// =========================================================
// TESTS — LocacionService
// =========================================================

func TestLocacionService_Registrar(t *testing.T) {
	casos := []struct {
		nombre     string
		entrada    modelos.Locacion
		debeFallar bool
	}{
		{
			nombre: "carrito ID vacío -> falla",
			entrada: modelos.Locacion{
				Latitud:  -0.950,
				Longitud: -80.750,
				CarritoID: 0,
			},
			debeFallar: true,
		},
		{
			nombre: "coordenadas cero -> falla",
			entrada: modelos.Locacion{
				Latitud:   0,
				Longitud:  0,
				CarritoID: 1,
			},
			debeFallar: true,
		},
		{
			nombre: "locación válida -> éxito",
			entrada: modelos.Locacion{
				Latitud:   -0.950,
				Longitud:  -80.750,
				CarritoID: 1,
			},
			debeFallar: false,
		},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			repo := new(almacenMock)

			if !c.debeFallar {
				repo.On("RegistrarLocacion", c.entrada).Return(c.entrada)
			}

			svc := st.NewLocacionService(repo)
			registrada, err := svc.Registrar(c.entrada)

			if c.debeFallar {
				assert.Error(t, err)
				repo.AssertNotCalled(t, "RegistrarLocacion")
			} else {
				assert.NoError(t, err)
				assert.Equal(t, c.entrada.Latitud, registrada.Latitud)
			}
			repo.AssertExpectations(t)
		})
	}
}

func TestLocacionService_ObtenerUltima_Exitoso(t *testing.T) {
	repo := new(almacenMock)
	esperada := modelos.Locacion{
		Latitud:   -0.950,
		Longitud:  -80.750,
		CarritoID: 1,
	}

	repo.On("ObtenerUltimaLocacionPorCarrito", 1).Return(esperada, true)

	svc := st.NewLocacionService(repo)
	resultado, ok := svc.ObtenerUltimaDelCarrito(1)

	assert.True(t, ok)
	assert.Equal(t, -0.950, resultado.Latitud)
	repo.AssertExpectations(t)
}
func TestLocacionService_ObtenerUltima_NoEncontrado(t *testing.T) {
	repo := new(almacenMock)

	repo.On("ObtenerUltimaLocacionPorCarrito", 999).Return(modelos.Locacion{}, false)

	svc := st.NewLocacionService(repo)
	_, ok := svc.ObtenerUltimaDelCarrito(999)

	assert.False(t, ok)
	repo.AssertExpectations(t)
}

func TestLocacionService_Listar(t *testing.T) {
	repo := new(almacenMock)
	esperadas := []modelos.Locacion{
		{ID: 1, CarritoID: 1, Latitud: -0.950, Longitud: -80.750},
		{ID: 2, CarritoID: 2, Latitud: -0.921, Longitud: -80.735},
	}

	repo.On("ListarLocaciones").Return(esperadas)

	svc := st.NewLocacionService(repo)
	resultado := svc.Listar()

	assert.Len(t, resultado, 2)
	assert.Equal(t, 1, resultado[0].CarritoID)
	repo.AssertExpectations(t)
}
 