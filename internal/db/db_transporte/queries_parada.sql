-- name: ListarParadas :many
SELECT id_parada, nombre, latitud, longitud FROM paradas;

-- name: BuscarParadaPorID :one
SELECT id_parada, nombre, latitud, longitud FROM paradas WHERE id_parada = ?;

-- name: CrearParada :one
INSERT INTO paradas (nombre, latitud, longitud) VALUES (?, ?, ?) RETURNING id_parada, nombre, latitud, longitud;

-- name: ActualizarParada :one
UPDATE paradas SET nombre = ?, latitud = ?, longitud = ? WHERE id_parada = ? RETURNING id_parada, nombre, latitud, longitud;

-- name: BorrarParada :execrows
DELETE FROM paradas WHERE id_parada = ?;