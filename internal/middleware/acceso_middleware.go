package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggerAcceso registra detalles de auditoría para el módulo de accesos
func LoggerAcceso(siguiente http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		inicio := time.Now()
		log.Printf("[AUDITORÍA ACCESOS] Petición: %s %s desde %s", r.Method, r.URL.Path, r.RemoteAddr)

		siguiente(w, r)

		log.Printf("[AUDITORÍA ACCESOS] Petición procesada en %v", time.Since(inicio))
	}
}
