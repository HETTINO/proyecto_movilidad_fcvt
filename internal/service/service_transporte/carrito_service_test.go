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
func TestCarritoService_Listar(t *testing.T) {
	repo := new(almacenMock)
	esperados := []modelos.Carrito{
		{ID: 1, NombreCarrito: "Carrito 1", Capacidad: 8},
		{ID: 2, NombreCarrito: "Carrito 2", Capacidad: 6},
	}

	repo.On("ListarCarritos").Return(esperados)

	svc := st.NewCarritoService(repo)
	resultado := svc.Listar()

	assert.Len(t, resultado, 2)
	assert.Equal(t, "Carrito 1", resultado[0].NombreCarrito)
	repo.AssertExpectations(t)
}

func TestCarritoService_Actualizar_Exitoso(t *testing.T) {
	repo := new(almacenMock)
	datos := modelos.Carrito{NombreCarrito: "Carrito Actualizado", Capacidad: 10}
	esperado := datos
	esperado.ID = 1

	repo.On("ActualizarCarrito", 1, datos).Return(esperado, true)

	svc := st.NewCarritoService(repo)
	actualizado, encontrado, err := svc.Actualizar(1, datos)

	assert.NoError(t, err)
	assert.True(t, encontrado)
	assert.Equal(t, "Carrito Actualizado", actualizado.NombreCarrito)
	repo.AssertExpectations(t)
}

func TestCarritoService_Actualizar_NoEncontrado(t *testing.T) {
	repo := new(almacenMock)
	datos := modelos.Carrito{NombreCarrito: "Carrito X", Capacidad: 4}

	repo.On("ActualizarCarrito", 999, datos).Return(modelos.Carrito{}, false)

	svc := st.NewCarritoService(repo)
	_, encontrado, err := svc.Actualizar(999, datos)

	assert.Error(t, err)
	assert.False(t, encontrado)
	repo.AssertExpectations(t)
}

func TestCarritoService_Actualizar_DatosInvalidos(t *testing.T) {
	repo := new(almacenMock)
	datos := modelos.Carrito{NombreCarrito: "Válido", Capacidad: 0} // inválido

	svc := st.NewCarritoService(repo)
	_, encontrado, err := svc.Actualizar(1, datos)

	assert.Error(t, err)
	assert.False(t, encontrado)
	repo.AssertNotCalled(t, "ActualizarCarrito")
}

func TestCarritoService_Borrar_Exitoso(t *testing.T) {
	repo := new(almacenMock)
	repo.On("BorrarCarrito", 1).Return(true)

	svc := st.NewCarritoService(repo)
	err := svc.Borrar(1)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestCarritoService_Borrar_NoEncontrado(t *testing.T) {
	repo := new(almacenMock)
	repo.On("BorrarCarrito", 999).Return(false)

	svc := st.NewCarritoService(repo)
	err := svc.Borrar(999)

	assert.Error(t, err)
	repo.AssertExpectations(t)
}
