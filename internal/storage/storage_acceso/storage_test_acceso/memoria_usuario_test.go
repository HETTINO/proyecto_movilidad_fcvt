package storage_test_acceso

import (
	"testing"

	"proyecto_movilidad_fcvt/internal/modelos"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_acceso"
)

func TestMemoria_CrearYBuscarUsuario(t *testing.T) {
	m := storage.NuevoMemoriaAcceso()

	m.CrearUsuario(modelos.Usuario{
		Cedula:     "131555",
		Nombre:     "Shirley Juleidy",
		Contrasena: "123456",
		Email:      "shirley@test.com",
		Rol:        "Estudiante",
	})

	encontrado, ok := m.BuscarUsuarioPorCedula("131555")
	if !ok {
		t.Fatalf("no se encontró el usuario con cédula 131555")
	}

	if encontrado.Nombre != "Shirley Juleidy" {
		t.Errorf("nombre = %q; esperaba %q", encontrado.Nombre, "Shirley Juleidy")
	}
}

func TestMemoria_ActualizarYBorrarUsuario(t *testing.T) {
	m := storage.NuevoMemoriaAcceso()

	m.CrearUsuario(modelos.Usuario{
		Cedula:     "777",
		Nombre:     "Original",
		Contrasena: "123",
		Email:      "original@test.com",
		Rol:        "Docente",
	})

	_, ok := m.ActualizarUsuario("777", modelos.Usuario{
		Cedula:     "777",
		Nombre:     "Modificado",
		Contrasena: "123",
		Email:      "modificado@test.com",
		Rol:        "Docente",
	})

	if !ok {
		t.Fatalf("no se pudo actualizar el usuario")
	}

	if !m.BorrarUsuario("777") {
		t.Errorf("esperaba poder borrar el usuario")
	}
}
