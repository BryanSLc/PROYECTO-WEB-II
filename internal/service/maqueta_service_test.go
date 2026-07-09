package service

import (
	"proyecto/internal/models"
	"testing"
)

type mockAlmacenMaqueta struct {
	maquetas   []models.Maqueta
	evoluciones []models.EvolucionMaqueta
}

func (m *mockAlmacenMaqueta) CrearMaqueta(mq models.Maqueta) models.Maqueta {
	mq.ID = len(m.maquetas) + 1
	m.maquetas = append(m.maquetas, mq)
	return mq
}
func (m *mockAlmacenMaqueta) ListarMaquetas() []models.Maqueta { return m.maquetas }
func (m *mockAlmacenMaqueta) BuscarMaquetaPorID(id int) (models.Maqueta, bool) {
	for _, mq := range m.maquetas {
		if mq.ID == id {
			return mq, true
		}
	}
	return models.Maqueta{}, false
}
func (m *mockAlmacenMaqueta) ActualizarMaqueta(id int, datos models.Maqueta) (models.Maqueta, bool) {
	for i, mq := range m.maquetas {
		if mq.ID == id {
			datos.ID = id
			m.maquetas[i] = datos
			return datos, true
		}
	}
	return models.Maqueta{}, false
}
func (m *mockAlmacenMaqueta) EliminarMaqueta(id int) bool {
	for i, mq := range m.maquetas {
		if mq.ID == id {
			m.maquetas = append(m.maquetas[:i], m.maquetas[i+1:]...)
			return true
		}
	}
	return false
}
func (m *mockAlmacenMaqueta) AgregarEvolucion(e models.EvolucionMaqueta) models.EvolucionMaqueta {
	e.ID = len(m.evoluciones) + 1
	m.evoluciones = append(m.evoluciones, e)
	return e
}
func (m *mockAlmacenMaqueta) ListarEvolucionPorMaqueta(id int) []models.EvolucionMaqueta {
	var resultado []models.EvolucionMaqueta
	for _, e := range m.evoluciones {
		if e.MaquetaID == id {
			resultado = append(resultado, e)
		}
	}
	return resultado
}
func (m *mockAlmacenMaqueta) EliminarEvolucion(id int) bool {
	for i, e := range m.evoluciones {
		if e.ID == id {
			m.evoluciones = append(m.evoluciones[:i], m.evoluciones[i+1:]...)
			return true
		}
	}
	return false
}

func TestMaquetaService_Crear_Exitoso(t *testing.T) {
	mock := &mockAlmacenMaqueta{}
	svc := NuevoMaquetaService(mock)

	mq, err := svc.Crear(models.Maqueta{Titulo: "Casa Moderna"})
	if err != nil {
		t.Fatalf("no se esperaba error, obtuve: %v", err)
	}
	if mq.ID == 0 {
		t.Fatal("se esperaba un ID asignado")
	}
}

func TestMaquetaService_Crear_RechazaTituloVacio(t *testing.T) {
	mock := &mockAlmacenMaqueta{}
	svc := NuevoMaquetaService(mock)

	_, err := svc.Crear(models.Maqueta{})
	if err != ErrTituloObligatorio {
		t.Fatalf("esperaba ErrTituloObligatorio, obtuve: %v", err)
	}
}

func TestMaquetaService_Listar(t *testing.T) {
	mock := &mockAlmacenMaqueta{}
	svc := NuevoMaquetaService(mock)

	svc.Crear(models.Maqueta{Titulo: "Casa A"})
	svc.Crear(models.Maqueta{Titulo: "Casa B"})

	lista := svc.Listar()
	if len(lista) != 2 {
		t.Fatalf("esperaba 2 maquetas, obtuve %d", len(lista))
	}
}

