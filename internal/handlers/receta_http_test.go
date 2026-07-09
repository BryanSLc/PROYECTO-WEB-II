// Archivo: internal/handlers/receta_http_test.go
package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"proyecto/internal/models"
	"proyecto/internal/storage"
)

// fakeAlmacenReceta envuelve *storage.Memoria (que ya implementa Usuario,
// Maqueta, Evolucion y Receta en memoria) y agrega stubs vacíos para el
// resto de la interfaz storage.Almacen (Proveedor, Material, Ubicacion,
// Asesor, Contratacion, Servicio) para que el fake compile como Almacen
// completo. Estos stubs no se usan en los tests de receta.
type fakeAlmacenReceta struct {
	*storage.Memoria
}

func nuevoFakeAlmacenReceta() *fakeAlmacenReceta {
	return &fakeAlmacenReceta{Memoria: storage.NuevaMemoria()}
}

// --- Usuario: falta BuscarUsuarioPorEmail en Memoria ---
func (f *fakeAlmacenReceta) BuscarUsuarioPorEmail(email string) (models.Usuario, bool) {
	return models.Usuario{}, false
}

// --- Proveedores ---
func (f *fakeAlmacenReceta) CrearProveedor(models.Proveedor) models.Proveedor {
	return models.Proveedor{}
}
func (f *fakeAlmacenReceta) ListarProveedores() []models.Proveedor { return nil }
func (f *fakeAlmacenReceta) BuscarProveedorPorID(int) (models.Proveedor, bool) {
	return models.Proveedor{}, false
}
func (f *fakeAlmacenReceta) ActualizarProveedor(int, models.Proveedor) (models.Proveedor, bool) {
	return models.Proveedor{}, false
}
func (f *fakeAlmacenReceta) EliminarProveedor(int) bool { return false }

// --- Materiales ---
func (f *fakeAlmacenReceta) CrearMaterial(models.MaterialProveedor) models.MaterialProveedor {
	return models.MaterialProveedor{}
}
func (f *fakeAlmacenReceta) ListarMateriales() []models.MaterialProveedor { return nil }
func (f *fakeAlmacenReceta) BuscarMaterialPorID(int) (models.MaterialProveedor, bool) {
	return models.MaterialProveedor{}, false
}
func (f *fakeAlmacenReceta) ActualizarMaterial(int, models.MaterialProveedor) (models.MaterialProveedor, bool) {
	return models.MaterialProveedor{}, false
}
func (f *fakeAlmacenReceta) EliminarMaterial(int) bool { return false }

// --- Ubicaciones ---
func (f *fakeAlmacenReceta) CrearUbicacion(models.Ubicacion) models.Ubicacion {
	return models.Ubicacion{}
}
func (f *fakeAlmacenReceta) ListarUbicaciones() []models.Ubicacion { return nil }
func (f *fakeAlmacenReceta) BuscarUbicacionPorID(int) (models.Ubicacion, bool) {
	return models.Ubicacion{}, false
}
func (f *fakeAlmacenReceta) ActualizarUbicacion(int, models.Ubicacion) (models.Ubicacion, bool) {
	return models.Ubicacion{}, false
}
func (f *fakeAlmacenReceta) EliminarUbicacion(int) bool { return false }

// --- Asesores ---
func (f *fakeAlmacenReceta) CrearAsesor(models.Asesor) models.Asesor { return models.Asesor{} }
func (f *fakeAlmacenReceta) ListarAsesores() []models.Asesor         { return nil }
func (f *fakeAlmacenReceta) BuscarAsesorPorID(int) (models.Asesor, bool) {
	return models.Asesor{}, false
}
func (f *fakeAlmacenReceta) ActualizarAsesor(int, models.Asesor) (models.Asesor, bool) {
	return models.Asesor{}, false
}
func (f *fakeAlmacenReceta) EliminarAsesor(int) bool { return false }

