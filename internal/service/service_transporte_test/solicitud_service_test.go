package servicetransporte_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	modelos "proyecto_movilidad_fcvt/internal/models"
	st "proyecto_movilidad_fcvt/internal/service/service_transporte"
)

// =========================================================
// TESTS — SolicitudService
// =========================================================

func TestSolicitudService_Crear(t *testing.T) {

	casos := []struct {
		nombre        string
		entrada       modelos.Solicitud
		debeFallar    bool
		debePersistir bool
	}{
		{
			nombre: "cedula_usuario vacía -> no persiste",
			entrada: modelos.Solicitud{
				CedulaUsuario: "",
				CantPersonas:  2,
				PuntoDestino:  "Centro",
			},
			debeFallar:    true,
			debePersistir: false,
		},
		{
			nombre: "cant_personas inválida -> no persiste",
			entrada: modelos.Solicitud{
				CedulaUsuario: "0102030405",
				CantPersonas:  0,
				PuntoDestino:  "Centro",
			},
			debeFallar:    true,
			debePersistir: false,
		},
		{
			nombre: "punto_destino vacío -> no persiste",
			entrada: modelos.Solicitud{
				CedulaUsuario: "0102030405",
				CantPersonas:  2,
				PuntoDestino:  "",
			},
			debeFallar:    true,
			debePersistir: false,
		},
		{
			nombre: "solicitud válida -> se persiste",
			entrada: modelos.Solicitud{
				CedulaUsuario: "0102030405",
				CantPersonas:  2,
				ParadaOrigen:  1,
				PuntoDestino:  "Centro",
			},
			debeFallar:    false,
			debePersistir: true,
		},
	}

	for _, c := range casos {

		t.Run(c.nombre, func(t *testing.T) {

			repo := new(almacenMock)

			if c.debePersistir {

				enviado := c.entrada
				enviado.Estado = "pendiente"

				guardado := enviado
				guardado.ID = 1

				repo.
					On("CrearSolicitud", enviado).
					Return(guardado)

			}

			svc := st.NewSolicitudService(repo)

			creado, err := svc.Crear(c.entrada)

			if c.debeFallar {

				assert.Error(t, err)

				repo.AssertNotCalled(t, "CrearSolicitud")

			} else {

				assert.NoError(t, err)

				assert.Equal(t, 1, creado.ID)

				assert.Equal(t, "pendiente", creado.Estado)

			}

		})

	}

}

func TestSolicitudService_Obtener_NoEncontrado(t *testing.T) {

	repo := new(almacenMock)

	repo.
		On("BuscarSolicitudPorID", 999).
		Return(modelos.Solicitud{}, false)

	svc := st.NewSolicitudService(repo)

	_, ok := svc.Obtener(999)

	assert.False(t, ok)

	repo.AssertExpectations(t)

}

func TestSolicitudService_Obtener_Exitoso(t *testing.T) {

	repo := new(almacenMock)

	esperado := modelos.Solicitud{
		ID:            1,
		CedulaUsuario: "0102030405",
		CantPersonas:  2,
		ParadaOrigen:  1,
		PuntoDestino:  "Centro",
		Estado:        "pendiente",
	}

	repo.
		On("BuscarSolicitudPorID", 1).
		Return(esperado, true)

	svc := st.NewSolicitudService(repo)

	resultado, ok := svc.Obtener(1)

	assert.True(t, ok)

	assert.Equal(t, "0102030405", resultado.CedulaUsuario)

	assert.Equal(t, "Centro", resultado.PuntoDestino)

	repo.AssertExpectations(t)

}