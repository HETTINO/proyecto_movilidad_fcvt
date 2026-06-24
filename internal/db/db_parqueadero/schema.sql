CREATE TABLE parqueaderos (
    id_parqueadero INTEGER PRIMARY KEY AUTOINCREMENT,
    nombre         TEXT    NOT NULL,
    capacidad      INTEGER NOT NULL,
    tipo           TEXT    NOT NULL
);

CREATE TABLE espacios (
    id_espacio     INTEGER PRIMARY KEY AUTOINCREMENT,
    id_parqueadero INTEGER NOT NULL,
    numero         INTEGER NOT NULL,
    estado         TEXT    NOT NULL,
    tipo_espacio   TEXT    NOT NULL,
    FOREIGN KEY (id_parqueadero) REFERENCES parqueaderos(id_parqueadero)
);

CREATE TABLE ocupacion (
    id_ocupacion   INTEGER PRIMARY KEY AUTOINCREMENT,
    placa_vehiculo TEXT    NOT NULL CHECK(length(placa_vehiculo) <= 10),
    id_espacio     INTEGER NOT NULL,
    id_acceso      INTEGER NOT NULL,
    hora_inicio    DATETIME NOT NULL,
    hora_fin       DATETIME,
    FOREIGN KEY (id_espacio) REFERENCES espacios(id_espacio)
);