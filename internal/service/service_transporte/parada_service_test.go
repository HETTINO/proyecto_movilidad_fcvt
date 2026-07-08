package servicetransporte_test

import (
	"github.com/stretchr/testify/assert"
	modelos "proyecto_movilidad_fcvt/internal/modelos"
	st "proyecto_movilidad_fcvt/internal/service/service_transporte"
	"testing"
)

func TestParadaService_Crear(t *testing.T) {
	casos := []struct {
		nombre        string
		entrada       modelos.Parada
		debeFallar    bool
		debePersistir bool
	}{
		{
			nombre:        "nombre vacío -> falla",
			entrada:       modelos.Parada{Nombre: "", Latitud: -0.950, Longitud: -80.750},
			debeFallar:    true,
			debePersistir: false,
		},
		{
			nombre:        "coordenadas cero -> falla",
			entrada:       modelos.Parada{Nombre: "Parada 1", Latitud: 0, Longitud: 0},
			debeFallar:    true,
			debePersistir: false,
		},
		{
			nombre:        "parada válida -> se persiste",
			entrada:       modelos.Parada{Nombre: "Parada ULEAM", Latitud: -0.950, Longitud: -80.750},
			debeFallar:    false,
			debePersistir: true,
		},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			repo := new(almacenMock)

			if c.debePersistir {
				guardada := c.entrada
				guardada.IDParada = 1
				repo.On("CrearParada", c.entrada).Return(guardada)
			}

			svc := st.NewParadaService(repo)
			creada, err := svc.Crear(c.entrada)

			if c.debeFallar {
				assert.Error(t, err)
				repo.AssertNotCalled(t, "CrearParada")
			} else {
				assert.NoError(t, err)
				assert.Equal(t, 1, creada.IDParada)
				repo.AssertCalled(t, "CrearParada", c.entrada)
			}
		})
	}
}

func TestParadaService_Obtener_NoEncontrado(t *testing.T) {
	repo := new(almacenMock)
	repo.On("BuscarParadaPorID", 999).Return(modelos.Parada{}, false)
	svc := st.NewParadaService(repo)

	_, ok := svc.Obtener(999)

	assert.False(t, ok)
	repo.AssertExpectations(t)
}

func TestParadaService_Obtener_Exitoso(t *testing.T) {
	repo := new(almacenMock)
	esperada := modelos.Parada{IDParada: 1, Nombre: "Parada Paraninfo"}

	repo.On("BuscarParadaPorID", 1).Return(esperada, true)
	svc := st.NewParadaService(repo)

	resultado, ok := svc.Obtener(1)

	assert.True(t, ok)
	assert.Equal(t, "Parada Paraninfo", resultado.Nombre)
	repo.AssertExpectations(t)
}

func TestParadaService_Borrar(t *testing.T) {
	repo := new(almacenMock)

	t.Run("Borrar exitoso", func(t *testing.T) {
		repo.On("BorrarParada", 1).Return(true)
		svc := st.NewParadaService(repo)

		err := svc.Borrar(1)

		assert.NoError(t, err)
	})

	t.Run("Borrar no encontrado", func(t *testing.T) {
		repo.On("BorrarParada", 99).Return(false)
		svc := st.NewParadaService(repo)

		err := svc.Borrar(99)

		assert.Error(t, err)
	})
}
