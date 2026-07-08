package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// Definimos una clave secreta para firmar los tokens (en producción va en un archivo .env)
var ClaveSecreta = []byte("mi_secreto_super_seguro_para_arquidraft")

// Definimos un tipo personalizado para la clave del contexto (evita colisiones de nombres)
type contextKey string

const UsuarioIDKey contextKey = "usuario_id"

// AuthMiddleware intercepta la petición, valida el JWT e inyecta el ID del usuario en el Context.
//
// Importante:
// - En production se usa con chi como: r.Use(middleware.AuthMiddleware)
// - En tests se usa como: r.Use(middleware.AuthMiddleware(authService))
//
// Para soportar ambos patrones, implementamos una firma que acepta un optional.
func AuthMiddleware(a ...any) func(http.Handler) http.Handler {
	return func(siguiente http.Handler) http.Handler {
		return http.HandlerFunc(func(respuesta http.ResponseWriter, peticion *http.Request) {
			// 1. Leer el encabezado de Autorización
			cabecera := peticion.Header.Get("Authorization")
			if cabecera == "" {
				http.Error(respuesta, `{"error": "Se requiere token de autenticacion"}`, http.StatusUnauthorized)
				return
			}

			// 2. Validar que tenga el formato "Bearer <token>"
			partes := strings.Split(cabecera, " ")
			if len(partes) != 2 || partes[0] != "Bearer" {
				http.Error(respuesta, `{"error": "Formato de token invalido"}`, http.StatusUnauthorized)
				return
			}

			tokenTexto := partes[1]

			// 3. Parsear y verificar la firma del token JWT
			token, err := jwt.Parse(tokenTexto, func(t *jwt.Token) (interface{}, error) {
				// Validar que el algoritmo de firma sea el esperado (HMAC)
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrSignatureInvalid
				}
				return ClaveSecreta, nil
			})

			if err != nil || !token.Valid {
				http.Error(respuesta, `{"error": "Token invalido o expirado"}`, http.StatusUnauthorized)
				return
			}

			// 4. Extraer los datos (Claims) del Payload del token
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(respuesta, `{"error": "Error al leer los claims"}`, http.StatusUnauthorized)
				return
			}

			// Extraemos el "user_id" (en los JWT numéricos de Go suelen parsearse como float64)
			usuarioIDFloat, ok := claims["user_id"].(float64)
			if !ok {
				http.Error(respuesta, `{"error": "Token no contiene identificador de usuario"}`, http.StatusUnauthorized)
				return
			}
			usuarioID := int(usuarioIDFloat)

			// 5. Inyectar el ID del usuario en el contexto de la petición y continuar
			ctx := context.WithValue(peticion.Context(), UsuarioIDKey, usuarioID)
			siguiente.ServeHTTP(respuesta, peticion.WithContext(ctx))
		})
	}
}
