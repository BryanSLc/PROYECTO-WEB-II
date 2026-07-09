package storage

import (
	"proyecto/internal/models"
	"testing"
)

// ==========================================
//          TESTS DE SERVICIO
// ==========================================

func TestSQLite_Servicio_CRUD(t *testing.T) {
	a := nuevoAlmacenTest(t)

	serv := a.CrearServicio(models.Servicio{
		Titulo:         "Diseño de fachada",
		Descripcion:    "Renderizado 3D de fachada residencial",
		Precio:         150.00,
		Disponibilidad: "lunes a viernes",
		IDasesor:       1,
	})
	if serv.IDservicio == 0 {
		t.Fatal("esperaba ID asignado al crear servicio")
	}

	lista := a.ListarServicios()
	if len(lista) == 0 {
		t.Fatal("esperaba al menos 1 servicio en la lista")
	}

	encontrado, ok := a.BuscarServicioPorID(serv.IDservicio)
	if !ok {
		t.Fatal("esperaba encontrar el servicio recién creado")
	}
	if encontrado.Titulo != "Diseño de fachada" {
		t.Fatalf("esperaba 'Diseño de fachada', obtuve '%s'", encontrado.Titulo)
	}
	if encontrado.Precio != 150.00 {
		t.Fatalf("esperaba precio 150.00, obtuve %v", encontrado.Precio)
	}

	_, ok = a.BuscarServicioPorID(999)
	if ok {
		t.Fatal("no debía encontrar servicio inexistente")
	}

	actualizado, ok := a.ActualizarServicio(serv.IDservicio, models.Servicio{
		Titulo:         "Diseño de fachada Premium",
		Descripcion:    "Renderizado 3D + planos",
		Precio:         220.00,
		Disponibilidad: "todos los días",
		IDasesor:       2,
	})
	if !ok {
		t.Fatal("esperaba actualizar servicio")
	}
	if actualizado.Titulo != "Diseño de fachada Premium" {
		t.Fatalf("esperaba 'Diseño de fachada Premium', obtuve '%s'", actualizado.Titulo)
	}
	if actualizado.Precio != 220.00 {
		t.Fatalf("esperaba precio 220.00, obtuve %v", actualizado.Precio)
	}
	if actualizado.IDasesor != 2 {
		t.Fatalf("esperaba IDasesor 2, obtuve %d", actualizado.IDasesor)
	}

	_, ok = a.ActualizarServicio(999, models.Servicio{Titulo: "X"})
	if ok {
		t.Fatal("no debía actualizar servicio inexistente")
	}

	if !a.EliminarServicio(serv.IDservicio) {
		t.Fatal("esperaba eliminar servicio")
	}
	if a.EliminarServicio(999) {
		t.Fatal("no debía eliminar servicio inexistente")
	}
}
