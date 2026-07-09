package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"proyecto/internal/models"
	"proyecto/internal/service"
)

type fakeRepoMaqueta struct {
	maquetas    []models.Maqueta
	evoluciones []models.EvolucionMaqueta
}

func (f *fakeRepoMaqueta) CrearMaqueta(m models.Maqueta) models.Maqueta {
	m.ID = len(f.maquetas) + 1
	f.maquetas = append(f.maquetas, m)
	return m
}
func (f *fakeRepoMaqueta) ListarMaquetas() []models.Maqueta { return f.maquetas }
func (f *fakeRepoMaqueta) BuscarMaquetaPorID(id int) (models.Maqueta, bool) {
	for _, m := range f.maquetas {
		if m.ID == id {
			return m, true
		}
	}
	return models.Maqueta{}, false
}
func (f *fakeRepoMaqueta) ActualizarMaqueta(id int, datos models.Maqueta) (models.Maqueta, bool) {
	for i, m := range f.maquetas {
		if m.ID == id {
			datos.ID = id
			f.maquetas[i] = datos
			return datos, true
		}
	}
	return models.Maqueta{}, false
}
func (f *fakeRepoMaqueta) EliminarMaqueta(id int) bool {
	for i, m := range f.maquetas {
		if m.ID == id {
			f.maquetas = append(f.maquetas[:i], f.maquetas[i+1:]...)
			return true
		}
	}
	return false
}
func (f *fakeRepoMaqueta) AgregarEvolucion(e models.EvolucionMaqueta) models.EvolucionMaqueta {
	e.ID = len(f.evoluciones) + 1
	f.evoluciones = append(f.evoluciones, e)
	return e
}
func (f *fakeRepoMaqueta) ListarEvolucionPorMaqueta(id int) []models.EvolucionMaqueta {
	var r []models.EvolucionMaqueta
	for _, e := range f.evoluciones {
		if e.MaquetaID == id {
			r = append(r, e)
		}
	}
	return r
}
func (f *fakeRepoMaqueta) EliminarEvolucion(id int) bool {
	for i, e := range f.evoluciones {
		if e.ID == id {
			f.evoluciones = append(f.evoluciones[:i], f.evoluciones[i+1:]...)
			return true
		}
	}
	return false
}

