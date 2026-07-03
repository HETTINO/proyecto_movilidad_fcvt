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