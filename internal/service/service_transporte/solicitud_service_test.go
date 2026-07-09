package servicetransporte_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	modelos "proyecto_movilidad_fcvt/internal/modelos"
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
func TestSolicitudService_Listar(t *testing.T) {

	repo := new(almacenMock)

	esperadas := []modelos.Solicitud{
		{ID: 1, CedulaUsuario: "0102030405", CantPersonas: 2, PuntoDestino: "Centro"},
		{ID: 2, CedulaUsuario: "1122334455", CantPersonas: 1, PuntoDestino: "Norte"},
	}

	repo.
		On("ListarSolicitudes").
		Return(esperadas)

	svc := st.NewSolicitudService(repo)

	resultado := svc.Listar()

	assert.Len(t, resultado, 2)
	assert.Equal(t, "Centro", resultado[0].PuntoDestino)

	repo.AssertExpectations(t)

}

func TestSolicitudService_Actualizar_Exitoso(t *testing.T) {

	repo := new(almacenMock)

	datos := modelos.Solicitud{
		CedulaUsuario: "0102030405",
		CantPersonas:  2,
		ParadaOrigen:  1,
		PuntoDestino:  "Centro",
	}

	esperado := datos
	esperado.ID = 1

	repo.
		On("ActualizarSolicitud", 1, datos).
		Return(esperado, true)

	svc := st.NewSolicitudService(repo)

	actualizado, encontrado, err := svc.Actualizar(1, datos)

	assert.NoError(t, err)
	assert.True(t, encontrado)
	assert.Equal(t, 1, actualizado.ID)

	repo.AssertExpectations(t)

}
func TestSolicitudService_Actualizar_NoEncontrado(t *testing.T) {

	repo := new(almacenMock)

	datos := modelos.Solicitud{
		CedulaUsuario: "0102030405",
		CantPersonas:  2,
		PuntoDestino:  "Centro",
	}

	repo.
		On("ActualizarSolicitud", 999, datos).
		Return(modelos.Solicitud{}, false)

	svc := st.NewSolicitudService(repo)

	_, encontrado, err := svc.Actualizar(999, datos)

	assert.Error(t, err)
	assert.False(t, encontrado)

	repo.AssertExpectations(t)

}

func TestSolicitudService_Actualizar_DatosInvalidos(t *testing.T) {

	repo := new(almacenMock)

	datos := modelos.Solicitud{
		CedulaUsuario: "", // inválido: campo requerido vacío
		CantPersonas:  2,
		PuntoDestino:  "Centro",
	}

	svc := st.NewSolicitudService(repo)

	_, encontrado, err := svc.Actualizar(1, datos)

	assert.Error(t, err)
	assert.False(t, encontrado)

	// La validación debe fallar ANTES de tocar el repositorio
	repo.AssertNotCalled(t, "ActualizarSolicitud")

}

func TestSolicitudService_Borrar_Exitoso(t *testing.T) {

	repo := new(almacenMock)

	repo.
		On("BorrarSolicitud", 1).
		Return(true)

	svc := st.NewSolicitudService(repo)

	err := svc.Borrar(1)

	assert.NoError(t, err)

	repo.AssertExpectations(t)

}

func TestSolicitudService_Borrar_NoEncontrado(t *testing.T) {

	repo := new(almacenMock)

	repo.
		On("BorrarSolicitud", 999).
		Return(false)

	svc := st.NewSolicitudService(repo)

	err := svc.Borrar(999)

	assert.Error(t, err)

	repo.AssertExpectations(t)

}