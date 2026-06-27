package middleware

import (
	"context"
	"net/http"
	"strings"

	"proyecto/internal/service"
)

// Tipo personalizado para la clave del contexto (evita colisiones de nombres)
type contextKey string

const UsuarioIDKey contextKey = "usuario_id"

// AuthMiddleware recibe el AuthService inyectado y devuelve el middleware real.
// Así la lógica de verificación del JWT vive en el service (igual que Registrar/Login),
// y el middleware solo se encarga de leer el header y guardar el ID en el contexto.
func AuthMiddleware(authService *service.AuthService) func(http.Handler) http.Handler {
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

			// 3. Delegar la verificación al AuthService (igual que generarToken vive en el service)
			usuarioID, err := authService.VerificarToken(partes[1])
			if err != nil {
				http.Error(respuesta, `{"error": "Token invalido o expirado"}`, http.StatusUnauthorized)
				return
			}

			// 4. Inyectar el ID del usuario en el contexto de la petición y continuar
			ctx := context.WithValue(peticion.Context(), UsuarioIDKey, usuarioID)
			siguiente.ServeHTTP(respuesta, peticion.WithContext(ctx))
		})
	}
}
