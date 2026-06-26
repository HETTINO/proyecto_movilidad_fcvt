package service_parqueadero_test

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	sp "proyecto_movilidad_fcvt/internal/service/service_parqueadero"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParqueaderoService_Crear(t *testing.T) {
	casos := []struct {
		nombre        string
		entrada       modelos.Parqueadero
		debeFallar    bool
		debePersistir bool
	}{
		{
			nombre:        "nombre vacío -> no persiste",
			entrada:       modelos.Parqueadero{Nombre: "   "},
			debeFallar:    true,
			debePersistir: false,
		},
		{
			nombre:        "parqueadero válido -> se persiste",
			entrada:       modelos.Parqueadero{Nombre: "Norte"},
			debeFallar:    false,
			debePersistir: true,
		},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {

			repo := new(parqueaderoRepoMock)

			if c.debePersistir {
				guardado := c.entrada
				guardado.IDParqueadero = 1

				repo.
					On("CrearParqueadero", c.entrada).
					Return(guardado)
			}

			svc := sp.NewParqueaderoService(repo)

			creado, err := svc.Crear(c.entrada)

			if c.debeFallar {
				assert.Error(t, err)
				repo.AssertNotCalled(t, "CrearParqueadero")
			} else {
				assert.NoError(t, err)
				assert.Equal(t, 1, creado.IDParqueadero)
				repo.AssertCalled(t, "CrearParqueadero", c.entrada)
			}
		})
	}
}

func TestParqueaderoService_Obtener_NoEncontrado(t *testing.T) {
	repo := new(parqueaderoRepoMock)
	repo.On("BuscarParqueaderoPorID", 999).Return(modelos.Parqueadero{}, false)
	svc := sp.NewParqueaderoService(repo)

	_, ok := svc.Obtener(999)

	assert.False(t, ok)
	repo.AssertExpectations(t)
}

func TestParqueaderoService_Obtener_Exitoso(t *testing.T) {
	repo := new(parqueaderoRepoMock)
	esperado := modelos.Parqueadero{IDParqueadero: 1, Nombre: "Norte", Capacidad: 50, Tipo: "cubierto"}
	repo.On("BuscarParqueaderoPorID", 1).Return(esperado, true)
	svc := sp.NewParqueaderoService(repo)

	resultado, ok := svc.Obtener(1)

	assert.True(t, ok)
	assert.Equal(t, "Norte", resultado.Nombre)
	repo.AssertExpectations(t)
}
