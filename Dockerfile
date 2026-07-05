# ============================================================================
# Build multi-stage: una etapa "builder" compila el binario y una etapa final
# minima solo lo copia. Resultado: imagen pequeña y sin el toolchain de Go.
# ============================================================================

# ---- Etapa 1: builder ----
FROM golang:1.26-alpine AS builder
WORKDIR /src

# Cachear dependencias: copiar primero los modulos y descargar.
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto del codigo y compilar.
COPY . .
# CGO_ENABLED=0 produce un binario estatico (glebarez/sqlite y gorm/postgres
# son Go puro, asi que no hace falta CGO). GOOS=linux para el runner.
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/servidor ./cmd/main.go

# ---- Etapa 2: runner (imagen final minima) ----
FROM alpine:3.20
# ca-certificates por si en el futuro se conecta por TLS; tzdata para zonas horarias
# (útil porque usas TimeZone=America/Guayaquil en el DSN de Postgres).
RUN apk add --no-cache ca-certificates tzdata
# Usuario no-root por seguridad.
RUN adduser -D -u 10001 appuser
WORKDIR /app
COPY --from=builder /bin/servidor /app/servidor
USER appuser
EXPOSE 8080
ENTRYPOINT ["/app/servidor"]