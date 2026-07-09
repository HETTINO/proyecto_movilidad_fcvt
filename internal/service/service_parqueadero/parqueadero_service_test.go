package serviceparqueadero_test

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	service "proyecto_movilidad_fcvt/internal/service"
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
					Return(guardado, nil)
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

func TestParqueaderoService_Actualizar_Exitoso(t *testing.T) {
	repo := new(parqueaderoRepoMock)
	datos := modelos.Parqueadero{Nombre: "Norte Renovado", Capacidad: 80}
	actualizado := datos
	actualizado.IDParqueadero = 1

	repo.On("ActualizarParqueadero", 1, datos).Return(actualizado, true)

	svc := sp.NewParqueaderoService(repo)
	resultado, ok, err := svc.Actualizar(1, datos)

	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, "Norte Renovado", resultado.Nombre)
	repo.AssertExpectations(t)
}

func TestParqueaderoService_Actualizar_NoEncontrado(t *testing.T) {
	repo := new(parqueaderoRepoMock)
	datos := modelos.Parqueadero{Nombre: "Norte"}
	repo.On("ActualizarParqueadero", 999, datos).Return(modelos.Parqueadero{}, false)

	svc := sp.NewParqueaderoService(repo)
	_, ok, err := svc.Actualizar(999, datos)

	assert.False(t, ok)
	assert.ErrorIs(t, err, service.ErrNoEncontrado)
}

func TestParqueaderoService_Actualizar_NombreVacio(t *testing.T) {
	repo := new(parqueaderoRepoMock)
	svc := sp.NewParqueaderoService(repo)

	_, ok, err := svc.Actualizar(1, modelos.Parqueadero{Nombre: "   "})

	assert.False(t, ok)
	assert.ErrorIs(t, err, service.ErrNombreVacio)
	repo.AssertNotCalled(t, "ActualizarParqueadero", 1, modelos.Parqueadero{})
}

func TestParqueaderoService_Borrar_Exitoso(t *testing.T) {
	repo := new(parqueaderoRepoMock)
	repo.On("BorrarParqueadero", 1).Return(true)

	svc := sp.NewParqueaderoService(repo)
	err := svc.Borrar(1)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestParqueaderoService_Borrar_NoEncontrado(t *testing.T) {
	repo := new(parqueaderoRepoMock)
	repo.On("BorrarParqueadero", 999).Return(false)

	svc := sp.NewParqueaderoService(repo)
	err := svc.Borrar(999)

	assert.ErrorIs(t, err, service.ErrNoEncontrado)
}

func TestParqueaderoService_Listar(t *testing.T) {
	repo := new(parqueaderoRepoMock)
	esperado := []modelos.Parqueadero{{IDParqueadero: 1, Nombre: "Norte"}}
	repo.On("ListarParqueaderos").Return(esperado)

	svc := sp.NewParqueaderoService(repo)
	resultado := svc.Listar()

	assert.Equal(t, esperado, resultado)
}
