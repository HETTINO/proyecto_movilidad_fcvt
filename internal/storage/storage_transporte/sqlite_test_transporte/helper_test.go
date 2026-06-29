package sqlite_test_trasnporte

import (
	modelos "proyecto_movilidad_fcvt/internal/modelos"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_transporte"
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
		&modelos.Solicitud{},
		&modelos.Carrito{},
		&modelos.Parada{},
		&modelos.Ruta{},
		&modelos.Locacion{},
	)
	require.NoError(t, err)
	return db
}

func nuevoRepo(t *testing.T) *storage.AlmacenSQLite {
	return storage.NuevoAlmacenSQLite(abrirDBMemoria(t))
}
