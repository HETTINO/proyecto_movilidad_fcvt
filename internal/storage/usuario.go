package storage

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	"time"

	"gorm.io/gorm"
)

type UsuarioGorm struct {
	db *gorm.DB
}

func NewUsuarioGORM(db *gorm.DB) *UsuarioGorm {
	return &UsuarioGorm{db: db}
}
func (u *UsuarioGorm) CrearUsuario(usuario modelos.Usuario) (modelos.Usuario, error) {
	usuario.Creadoen = time.Now()
	if err := u.db.Create(&usuario).Error; err != nil {
		return modelos.Usuario{}, err
	}
	return usuario, nil
}
func (r *UsuarioGorm) BuscarUsuarioPorEmail(email string) (modelos.Usuario, bool) {
	var u modelos.Usuario
	if err := r.db.Where("email = ?", email).First(&u).Error; err != nil {
		return modelos.Usuario{}, false
	}
	return u, true
}
