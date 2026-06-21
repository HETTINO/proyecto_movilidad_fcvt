package modelos

import "time"

type Ruta struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}

type Parada struct {
	IDParada int     `json:"id_parada"`
	Nombre   string  `json:"nombre"`
	Latitud  float64 `json:"latitud"`
	Longitu  float64 `json:"longitu"`
	RutaID   int     `json:"ruta_id"`
}

type Carrito struct {
	ID            int    `json:"id"`
	NombreCarrito string `json:"nombre_carrito"`
	Capacidad     int    `json:"capacidad"`
	Estado        string `json:"estado"`
	RutaID        int    `json:"ruta_id"`
}

type Locacion struct {
	ID        int       `json:"id"`
	Latitud   float64   `json:"latitud"`
	Longitud  float64   `json:"longitud"`
	TimeStamp time.Time `json:"time_stamp"`
	CarritoID string    `json:"carrito_id"`
}

type Solicitud struct {
	ID            int     `json:"id"`
	CedulaUsuario string  `json:"cedula_usuario"`
	CantPersonas  int     `json:"cant_personas"`
	PuntoDestino  string  `json:"punto_destino"`
	Estado        string  `json:"estado"`
	IDCarrito     *string `json:"id_carrito,omitempty"`
}
