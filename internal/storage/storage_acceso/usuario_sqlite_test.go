package storage_acceso_test

import (
	"testing"

	"proyecto_movilidad_fcvt/internal/modelos"
)

func TestUsuarioSQLite(t *testing.T) {
	repo := nuevoRepo(t)

	// CREATE
	user := modelos.Usuario{
		Cedula:     "131555",
		Nombre:     "Shirley Juleidy",
		Contrasena: "1234",
		Email:      "shirley@example.com",
		Rol:        "admin",
	}

	creado := repo.CrearUsuario(user)

	if creado.Cedula != "131555" {
		t.Fatalf("cedula incorrecta")
	}

	// LISTAR
	lista := repo.ListarUsuarios()

	if len(lista) == 0 {
		t.Fatalf("se esperaba al menos 1 usuario")
	}

	if lista[0].Nombre != "Shirley Juleidy" {
		t.Errorf("nombre incorrecto")
	}

	// BUSCAR
	encontrado, ok := repo.BuscarUsuarioPorCedula("131555")
	if !ok {
		t.Fatalf("usuario no encontrado")
	}

	if encontrado.Cedula != "131555" {
		t.Errorf("cedula no coincide")
	}

	// ACTUALIZAR
	_, ok = repo.ActualizarUsuario("131555", modelos.Usuario{
		Cedula:     "131555",
		Nombre:     "Shirley Actualizada",
		Contrasena: "1234",
		Email:      "shirley@example.com",
		Rol:        "admin",
	})

	if !ok {
		t.Fatalf("no se pudo actualizar usuario")
	}

	// DELETE
	if !repo.BorrarUsuario("131555") {
		t.Errorf("no se pudo eliminar usuario")
	}
}
