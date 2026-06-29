package models

import "time"

type Ruta struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Nombre      string `json:"nombre" gorm:"type:varchar(100)"`
	Descripcion string `json:"descripcion" gorm:"type:text"`
}

type Parada struct {
	IDParada int     `json:"id_parada" gorm:"primaryKey"`
	Nombre   string  `json:"nombre" gorm:"type:varchar(100)"`
	Latitud  float64 `json:"latitud"`
	Longitud float64 `json:"longitud"`
}

type Carrito struct {
	ID            int    `json:"id" gorm:"primaryKey"`
	NombreCarrito string `json:"nombre_carrito" gorm:"type:varchar(100)"`
	Capacidad     int    `json:"capacidad"`
	Estado        string `json:"estado" gorm:"type:varchar(50)"`
	RutaID        int    `json:"ruta_id"`
}

type Locacion struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Latitud   float64   `json:"latitud"`
	Longitud  float64   `json:"longitud"`
	TimeStamp time.Time `json:"time_stamp"`
	CarritoID int       `json:"carrito_id"`
}

type Solicitud struct {
	ID            int    `json:"id" gorm:"primaryKey"`
	CedulaUsuario string `json:"cedula_usuario" gorm:"type:varchar(10)"`
	CantPersonas  int    `json:"cant_personas"`
	ParadaOrigen  int    `json:"parada_origen"`
	PuntoDestino  string `json:"punto_destino" gorm:"type:varchar(100)"`
	Estado        string `json:"estado" orm:"type:varchar(50)"`
	IDCarrito     *int   `json:"id_carrito,omitempty"`
}
