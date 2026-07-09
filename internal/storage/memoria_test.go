package storage

import (
	"proyecto/internal/models"
	"testing"
)

func TestMemoria_Usuario_CRUD(t *testing.T) {
	m := NuevaMemoria()

	// Crear
	u := m.CrearUsuario(models.Usuario{Nombre: "Ana", Email: "ana@test.com"})
	if u.ID == 0 {
		t.Fatal("esperaba ID asignado")
	}

	// Listar
	lista := m.ListarUsuarios()
	if len(lista) != 1 {
		t.Fatalf("esperaba 1 usuario, obtuve %d", len(lista))
	}

	// Buscar por ID
	encontrado, ok := m.BuscarUsuarioPorID(u.ID)
	if !ok {
		t.Fatal("esperaba encontrar usuario")
	}
	if encontrado.Nombre != "Ana" {
		t.Fatalf("esperaba 'Ana', obtuve '%s'", encontrado.Nombre)
	}

	// Buscar por ID inexistente
	_, ok = m.BuscarUsuarioPorID(999)
	if ok {
		t.Fatal("no debía encontrar usuario inexistente")
	}

	// Actualizar
	actualizado, ok := m.ActualizarUsuario(u.ID, models.Usuario{Nombre: "Ana Lopez", Email: "ana@test.com"})
	if !ok {
		t.Fatal("esperaba actualizar usuario")
	}
	if actualizado.Nombre != "Ana Lopez" {
		t.Fatalf("esperaba 'Ana Lopez', obtuve '%s'", actualizado.Nombre)
	}

	// Actualizar inexistente
	_, ok = m.ActualizarUsuario(999, models.Usuario{Nombre: "X"})
	if ok {
		t.Fatal("no debía actualizar usuario inexistente")
	}

	// Eliminar
	if !m.EliminarUsuario(u.ID) {
		t.Fatal("esperaba eliminar usuario")
	}
	if m.EliminarUsuario(999) {
		t.Fatal("no debía eliminar usuario inexistente")
	}
}

func TestMemoria_Maqueta_CRUD(t *testing.T) {
	m := NuevaMemoria()

	// Crear
	mq := m.CrearMaqueta(models.Maqueta{Titulo: "Casa A"})
	if mq.ID == 0 {
		t.Fatal("esperaba ID asignado")
	}

	// Listar
	lista := m.ListarMaquetas()
	if len(lista) != 1 {
		t.Fatalf("esperaba 1 maqueta, obtuve %d", len(lista))
	}

	// Buscar por ID
	encontrada, ok := m.BuscarMaquetaPorID(mq.ID)
	if !ok {
		t.Fatal("esperaba encontrar maqueta")
	}
	if encontrada.Titulo != "Casa A" {
		t.Fatalf("esperaba 'Casa A', obtuve '%s'", encontrada.Titulo)
	}

	// Buscar inexistente
	_, ok = m.BuscarMaquetaPorID(999)
	if ok {
		t.Fatal("no debía encontrar maqueta inexistente")
	}

	// Actualizar
	actualizada, ok := m.ActualizarMaqueta(mq.ID, models.Maqueta{Titulo: "Casa B"})
	if !ok {
		t.Fatal("esperaba actualizar maqueta")
	}
	if actualizada.Titulo != "Casa B" {
		t.Fatalf("esperaba 'Casa B', obtuve '%s'", actualizada.Titulo)
	}

	// Actualizar inexistente
	_, ok = m.ActualizarMaqueta(999, models.Maqueta{Titulo: "X"})
	if ok {
		t.Fatal("no debía actualizar maqueta inexistente")
	}

	// Eliminar
	if !m.EliminarMaqueta(mq.ID) {
		t.Fatal("esperaba eliminar maqueta")
	}
	if m.EliminarMaqueta(999) {
		t.Fatal("no debía eliminar maqueta inexistente")
	}
}

func TestMemoria_Evolucion_CRUD(t *testing.T) {
	m := NuevaMemoria()

	mq := m.CrearMaqueta(models.Maqueta{Titulo: "Casa A"})

	// Agregar evolución
	e := m.AgregarEvolucion(models.EvolucionMaqueta{MaquetaID: mq.ID, Titulo: "Avance 1", Paso: 1})
	if e.ID == 0 {
		t.Fatal("esperaba ID asignado")
	}

	// Listar por maqueta
	historial := m.ListarEvolucionPorMaqueta(mq.ID)
	if len(historial) != 1 {
		t.Fatalf("esperaba 1 evolución, obtuve %d", len(historial))
	}

	// Listar de maqueta sin evoluciones
	historial2 := m.ListarEvolucionPorMaqueta(999)
	if len(historial2) != 0 {
		t.Fatal("esperaba lista vacía")
	}

	// Eliminar
	if !m.EliminarEvolucion(e.ID) {
		t.Fatal("esperaba eliminar evolución")
	}
	if m.EliminarEvolucion(999) {
		t.Fatal("no debía eliminar evolución inexistente")
	}
}

func TestMemoria_Receta_CRUD(t *testing.T) {
	m := NuevaMemoria()

	// Crear
	r := m.CrearReceta(models.Receta{Titulo: "Receta A", MaquetaID: 1})
	if r.ID == 0 {
		t.Fatal("esperaba ID asignado")
	}

	// Listar
	lista := m.ListarRecetas()
	if len(lista) != 1 {
		t.Fatalf("esperaba 1 receta, obtuve %d", len(lista))
	}

	// Listar por maqueta
	porMaqueta := m.ListarRecetasPorMaqueta(1)
	if len(porMaqueta) != 1 {
		t.Fatalf("esperaba 1 receta por maqueta, obtuve %d", len(porMaqueta))
	}

	// Buscar por ID
	encontrada, ok := m.BuscarRecetaPorID(r.ID)
	if !ok {
		t.Fatal("esperaba encontrar receta")
	}
	if encontrada.Titulo != "Receta A" {
		t.Fatalf("esperaba 'Receta A', obtuve '%s'", encontrada.Titulo)
	}

	// Buscar inexistente
	_, ok = m.BuscarRecetaPorID(999)
	if ok {
		t.Fatal("no debía encontrar receta inexistente")
	}

	// Actualizar
	actualizada, ok := m.ActualizarReceta(r.ID, models.Receta{Titulo: "Receta B", MaquetaID: 1})
	if !ok {
		t.Fatal("esperaba actualizar receta")
	}
	if actualizada.Titulo != "Receta B" {
		t.Fatalf("esperaba 'Receta B', obtuve '%s'", actualizada.Titulo)
	}

	// Actualizar inexistente
	_, ok = m.ActualizarReceta(999, models.Receta{Titulo: "X"})
	if ok {
		t.Fatal("no debía actualizar receta inexistente")
	}

	// Eliminar
	if !m.EliminarReceta(r.ID) {
		t.Fatal("esperaba eliminar receta")
	}
	if m.EliminarReceta(999) {
		t.Fatal("no debía eliminar receta inexistente")
	}
}
