package modelos

import "time"

type Parqueadero struct {
	IDParqueadero int       `gorm:"primaryKey;autoIncrement" json:"id_parqueadero"`
	Nombre        string    `gorm:"not null" json:"nombre"`
	Capacidad     int       `gorm:"not null" json:"capacidad"`
	Tipo          string    `gorm:"not null" json:"tipo"`
	Espacios      []Espacio `gorm:"foreignKey:IDParqueadero"`
}

type Espacio struct {
	IDEspacio     int         `gorm:"primaryKey;autoIncrement" json:"id_espacio"`
	IDParqueadero int         `gorm:"not null" json:"id_parqueadero"`
	Numero        int         `gorm:"not null" json:"numero"`
	Estado        string      `gorm:"not null" json:"estado"`
	TipoEspacio   string      `gorm:"not null" json:"tipo_espacio"`
	Parqueadero   Parqueadero `gorm:"foreignKey:IDParqueadero"`
}

type Ocupacion struct {
	IDOcupacion   int        `gorm:"primaryKey;autoIncrement" json:"id_ocupacion"`
	PlacaVehiculo string     `gorm:"size:10;not null" json:"placa_vehiculo"`
	IDEspacio     int        `gorm:"not null" json:"id_espacio"`
	IDAcceso      int        `gorm:"not null" json:"id_acceso"`
	HoraInicio    time.Time  `gorm:"not null" json:"hora_inicio"`
	HoraFin       *time.Time `json:"hora_fin,omitempty"`
	Espacio       Espacio    `gorm:"foreignKey:IDEspacio"`
}
