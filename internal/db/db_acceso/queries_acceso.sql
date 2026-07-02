-- name: ListarAccesos :many
SELECT ID_acceso, Placa_vehiculo, ID_puntoacceso, Tiempo_entrada, Tiempo_salida, Estado, Observaciones
FROM Acceso
ORDER BY Tiempo_entrada DESC;

-- name: ObtenerAccesoPorID :one
SELECT ID_acceso, Placa_vehiculo, ID_puntoacceso, Tiempo_entrada, Tiempo_salida, Estado, Observaciones
FROM Acceso
WHERE ID_acceso = ? LIMIT 1;

-- name: RegistrarAccesoEntrada :one
INSERT INTO Acceso (Placa_vehiculo, ID_puntoacceso, Tiempo_entrada, Estado, Observaciones)
VALUES (?, ?, ?, ?, ?)
RETURNING id_acceso, placa_vehiculo, id_puntoacceso, tiempo_entrada, tiempo_salida, estado, observaciones;

-- name: RegistrarAccesoSalida :one
UPDATE Acceso
SET Tiempo_salida = ?, Estado = ?, Observaciones = ?
WHERE ID_acceso = ?
RETURNING id_acceso, placa_vehiculo, id_puntoacceso, tiempo_entrada, tiempo_salida, estado, observaciones;

-- name: ActualizarAcceso :one
UPDATE Acceso
SET Placa_vehiculo = ?, ID_puntoacceso = ?, Tiempo_entrada = ?, Tiempo_salida = ?, Estado = ?, Observaciones = ?
WHERE ID_acceso = ?
RETURNING id_acceso, placa_vehiculo, id_puntoacceso, tiempo_entrada, tiempo_salida, estado, observaciones;

-- name: BorrarAcceso :execrows
DELETE FROM Acceso
WHERE ID_acceso = ?;