package middleware

import (
	"context"
	"net/http"
	"proyecto_movilidad_fcvt/internal/service"
	"strings"
)

type claveContext string

const ClaveUsuarioID claveContext = "usuarioID"

func Auth(auth *service.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			encabezado := r.Header.Get("Authorization")
			partes := strings.SplitN(encabezado, " ", 2)
			if len(partes) != 2 || partes[0] != "Bearer" {
				respondeNoAutorizado(w)
				return
			}
			usuarioid, err := auth.ValidarToken(partes[1])
			if err != nil {
				respondeNoAutorizado(w)
				return
			}
			ctx := context.WithValue(r.Context(), ClaveUsuarioID, usuarioid)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func respondeNoAutorizado(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	_, _ = w.Write([]byte(`{"error":"Token inexistente o invalido"}`))
}

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error":"token requerido"}`))
			return
		}

		next.ServeHTTP(w, r)
	})
}
