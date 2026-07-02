package service

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	"proyecto_movilidad_fcvt/internal/storage/storage_acceso"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretJWT = []byte("cualquier_cosa_secreta")

const duracionToken = 24 * time.Hour

type Claims struct {
	Cedula string `json:"cedula"` // Usamos la Cédula como identificador en tu módulo
	jwt.RegisteredClaims
}

type AuthService struct {
	repo storage_acceso.UsuarioRepository
}

func NewAuthService(repo storage_acceso.UsuarioRepository) *AuthService {
	return &AuthService{repo: repo}
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
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duracionToken)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretJWT)
}

func (s *AuthService) ValidarToken(tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrCredencialesInvalidas
		}
		return secretJWT, nil
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
