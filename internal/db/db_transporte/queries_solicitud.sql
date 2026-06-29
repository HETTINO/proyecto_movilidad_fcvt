-- name: ListarSolicitudes :many
SELECT id, cedula_usuario, cant_personas, parada_origen, punto_destino, estado, id_carrito FROM solicitudes;

-- name: BuscarSolicitudPorID :one
SELECT id, cedula_usuario, cant_personas, parada_origen, punto_destino, estado, id_carrito FROM solicitudes WHERE id = ?;

-- name: CrearSolicitud :one
INSERT INTO solicitudes (cedula_usuario, cant_personas, parada_origen, punto_destino, estado, id_carrito) VALUES (?, ?, ?, ?, ?, ?) RETURNING id, cedula_usuario, cant_personas, parada_origen, punto_destino, estado, id_carrito;

-- name: ActualizarSolicitud :one
UPDATE solicitudes SET cedula_usuario = ?, cant_personas = ?, parada_origen = ?, punto_destino = ?, estado = ?, id_carrito = ? WHERE id = ? RETURNING id, cedula_usuario, cant_personas, parada_origen, punto_destino, estado, id_carrito;

-- name: BorrarSolicitud :execrows
DELETE FROM solicitudes WHERE id = ?;