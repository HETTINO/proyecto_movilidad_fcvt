
-- name: ListarOcupaciones :many
SELECT id_ocupacion, placa_vehiculo, id_espacio, id_acceso, hora_inicio, hora_fin
FROM ocupacion;

-- name: BuscarOcupacionPorID :one
SELECT id_ocupacion, placa_vehiculo, id_espacio, id_acceso, hora_inicio, hora_fin
FROM ocupacion
WHERE id_ocupacion = ?;

-- name: ListarOcupacionesActivasPorEspacio :many
SELECT id_ocupacion, placa_vehiculo, id_espacio, id_acceso, hora_inicio, hora_fin
FROM ocupacion
WHERE id_espacio = ? AND hora_fin IS NULL;

-- name: CrearOcupacion :one
INSERT INTO ocupacion (placa_vehiculo, id_espacio, id_acceso, hora_inicio)
VALUES (?, ?, ?, ?)
RETURNING id_ocupacion, placa_vehiculo, id_espacio, id_acceso, hora_inicio, hora_fin;

-- name: CerrarOcupacion :one
UPDATE ocupacion
SET hora_fin = ?
WHERE id_ocupacion = ?
RETURNING id_ocupacion, placa_vehiculo, id_espacio, id_acceso, hora_inicio, hora_fin;

-- name: ActualizarOcupacion :one
UPDATE ocupacion
SET placa_vehiculo = ?, id_espacio = ?, id_acceso = ?, hora_inicio = ?, hora_fin = ?
WHERE id_ocupacion = ?
RETURNING id_ocupacion, placa_vehiculo, id_espacio, id_acceso, hora_inicio, hora_fin;

-- name: BorrarOcupacion :execrows
DELETE FROM ocupacion WHERE id_ocupacion = ?;