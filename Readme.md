Proyecto Movilidad FCVT

API REST en Go para la gestión de movilidad universitaria: acceso vehicular (usuarios, vehículos y puntos de acceso), parqueadero (espacios y ocupaciones) y transporte interno (rutas, carritos, paradas y solicitudes).

Proyecto semestral — TDI-601 Aplicaciones Web II, ULEAM.

Stack


Lenguaje: Go 1.26
Router: Chi
ORM: GORM (SQLite en desarrollo, PostgreSQL en producción)
Auth: JWT (golang-jwt)
Tests: testify (mocks + asserts)
Contenedores: Docker + Docker Compose
CI/CD: GitHub Actions (lint → test → build)

Arquitectura

El proyecto está organizado en 3 módulos de dominio, cada uno con su propia capa de repositorio, servicio y handler:

cmd/main.go              → arranque, wiring de dependencias, migraciones y rutas
internal/
├── config/               → carga de variables de entorno (.env)
├── httpserver/            → servidor HTTP con graceful shutdown
├── middleware/            → auth (JWT) y CORS
├── service/               → AuthService compartido (login/registro)
├── modelos/               → structs GORM de las 3 entidades del dominio
│
├── handlers | service | storage _acceso        → Usuario, Vehículo, Punto de acceso, Acceso
├── handlers | service | storage _parqueadero    → Parqueadero, Espacio, Ocupación
└── handlers | service | storage _transporte     → Ruta, Carrito, Parada, Locación, Solicitud

Cada módulo sigue el flujo handler → service → repository (interface), lo que permite testear los servicios con mocks sin depender de una base de datos real.