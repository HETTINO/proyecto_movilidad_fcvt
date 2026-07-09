package modelos

import "time"

type Ruta struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Nombre      string `json:"nombre" gorm:"type:varchar(100)"`
	Descripcion string `json:"descripcion" gorm:"type:text"`

	Paradas  []Parada  `json:"paradas,omitempty" gorm:"foreignKey:RutaID"` // ← esto también
	Carritos []Carrito `json:"carritos,omitempty" gorm:"foreignKey:RutaID"`
}

type Parada struct {
	IDParada int     `json:"id_parada" gorm:"primaryKey"`
	Nombre   string  `json:"nombre" gorm:"type:varchar(100)"`
	Latitud  float64 `json:"latitud"`
	Longitud float64 `json:"longitud"`
	RutaID   int     `json:"ruta_id"`
	Ruta     Ruta    `json:"ruta,omitempty" gorm:"foreignKey:RutaID;references:ID"`
}

type Carrito struct {
	ID            int         `json:"id" gorm:"primaryKey"`
	NombreCarrito string      `json:"nombre_carrito" gorm:"type:varchar(100)"`
	Capacidad     int         `json:"capacidad"`
	Estado        string      `json:"estado" gorm:"type:varchar(50)"`
	RutaID        int         `json:"ruta_id"`
	Locaciones    []Locacion  `json:"locaciones,omitempty" gorm:"foreignKey:CarritoID"`
	Solicitudes   []Solicitud `json:"solicitudes,omitempty" gorm:"foreignKey:IDCarrito"`
}

type Locacion struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Latitud   float64   `json:"latitud"`
	Longitud  float64   `json:"longitud"`
	TimeStamp time.Time `json:"time_stamp"`
	CarritoID int       `json:"carrito_id"`
	Carrito   Carrito   `json:"carrito,omitempty" gorm:"foreignKey:CarritoID;references:ID"`
}

type Solicitud struct {
	ID            int    `json:"id" gorm:"primaryKey"`
	CedulaUsuario string `json:"cedula_usuario" gorm:"type:varchar(10)"`
	CantPersonas  int    `json:"cant_personas"`
	ParadaOrigen  int    `json:"parada_origen"`
	Parada        Parada `json:"parada,omitempty" gorm:"foreignKey:ParadaOrigen;references:IDParada"` // ← falta esto
	PuntoDestino  string `json:"punto_destino" gorm:"type:varchar(100)"`
	Estado        string `json:"estado" gorm:"type:varchar(50)"`
	IDCarrito     *int   `json:"id_carrito,omitempty"`

	Carrito *Carrito `json:"carrito,omitempty" gorm:"foreignKey:IDCarrito;references:ID"`
}