// --- Contrataciones ---
func (f *fakeAlmacenReceta) CrearContratacion(models.Contratacion) models.Contratacion {
	return models.Contratacion{}
}
func (f *fakeAlmacenReceta) ListarContrataciones() []models.Contratacion { return nil }
func (f *fakeAlmacenReceta) BuscarContratacionPorID(int) (models.Contratacion, bool) {
	return models.Contratacion{}, false
}
func (f *fakeAlmacenReceta) ActualizarContratacion(int, models.Contratacion) (models.Contratacion, bool) {
	return models.Contratacion{}, false
}
func (f *fakeAlmacenReceta) EliminarContratacion(int) bool { return false }

// --- Servicios ---
func (f *fakeAlmacenReceta) CrearServicio(models.Servicio) models.Servicio {
	return models.Servicio{}
}
func (f *fakeAlmacenReceta) ListarServicios() []models.Servicio { return nil }
func (f *fakeAlmacenReceta) BuscarServicioPorID(int) (models.Servicio, bool) {
	return models.Servicio{}, false
}
func (f *fakeAlmacenReceta) ActualizarServicio(int, models.Servicio) (models.Servicio, bool) {
	return models.Servicio{}, false
}
func (f *fakeAlmacenReceta) EliminarServicio(int) bool { return false }

// ==========================================
//              TESTS DE RECETA
// ==========================================

