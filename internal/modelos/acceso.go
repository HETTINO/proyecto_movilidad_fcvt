package modelos

import "time"

// Usuario representa la tabla de usuarios en el sistema
type Usuario struct {
	Cedula     string `json:"cedula" gorm:"primaryKey;type:varchar(10)"` // Usamos varchar(10) como en tu Solicitud.CedulaUsuario
	Nombre     string `json:"nombre" gorm:"type:varchar(100)"`
	Contrasena string `json:"contrasena" gorm:"type:varchar(255)"`
	Email      string `json:"email" gorm:"type:varchar(100)"`
	Rol        string `json:"rol" gorm:"type:varchar(50)"`
}

// Vehiculo representa los datos del vehículo asociado
type Vehiculo struct {
	Placa        string `json:"placa" gorm:"primaryKey;type:varchar(10)"`
	IDUsuario    string `json:"id_usuario" gorm:"type:varchar(10)"` // Llave foránea hacia Usuario.Cedula
	TipoVehiculo string `json:"tipo_vehiculo" gorm:"type:varchar(50)"`
	Marca        string `json:"marca" gorm:"type:varchar(50)"`
	Modelo       string `json:"modelo" gorm:"type:varchar(50)"`
	Color        string `json:"color" gorm:"type:varchar(30)"`
	Año          int    `json:"año"`
}

// PuntoDeAcceso representa el lugar físico por donde se ingresa o sale
type PuntoDeAcceso struct {
	ID         int    `json:"id" gorm:"primaryKey"` // Cambiado a ID para seguir el estándar de Ruta, Carrito, Locación, etc.
	Frecuencia string `json:"frecuencia" gorm:"type:varchar(50)"`
	Ubicacion  string `json:"ubicacion" gorm:"type:varchar(100)"`
}

// Acceso representa el registro de eventos de entrada y salida
type Acceso struct {
	ID            int        `json:"id" gorm:"primaryKey"` // Cambiado a ID estándar
	PlacaVehiculo string     `json:"placa_vehiculo" gorm:"type:varchar(10)"`
	PuntoAccesoID int        `json:"punto_acceso_id"` // Estandarizado como tu "RutaID" o "CarritoID"
	TiempoEntrada time.Time  `json:"tiempo_entrada"`
	TiempoSalida  *time.Time `json:"tiempo_salida,omitempty"` // Puntero para admitir nulos (aún no sale)
	Estado        string     `json:"estado" gorm:"type:varchar(50)"`
	Observaciones string     `json:"observaciones" gorm:"type:text"` // Como tu Descripción en Ruta
}
