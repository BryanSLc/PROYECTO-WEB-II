// Archivo: internal/handlers/contratacion_http_test.go
package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

// contratacion.go usa storage.CreateContratacion / GetAllContrataciones / etc.,
// que son funciones de paquete respaldadas por variables globales (no una
// interfaz inyectable). Por eso estos tests no dependen de IDs fijos: leen el
// ID real devuelto por cada creación y lo reutilizan.

func extraerIDContratacion(t *testing.T, body []byte) int {
	t.Helper()
	var resultado map[string]any
	if err := json.Unmarshal(body, &resultado); err != nil {
		t.Fatalf("no se pudo decodificar la respuesta: %v. Body: %s", err, body)
	}
	idFloat, ok := resultado["id_contratacion"].(float64)
	if !ok {
		t.Fatalf("no se encontro id_contratacion en la respuesta: %s", body)
	}
	return int(idFloat)
}

func crearContratacionValida(t *testing.T) int {
	t.Helper()
	body := `{"estudiante":"Carlos Ramirez","fecha":"2026-07-08","estado":"pendiente","id_servicio":1}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/contrataciones", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	CrearContratacion(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("esperaba 201 al preparar contratacion base, obtuve %d. Body: %s", rec.Code, rec.Body.String())
	}
	return extraerIDContratacion(t, rec.Body.Bytes())
}

func TestCrearContratacion_HTTP_Exitoso(t *testing.T) {
	body := `{"estudiante":"Ana Torres","fecha":"2026-07-08","estado":"pendiente","id_servicio":1}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/contrataciones", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	CrearContratacion(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("esperaba 201, obtuve %d. Body: %s", rec.Code, rec.Body.String())
	}
}

