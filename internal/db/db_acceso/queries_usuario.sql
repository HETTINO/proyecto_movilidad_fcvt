-- name: ListarUsuarios :many
SELECT Cedula_int, Nombre_usuario, Contrasena, Email, Rol
FROM Usuario;

-- name: ObtenerUsuarioPorCedula :one
SELECT Cedula_int, Nombre_usuario, Contrasena, Email, Rol
FROM Usuario
WHERE Cedula_int = ? LIMIT 1;

-- name: RegistrarUsuario :one
INSERT INTO Usuario (Cedula_int, Nombre_usuario, Contrasena, Email, Rol)
VALUES (?, ?, ?, ?, ?)
RETURNING cedula_int, nombre_usuario, contrasena, email, rol;

-- name: ActualizarUsuario :one
UPDATE Usuario
SET Nombre_usuario = ?, Contrasena = ?, Email = ?, Rol = ?
WHERE Cedula_int = ?
RETURNING cedula_int, nombre_usuario, contrasena, email, rol;

-- name: BorrarUsuario :execrows
DELETE FROM Usuario
WHERE Cedula_int = ?;