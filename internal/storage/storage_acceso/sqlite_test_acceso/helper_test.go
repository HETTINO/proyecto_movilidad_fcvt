package sqlite_test_acceso

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_acceso"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func abrirDBMemoria(t *testing.T) *gorm.DB {
	t.Helper()

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)

	err = db.AutoMigrate(
		&modelos.Usuario{},
		&modelos.Vehiculo{},
		&modelos.PuntoDeAcceso{}, // ✅ CORRECTO
		&modelos.Acceso{},
	)
	require.NoError(t, err)

	return db
}

func nuevoRepo(t *testing.T) *storage.AlmacenSQLite {
	t.Helper()
	return storage.NuevoAlmacenSQLite(abrirDBMemoria(t))
}
