// Archivo: internal/service/auth.go
package service

import (
	"strings"
	"time"

	"proyecto/internal/models"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// IMPORTANTE: en producción esto debe venir de una variable de entorno,
// no hardcodeado. Por ahora lo dejamos igual que en tu middleware para
// que ambos firmen/verifiquen con la misma clave.
var claveSecretaJWT = []byte("mi_secreto_super_seguro_para_arquidraft")

var duracionToken = time.Hour * 24

// RepositorioUsuarios es la interfaz mínima que AuthService necesita del
// almacenamiento. *storage.SQLiteStorage ya implementa estos dos métodos
// (los tiene en sqlite.go), así que NO hay que tocar storage para nada:
// simplemente cumple la interfaz "gratis". Esto es lo que nos permite
// inyectar un mock en los tests sin tocar la base de datos real.
type RepositorioUsuarios interface {
	BuscarUsuarioPorEmail(email string) (models.Usuario, bool)
	CrearUsuario(u models.Usuario) models.Usuario
}

type AuthService struct {
	almacen RepositorioUsuarios
}

func NuevoAuthService(a RepositorioUsuarios) *AuthService {
	return &AuthService{almacen: a}
}

// Registrar crea un nuevo usuario validando email/contraseña y hasheando con bcrypt.
func (s *AuthService) Registrar(u models.Usuario) (models.Usuario, error) {
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))

	if strings.TrimSpace(u.Nombre) == "" {
		return models.Usuario{}, ErrNombreObligatorio
	}
	if u.Email == "" {
		return models.Usuario{}, ErrEmailObligatorio
	}
	if strings.TrimSpace(u.Password) == "" {
		return models.Usuario{}, ErrCredencialesInvalidas
	}

	// Regla de negocio real que vamos a probar con mock:
	// si el email ya existe, NUNCA debe llegar a CrearUsuario.
	if _, existe := s.almacen.BuscarUsuarioPorEmail(u.Email); existe {
		return models.Usuario{}, ErrEmailEnUso
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.Usuario{}, err
	}
	u.Password = string(hash)

	return s.almacen.CrearUsuario(u), nil
}

// Login verifica credenciales y devuelve un token JWT firmado.
func (s *AuthService) Login(email, password string) (string, error) {
	email = strings.TrimSpace(strings.ToLower(email))

	usuario, existe := s.almacen.BuscarUsuarioPorEmail(email)
	if !existe {
		return "", ErrCredencialesInvalidas
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usuario.Password), []byte(password)); err != nil {
		return "", ErrCredencialesInvalidas
	}

	return s.generarToken(usuario)
}

// generarToken usa jwt.MapClaims con la clave "user_id", igual que la espera
// tu AuthMiddleware actual (claims["user_id"].(float64)).
func (s *AuthService) generarToken(u models.Usuario) (string, error) {
	claims := jwt.MapClaims{
		"user_id": u.ID,
		"rol":     u.Rol,
		"exp":     jwt.NewNumericDate(time.Now().Add(duracionToken)).Unix(),
		"iat":     jwt.NewNumericDate(time.Now()).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(claveSecretaJWT)
}

// VerificarToken valida la firma y expiración del token, y devuelve el ID
// del usuario contenido en sus claims. El middleware llama a este método
// en lugar de parsear el JWT directamente.
func (s *AuthService) VerificarToken(tokenTexto string) (int, error) {
	token, err := jwt.Parse(tokenTexto, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrCredencialesInvalidas
		}
		return claveSecretaJWT, nil
	})
	if err != nil || !token.Valid {
		return 0, ErrCredencialesInvalidas
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, ErrCredencialesInvalidas
	}

	usuarioIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, ErrCredencialesInvalidas
	}

	return int(usuarioIDFloat), nil
}
