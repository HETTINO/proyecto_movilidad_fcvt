package storage_acceso

import (
	"log"
	"proyecto_movilidad_fcvt/internal/modelos"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Almacen struct {
	DB *gorm.DB
}

// NuevoAlmacen conecta SQLite y ejecuta las migraciones de tu módulo
func NuevoAlmacen(rutaDB string) *Almacen {
	db, err := gorm.Open(sqlite.Open(rutaDB), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error conectando BD de acceso: %v", err)
	}

	// Migración automática de tus 4 entidades
	err = db.AutoMigrate(
		&modelos.Usuario{},
		&modelos.Vehiculo{},
		&modelos.PuntoDeAcceso{},
		&modelos.Acceso{},
	)
	if err != nil {
		log.Fatalf("Error en migración de acceso: %v", err)
	}

	return &Almacen{DB: db}
}
