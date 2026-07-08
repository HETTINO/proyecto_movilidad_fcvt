# Proyecto Movilidad FCVT

API REST en Go para la gestión de movilidad universitaria: acceso vehicular (usuarios, vehículos y puntos de acceso), parqueadero (espacios y ocupaciones) y transporte interno (rutas, carritos, paradas y solicitudes).

Proyecto semestral — TDI-601 Aplicaciones Web II, ULEAM. Hito 3.

## Stack

- **Lenguaje:** Go 1.26
- **Router:** Chi
- **ORM:** GORM (SQLite en desarrollo, PostgreSQL en Docker)
- **Auth:** JWT (golang-jwt), roles `admin` / `usuario`
- **Tests:** testify (mocks + asserts)
- **Contenedores:** Docker (multi-stage) + Docker Compose (API + PostgreSQL)
- **CI/CD:** GitHub Actions — `lint` (gofmt + go vet + golangci-lint) → `test` (go test -race + cobertura) → `build` (binario + imagen Docker)

## Cómo correrlo

### Con Docker (recomendado — así se evalúa el gate G2)

```bash
git clone <url-del-repo>
cd proyecto_movilidad_fcvt
docker-compose up
```

No se requiere ningún paso manual adicional: `docker-compose.yml` ya inyecta las variables de entorno (`DB_DRIVER=postgres`, credenciales de la base, JWT) y espera a que PostgreSQL esté *healthy* antes de levantar la API.

### En local sin Docker (SQLite)

```bash
cp .env.example .env
go run ./cmd/main.go
```

En ambos casos la API queda disponible en `http://localhost:8080/api/v1`, y el seed de datos (usuarios, vehículos, parqueaderos, rutas de ejemplo) se carga automáticamente si la base está vacía.

### Colección de Postman

`/postman/movilidad_fcvt.postman_collection.json` — impórtala junto con el environment `base_url = http://localhost:8080/api/v1`. Incluye el flujo de Login (guarda el token automáticamente) y las peticiones de los 3 módulos.

## Arquitectura

### Estructura del repositorio

```
proyecto_movilidad_fcvt
├── .github
│   └── workflows
│       └── ci.yml                 # Pipeline: lint → test → build
├── cmd
│   └── main.go                    # Punto de entrada: migraciones, DI, router
├── internal
│   ├── config
│   │   └── config.go               # Carga de variables de entorno
│   ├── middleware
│   │   ├── auth.go                 # Middleware Auth + RequireRol
│   │   └── cors.go
│   ├── httpserver
│   │   └── httpserver.go           # Wrapper de http.Server + graceful shutdown
│   ├── modelos
│   │   ├── acceso.go                # Usuario, Vehiculo, PuntoDeAcceso, Acceso
│   │   ├── parqueadero.go           # Parqueadero, Espacio, Ocupacion
│   │   └── transporte.go            # Ruta, Parada, Carrito, Locacion, Solicitud
│   ├── handlers
│   │   ├── handler_acceso/          # HTTP: usuarios, vehículos, puntos de acceso
│   │   ├── handler_parqueadero/     # HTTP: parqueaderos, espacios, ocupaciones
│   │   └── handler_transporte/      # HTTP: rutas, carritos, paradas, solicitudes
│   ├── service
│   │   ├── auth.go                  # JWT: registro, login, validación de claims
│   │   ├── service_acceso/          # Reglas de negocio del módulo Acceso
│   │   ├── service_parqueadero/     # Reglas de negocio del módulo Parqueadero
│   │   └── service_transporte/      # Reglas de negocio del módulo Transporte
│   └── storage
│       ├── storage_acceso/          # Repositorios (interface + GORM) de Acceso
│       ├── storage_parqueadero/     # Repositorios (interface + GORM) de Parqueadero
│       └── storage_transporte/      # Repositorios (interface + GORM) de Transporte
├── postman
│   └── movilidad_fcvt.postman_collection.json
├── docker-compose.yml               # API + PostgreSQL + healthcheck
├── Dockerfile                        # Build multi-stage
├── go.mod / go.sum
├── .env.example
└── Readme.md
```

Cada módulo de dominio (`acceso`, `parqueadero`, `transporte`) repite la misma subestructura de 3 capas (`handlers/`, `service/`, `storage/`), lo que hace que el proyecto sea predecible: para entender cualquier módulo nuevo alcanza con reconocer este mismo patrón tres veces.

### Flujo de una request (handler → service → repository)

