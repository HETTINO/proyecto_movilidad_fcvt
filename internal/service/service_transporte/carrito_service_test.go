package servicetransporte_test

import (
	"github.com/stretchr/testify/assert"
	modelos "proyecto_movilidad_fcvt/internal/modelos"
	st "proyecto_movilidad_fcvt/internal/service/service_transporte"
	"testing"
)

// =========================================================
// TESTS — CarritoService
// =========================================================

func TestCarritoService_Crear(t *testing.T) {
	casos := []struct {
		nombre     string
		entrada    modelos.Carrito
		debeFallar bool
	}{
		{
			nombre: "nombre vacío -> no persiste",
			entrada: modelos.Carrito{
				NombreCarrito: "",
				Capacidad:     5,
			},
			debeFallar: true,
		},
		{
			nombre: "capacidad cero -> no persiste",
			entrada: modelos.Carrito{
				NombreCarrito: "Carrito Rectorado",
				Capacidad:     0,
			},
			debeFallar: true,
		},
		{
			nombre: "carrito válido -> se persiste",
			entrada: modelos.Carrito{
				NombreCarrito: "Carrito 1 - Rectorado",
				Capacidad:     3,
			},
			debeFallar: false,
		},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			repo := new(almacenMock)

			if !c.debeFallar {
				esperado := c.entrada
				esperado.ID = 1
				repo.On("CrearCarrito", c.entrada).Return(esperado)
			}

			svc := st.NewCarritoService(repo)
			creado, err := svc.Crear(c.entrada)

			if c.debeFallar {
				assert.Error(t, err)
				repo.AssertNotCalled(t, "CrearCarrito")
			} else {
				assert.NoError(t, err)
				assert.Equal(t, 1, creado.ID)
				assert.Equal(t, c.entrada.NombreCarrito, creado.NombreCarrito)
			}
			repo.AssertExpectations(t)
		})
	}
}

func TestCarritoService_Obtener_NoEncontrado(t *testing.T) {
	repo := new(almacenMock)
	repo.On("BuscarCarritoPorID", 999).Return(modelos.Carrito{}, false)

	svc := st.NewCarritoService(repo)
	_, ok := svc.Obtener(999)

	assert.False(t, ok)
	repo.AssertExpectations(t)
}

func TestCarritoService_Obtener_Exitoso(t *testing.T) {
	repo := new(almacenMock)
	esperado := modelos.Carrito{
		ID:            1,
		NombreCarrito: "Carrito Tasty",
		Capacidad:     3,
	}

	repo.On("BuscarCarritoPorID", 1).Return(esperado, true)

	svc := st.NewCarritoService(repo)
	resultado, ok := svc.Obtener(1)

	assert.True(t, ok)
	assert.Equal(t, "Carrito Tasty", resultado.NombreCarrito)
	repo.AssertExpectations(t)
}
