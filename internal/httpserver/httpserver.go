package httpserver

import (
	"context"
	"log"
	"net/http"
	"time"
)

// Server envuelve http.Server para exponer Iniciar/Apagar de forma simple.
type Server struct {
	srv *http.Server
}

// Option configura parámetros OPCIONALES del servidor HTTP.
type Option func(*http.Server)

func WithReadTimeout(d time.Duration) Option {
	return func(s *http.Server) { s.ReadTimeout = d }
}

func WithWriteTimeout(d time.Duration) Option {
	return func(s *http.Server) { s.WriteTimeout = d }
}

func WithIdleTimeout(d time.Duration) Option {
	return func(s *http.Server) { s.IdleTimeout = d }
}

// Nuevo crea el servidor con defaults razonables; opts es variádico,
// así que Nuevo(puerto, handler) sigue compilando sin opciones.
func Nuevo(puerto string, handler http.Handler, opts ...Option) *Server {
	s := &http.Server{
		Addr:         ":" + puerto,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	for _, opt := range opts {
		opt(s)
	}
	return &Server{srv: s}
}

// IniciarConGracefulShutdown arranca el servidor en background y bloquea
// hasta que ctx se cancele (Ctrl+C / SIGTERM), momento en el que apaga
// ordenadamente dando hasta `timeout` a las peticiones en curso.
func (s *Server) IniciarConGracefulShutdown(ctx context.Context, timeout time.Duration) {
	go func() {
		log.Printf("Servidor escuchando en http://localhost%s\n", s.srv.Addr)
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error en ListenAndServe: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("señal de apagado recibida, cerrando conexiones en curso...")

	ctxApagado, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := s.srv.Shutdown(ctxApagado); err != nil {
		log.Fatalf("error al apagar el servidor: %v", err)
	}
	log.Println("servidor apagado correctamente")
}
