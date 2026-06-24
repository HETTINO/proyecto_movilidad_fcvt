
-- name: ListarEspacios :many
SELECT id_espacio, id_parqueadero, numero, estado, tipo_espacio FROM espacios;

-- name: ListarEspaciosPorParqueadero :many
SELECT id_espacio, id_parqueadero, numero, estado, tipo_espacio FROM espacios
WHERE id_parqueadero = ?;

-- name: BuscarEspacioPorID :one
SELECT id_espacio, id_parqueadero, numero, estado, tipo_espacio FROM espacios
WHERE id_espacio = ?;

-- name: CrearEspacio :one
INSERT INTO espacios (id_parqueadero, numero, estado, tipo_espacio)
VALUES (?, ?, ?, ?)
RETURNING id_espacio, id_parqueadero, numero, estado, tipo_espacio;

-- name: ActualizarEspacio :one
UPDATE espacios
SET id_parqueadero = ?, numero = ?, estado = ?, tipo_espacio = ?
WHERE id_espacio = ?
RETURNING id_espacio, id_parqueadero, numero, estado, tipo_espacio;

-- name: BorrarEspacio :execrows
DELETE FROM espacios WHERE id_espacio = ?;