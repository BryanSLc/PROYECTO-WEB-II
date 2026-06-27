// Archivo: internal/handlers/auth_http_test.go
package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"

	"proyecto/internal/handlers/middleware"
	"proyecto/internal/models"
	"proyecto/internal/service"
)

// fakeRepositorioUsuarios es un doble en memoria: SÍ guarda los datos
// (a diferencia del mock del test de service, que solo cuenta llamadas),
// pero no es la base de datos real (no usa GORM ni SQLite).
type fakeRepositorioUsuarios struct {
	usuarios []models.Usuario
}

func (f *fakeRepositorioUsuarios) BuscarUsuarioPorEmail(email string) (models.Usuario, bool) {
	for _, u := range f.usuarios {
		if u.Email == email {
			return u, true
		}
	}
	return models.Usuario{}, false
}

func (f *fakeRepositorioUsuarios) CrearUsuario(u models.Usuario) models.Usuario {
	u.ID = len(f.usuarios) + 1
	f.usuarios = append(f.usuarios, u)
	return u
}

// TestRegistrar_HTTP_Exitoso prueba el handler completo (HTTP real, en
// memoria) usando httptest: un registro válido debe responder 201 Created.
func TestRegistrar_HTTP_Exitoso(t *testing.T) {
	fake := &fakeRepositorioUsuarios{}
	servidor := &Servidor{AuthService: service.NuevoAuthService(fake)}

	cuerpo := `{"nombre":"Carlos","apellido":"Ramirez","email":"carlos@uleam.edu.ec","password":"clave123","semestre":4,"telefono":"099","rol":"estudiante"}`
	peticion := httptest.NewRequest(http.MethodPost, "/api/v1/auth/registro", bytes.NewBufferString(cuerpo))
	grabadora := httptest.NewRecorder()

	servidor.Registrar(grabadora, peticion)

	if grabadora.Code != http.StatusCreated {
		t.Fatalf("se esperaba 201, se obtuvo %d. Body: %s", grabadora.Code, grabadora.Body.String())
	}

	if len(fake.usuarios) != 1 {
		t.Fatalf("se esperaba 1 usuario guardado en el fake, hay %d", len(fake.usuarios))
	}
}

// TestRutaProtegida_SinToken_Devuelve401 prueba el AuthMiddleware montado
// sobre una ruta protegida real (con chi), igual que en main.go: si la
// petición no incluye el header Authorization, debe responder 401.
func TestRutaProtegida_SinToken_Devuelve401(t *testing.T) {
	fake := &fakeRepositorioUsuarios{}
	authService := service.NuevoAuthService(fake)

	enrutador := chi.NewRouter()
	enrutador.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(authService))
		r.Get("/api/v1/maquetas", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK) // nunca debería llegar aquí
		})
	})

	peticion := httptest.NewRequest(http.MethodGet, "/api/v1/maquetas", nil)
	// Importante: NO se agrega el header Authorization.
	grabadora := httptest.NewRecorder()

	enrutador.ServeHTTP(grabadora, peticion)

	if grabadora.Code != http.StatusUnauthorized {
		t.Fatalf("se esperaba 401, se obtuvo %d. Body: %s", grabadora.Code, grabadora.Body.String())
	}

	if !strings.Contains(grabadora.Body.String(), "token") {
		t.Fatalf("se esperaba un mensaje mencionando el token, se obtuvo: %s", grabadora.Body.String())
	}
}
