package servicetransporte_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	modelos "proyecto_movilidad_fcvt/internal/modelos"
	st "proyecto_movilidad_fcvt/internal/service/service_transporte"
)

// =========================================================
// TESTS — RutaService
// =========================================================

func TestRutaService_Crear(t *testing.T) {
	casos := []struct {
		nombre      string
		entrada     modelos.Ruta
		debeFallar  bool
	}{
		{
			nombre: "nombre vacío -> no persiste",
			entrada: modelos.Ruta{
				Nombre:      "",
				Descripcion: "Sin nombre",
			},
			debeFallar: true,
		},
		{
			nombre: "ruta válida -> se persiste",
			entrada: modelos.Ruta{
				Nombre:      "Ruta Norte",
				Descripcion: "Recorrido por el norte",
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
				repo.On("CrearRuta", c.entrada).Return(esperado)
			}

			svc := st.NewRutaService(repo)
			creado, err := svc.Crear(c.entrada)

			if c.debeFallar {
				assert.Error(t, err)
				repo.AssertNotCalled(t, "CrearRuta")
			} else {
				assert.NoError(t, err)
				assert.Equal(t, 1, creado.ID)
				assert.Equal(t, c.entrada.Nombre, creado.Nombre)
			}
			repo.AssertExpectations(t)
		})
	}
}

func TestRutaService_Obtener_NoEncontrado(t *testing.T) {
	repo := new(almacenMock)
	repo.On("BuscarRutaPorID", 999).Return(modelos.Ruta{}, false)

	svc := st.NewRutaService(repo)
	_, ok := svc.Obtener(999)

	assert.False(t, ok)
	repo.AssertExpectations(t)
}

func TestRutaService_Obtener_Exitoso(t *testing.T) {
	repo := new(almacenMock)
	esperado := modelos.Ruta{
		ID:     1,
		Nombre: "Ruta Sur",
	}

	repo.On("BuscarRutaPorID", 1).Return(esperado, true)

	svc := st.NewRutaService(repo)
	resultado, ok := svc.Obtener(1)

	assert.True(t, ok)
	assert.Equal(t, "Ruta Sur", resultado.Nombre)
	repo.AssertExpectations(t)
}
func TestRutaService_Listar(t *testing.T) {
	repo := new(almacenMock)
	esperadas := []modelos.Ruta{
		{ID: 1, Nombre: "Ruta Norte"},
		{ID: 2, Nombre: "Ruta Sur"},
	}

	repo.On("ListarRutas").Return(esperadas)

	svc := st.NewRutaService(repo)
	resultado := svc.Listar()

	assert.Len(t, resultado, 2)
	assert.Equal(t, "Ruta Norte", resultado[0].Nombre)
	repo.AssertExpectations(t)
}

func TestRutaService_Actualizar_Exitoso(t *testing.T) {
	repo := new(almacenMock)
	datos := modelos.Ruta{Nombre: "Ruta Actualizada", Descripcion: "Nueva desc"}
	esperado := datos
	esperado.ID = 1

	repo.On("ActualizarRuta", 1, datos).Return(esperado, true)

	svc := st.NewRutaService(repo)
	actualizado, encontrado, err := svc.Actualizar(1, datos)

	assert.NoError(t, err)
	assert.True(t, encontrado)
	assert.Equal(t, "Ruta Actualizada", actualizado.Nombre)
	repo.AssertExpectations(t)
}


func TestRutaService_Actualizar_DatosInvalidos(t *testing.T) {
	repo := new(almacenMock)
	datos := modelos.Ruta{Nombre: ""} // inválido

	svc := st.NewRutaService(repo)
	_, encontrado, err := svc.Actualizar(1, datos)

	assert.Error(t, err)
	assert.False(t, encontrado)
	repo.AssertNotCalled(t, "ActualizarRuta")
}

func TestRutaService_Borrar_Exitoso(t *testing.T) {
	repo := new(almacenMock)
	repo.On("BorrarRuta", 1).Return(true)

	svc := st.NewRutaService(repo)
	err := svc.Borrar(1)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestRutaService_Borrar_NoEncontrado(t *testing.T) {
	repo := new(almacenMock)
	repo.On("BorrarRuta", 999).Return(false)

	svc := st.NewRutaService(repo)
	err := svc.Borrar(999)

	assert.Error(t, err)
	repo.AssertExpectations(t)
}