func TestMaquetaService_BuscarPorID_Existe(t *testing.T) {
	mock := &mockAlmacenMaqueta{}
	svc := NuevoMaquetaService(mock)

	creada, _ := svc.Crear(models.Maqueta{Titulo: "Casa Moderna"})
	encontrada, err := svc.BuscarPorID(creada.ID)
	if err != nil {
		t.Fatalf("no se esperaba error, obtuve: %v", err)
	}
	if encontrada.Titulo != "Casa Moderna" {
		t.Fatalf("esperaba 'Casa Moderna', obtuve '%s'", encontrada.Titulo)
	}
}

func TestMaquetaService_BuscarPorID_NoExiste(t *testing.T) {
	mock := &mockAlmacenMaqueta{}
	svc := NuevoMaquetaService(mock)

	_, err := svc.BuscarPorID(999)
	if err != ErrMaquetaNoEncontrada {
		t.Fatalf("esperaba ErrMaquetaNoEncontrada, obtuve: %v", err)
	}
}

func TestMaquetaService_Actualizar_Exitoso(t *testing.T) {
	mock := &mockAlmacenMaqueta{}
	svc := NuevoMaquetaService(mock)

	creada, _ := svc.Crear(models.Maqueta{Titulo: "Casa A"})
	actualizada, err := svc.Actualizar(creada.ID, models.Maqueta{Titulo: "Casa B"})
	if err != nil {
		t.Fatalf("no se esperaba error, obtuve: %v", err)
	}
	if actualizada.Titulo != "Casa B" {
		t.Fatalf("esperaba 'Casa B', obtuve '%s'", actualizada.Titulo)
	}
}

func TestMaquetaService_Actualizar_RechazaTituloVacio(t *testing.T) {
	mock := &mockAlmacenMaqueta{}
	svc := NuevoMaquetaService(mock)

	creada, _ := svc.Crear(models.Maqueta{Titulo: "Casa A"})
	_, err := svc.Actualizar(creada.ID, models.Maqueta{})
	if err != ErrTituloObligatorio {
		t.Fatalf("esperaba ErrTituloObligatorio, obtuve: %v", err)
	}
}

func TestMaquetaService_Actualizar_NoExiste(t *testing.T) {
	mock := &mockAlmacenMaqueta{}
	svc := NuevoMaquetaService(mock)

	_, err := svc.Actualizar(999, models.Maqueta{Titulo: "Casa B"})
	if err != ErrMaquetaNoEncontrada {
		t.Fatalf("esperaba ErrMaquetaNoEncontrada, obtuve: %v", err)
	}
}

func TestMaquetaService_Eliminar_Exitoso(t *testing.T) {
	mock := &mockAlmacenMaqueta{}
	svc := NuevoMaquetaService(mock)

	creada, _ := svc.Crear(models.Maqueta{Titulo: "Casa A"})
	err := svc.Eliminar(creada.ID)
	if err != nil {
		t.Fatalf("no se esperaba error, obtuve: %v", err)
	}
}

func TestMaquetaService_Eliminar_NoExiste(t *testing.T) {
	mock := &mockAlmacenMaqueta{}
	svc := NuevoMaquetaService(mock)

	err := svc.Eliminar(999)
	if err != ErrMaquetaNoEncontrada {
		t.Fatalf("esperaba ErrMaquetaNoEncontrada, obtuve: %v", err)
	}
}

func TestMaquetaService_AgregarEvolucion_Exitoso(t *testing.T) {
	mock := &mockAlmacenMaqueta{}
	svc := NuevoMaquetaService(mock)

	creada, _ := svc.Crear(models.Maqueta{Titulo: "Casa A"})
	e, err := svc.AgregarEvolucion(models.EvolucionMaqueta{
		MaquetaID: creada.ID,
		Titulo:    "Avance 1",
		Paso:      1,
	})
	if err != nil {
		t.Fatalf("no se esperaba error, obtuve: %v", err)
	}
	if e.ID == 0 {
		t.Fatal("se esperaba un ID asignado")
	}
}