```
Request HTTP
     │
     ▼
┌─────────────────────┐
│      Handler         │  internal/handlers/handler_<modulo>/
│  (HTTP, JSON, chi)    │  Decodifica el request, llama al Service, responde JSON
└──────────┬───────────┘
           ▼
┌─────────────────────┐
│      Service          │  internal/service/service_<modulo>/
│  (reglas de negocio)   │  Valida datos, aplica reglas, traduce errores de dominio
└──────────┬───────────┘
           ▼
┌─────────────────────┐
│    Repository          │  internal/storage/storage_<modulo>/
│  (interface + GORM)    │  CRUD contra la base (SQLite o Postgres, vía GORM)
└──────────┬───────────┘
           ▼
        Base de datos
```

Cada capa depende de la interface de la capa inferior (no de su implementación concreta), y la inyección de dependencias se arma en `cmd/main.go`: ahí se instancian los repositorios, se inyectan en los servicios, y los servicios en los handlers.

Los middlewares (`internal/middleware`) se aplican a nivel de router en `main.go`:
- `Auth`: valida el JWT y agrega `cédula` y `rol` al contexto de la request.
- `RequireRol("admin")`: restringe rutas puntuales (p. ej. borrar usuarios, gestionar puntos de acceso) a un rol específico.

## Módulos y endpoints

Todas las rutas están bajo el prefijo `/api/v1` y requieren `Authorization: Bearer <token>` excepto `auth/register` y `auth/login`.

### Auth (compartido, sin responsable único)

| Método | Ruta | Descripción |
|---|---|---|
| POST | `/auth/register` | Registra un usuario |
| POST | `/auth/login` | Devuelve el JWT |

### Acceso — responsable: **Shirley Cedeño**

| Método | Ruta | Descripción |
|---|---|---|
| GET/POST | `/usuarios` | Listar / crear usuarios |
| GET/PUT | `/usuarios/{id}` | Obtener / actualizar usuario |
| DELETE | `/usuarios/{id}` | Borrar usuario — solo rol `admin` |
| GET/POST | `/vehiculos` | Listar / crear vehículos |
| GET/PUT/DELETE | `/vehiculos/{placa}` | Obtener / actualizar / borrar vehículo |
| GET | `/puntos-acceso` | Listar puntos de acceso |
| GET | `/puntos-acceso/{id}` | Obtener punto de acceso |
| POST/PUT/DELETE | `/puntos-acceso` / `/puntos-acceso/{id}` | Gestionar puntos de acceso — solo rol `admin` |
| GET/POST | `/accesos` | Listar / registrar accesos (entrada/salida) |
| GET/PUT/DELETE | `/accesos/{id}` | Obtener / actualizar / borrar acceso |

### Parqueadero — responsable: **Eduardo López**

| Método | Ruta | Descripción |
|---|---|---|
| GET/POST | `/parqueaderos` | Listar / crear parqueaderos |
| GET/PUT/DELETE | `/parqueaderos/{id}` | Obtener / actualizar / borrar parqueadero |
| GET/POST | `/espacios` | Listar / crear espacios |
| GET/PUT/DELETE | `/espacios/{id}` | Obtener / actualizar / borrar espacio (bloqueado si tiene ocupaciones activas) |
| GET/POST | `/ocupaciones` | Listar / crear ocupaciones |
| GET/PUT/DELETE | `/ocupaciones/{id}` | Obtener / actualizar / borrar ocupación |
| PATCH | `/ocupaciones/{id}/liberar` | Cierra la ocupación (asigna `hora_fin`) |

### Transporte — responsable: **Cristina Cedeño**

| Método | Ruta | Descripción |
|---|---|---|
| GET/POST | `/rutas` | Listar / crear rutas |
| GET/PUT/DELETE | `/rutas/{id}` | Obtener / actualizar / borrar ruta |
| GET/POST | `/carritos` | Listar / crear carritos |
| GET/PUT/DELETE | `/carritos/{id}` | Obtener / actualizar / borrar carrito |
| GET/POST | `/paradas` | Listar / crear paradas |
| GET/PUT/DELETE | `/paradas/{id}` | Obtener / actualizar / borrar parada |
| GET/POST | `/locaciones` | Listar / registrar ubicación de un carrito |
| GET | `/locaciones/carrito/{id}` | Última ubicación conocida del carrito |
| GET | `/tiempo-estimado` | Tiempo estimado de llegada |
| GET/POST | `/solicitudes` | Listar / crear solicitudes de transporte |
| GET/PUT/DELETE | `/solicitudes/{id}` | Obtener / actualizar / borrar solicitud |

## Tests

Cada módulo cuenta con tests unitarios de servicio (mocks de repositorio con testify) y tests de integración de handlers. Para correrlos:

```bash
go test ./... -race -coverprofile=coverage.out -covermode=atomic
go tool cover -func=coverage.out | tail -1
```

## Autores

| Integrante | Módulo |
|---|---|
| Shirley Cedeño | Acceso |
| Eduardo López | Parqueadero |
| Cristina Cedeño | Transporte |
