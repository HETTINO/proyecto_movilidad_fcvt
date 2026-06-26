package storage_parqueadero

import (
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"proyecto_movilidad_fcvt/internal/modelos"
)

func abrirDBMemoria(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err, "no se pudo abrir SQLite :memory:")

	err = db.AutoMigrate(
		&modelos.Parqueadero{},
		&modelos.Espacio{},
		&modelos.Ocupacion{},
	)
	require.NoError(t, err, "AutoMigrate falló")
	return db
}

// =========================================================
// TEST 1 — Crear → Listar lo refleja
// =========================================================

func TestSQLite_CrearYListarParqueadero(t *testing.T) {
	db := abrirDBMemoria(t)
	repo := NuevoAlmacenSQLite(db)

	repo.CrearParqueadero(modelos.Parqueadero{
		Nombre:    "Parqueadero Norte",
		Capacidad: 50,
		Tipo:      "cubierto",
	})

	lista := repo.ListarParqueaderos()

	assert.Len(t, lista, 1, "debería haber exactamente 1 parqueadero")
	assert.Equal(t, "Parqueadero Norte", lista[0].Nombre)
}

// =========================================================
// TEST 2 — Crear → BuscarPorID lo encuentra
// =========================================================

func TestSQLite_CrearYBuscarParqueaderoPorID(t *testing.T) {
	db := abrirDBMemoria(t)
	repo := NuevoAlmacenSQLite(db)

	creado := repo.CrearParqueadero(modelos.Parqueadero{
		Nombre:    "Parqueadero Sur",
		Capacidad: 30,
		Tipo:      "abierto",
	})

	assert.NotZero(t, creado.IDParqueadero, "GORM debe asignar un ID")

	encontrado, ok := repo.BuscarParqueaderoPorID(creado.IDParqueadero)

	assert.True(t, ok, "debería encontrar el parqueadero recién creado")
	assert.Equal(t, "Parqueadero Sur", encontrado.Nombre)
	assert.Equal(t, 30, encontrado.Capacidad)
}

// =========================================================
// TEST 3 — BuscarPorID inexistente devuelve false
// =========================================================

func TestSQLite_BuscarParqueaderoInexistente(t *testing.T) {
	db := abrirDBMemoria(t)
	repo := NuevoAlmacenSQLite(db)

	_, ok := repo.BuscarParqueaderoPorID(999)

	assert.False(t, ok, "un ID que no existe debe devolver false")
}

func TestSQLite_CrearYBuscarEspacio(t *testing.T) {
	db := abrirDBMemoria(t)
	repo := NuevoAlmacenSQLite(db)

	// Primero necesitas un parqueadero porque Espacio tiene FK
	park := repo.CrearParqueadero(modelos.Parqueadero{
		Nombre:    "Parqueadero Central",
		Capacidad: 20,
		Tipo:      "cubierto",
	})

	creado := repo.CrearEspacio(modelos.Espacio{
		IDParqueadero: park.IDParqueadero,
		Numero:        1,
		Estado:        "libre",
		TipoEspacio:   "auto",
	})

	assert.NotZero(t, creado.IDEspacio)

	encontrado, ok := repo.BuscarEspacioPorID(creado.IDEspacio)
	assert.True(t, ok)
	assert.Equal(t, "libre", encontrado.Estado)
}
