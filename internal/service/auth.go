package service

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	"proyecto_movilidad_fcvt/internal/storage/storage_acceso"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// valores por defecto: se usan si nadie pasa una Option
var secretoPorDefecto = []byte("cualquier_cosa_secreta")

const duracionPorDefecto = 24 * time.Hour

type Claims struct {
	Cedula string `json:"cedula"` // Usamos la Cédula como identificador en tu módulo
	jwt.RegisteredClaims
}

type AuthService struct {
	repo        storage_acceso.UsuarioRepository
	secretoJWT  []byte
	duracionJWT time.Duration
}

// AuthOption configura parámetros OPCIONALES de AuthService.
type AuthOption func(*AuthService)

// WithSecreto permite inyectar el secreto JWT (p.ej. desde config/.env)
// en vez de usar el valor global hardcodeado.
func WithSecreto(secreto []byte) AuthOption {
	return func(s *AuthService) {
		if len(secreto) > 0 {
			s.secretoJWT = secreto
		}
	}
}

// WithDuracion permite configurar cuánto dura el token.
func WithDuracion(d time.Duration) AuthOption {
	return func(s *AuthService) {
		if d > 0 {
			s.duracionJWT = d
		}
	}
}

// NewAuthService sigue aceptando NewAuthService(repo) sin romper nada:
// opts es variádico, así que las llamadas existentes compilan igual.
func NewAuthService(repo storage_acceso.UsuarioRepository, opts ...AuthOption) *AuthService {
	s := &AuthService{
		repo:        repo,
		secretoJWT:  secretoPorDefecto,
		duracionJWT: duracionPorDefecto,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *AuthService) Registrar(cedula, nombre, email, password, rol string) (modelos.Usuario, error) {
	cedula = strings.TrimSpace(cedula)
	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)

	if cedula == "" || email == "" || password == "" {
		return modelos.Usuario{}, ErrCredencialesInvalidas
	}

	// Buscamos si ya existe por Cédula en tu MemoriaAcceso
	if _, existe := s.repo.BuscarUsuarioPorCedula(cedula); existe {
		return modelos.Usuario{}, ErrEmailenUso
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return modelos.Usuario{}, err
	}

	nuevoUsuario := modelos.Usuario{
		Cedula:     cedula,
		Nombre:     strings.TrimSpace(nombre),
		Email:      email,
		Contrasena: string(hash), // Guardamos el hash seguro
		Rol:        strings.TrimSpace(rol),
	}

	return s.repo.CrearUsuario(nuevoUsuario), nil
}

func (s *AuthService) Login(cedula, password string) (string, error) {
	cedula = strings.TrimSpace(cedula)
	password = strings.TrimSpace(password)

	if cedula == "" || password == "" {
		return "", ErrCredencialesInvalidas
	}

	u, existe := s.repo.BuscarUsuarioPorCedula(cedula)
	if !existe {
		return "", ErrCredencialesInvalidas
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Contrasena), []byte(password)); err != nil {
		return "", ErrCredencialesInvalidas
	}

	return s.generarToken(u)
}

func (s *AuthService) generarToken(u modelos.Usuario) (string, error) {
	claims := &Claims{
		Cedula: u.Cedula,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.duracionJWT)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secretoJWT)
}

func (s *AuthService) ValidarToken(tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrCredencialesInvalidas
		}
		return s.secretoJWT, nil
	})
	if err != nil || !token.Valid {
		return "", ErrCredencialesInvalidas
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return "", ErrCredencialesInvalidas
	}
	return claims.Cedula, nil
}
