package serviceparqueadero_test
import (
	"testing"

	"github.com/stretchr/testify/assert"

	"proyecto_movilidad_fcvt/internal/modelos"
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

			svc := sp.NewEspacioService(repo)

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

	svc := sp.NewEspacioService(repo)

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

	svc := sp.NewEspacioService(repo)

	resultado, ok := svc.Obtener(1)

	assert.True(t, ok)

	assert.Equal(t, 10, resultado.Numero)

	assert.Equal(t, "libre", resultado.Estado)

	repo.AssertExpectations(t)

}
