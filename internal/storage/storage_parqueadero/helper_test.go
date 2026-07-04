package storage_parqueadero_test

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_parqueadero"
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
		&modelos.Parqueadero{},
		&modelos.Espacio{},
		&modelos.Ocupacion{},
	)
	require.NoError(t, err)
	return db
}

func nuevoRepo(t *testing.T) *storage.AlmacenSQLite {
	return storage.NuevoAlmacenSQLite(abrirDBMemoria(t))
}
