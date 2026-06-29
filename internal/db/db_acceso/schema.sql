-- Tabla 1: Usuario
CREATE TABLE Usuario (
    Cedula_int INTEGER PRIMARY KEY, 
    Nombre_usuario TEXT NOT NULL,
    Contrasena TEXT NOT NULL, -- Cambiado de Contraseña a Contrasena
    Email TEXT NOT NULL UNIQUE,
    Rol TEXT NOT NULL
);

-- Tabla 2: Vehiculos
CREATE TABLE Vehiculos (
    placa_vehiculo TEXT PRIMARY KEY, 
    id_usuario INTEGER NOT NULL,
    Tipo_usuario TEXT NOT NULL,
    Marca TEXT NOT NULL,
    Modelo TEXT NOT NULL,
    Color TEXT NOT NULL,
    Anio INTEGER NOT NULL, -- Cambiado de Año a Anio
    FOREIGN KEY (id_usuario) REFERENCES Usuario(Cedula_int) ON DELETE CASCADE
);

-- Tabla 3: Punto_de_acceso
CREATE TABLE Punto_de_acceso (
    ID_puntoacceso INTEGER PRIMARY KEY AUTOINCREMENT,
    Tipo_acceso TEXT NOT NULL,
    Ubicacion TEXT NOT NULL
);

-- Tabla 4: Acceso
CREATE TABLE Acceso (
    ID_acceso INTEGER PRIMARY KEY AUTOINCREMENT,
    Placa_vehiculo TEXT NOT NULL,
    ID_puntoacceso INTEGER NOT NULL,
    Tiempo_entrada DATETIME NOT NULL,
    Tiempo_salida DATETIME, 
    Estado TEXT NOT NULL,
    Observaciones TEXT,
    FOREIGN KEY (Placa_vehiculo) REFERENCES Vehiculos(placa_vehiculo),
    FOREIGN KEY (ID_puntoacceso) REFERENCES Punto_de_acceso(ID_puntoacceso)
);