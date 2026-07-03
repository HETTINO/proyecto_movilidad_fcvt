package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Config centraliza TODO lo que antes estaba disperso o hardcodeado
// en main.go y en auth.go: puerto, ruta/driver de la DB y el secreto JWT.
type Config struct {
	Puerto      string
	DBDriver    string // "sqlite" (local) o "postgres" (docker)
	RutaDB      string // usado si DBDriver == "sqlite"
	PostgresDSN string // usado si DBDriver == "postgres"
	JWTSecreto  []byte
	JWTDuracion time.Duration
}

// Cargar lee el .env (si existe) y arma el Config con valores
// por defecto sensatos cuando alguna variable no está definida.
func Cargar() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("no se encontró .env, usando variables de entorno del sistema / defaults")
	}

	cfg := Config{
		Puerto:      obtenerEnv("PUERTO", "8080"),
		DBDriver:    obtenerEnv("DB_DRIVER", "sqlite"),
		RutaDB:      obtenerEnv("RUTA_DB", "parqueadero.db"),
		PostgresDSN: obtenerEnv("POSTGRES_DSN", ""),
		JWTSecreto:  []byte(obtenerEnv("JWT_SECRETO", "cualquier_cosa_secreta")),
		JWTDuracion: obtenerEnvDuracion("JWT_DURACION", 24*time.Hour),
	}

	return cfg
}

func obtenerEnv(clave, porDefecto string) string {
	if valor := os.Getenv(clave); valor != "" {
		return valor
	}
	return porDefecto
}

func obtenerEnvDuracion(clave string, porDefecto time.Duration) time.Duration {
	if valor := os.Getenv(clave); valor != "" {
		if d, err := time.ParseDuration(valor); err == nil {
			return d
		}
		log.Printf("valor inválido para %s, usando default", clave)
	}
	return porDefecto
}
