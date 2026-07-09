package serviceparqueadero_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"proyecto_movilidad_fcvt/internal/modelos"
	service "proyecto_movilidad_fcvt/internal/service"
	sp "proyecto_movilidad_fcvt/internal/service/service_parqueadero"
)

// =========================================================
// TESTS — EspacioService
// =========================================================

func TestEspacioService_Crear(t *testing.T) {

	casos := []struct {
		nombre        string
		entrada       modelos.Espacio
		debeFallar    bool
		debePersistir bool
	}{
		{
			nombre: "id_parqueadero vacío -> no persiste",
			entrada: modelos.Espacio{
				IDParqueadero: 0,
			},
			debeFallar:    true,
			debePersistir: false,
		},
		{
			nombre: "espacio válido -> se persiste",
			entrada: modelos.Espacio{
				IDParqueadero: 1,
				Numero:        20,
				Estado:        "libre",
				TipoEspacio:   "auto",
			},
			debeFallar:    false,
			debePersistir: true,
		},
	}

	for _, c := range casos {

		t.Run(c.nombre, func(t *testing.T) {

			repo := new(espacioRepoMock)

			if c.debePersistir {

				guardado := c.entrada
				guardado.IDEspacio = 1

				repo.
					On("CrearEspacio", c.entrada).
					Return(guardado)

			}

			ocupacionRepo := new(ocupacionRepoMock)
			svc := sp.NewEspacioService(repo, ocupacionRepo)

			creado, err := svc.Crear(c.entrada)

			if c.debeFallar {

				assert.Error(t, err)

				repo.AssertNotCalled(t, "CrearEspacio")

			} else {

				assert.NoError(t, err)

				assert.Equal(t, 1, creado.IDEspacio)

				repo.AssertCalled(t, "CrearEspacio", c.entrada)

			}

		})

	}

}

func TestEspacioService_Obtener_NoEncontrado(t *testing.T) {

	repo := new(espacioRepoMock)

	repo.
		On("BuscarEspacioPorID", 999).
		Return(modelos.Espacio{}, false)

	svc := sp.NewEspacioService(repo, new(ocupacionRepoMock))

	_, ok := svc.Obtener(999)

	assert.False(t, ok)

	repo.AssertExpectations(t)

}

func TestEspacioService_Obtener_Exitoso(t *testing.T) {

	repo := new(espacioRepoMock)

	esperado := modelos.Espacio{
		IDEspacio:     1,
		IDParqueadero: 1,
		Numero:        10,
		Estado:        "libre",
		TipoEspacio:   "auto",
	}

	repo.
		On("BuscarEspacioPorID", 1).
		Return(esperado, true)

	svc := sp.NewEspacioService(repo, new(ocupacionRepoMock))

	resultado, ok := svc.Obtener(1)

	assert.True(t, ok)

	assert.Equal(t, 10, resultado.Numero)

	assert.Equal(t, "libre", resultado.Estado)

	repo.AssertExpectations(t)

}

func TestEspacioService_Actualizar_Exitoso(t *testing.T) {
	repo := new(espacioRepoMock)
	datos := modelos.Espacio{IDParqueadero: 1, Numero: 5, Estado: "ocupado"}
	actualizado := datos
	actualizado.IDEspacio = 1

	repo.On("ActualizarEspacio", 1, datos).Return(actualizado, true)

	svc := sp.NewEspacioService(repo, new(ocupacionRepoMock))
	resultado, ok, err := svc.Actualizar(1, datos)

	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, "ocupado", resultado.Estado)
}

func TestEspacioService_Actualizar_NoEncontrado(t *testing.T) {
	repo := new(espacioRepoMock)
	datos := modelos.Espacio{IDParqueadero: 1}
	repo.On("ActualizarEspacio", 999, datos).Return(modelos.Espacio{}, false)

	svc := sp.NewEspacioService(repo, new(ocupacionRepoMock))
	_, ok, err := svc.Actualizar(999, datos)

	assert.False(t, ok)
	assert.ErrorIs(t, err, service.ErrNoEncontrado)
}

func TestEspacioService_Actualizar_IDParqueaderoVacio(t *testing.T) {
	repo := new(espacioRepoMock)
	svc := sp.NewEspacioService(repo, new(ocupacionRepoMock))

	_, ok, err := svc.Actualizar(1, modelos.Espacio{IDParqueadero: 0})

	assert.False(t, ok)
	assert.ErrorIs(t, err, service.ErrCampoRequerido)
	repo.AssertNotCalled(t, "ActualizarEspacio", 1, modelos.Espacio{})
}

func TestEspacioService_Borrar_Exitoso(t *testing.T) {
	repo := new(espacioRepoMock)
	ocupacionRepo := new(ocupacionRepoMock)

	repo.On("BuscarEspacioPorID", 1).Return(modelos.Espacio{IDEspacio: 1}, true)
	ocupacionRepo.On("ListarOcupacionesActivas", 1).Return([]modelos.Ocupacion{})
	repo.On("BorrarEspacio", 1).Return(true)

	svc := sp.NewEspacioService(repo, ocupacionRepo)
	err := svc.Borrar(1)

	assert.NoError(t, err)
	repo.AssertCalled(t, "BorrarEspacio", 1)
}

func TestEspacioService_Borrar_NoEncontrado(t *testing.T) {
	repo := new(espacioRepoMock)
	ocupacionRepo := new(ocupacionRepoMock)

	repo.On("BuscarEspacioPorID", 999).Return(modelos.Espacio{}, false)

	svc := sp.NewEspacioService(repo, ocupacionRepo)
	err := svc.Borrar(999)

	assert.ErrorIs(t, err, service.ErrNoEncontrado)
	repo.AssertNotCalled(t, "BorrarEspacio", 999)
	ocupacionRepo.AssertNotCalled(t, "ListarOcupacionesActivas", 999)
}

// TestEspacioService_Borrar_ConOcupacionesActivas es el caso que reprodujo el
// bug original: el espacio 1 tenía ocupaciones sin HoraFin apuntándole y
// aun así el DELETE devolvía 204. Ahora debe bloquearse con un error de
// dominio (mapeado a 409 Conflict en el handler) y BorrarEspacio jamás debe
// llegar a ejecutarse.
func TestEspacioService_Borrar_ConOcupacionesActivas(t *testing.T) {
	repo := new(espacioRepoMock)
	ocupacionRepo := new(ocupacionRepoMock)

	repo.On("BuscarEspacioPorID", 1).Return(modelos.Espacio{IDEspacio: 1}, true)
	ocupacionRepo.On("ListarOcupacionesActivas", 1).Return([]modelos.Ocupacion{
		{IDOcupacion: 1, IDEspacio: 1},
		{IDOcupacion: 2, IDEspacio: 1},
	})

	svc := sp.NewEspacioService(repo, ocupacionRepo)
	err := svc.Borrar(1)

	assert.ErrorIs(t, err, service.ErrEspacioConOcupacionesActivas)
	repo.AssertNotCalled(t, "BorrarEspacio", 1)
}

func TestEspacioService_Listar(t *testing.T) {
	repo := new(espacioRepoMock)
	esperado := []modelos.Espacio{{IDEspacio: 1, IDParqueadero: 1}}
	repo.On("ListarEspacios").Return(esperado)

	svc := sp.NewEspacioService(repo, new(ocupacionRepoMock))
	resultado := svc.Listar()

	assert.Equal(t, esperado, resultado)
}
