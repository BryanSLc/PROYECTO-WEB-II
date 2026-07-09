// Archivo: internal/handlers/servicio_http_test.go
package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func extraerIDServicio(t *testing.T, body []byte) int {
	t.Helper()
	var resultado map[string]any
	if err := json.Unmarshal(body, &resultado); err != nil {
		t.Fatalf("no se pudo decodificar la respuesta: %v. Body: %s", err, body)
	}
	idFloat, ok := resultado["id_servicio"].(float64)
	if !ok {
		t.Fatalf("no se encontro id_servicio en la respuesta: %s", body)
	}
	return int(idFloat)
}

func crearServicioValido(t *testing.T) int {
	t.Helper()
	body := `{"titulo":"Asesoria estructural","descripcion":"Revision de planos","precio":25.5,"disponibilidad":"lunes a viernes","id_asesor":1}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/servicios", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	CrearServicio(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("esperaba 201 al preparar servicio base, obtuve %d. Body: %s", rec.Code, rec.Body.String())
	}
	return extraerIDServicio(t, rec.Body.Bytes())
}

func TestCrearServicio_HTTP_Exitoso(t *testing.T) {
	body := `{"titulo":"Diseño 3D","descripcion":"Modelado del proyecto","precio":40,"disponibilidad":"fines de semana","id_asesor":1}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/servicios", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	CrearServicio(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("esperaba 201, obtuve %d. Body: %s", rec.Code, rec.Body.String())
	}
}

