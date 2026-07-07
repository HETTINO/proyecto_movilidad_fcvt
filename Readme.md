# Proyecto Movilidad FCVT

API REST en Go para la gestión de movilidad universitaria: acceso vehicular (usuarios, vehículos y puntos de acceso), parqueadero (espacios y ocupaciones) y transporte interno (rutas, carritos, paradas y solicitudes).

Proyecto semestral — TDI-601 Aplicaciones Web II, ULEAM.

## Stack

- **Lenguaje:** Go 1.26
- **Router:** Chi
- **ORM:** GORM (SQLite en desarrollo, PostgreSQL en producción)
- **Auth:** JWT (golang-jwt)
- **Tests:** testify (mocks + asserts)
- **Contenedores:** Docker + Docker Compose
- **CI/CD:** GitHub Actions (lint → test → build)

## Cómo correrlo

```bash
git clone <url-del-repo>
cd proyecto_movilidad_fcvt
cp .env.example .env
docker-compose up
```

La API queda disponible en `http://localhost:8080/api/v1`. El seed de datos se carga automáticamente al iniciar (usuarios, vehículos, parqueaderos, rutas de ejemplo).

Colección de Postman: `/postman/movilidad_fcvt.postman_collection.json` — importa y usa el environment con `base_url = http://localhost:8080/api/v1`.

## Arquitectura

El proyecto está organizado en 3 módulos de dominio, cada uno con su propia capa de repositorio, servicio y handler:

## Autores

- [Shirley Cedeño] — módulo Acceso
- [Eduardo Lopez] — módulo Parqueadero
- [Cristina Cedeño] — módulo Transporte