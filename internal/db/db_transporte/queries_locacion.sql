-- name: ListarLocaciones :many
SELECT id, latitud, longitud, time_stamp, carrito_id FROM locaciones;

-- name: RegistrarLocacion :one
INSERT INTO locaciones (latitud, longitud, time_stamp, carrito_id) VALUES (?, ?, ?, ?) RETURNING id, latitud, longitud, time_stamp, carrito_id;

-- name: ObtenerUltimaLocacionPorCarrito :one
SELECT id, latitud, longitud, time_stamp, carrito_id FROM locaciones WHERE carrito_id = ? ORDER BY time_stamp DESC LIMIT 1;