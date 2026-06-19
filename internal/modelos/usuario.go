package modelos

import "time"

type Usuario struct {
	ID       uint      `json:"id" gorm:"primaryKey"`
	Email    string    `json:"email" gorm:"not null;uniqueIndex"`
	Password string    `json:"-" gorm:"not null"`
	Creadoen time.Time `json:"creado_at"`
}