func TestCrearMaqueta_HTTP_Exitoso(t *testing.T) {
	fake := &fakeRepoMaqueta{}
	srv := &Servidor{MaquetaService: service.NuevoMaquetaService(fake)}

	body := `{"titulo":"Casa Moderna","descripcion":"Proyecto residencial"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/maquetas", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	srv.CrearMaqueta(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("esperaba 201, obtuve %d. Body: %s", rec.Code, rec.Body.String())
	}
}

func TestCrearMaqueta_HTTP_SinTitulo(t *testing.T) {
	fake := &fakeRepoMaqueta{}
	srv := &Servidor{MaquetaService: service.NuevoMaquetaService(fake)}

	body := `{"descripcion":"Sin titulo"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/maquetas", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	srv.CrearMaqueta(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestObtenerMaquetas_HTTP(t *testing.T) {
	fake := &fakeRepoMaqueta{}
	svc := service.NuevoMaquetaService(fake)
	svc.Crear(models.Maqueta{Titulo: "Casa A"})
	srv := &Servidor{MaquetaService: svc}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/maquetas", nil)
	rec := httptest.NewRecorder()

	srv.ObtenerMaquetas(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperaba 200, obtuve %d", rec.Code)
	}
}

func TestObtenerMaquetaPorID_HTTP_Existe(t *testing.T) {
	fake := &fakeRepoMaqueta{}
	svc := service.NuevoMaquetaService(fake)
	svc.Crear(models.Maqueta{Titulo: "Casa A"})
	srv := &Servidor{MaquetaService: svc}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/maquetas/1", nil)
	req = chiCtxConID(req, "1")
	rec := httptest.NewRecorder()

	srv.ObtenerMaquetaPorID(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperaba 200, obtuve %d", rec.Code)
	}
}

func TestObtenerMaquetaPorID_HTTP_NoExiste(t *testing.T) {
	fake := &fakeRepoMaqueta{}
	srv := &Servidor{MaquetaService: service.NuevoMaquetaService(fake)}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/maquetas/999", nil)
	req = chiCtxConID(req, "999")
	rec := httptest.NewRecorder()

	srv.ObtenerMaquetaPorID(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("esperaba 404, obtuve %d", rec.Code)
	}
}

func TestActualizarMaqueta_HTTP_Exitoso(t *testing.T) {
	fake := &fakeRepoMaqueta{}
	svc := service.NuevoMaquetaService(fake)
	svc.Crear(models.Maqueta{Titulo: "Casa A"})
	srv := &Servidor{MaquetaService: svc}

	body := `{"titulo":"Casa B"}`
	req := httptest.NewRequest(http.MethodPut, "/api/v1/maquetas/1", bytes.NewBufferString(body))
	req = chiCtxConID(req, "1")
	rec := httptest.NewRecorder()

	srv.ActualizarMaqueta(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperaba 200, obtuve %d. Body: %s", rec.Code, rec.Body.String())
	}
}

func TestActualizarMaqueta_HTTP_NoExiste(t *testing.T) {
	fake := &fakeRepoMaqueta{}
	srv := &Servidor{MaquetaService: service.NuevoMaquetaService(fake)}

	body := `{"titulo":"Casa B"}`
	req := httptest.NewRequest(http.MethodPut, "/api/v1/maquetas/999", bytes.NewBufferString(body))
	req = chiCtxConID(req, "999")
	rec := httptest.NewRecorder()

	srv.ActualizarMaqueta(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("esperaba 404, obtuve %d", rec.Code)
	}
}

func TestEliminarMaqueta_HTTP_Exitoso(t *testing.T) {
	fake := &fakeRepoMaqueta{}
	svc := service.NuevoMaquetaService(fake)
	svc.Crear(models.Maqueta{Titulo: "Casa A"})
	srv := &Servidor{MaquetaService: svc}

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/maquetas/1", nil)
	req = chiCtxConID(req, "1")
	rec := httptest.NewRecorder()

	srv.EliminarMaqueta(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperaba 200, obtuve %d", rec.Code)
	}
}

func TestEliminarMaqueta_HTTP_NoExiste(t *testing.T) {
	fake := &fakeRepoMaqueta{}
	srv := &Servidor{MaquetaService: service.NuevoMaquetaService(fake)}

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/maquetas/999", nil)
	req = chiCtxConID(req, "999")
	rec := httptest.NewRecorder()

	srv.EliminarMaqueta(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("esperaba 404, obtuve %d", rec.Code)
	}
}

func TestAgregarEvolucion_HTTP_Exitoso(t *testing.T) {
	fake := &fakeRepoMaqueta{}
	svc := service.NuevoMaquetaService(fake)
	svc.Crear(models.Maqueta{Titulo: "Casa A"})
	srv := &Servidor{MaquetaService: svc}

	body := `{"maqueta_id":1,"titulo":"Avance 1","paso":1}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/maquetas/evolucion", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	srv.AgregarEvolucionMaqueta(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("esperaba 201, obtuve %d. Body: %s", rec.Code, rec.Body.String())
	}
}

func TestAgregarEvolucion_HTTP_SinMaquetaID(t *testing.T) {
	fake := &fakeRepoMaqueta{}
	srv := &Servidor{MaquetaService: service.NuevoMaquetaService(fake)}

	body := `{"titulo":"Avance 1","paso":1}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/maquetas/evolucion", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	srv.AgregarEvolucionMaqueta(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestObtenerEvolucionPorMaqueta_HTTP(t *testing.T) {
	fake := &fakeRepoMaqueta{}
	svc := service.NuevoMaquetaService(fake)
	svc.Crear(models.Maqueta{Titulo: "Casa A"})
	srv := &Servidor{MaquetaService: svc}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/maquetas/1/evolucion", nil)
	req = chiCtxConID(req, "1")
	rec := httptest.NewRecorder()

	srv.ObtenerEvolucionPorMaqueta(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperaba 200, obtuve %d", rec.Code)
	}
}

func TestEliminarEvolucion_HTTP_NoExiste(t *testing.T) {
	fake := &fakeRepoMaqueta{}
	srv := &Servidor{MaquetaService: service.NuevoMaquetaService(fake)}

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/maquetas/evolucion/999", nil)
	req = chiCtxConID(req, "999")
	rec := httptest.NewRecorder()

	srv.EliminarEvolucion(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("esperaba 404, obtuve %d", rec.Code)
	}
}
