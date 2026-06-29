CREATE TABLE rutas (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  nombre TEXT NOT NULL,
  descripcion TEXT
);

CREATE TABLE carritos (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  nombre_carrito TEXT NOT NULL,
  capacidad INTEGER NOT NULL,
  estado TEXT,
  ruta_id INTEGER
);

CREATE TABLE paradas (
  id_parada INTEGER PRIMARY KEY AUTOINCREMENT,
  nombre TEXT NOT NULL,
  latitud REAL,
  longitud REAL
);

CREATE TABLE locaciones (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  latitud REAL NOT NULL,
  longitud REAL NOT NULL,
  time_stamp DATETIME,
  carrito_id INTEGER NOT NULL
);

CREATE TABLE solicitudes (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  cedula_usuario TEXT NOT NULL,
  cant_personas INTEGER,
  parada_origen INTEGER,
  punto_destino TEXT,
  estado TEXT,
  id_carrito INTEGER
);