func TestCrearServicio_HTTP_BodyInvalido(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/api/v1/servicios", bytes.NewBufferString(`{invalido`))
	rec := httptest.NewRecorder()

	CrearServicio(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestCrearServicio_HTTP_SinTitulo(t *testing.T) {
	body := `{"descripcion":"Sin titulo","precio":10,"disponibilidad":"lunes","id_asesor":1}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/servicios", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	CrearServicio(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestCrearServicio_HTTP_SinDescripcion(t *testing.T) {
	body := `{"titulo":"Asesoria","precio":10,"disponibilidad":"lunes","id_asesor":1}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/servicios", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	CrearServicio(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestCrearServicio_HTTP_SinDisponibilidad(t *testing.T) {
	body := `{"titulo":"Asesoria","descripcion":"desc","precio":10,"id_asesor":1}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/servicios", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	CrearServicio(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestCrearServicio_HTTP_SinIDAsesor(t *testing.T) {
	body := `{"titulo":"Asesoria","descripcion":"desc","precio":10,"disponibilidad":"lunes"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/servicios", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	CrearServicio(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestCrearServicio_HTTP_PrecioNegativo(t *testing.T) {
	body := `{"titulo":"Asesoria","descripcion":"desc","precio":-5,"disponibilidad":"lunes","id_asesor":1}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/servicios", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	CrearServicio(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestObtenerServicios_HTTP(t *testing.T) {
	crearServicioValido(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/servicios", nil)
	rec := httptest.NewRecorder()

	ObtenerServicios(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperaba 200, obtuve %d", rec.Code)
	}
}

func TestObtenerServicioPorID_HTTP_Existe(t *testing.T) {
	id := crearServicioValido(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/servicios/x", nil)
	req = chiCtxConID(req, strconv.Itoa(id))
	rec := httptest.NewRecorder()

	ObtenerServicioPorID(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperaba 200, obtuve %d. Body: %s", rec.Code, rec.Body.String())
	}
}

func TestObtenerServicioPorID_HTTP_NoExiste(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/servicios/999999", nil)
	req = chiCtxConID(req, "999999")
	rec := httptest.NewRecorder()

	ObtenerServicioPorID(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("esperaba 404, obtuve %d", rec.Code)
	}
}

func TestObtenerServicioPorID_HTTP_IDInvalido(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/servicios/abc", nil)
	req = chiCtxConID(req, "abc")
	rec := httptest.NewRecorder()

	ObtenerServicioPorID(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestActualizarServicio_HTTP_Exitoso(t *testing.T) {
	id := crearServicioValido(t)

	body := `{"titulo":"Actualizado","descripcion":"Nueva desc","precio":30,"disponibilidad":"martes","id_asesor":1}`
	req := httptest.NewRequest(http.MethodPut, "/api/v1/servicios/x", bytes.NewBufferString(body))
	req = chiCtxConID(req, strconv.Itoa(id))
	rec := httptest.NewRecorder()

	ActualizarServicio(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperaba 200, obtuve %d. Body: %s", rec.Code, rec.Body.String())
	}
}

func TestActualizarServicio_HTTP_NoExiste(t *testing.T) {
	body := `{"titulo":"X","descripcion":"desc","precio":10,"disponibilidad":"lunes","id_asesor":1}`
	req := httptest.NewRequest(http.MethodPut, "/api/v1/servicios/999999", bytes.NewBufferString(body))
	req = chiCtxConID(req, "999999")
	rec := httptest.NewRecorder()

	ActualizarServicio(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("esperaba 404, obtuve %d", rec.Code)
	}
}

func TestActualizarServicio_HTTP_IDInvalido(t *testing.T) {
	body := `{"titulo":"X","descripcion":"desc","precio":10,"disponibilidad":"lunes","id_asesor":1}`
	req := httptest.NewRequest(http.MethodPut, "/api/v1/servicios/abc", bytes.NewBufferString(body))
	req = chiCtxConID(req, "abc")
	rec := httptest.NewRecorder()

	ActualizarServicio(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestActualizarServicio_HTTP_BodyInvalido(t *testing.T) {
	id := crearServicioValido(t)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/servicios/x", bytes.NewBufferString(`{invalido`))
	req = chiCtxConID(req, strconv.Itoa(id))
	rec := httptest.NewRecorder()

	ActualizarServicio(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestActualizarServicio_HTTP_SinTitulo(t *testing.T) {
	id := crearServicioValido(t)

	body := `{"descripcion":"desc","precio":10,"disponibilidad":"lunes","id_asesor":1}`
	req := httptest.NewRequest(http.MethodPut, "/api/v1/servicios/x", bytes.NewBufferString(body))
	req = chiCtxConID(req, strconv.Itoa(id))
	rec := httptest.NewRecorder()

	ActualizarServicio(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestActualizarServicio_HTTP_SinDescripcion(t *testing.T) {
	id := crearServicioValido(t)

	body := `{"titulo":"X","precio":10,"disponibilidad":"lunes","id_asesor":1}`
	req := httptest.NewRequest(http.MethodPut, "/api/v1/servicios/x", bytes.NewBufferString(body))
	req = chiCtxConID(req, strconv.Itoa(id))
	rec := httptest.NewRecorder()

	ActualizarServicio(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestActualizarServicio_HTTP_SinDisponibilidad(t *testing.T) {
	id := crearServicioValido(t)

	body := `{"titulo":"X","descripcion":"desc","precio":10,"id_asesor":1}`
	req := httptest.NewRequest(http.MethodPut, "/api/v1/servicios/x", bytes.NewBufferString(body))
	req = chiCtxConID(req, strconv.Itoa(id))
	rec := httptest.NewRecorder()

	ActualizarServicio(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestActualizarServicio_HTTP_SinIDAsesor(t *testing.T) {
	id := crearServicioValido(t)

	body := `{"titulo":"X","descripcion":"desc","precio":10,"disponibilidad":"lunes"}`
	req := httptest.NewRequest(http.MethodPut, "/api/v1/servicios/x", bytes.NewBufferString(body))
	req = chiCtxConID(req, strconv.Itoa(id))
	rec := httptest.NewRecorder()

	ActualizarServicio(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestActualizarServicio_HTTP_PrecioNegativo(t *testing.T) {
	id := crearServicioValido(t)

	body := `{"titulo":"X","descripcion":"desc","precio":-1,"disponibilidad":"lunes","id_asesor":1}`
	req := httptest.NewRequest(http.MethodPut, "/api/v1/servicios/x", bytes.NewBufferString(body))
	req = chiCtxConID(req, strconv.Itoa(id))
	rec := httptest.NewRecorder()

	ActualizarServicio(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}

func TestEliminarServicio_HTTP_Exitoso(t *testing.T) {
	id := crearServicioValido(t)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/servicios/x", nil)
	req = chiCtxConID(req, strconv.Itoa(id))
	rec := httptest.NewRecorder()

	EliminarServicio(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperaba 200, obtuve %d. Body: %s", rec.Code, rec.Body.String())
	}
}

func TestEliminarServicio_HTTP_NoExiste(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/api/v1/servicios/999999", nil)
	req = chiCtxConID(req, "999999")
	rec := httptest.NewRecorder()

	EliminarServicio(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("esperaba 404, obtuve %d", rec.Code)
	}
}

func TestEliminarServicio_HTTP_IDInvalido(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/api/v1/servicios/abc", nil)
	req = chiCtxConID(req, "abc")
	rec := httptest.NewRecorder()

	EliminarServicio(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuve %d", rec.Code)
	}
}