func TestCrearContratacion_HTTP_BodyInvalido(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/api/v1/contrataciones", bytes.NewBufferString(`{invalido`))
	rec := httptest.NewRecorder()

	CrearContratacion(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestCrearContratacion_HTTP_SinEstudiante(t *testing.T) {
	body := `{"fecha":"2026-07-08","estado":"pendiente","id_servicio":1}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/contrataciones", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	CrearContratacion(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestCrearContratacion_HTTP_SinFecha(t *testing.T) {
	body := `{"estudiante":"Ana","estado":"pendiente","id_servicio":1}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/contrataciones", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	CrearContratacion(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestCrearContratacion_HTTP_SinEstado(t *testing.T) {
	body := `{"estudiante":"Ana","fecha":"2026-07-08","id_servicio":1}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/contrataciones", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	CrearContratacion(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestCrearContratacion_HTTP_SinIDServicio(t *testing.T) {
	body := `{"estudiante":"Ana","fecha":"2026-07-08","estado":"pendiente"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/contrataciones", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	CrearContratacion(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestObtenerContrataciones_HTTP(t *testing.T) {
	crearContratacionValida(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/contrataciones", nil)
	rec := httptest.NewRecorder()

	ObtenerContrataciones(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperaba 200, obtuve %d", rec.Code)
	}
}

func TestObtenerContratacionPorID_HTTP_Existe(t *testing.T) {
	id := crearContratacionValida(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/contrataciones/x", nil)
	req = chiCtxConID(req, strconv.Itoa(id))
	rec := httptest.NewRecorder()

	ObtenerContratacionPorID(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperaba 200, obtuve %d. Body: %s", rec.Code, rec.Body.String())
	}
}

func TestObtenerContratacionPorID_HTTP_NoExiste(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/contrataciones/999999", nil)
	req = chiCtxConID(req, "999999")
	rec := httptest.NewRecorder()

	ObtenerContratacionPorID(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("esperaba 404, obtuve %d", rec.Code)
	}
}

func TestObtenerContratacionPorID_HTTP_IDInvalido(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/contrataciones/abc", nil)
	req = chiCtxConID(req, "abc")
	rec := httptest.NewRecorder()

	ObtenerContratacionPorID(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestActualizarContratacion_HTTP_Exitoso(t *testing.T) {
	id := crearContratacionValida(t)

	body := `{"estudiante":"Ana Actualizada","fecha":"2026-07-09","estado":"confirmado","id_servicio":1}`
	req := httptest.NewRequest(http.MethodPut, "/api/v1/contrataciones/x", bytes.NewBufferString(body))
	req = chiCtxConID(req, strconv.Itoa(id))
	rec := httptest.NewRecorder()

	ActualizarContratacion(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperaba 200, obtuve %d. Body: %s", rec.Code, rec.Body.String())
	}
}

func TestActualizarContratacion_HTTP_NoExiste(t *testing.T) {
	body := `{"estudiante":"Ana","fecha":"2026-07-08","estado":"pendiente","id_servicio":1}`
	req := httptest.NewRequest(http.MethodPut, "/api/v1/contrataciones/999999", bytes.NewBufferString(body))
	req = chiCtxConID(req, "999999")
	rec := httptest.NewRecorder()

	ActualizarContratacion(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("esperaba 404, obtuve %d", rec.Code)
	}
}

func TestActualizarContratacion_HTTP_IDInvalido(t *testing.T) {
	body := `{"estudiante":"Ana","fecha":"2026-07-08","estado":"pendiente","id_servicio":1}`
	req := httptest.NewRequest(http.MethodPut, "/api/v1/contrataciones/abc", bytes.NewBufferString(body))
	req = chiCtxConID(req, "abc")
	rec := httptest.NewRecorder()

	ActualizarContratacion(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestActualizarContratacion_HTTP_BodyInvalido(t *testing.T) {
	id := crearContratacionValida(t)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/contrataciones/x", bytes.NewBufferString(`{invalido`))
	req = chiCtxConID(req, strconv.Itoa(id))
	rec := httptest.NewRecorder()

	ActualizarContratacion(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestActualizarContratacion_HTTP_SinEstudiante(t *testing.T) {
	id := crearContratacionValida(t)

	body := `{"fecha":"2026-07-08","estado":"pendiente","id_servicio":1}`
	req := httptest.NewRequest(http.MethodPut, "/api/v1/contrataciones/x", bytes.NewBufferString(body))
	req = chiCtxConID(req, strconv.Itoa(id))
	rec := httptest.NewRecorder()

	ActualizarContratacion(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestActualizarContratacion_HTTP_SinFecha(t *testing.T) {
	id := crearContratacionValida(t)

	body := `{"estudiante":"Ana","estado":"pendiente","id_servicio":1}`
	req := httptest.NewRequest(http.MethodPut, "/api/v1/contrataciones/x", bytes.NewBufferString(body))
	req = chiCtxConID(req, strconv.Itoa(id))
	rec := httptest.NewRecorder()

	ActualizarContratacion(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestActualizarContratacion_HTTP_SinEstado(t *testing.T) {
	id := crearContratacionValida(t)

	body := `{"estudiante":"Ana","fecha":"2026-07-08","id_servicio":1}`
	req := httptest.NewRequest(http.MethodPut, "/api/v1/contrataciones/x", bytes.NewBufferString(body))
	req = chiCtxConID(req, strconv.Itoa(id))
	rec := httptest.NewRecorder()

	ActualizarContratacion(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestActualizarContratacion_HTTP_SinIDServicio(t *testing.T) {
	id := crearContratacionValida(t)

	body := `{"estudiante":"Ana","fecha":"2026-07-08","estado":"pendiente"}`
	req := httptest.NewRequest(http.MethodPut, "/api/v1/contrataciones/x", bytes.NewBufferString(body))
	req = chiCtxConID(req, strconv.Itoa(id))
	rec := httptest.NewRecorder()

	ActualizarContratacion(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestEliminarContratacion_HTTP_Exitoso(t *testing.T) {
	id := crearContratacionValida(t)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/contrataciones/x", nil)
	req = chiCtxConID(req, strconv.Itoa(id))
	rec := httptest.NewRecorder()

	EliminarContratacion(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperaba 200, obtuve %d. Body: %s", rec.Code, rec.Body.String())
	}
}

func TestEliminarContratacion_HTTP_NoExiste(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/api/v1/contrataciones/999999", nil)
	req = chiCtxConID(req, "999999")
	rec := httptest.NewRecorder()

	EliminarContratacion(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("esperaba 404, obtuve %d", rec.Code)
	}
}

func TestEliminarContratacion_HTTP_IDInvalido(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/api/v1/contrataciones/abc", nil)
	req = chiCtxConID(req, "abc")
	rec := httptest.NewRecorder()

	EliminarContratacion(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}
