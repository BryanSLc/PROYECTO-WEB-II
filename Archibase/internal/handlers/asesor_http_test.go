package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"proyecto/internal/handlers/middleware"
	"proyecto/internal/models"
	"proyecto/internal/service"

	"github.com/go-chi/chi/v5"
)

// fakeRepoAsesores es un doble en memoria: sí guarda datos, pero no usa GORM ni SQLite
type fakeRepoAsesores struct{ lista []models.Asesor }

func (f *fakeRepoAsesores) CrearAsesor(a models.Asesor) models.Asesor {
	a.IDasesor = len(f.lista) + 1
	f.lista = append(f.lista, a)
	return a
}
func (f *fakeRepoAsesores) ListarAsesores() []models.Asesor { return f.lista }
func (f *fakeRepoAsesores) BuscarAsesorPorID(int) (models.Asesor, bool) {
	return models.Asesor{}, false
}
func (f *fakeRepoAsesores) ActualizarAsesor(int, models.Asesor) (models.Asesor, bool) {
	return models.Asesor{}, false
}
func (f *fakeRepoAsesores) EliminarAsesor(int) bool { return false }

// TestCrearAsesor_HTTP_Exitoso: POST con datos válidos → 201 Created
func TestCrearAsesor_HTTP_Exitoso(t *testing.T) {
	fake := &fakeRepoAsesores{}
	srv := &Servidor{AsesorService: service.NuevoAsesorService(fake)}

	body := `{"nombre":"Ing. Solano","especialidad":"Diseño","experiencia":"5 años","contacto":"099","modalidad":"virtual"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/asesores", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	srv.CrearAsesor(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("esperaba 201, obtuve %d. Body: %s", rec.Code, rec.Body.String())
	}
	if len(fake.lista) != 1 {
		t.Fatalf("esperaba 1 asesor guardado, hay %d", len(fake.lista))
	}
}

// TestRutaAsesores_SinToken_Devuelve401: ruta protegida sin Authorization → 401
func TestRutaAsesores_SinToken_Devuelve401(t *testing.T) {
	fakeAuth := &fakeRepositorioUsuarios{} // reutiliza el fake de auth_http_test.go
	authService := service.NuevoAuthService(fakeAuth)

	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(authService))
		r.Get("/api/v1/asesores", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	})

	req := httptest.NewRequest(http.MethodGet, "/api/v1/asesores", nil)
	// Sin header Authorization
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("esperaba 401, obtuve %d", rec.Code)
	}
}
