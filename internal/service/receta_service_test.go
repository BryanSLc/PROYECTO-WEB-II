package service

import (
	"proyecto/internal/models"
	"testing"
)

type mockAlmacenReceta struct {
	recetas []models.Receta
}

func (m *mockAlmacenReceta) CrearReceta(r models.Receta) models.Receta {
	r.ID = len(m.recetas) + 1
	m.recetas = append(m.recetas, r)
	return r
}
func (m *mockAlmacenReceta) ListarRecetas() []models.Receta { return m.recetas }
func (m *mockAlmacenReceta) ListarRecetasPorMaqueta(maquetaID int) []models.Receta {
	var resultado []models.Receta
	for _, r := range m.recetas {
		if r.MaquetaID == maquetaID {
			resultado = append(resultado, r)
		}
	}
	return resultado
}
func (m *mockAlmacenReceta) BuscarRecetaPorID(id int) (models.Receta, bool) {
	for _, r := range m.recetas {
		if r.ID == id {
			return r, true
		}
	}
	return models.Receta{}, false
}
func (m *mockAlmacenReceta) ActualizarReceta(id int, datos models.Receta) (models.Receta, bool) {
	for i, r := range m.recetas {
		if r.ID == id {
			datos.ID = id
			m.recetas[i] = datos
			return datos, true
		}
	}
	return models.Receta{}, false
}
func (m *mockAlmacenReceta) EliminarReceta(id int) bool {
	for i, r := range m.recetas {
		if r.ID == id {
			m.recetas = append(m.recetas[:i], m.recetas[i+1:]...)
			return true
		}
	}
	return false
}

func TestRecetaService_Crear_Exitoso(t *testing.T) {
	mock := &mockAlmacenReceta{}
	svc := NuevoRecetaService(mock)

	r, err := svc.Crear(models.Receta{Titulo: "Receta A", MaquetaID: 1})
	if err != nil {
		t.Fatalf("no se esperaba error, obtuve: %v", err)
	}
	if r.ID == 0 {
		t.Fatal("se esperaba un ID asignado")
	}
}

func TestRecetaService_Crear_RechazaTituloVacio(t *testing.T) {
	mock := &mockAlmacenReceta{}
	svc := NuevoRecetaService(mock)

	_, err := svc.Crear(models.Receta{MaquetaID: 1})
	if err != ErrTituloObligatorio {
		t.Fatalf("esperaba ErrTituloObligatorio, obtuve: %v", err)
	}
}

func TestRecetaService_Listar(t *testing.T) {
	mock := &mockAlmacenReceta{}
	svc := NuevoRecetaService(mock)

	svc.Crear(models.Receta{Titulo: "Receta A", MaquetaID: 1})
	svc.Crear(models.Receta{Titulo: "Receta B", MaquetaID: 1})

	lista := svc.Listar()
	if len(lista) != 2 {
		t.Fatalf("esperaba 2 recetas, obtuve %d", len(lista))
	}
}

func TestRecetaService_ListarPorMaqueta(t *testing.T) {
	mock := &mockAlmacenReceta{}
	svc := NuevoRecetaService(mock)

	svc.Crear(models.Receta{Titulo: "Receta A", MaquetaID: 1})
	svc.Crear(models.Receta{Titulo: "Receta B", MaquetaID: 2})

	lista := svc.ListarPorMaqueta(1)
	if len(lista) != 1 {
		t.Fatalf("esperaba 1 receta para maqueta 1, obtuve %d", len(lista))
	}
}

func TestRecetaService_BuscarPorID_Existe(t *testing.T) {
	mock := &mockAlmacenReceta{}
	svc := NuevoRecetaService(mock)

	creada, _ := svc.Crear(models.Receta{Titulo: "Receta A", MaquetaID: 1})
	encontrada, err := svc.BuscarPorID(creada.ID)
	if err != nil {
		t.Fatalf("no se esperaba error, obtuve: %v", err)
	}
	if encontrada.Titulo != "Receta A" {
		t.Fatalf("esperaba 'Receta A', obtuve '%s'", encontrada.Titulo)
	}
}

func TestRecetaService_BuscarPorID_NoExiste(t *testing.T) {
	mock := &mockAlmacenReceta{}
	svc := NuevoRecetaService(mock)

	_, err := svc.BuscarPorID(999)
	if err != ErrRecetaNoEncontrada {
		t.Fatalf("esperaba ErrRecetaNoEncontrada, obtuve: %v", err)
	}
}

func TestRecetaService_Actualizar_Exitoso(t *testing.T) {
	mock := &mockAlmacenReceta{}
	svc := NuevoRecetaService(mock)

	creada, _ := svc.Crear(models.Receta{Titulo: "Receta A", MaquetaID: 1})
	actualizada, err := svc.Actualizar(creada.ID, models.Receta{Titulo: "Receta B", MaquetaID: 1})
	if err != nil {
		t.Fatalf("no se esperaba error, obtuve: %v", err)
	}
	if actualizada.Titulo != "Receta B" {
		t.Fatalf("esperaba 'Receta B', obtuve '%s'", actualizada.Titulo)
	}
}

func TestRecetaService_Actualizar_RechazaTituloVacio(t *testing.T) {
	mock := &mockAlmacenReceta{}
	svc := NuevoRecetaService(mock)

	creada, _ := svc.Crear(models.Receta{Titulo: "Receta A", MaquetaID: 1})
	_, err := svc.Actualizar(creada.ID, models.Receta{MaquetaID: 1})
	if err != ErrTituloObligatorio {
		t.Fatalf("esperaba ErrTituloObligatorio, obtuve: %v", err)
	}
}

func TestRecetaService_Actualizar_NoExiste(t *testing.T) {
	mock := &mockAlmacenReceta{}
	svc := NuevoRecetaService(mock)

	_, err := svc.Actualizar(999, models.Receta{Titulo: "Receta B"})
	if err != ErrRecetaNoEncontrada {
		t.Fatalf("esperaba ErrRecetaNoEncontrada, obtuve: %v", err)
	}
}

func TestRecetaService_Eliminar_Exitoso(t *testing.T) {
	mock := &mockAlmacenReceta{}
	svc := NuevoRecetaService(mock)

	creada, _ := svc.Crear(models.Receta{Titulo: "Receta A", MaquetaID: 1})
	err := svc.Eliminar(creada.ID)
	if err != nil {
		t.Fatalf("no se esperaba error, obtuve: %v", err)
	}
}

func TestRecetaService_Eliminar_NoExiste(t *testing.T) {
	mock := &mockAlmacenReceta{}
	svc := NuevoRecetaService(mock)

	err := svc.Eliminar(999)
	if err != ErrRecetaNoEncontrada {
		t.Fatalf("esperaba ErrRecetaNoEncontrada, obtuve: %v", err)
	}
}
