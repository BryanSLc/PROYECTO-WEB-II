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

// ==========================================
//         FAKES DEL MÓDULO 3
// ==========================================

type fakeRepoProveedores struct{ lista []models.Proveedor }

func (f *fakeRepoProveedores) CrearProveedor(p models.Proveedor) models.Proveedor {
	p.ID = len(f.lista) + 1
	f.lista = append(f.lista, p)
	return p
}
func (f *fakeRepoProveedores) ListarProveedores() []models.Proveedor { return f.lista }
func (f *fakeRepoProveedores) BuscarProveedorPorID(int) (models.Proveedor, bool) {
	return models.Proveedor{}, false
}
func (f *fakeRepoProveedores) ActualizarProveedor(int, models.Proveedor) (models.Proveedor, bool) {
	return models.Proveedor{}, false
}
func (f *fakeRepoProveedores) EliminarProveedor(int) bool { return false }

type fakeRepoMaterialesHandler struct{ lista []models.MaterialProveedor }

func (f *fakeRepoMaterialesHandler) CrearMaterial(m models.MaterialProveedor) models.MaterialProveedor {
	m.ID = len(f.lista) + 1
	f.lista = append(f.lista, m)
	return m
}
func (f *fakeRepoMaterialesHandler) ListarMateriales() []models.MaterialProveedor { return f.lista }
func (f *fakeRepoMaterialesHandler) BuscarMaterialPorID(int) (models.MaterialProveedor, bool) {
	return models.MaterialProveedor{}, false
}
func (f *fakeRepoMaterialesHandler) ActualizarMaterial(int, models.MaterialProveedor) (models.MaterialProveedor, bool) {
	return models.MaterialProveedor{}, false
}
func (f *fakeRepoMaterialesHandler) EliminarMaterial(int) bool { return false }

type fakeRepoUbicacionesHandler struct{ lista []models.Ubicacion }

func (f *fakeRepoUbicacionesHandler) CrearUbicacion(u models.Ubicacion) models.Ubicacion {
	u.ID = len(f.lista) + 1
	f.lista = append(f.lista, u)
	return u
}
func (f *fakeRepoUbicacionesHandler) ListarUbicaciones() []models.Ubicacion { return f.lista }
func (f *fakeRepoUbicacionesHandler) BuscarUbicacionPorID(int) (models.Ubicacion, bool) {
	return models.Ubicacion{}, false
}
func (f *fakeRepoUbicacionesHandler) ActualizarUbicacion(int, models.Ubicacion) (models.Ubicacion, bool) {
	return models.Ubicacion{}, false
}
func (f *fakeRepoUbicacionesHandler) EliminarUbicacion(int) bool { return false }

// ==========================================
//         TESTS DE PROVEEDOR
// ==========================================

// TestCrearProveedor_HTTP_Exitoso: POST con datos válidos → 201 Created
func TestCrearProveedor_HTTP_Exitoso(t *testing.T) {
	fake := &fakeRepoProveedores{}
	srv := &Servidor{ProveedorService: service.NuevoProveedorService(fake)}

	body := `{"nombre":"Ferretería Central","ciudad":"Manta","provincia":"Manabí","direccion":"Av. 4 de Noviembre","telefono":"052123456"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/proveedores", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	srv.CrearProveedor(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("esperaba 201, obtuve %d. Body: %s", rec.Code, rec.Body.String())
	}
	if len(fake.lista) != 1 {
		t.Fatalf("esperaba 1 proveedor guardado, hay %d", len(fake.lista))
	}
}

// TestRutaProveedores_SinToken_Devuelve401: ruta protegida sin Authorization → 401
func TestRutaProveedores_SinToken_Devuelve401(t *testing.T) {
	fakeAuth := &fakeRepositorioUsuarios{}
	authService := service.NuevoAuthService(fakeAuth)

	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(authService))
		r.Get("/api/v1/proveedores", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	})

	req := httptest.NewRequest(http.MethodGet, "/api/v1/proveedores", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("esperaba 401, obtuve %d", rec.Code)
	}
}

// ==========================================
//         TESTS DE MATERIAL
// ==========================================

// TestCrearMaterial_HTTP_Exitoso: POST con datos válidos → 201 Created
func TestCrearMaterial_HTTP_Exitoso(t *testing.T) {
	fake := &fakeRepoMaterialesHandler{}
	srv := &Servidor{MaterialService: service.NuevoMaterialService(fake)}

	body := `{"nombre":"MDF 15mm","categoria":"Madera","precio_referencial":12.50,"id_proveedor":1}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/materiales", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	srv.CrearMaterial(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("esperaba 201, obtuve %d. Body: %s", rec.Code, rec.Body.String())
	}
	if len(fake.lista) != 1 {
		t.Fatalf("esperaba 1 material guardado, hay %d", len(fake.lista))
	}
}

// TestRutaMateriales_SinToken_Devuelve401: ruta protegida sin Authorization → 401
func TestRutaMateriales_SinToken_Devuelve401(t *testing.T) {
	fakeAuth := &fakeRepositorioUsuarios{}
	authService := service.NuevoAuthService(fakeAuth)

	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(authService))
		r.Get("/api/v1/materiales", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	})

	req := httptest.NewRequest(http.MethodGet, "/api/v1/materiales", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("esperaba 401, obtuve %d", rec.Code)
	}
}

// ==========================================
//         TESTS DE UBICACIÓN
// ==========================================

// TestCrearUbicacion_HTTP_Exitoso: POST con datos válidos → 201 Created
func TestCrearUbicacion_HTTP_Exitoso(t *testing.T) {
	fake := &fakeRepoUbicacionesHandler{}
	srv := &Servidor{UbicacionService: service.NuevoUbicacionService(fake)}

	body := `{"provincia":"Manabí","ciudad":"Manta"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/ubicaciones", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	srv.CrearUbicacion(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("esperaba 201, obtuve %d. Body: %s", rec.Code, rec.Body.String())
	}
	if len(fake.lista) != 1 {
		t.Fatalf("esperaba 1 ubicación guardada, hay %d", len(fake.lista))
	}
}

// TestRutaUbicaciones_SinToken_Devuelve401: ruta protegida sin Authorization → 401
func TestRutaUbicaciones_SinToken_Devuelve401(t *testing.T) {
	fakeAuth := &fakeRepositorioUsuarios{}
	authService := service.NuevoAuthService(fakeAuth)

	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(authService))
		r.Get("/api/v1/ubicaciones", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	})

	req := httptest.NewRequest(http.MethodGet, "/api/v1/ubicaciones", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("esperaba 401, obtuve %d", rec.Code)
	}
}