func TestMaquetaService_AgregarEvolucion_SinMaquetaID(t *testing.T) {
	mock := &mockAlmacenMaqueta{}
	svc := NuevoMaquetaService(mock)

	_, err := svc.AgregarEvolucion(models.EvolucionMaqueta{Titulo: "Avance", Paso: 1})
	if err != ErrIDMaquetaObligatorio {
		t.Fatalf("esperaba ErrIDMaquetaObligatorio, obtuve: %v", err)
	}
}

func TestMaquetaService_AgregarEvolucion_SinTitulo(t *testing.T) {
	mock := &mockAlmacenMaqueta{}
	svc := NuevoMaquetaService(mock)

	creada, _ := svc.Crear(models.Maqueta{Titulo: "Casa A"})
	_, err := svc.AgregarEvolucion(models.EvolucionMaqueta{MaquetaID: creada.ID, Paso: 1})
	if err != ErrTituloAvanceObligatorio {
		t.Fatalf("esperaba ErrTituloAvanceObligatorio, obtuve: %v", err)
	}
}

func TestMaquetaService_AgregarEvolucion_PasoInvalido(t *testing.T) {
	mock := &mockAlmacenMaqueta{}
	svc := NuevoMaquetaService(mock)

	creada, _ := svc.Crear(models.Maqueta{Titulo: "Casa A"})
	_, err := svc.AgregarEvolucion(models.EvolucionMaqueta{MaquetaID: creada.ID, Titulo: "Avance", Paso: 0})
	if err != ErrPasoInvalido {
		t.Fatalf("esperaba ErrPasoInvalido, obtuve: %v", err)
	}
}

func TestMaquetaService_ListarEvolucion_Exitoso(t *testing.T) {
	mock := &mockAlmacenMaqueta{}
	svc := NuevoMaquetaService(mock)

	creada, _ := svc.Crear(models.Maqueta{Titulo: "Casa A"})
	svc.AgregarEvolucion(models.EvolucionMaqueta{MaquetaID: creada.ID, Titulo: "Avance 1", Paso: 1})
	svc.AgregarEvolucion(models.EvolucionMaqueta{MaquetaID: creada.ID, Titulo: "Avance 2", Paso: 2})

	lista, err := svc.ListarEvolucion(creada.ID)
	if err != nil {
		t.Fatalf("no se esperaba error, obtuve: %v", err)
	}
	if len(lista) != 2 {
		t.Fatalf("esperaba 2 evoluciones, obtuve %d", len(lista))
	}
}

func TestMaquetaService_ListarEvolucion_MaquetaNoExiste(t *testing.T) {
	mock := &mockAlmacenMaqueta{}
	svc := NuevoMaquetaService(mock)

	_, err := svc.ListarEvolucion(999)
	if err != ErrMaquetaNoEncontrada {
		t.Fatalf("esperaba ErrMaquetaNoEncontrada, obtuve: %v", err)
	}
}

func TestMaquetaService_EliminarEvolucion_Exitoso(t *testing.T) {
	mock := &mockAlmacenMaqueta{}
	svc := NuevoMaquetaService(mock)

	creada, _ := svc.Crear(models.Maqueta{Titulo: "Casa A"})
	e, _ := svc.AgregarEvolucion(models.EvolucionMaqueta{MaquetaID: creada.ID, Titulo: "Avance 1", Paso: 1})

	err := svc.EliminarEvolucion(e.ID)
	if err != nil {
		t.Fatalf("no se esperaba error, obtuve: %v", err)
	}
}

func TestMaquetaService_EliminarEvolucion_NoExiste(t *testing.T) {
	mock := &mockAlmacenMaqueta{}
	svc := NuevoMaquetaService(mock)

	err := svc.EliminarEvolucion(999)
	if err != ErrEvolucionNoEncontrada {
		t.Fatalf("esperaba ErrEvolucionNoEncontrada, obtuve: %v", err)
	}
}
