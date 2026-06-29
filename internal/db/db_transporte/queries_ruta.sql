-- name: ListarRutas :many
SELECT id, nombre, descripcion FROM rutas;

-- name: BuscarRutaPorID :one
SELECT id, nombre, descripcion FROM rutas WHERE id = ?;

-- name: CrearRuta :one
INSERT INTO rutas (nombre, descripcion) VALUES (?, ?) RETURNING id, nombre, descripcion;

-- name: ActualizarRuta :one
UPDATE rutas SET nombre = ?, descripcion = ? WHERE id = ? RETURNING id, nombre, descripcion;

-- name: BorrarRuta :execrows
DELETE FROM rutas WHERE id = ?;