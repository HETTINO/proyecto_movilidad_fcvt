package service

import (
	"proyecto_movilidad_fcvt/internal/modelos"
	storage "proyecto_movilidad_fcvt/internal/storage/storage_parqueadero"

	"strings"

	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretJWT = []byte("cualquier_cosa_secreta")

const duracionToken = 24 * time.Hour

type Claims struct {
	UsuarioID int `json:"uid"`
	jwt.RegisteredClaims
}

type AuthService struct {
	repo storage.UsuarioRepository
}

func NewAuthService(repo storage.UsuarioRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Registrar(email, password string) (modelos.Usuario, error) {
	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)
	if email == "" || password == "" {
		return modelos.Usuario{}, ErrCredencialesInvalidas
	}
	if _, existe := s.repo.BuscarUsuarioPorEmail(email); existe {
		return modelos.Usuario{}, ErrEmailenUso
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return modelos.Usuario{}, err
	}

	return s.repo.CrearUsuario(modelos.Usuario{
		Email:    email,
		Password: string(hash),
	})
}

func (s *AuthService) Login(email, password string) (string, error) {
	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)
	if email == "" || password == "" {
		return "", ErrCredencialesInvalidas
	}
	u, existe := s.repo.BuscarUsuarioPorEmail(email)
	if !existe {
		return "", ErrCredencialesInvalidas
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return "", ErrCredencialesInvalidas
	}

	return s.generarToken(u)
}

func (s *AuthService) generarToken(u modelos.Usuario) (string, error) {
	claims := &Claims{
		UsuarioID: int(u.ID),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duracionToken)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretJWT)
}

func (s *AuthService) ValidarToken(tokenStr string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrCredencialesInvalidas
		}
		return secretJWT, nil
	})
	if err != nil || !token.Valid {
		return 0, ErrCredencialesInvalidas
	}
	Claims, ok := token.Claims.(*Claims)
	if !ok {
		return 0, ErrCredencialesInvalidas
	}
	return Claims.UsuarioID, nil
}
