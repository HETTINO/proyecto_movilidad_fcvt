package middleware

import (
	"context"
	"net/http"
	"proyecto_movilidad_fcvt/internal/service"
	"strings"
)

type claveContext string

const ClaveUsuarioID claveContext = "usuarioID"
const ClaveRol claveContext = "rol"

func Auth(auth *service.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			encabezado := r.Header.Get("Authorization")
			partes := strings.SplitN(encabezado, " ", 2)
			if len(partes) != 2 || partes[0] != "Bearer" {
				respondeNoAutorizado(w)
				return
			}
			claims, err := auth.ValidarTokenClaims(partes[1])
			if err != nil {
				respondeNoAutorizado(w)
				return
			}
			ctx := context.WithValue(r.Context(), ClaveUsuarioID, claims.Cedula)
			ctx = context.WithValue(ctx, ClaveRol, claims.Rol)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// RequireRol se coloca DESPUÉS de Auth en la cadena de middlewares y exige que
// el rol guardado en el contexto (por Auth) esté entre los roles permitidos.
// Uso: r.Use(middleware.Auth(authService)); r.Use(middleware.RequireRol("admin"))
func RequireRol(rolesPermitidos ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rol, _ := r.Context().Value(ClaveRol).(string)

			for _, permitido := range rolesPermitidos {
				if rol == permitido {
					next.ServeHTTP(w, r)
					return
				}
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			_, _ = w.Write([]byte(`{"error":"no tienes permisos para realizar esta accion"}`))
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
			_, _ = w.Write([]byte(`{"error":"token requerido"}`))
			return
		}

		next.ServeHTTP(w, r)
	})
}
