
-- name: ListarParqueaderos :many
SELECT id_parqueadero, nombre, capacidad, tipo FROM parqueaderos;

-- name: BuscarParqueaderoPorID :one
SELECT id_parqueadero, nombre, capacidad, tipo FROM parqueaderos
WHERE id_parqueadero = ?;

-- name: CrearParqueadero :one
INSERT INTO parqueaderos (nombre, capacidad, tipo)
VALUES (?, ?, ?)
RETURNING id_parqueadero, nombre, capacidad, tipo;

-- name: ActualizarParqueadero :one
UPDATE parqueaderos
SET nombre = ?, capacidad = ?, tipo = ?
WHERE id_parqueadero = ?
RETURNING id_parqueadero, nombre, capacidad, tipo;

-- name: BorrarParqueadero :execrows
DELETE FROM parqueaderos WHERE id_parqueadero = ?;