func TestCrearReceta_HTTP_Exitoso(t *testing.T) {
	fake := nuevoFakeAlmacenReceta()
	srv := &Servidor{Almacen: fake}

	body := `{"titulo":"Receta A","maqueta_id":1,"descripcion":"Paso a paso"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/recetas", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	srv.CrearReceta(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("esperaba 201, obtuve %d. Body: %s", rec.Code, rec.Body.String())
	}
}

func TestCrearReceta_HTTP_SinTitulo(t *testing.T) {
	fake := nuevoFakeAlmacenReceta()
	srv := &Servidor{Almacen: fake}

	body := `{"maqueta_id":1,"descripcion":"Sin titulo"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/recetas", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	srv.CrearReceta(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestCrearReceta_HTTP_BodyInvalido(t *testing.T) {
	fake := nuevoFakeAlmacenReceta()
	srv := &Servidor{Almacen: fake}

	req := httptest.NewRequest(http.MethodPost, "/api/v1/recetas", bytes.NewBufferString(`{invalido`))
	rec := httptest.NewRecorder()

	srv.CrearReceta(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestObtenerRecetas_HTTP_Todas(t *testing.T) {
	fake := nuevoFakeAlmacenReceta()
	fake.CrearReceta(models.Receta{Titulo: "Receta A", MaquetaID: 1})
	srv := &Servidor{Almacen: fake}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/recetas", nil)
	rec := httptest.NewRecorder()

	srv.ObtenerRecetas(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperaba 200, obtuve %d", rec.Code)
	}
}

func TestObtenerRecetas_HTTP_PorMaquetaID(t *testing.T) {
	fake := nuevoFakeAlmacenReceta()
	fake.CrearReceta(models.Receta{Titulo: "Receta A", MaquetaID: 1})
	srv := &Servidor{Almacen: fake}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/recetas?maqueta_id=1", nil)
	rec := httptest.NewRecorder()

	srv.ObtenerRecetas(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperaba 200, obtuve %d", rec.Code)
	}
}

func TestObtenerRecetas_HTTP_MaquetaIDInvalido(t *testing.T) {
	fake := nuevoFakeAlmacenReceta()
	srv := &Servidor{Almacen: fake}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/recetas?maqueta_id=abc", nil)
	rec := httptest.NewRecorder()

	srv.ObtenerRecetas(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestObtenerRecetaPorID_HTTP_Existe(t *testing.T) {
	fake := nuevoFakeAlmacenReceta()
	creada := fake.CrearReceta(models.Receta{Titulo: "Receta A", MaquetaID: 1})
	srv := &Servidor{Almacen: fake}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/recetas/1", nil)
	req = chiCtxConID(req, strconv.Itoa(creada.ID))
	rec := httptest.NewRecorder()

	srv.ObtenerRecetaPorID(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperaba 200, obtuve %d", rec.Code)
	}
}

func TestObtenerRecetaPorID_HTTP_NoExiste(t *testing.T) {
	fake := nuevoFakeAlmacenReceta()
	srv := &Servidor{Almacen: fake}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/recetas/999", nil)
	req = chiCtxConID(req, "999")
	rec := httptest.NewRecorder()

	srv.ObtenerRecetaPorID(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("esperaba 404, obtuve %d", rec.Code)
	}
}

func TestObtenerRecetaPorID_HTTP_IDInvalido(t *testing.T) {
	fake := nuevoFakeAlmacenReceta()
	srv := &Servidor{Almacen: fake}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/recetas/abc", nil)
	req = chiCtxConID(req, "abc")
	rec := httptest.NewRecorder()

	srv.ObtenerRecetaPorID(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestActualizarReceta_HTTP_Exitoso(t *testing.T) {
	fake := nuevoFakeAlmacenReceta()
	creada := fake.CrearReceta(models.Receta{Titulo: "Receta A", MaquetaID: 1})
	srv := &Servidor{Almacen: fake}

	body := `{"titulo":"Receta B","maqueta_id":1}`
	req := httptest.NewRequest(http.MethodPut, "/api/v1/recetas/1", bytes.NewBufferString(body))
	req = chiCtxConID(req, strconv.Itoa(creada.ID))
	rec := httptest.NewRecorder()

	srv.ActualizarReceta(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperaba 200, obtuve %d. Body: %s", rec.Code, rec.Body.String())
	}
}

func TestActualizarReceta_HTTP_NoExiste(t *testing.T) {
	fake := nuevoFakeAlmacenReceta()
	srv := &Servidor{Almacen: fake}

	body := `{"titulo":"Receta B","maqueta_id":1}`
	req := httptest.NewRequest(http.MethodPut, "/api/v1/recetas/999", bytes.NewBufferString(body))
	req = chiCtxConID(req, "999")
	rec := httptest.NewRecorder()

	srv.ActualizarReceta(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("esperaba 404, obtuve %d", rec.Code)
	}
}

func TestActualizarReceta_HTTP_SinTitulo(t *testing.T) {
	fake := nuevoFakeAlmacenReceta()
	creada := fake.CrearReceta(models.Receta{Titulo: "Receta A", MaquetaID: 1})
	srv := &Servidor{Almacen: fake}

	body := `{"maqueta_id":1}`
	req := httptest.NewRequest(http.MethodPut, "/api/v1/recetas/1", bytes.NewBufferString(body))
	req = chiCtxConID(req, strconv.Itoa(creada.ID))
	rec := httptest.NewRecorder()

	srv.ActualizarReceta(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestActualizarReceta_HTTP_IDInvalido(t *testing.T) {
	fake := nuevoFakeAlmacenReceta()
	srv := &Servidor{Almacen: fake}

	body := `{"titulo":"Receta B","maqueta_id":1}`
	req := httptest.NewRequest(http.MethodPut, "/api/v1/recetas/abc", bytes.NewBufferString(body))
	req = chiCtxConID(req, "abc")
	rec := httptest.NewRecorder()

	srv.ActualizarReceta(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestEliminarReceta_HTTP_Exitoso(t *testing.T) {
	fake := nuevoFakeAlmacenReceta()
	creada := fake.CrearReceta(models.Receta{Titulo: "Receta A", MaquetaID: 1})
	srv := &Servidor{Almacen: fake}

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/recetas/1", nil)
	req = chiCtxConID(req, strconv.Itoa(creada.ID))
	rec := httptest.NewRecorder()

	srv.EliminarReceta(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperaba 200, obtuve %d", rec.Code)
	}
}

func TestEliminarReceta_HTTP_NoExiste(t *testing.T) {
	fake := nuevoFakeAlmacenReceta()
	srv := &Servidor{Almacen: fake}

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/recetas/999", nil)
	req = chiCtxConID(req, "999")
	rec := httptest.NewRecorder()

	srv.EliminarReceta(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("esperaba 404, obtuve %d", rec.Code)
	}
}

func TestEliminarReceta_HTTP_IDInvalido(t *testing.T) {
	fake := nuevoFakeAlmacenReceta()
	srv := &Servidor{Almacen: fake}

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/recetas/abc", nil)
	req = chiCtxConID(req, "abc")
	rec := httptest.NewRecorder()

	srv.EliminarReceta(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

