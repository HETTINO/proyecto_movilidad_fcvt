-- name: ListarVehiculos :many
SELECT placa_vehiculo, id_usuario, Tipo_usuario, Marca, Modelo, Color, Anio
FROM Vehiculos;

-- name: ObtenerVehiculoPorPlaca :one
SELECT placa_vehiculo, id_usuario, Tipo_usuario, Marca, Modelo, Color, Anio
FROM Vehiculos
WHERE placa_vehiculo = ? LIMIT 1;

-- name: RegistrarVehiculo :one
INSERT INTO Vehiculos (placa_vehiculo, id_usuario, Tipo_usuario, Marca, Modelo, Color, Anio)
VALUES (?, ?, ?, ?, ?, ?, ?)
RETURNING placa_vehiculo, id_usuario, tipo_usuario, marca, modelo, color, anio;

-- name: ActualizarVehiculo :one
UPDATE Vehiculos
SET id_usuario = ?, Tipo_usuario = ?, Marca = ?, Modelo = ?, Color = ?, Anio = ?
WHERE placa_vehiculo = ?
RETURNING placa_vehiculo, id_usuario, tipo_usuario, marca, modelo, color, anio;

-- name: BorrarVehiculo :execrows
DELETE FROM Vehiculos
WHERE placa_vehiculo = ?;