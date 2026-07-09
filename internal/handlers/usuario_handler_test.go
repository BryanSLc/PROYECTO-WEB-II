package handlers

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"proyecto/internal/models"
	"proyecto/internal/service"

	"github.com/go-chi/chi/v5"
)

type fakeRepoUsuario struct {
	usuarios []models.Usuario
}

func (f *fakeRepoUsuario) CrearUsuario(u models.Usuario) models.Usuario {
	u.ID = len(f.usuarios) + 1
	f.usuarios = append(f.usuarios, u)
	return u
}
func (f *fakeRepoUsuario) ListarUsuarios() []models.Usuario { return f.usuarios }
func (f *fakeRepoUsuario) BuscarUsuarioPorID(id int) (models.Usuario, bool) {
	for _, u := range f.usuarios {
		if u.ID == id {
			return u, true
		}
	}
	return models.Usuario{}, false
}
func (f *fakeRepoUsuario) ActualizarUsuario(id int, datos models.Usuario) (models.Usuario, bool) {
	for i, u := range f.usuarios {
		if u.ID == id {
			datos.ID = id
			f.usuarios[i] = datos
			return datos, true
		}
	}
	return models.Usuario{}, false
}
func (f *fakeRepoUsuario) EliminarUsuario(id int) bool {
	for i, u := range f.usuarios {
		if u.ID == id {
			f.usuarios = append(f.usuarios[:i], f.usuarios[i+1:]...)
			return true
		}
	}
	return false
}

func chiCtxConID(r *http.Request, id string) *http.Request {
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", id)
	return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
}

func TestCrearUsuario_HTTP_Exitoso(t *testing.T) {
	fake := &fakeRepoUsuario{}
	srv := &Servidor{UsuarioService: service.NuevoUsuarioService(fake)}

	body := `{"nombre":"Ana","email":"ana@test.com","password":"clave123"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/usuarios", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	srv.CrearUsuario(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("esperaba 201, obtuve %d. Body: %s", rec.Code, rec.Body.String())
	}
}

func TestCrearUsuario_HTTP_SinNombre(t *testing.T) {
	fake := &fakeRepoUsuario{}
	srv := &Servidor{UsuarioService: service.NuevoUsuarioService(fake)}

	body := `{"email":"ana@test.com"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/usuarios", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	srv.CrearUsuario(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestCrearUsuario_HTTP_SinEmail(t *testing.T) {
	fake := &fakeRepoUsuario{}
	srv := &Servidor{UsuarioService: service.NuevoUsuarioService(fake)}

	body := `{"nombre":"Ana"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/usuarios", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	srv.CrearUsuario(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestObtenerUsuarios_HTTP(t *testing.T) {
	fake := &fakeRepoUsuario{}
	svc := service.NuevoUsuarioService(fake)
	svc.Crear(models.Usuario{Nombre: "Ana", Email: "ana@test.com"})
	srv := &Servidor{UsuarioService: svc}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/usuarios", nil)
	rec := httptest.NewRecorder()

	srv.ObtenerUsuarios(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperaba 200, obtuve %d", rec.Code)
	}
}

func TestObtenerUsuarioPorID_HTTP_Existe(t *testing.T) {
	fake := &fakeRepoUsuario{}
	svc := service.NuevoUsuarioService(fake)
	svc.Crear(models.Usuario{Nombre: "Ana", Email: "ana@test.com"})
	srv := &Servidor{UsuarioService: svc}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/usuarios/1", nil)
	req = chiCtxConID(req, "1")
	rec := httptest.NewRecorder()

	srv.ObtenerUsuarioPorID(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperaba 200, obtuve %d", rec.Code)
	}
}

func TestObtenerUsuarioPorID_HTTP_NoExiste(t *testing.T) {
	fake := &fakeRepoUsuario{}
	srv := &Servidor{UsuarioService: service.NuevoUsuarioService(fake)}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/usuarios/999", nil)
	req = chiCtxConID(req, "999")
	rec := httptest.NewRecorder()

	srv.ObtenerUsuarioPorID(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("esperaba 404, obtuve %d", rec.Code)
	}
}

func TestActualizarUsuario_HTTP_Exitoso(t *testing.T) {
	fake := &fakeRepoUsuario{}
	svc := service.NuevoUsuarioService(fake)
	svc.Crear(models.Usuario{Nombre: "Ana", Email: "ana@test.com"})
	srv := &Servidor{UsuarioService: svc}

	body := `{"nombre":"Ana Lopez","email":"ana@test.com"}`
	req := httptest.NewRequest(http.MethodPut, "/api/v1/usuarios/1", bytes.NewBufferString(body))
	req = chiCtxConID(req, "1")
	rec := httptest.NewRecorder()

	srv.ActualizarUsuario(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperaba 200, obtuve %d. Body: %s", rec.Code, rec.Body.String())
	}
}

func TestActualizarUsuario_HTTP_NoExiste(t *testing.T) {
	fake := &fakeRepoUsuario{}
	srv := &Servidor{UsuarioService: service.NuevoUsuarioService(fake)}

	body := `{"nombre":"Ana","email":"ana@test.com"}`
	req := httptest.NewRequest(http.MethodPut, "/api/v1/usuarios/999", bytes.NewBufferString(body))
	req = chiCtxConID(req, "999")
	rec := httptest.NewRecorder()

	srv.ActualizarUsuario(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("esperaba 404, obtuve %d", rec.Code)
	}
}

func TestEliminarUsuario_HTTP_Exitoso(t *testing.T) {
	fake := &fakeRepoUsuario{}
	svc := service.NuevoUsuarioService(fake)
	svc.Crear(models.Usuario{Nombre: "Ana", Email: "ana@test.com"})
	srv := &Servidor{UsuarioService: svc}

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/usuarios/1", nil)
	req = chiCtxConID(req, "1")
	rec := httptest.NewRecorder()

	srv.EliminarUsuario(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperaba 200, obtuve %d", rec.Code)
	}
}

func TestEliminarUsuario_HTTP_NoExiste(t *testing.T) {
	fake := &fakeRepoUsuario{}
	srv := &Servidor{UsuarioService: service.NuevoUsuarioService(fake)}

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/usuarios/999", nil)
	req = chiCtxConID(req, "999")
	rec := httptest.NewRecorder()

	srv.EliminarUsuario(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("esperaba 404, obtuve %d", rec.Code)
	}
}
