-- name: ListarCarritos :many
SELECT id, nombre_carrito, capacidad, estado, ruta_id FROM carritos;

-- name: BuscarCarritoPorID :one
SELECT id, nombre_carrito, capacidad, estado, ruta_id FROM carritos WHERE id = ?;

-- name: CrearCarrito :one
INSERT INTO carritos (nombre_carrito, capacidad, estado, ruta_id) VALUES (?, ?, ?, ?) RETURNING id, nombre_carrito, capacidad, estado, ruta_id;

-- name: ActualizarCarrito :one
UPDATE carritos SET nombre_carrito = ?, capacidad = ?, estado = ?, ruta_id = ? WHERE id = ? RETURNING id, nombre_carrito, capacidad, estado, ruta_id;

-- name: BorrarCarrito :execrows
DELETE FROM carritos WHERE id = ?;

-- name: ListarCarritosPorRuta :many
SELECT id, nombre_carrito, capacidad, estado, ruta_id FROM carritos WHERE ruta_id = ?;