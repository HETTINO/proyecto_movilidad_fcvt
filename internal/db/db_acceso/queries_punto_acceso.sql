-- name: CrearPuntoAcceso :one
INSERT INTO Punto_de_acceso (Tipo_acceso, Ubicacion)
VALUES (?, ?)
RETURNING id_puntoacceso, tipo_acceso, ubicacion;

-- name: ListarPuntosAcceso :many
SELECT ID_puntoacceso, Tipo_acceso, Ubicacion
FROM Punto_de_acceso;

-- name: ObtenerPuntoAccesoPorId :one
SELECT ID_puntoacceso, Tipo_acceso, Ubicacion
FROM Punto_de_acceso
WHERE ID_puntoacceso = ? LIMIT 1;

-- name: ActualizarPuntoAcceso :one
UPDATE Punto_de_acceso
SET Tipo_acceso = ?, Ubicacion = ?
WHERE ID_puntoacceso = ?
RETURNING id_puntoacceso, tipo_acceso, ubicacion;

-- name: BorrarPuntoAcceso :execrows
DELETE FROM Punto_de_acceso
WHERE ID_puntoacceso = ?;