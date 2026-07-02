package storage_acceso_test

import (
	"testing"
	"time"

	"proyecto_movilidad_fcvt/internal/modelos"
)

func TestVehiculoSQLite(t *testing.T) {
	repo := nuevoRepo(t)

	// =========================
	// USUARIO (Cedula es string)
	// =========================
	user := modelos.Usuario{
		Cedula:     "131555",
		Nombre:     "Shirley Juleidy",
		Contrasena: "1234",
		Email:      "shirley@example.com",
		Rol:        "admin",
	}

	usuarioCreado := repo.CrearUsuario(user)
	if usuarioCreado.Cedula == "" {
		t.Fatal("no se creó el usuario")
	}

	// =========================
	// VEHÍCULO (según tus entidades)
	// =========================
	v := modelos.Vehiculo{
		Placa:        "ABC-1234",
		IDUsuario:    usuarioCreado.Cedula,
		TipoVehiculo: "Auto",
		Marca:        "Toyota",
		Modelo:       "Corolla",
		Color:        "Rojo",
		Año:          2020,
	}

	creado := repo.CrearVehiculo(v)

	if creado.Placa != "ABC-1234" {
		t.Fatalf("placa incorrecta: %v", creado.Placa)
	}

	lista := repo.ListarVehiculos()
	if len(lista) == 0 {
		t.Fatal("esperaba al menos 1 vehículo")
	}

	if lista[0].Placa != "ABC-1234" {
		t.Errorf("vehículo no coincide")
	}

	// =========================
	// (Opcional) ACCESO básico
	// =========================
	now := time.Now()

	acceso := modelos.Acceso{
		PlacaVehiculo: "ABC-1234",
		PuntoAccesoID: 1,
		TiempoEntrada: now,
		Estado:        "activo",
	}

	_ = repo.CrearAcceso(acceso)
}
