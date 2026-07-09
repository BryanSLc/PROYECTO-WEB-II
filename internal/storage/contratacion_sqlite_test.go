package storage

import (
	"proyecto/internal/models"
	"testing"
)

// ==========================================
//          TESTS DE CONTRATACION
// ==========================================

func TestSQLite_Contratacion_CRUD(t *testing.T) {
	a := nuevoAlmacenTest(t)

	serv := a.CrearServicio(models.Servicio{
		Titulo:         "Asesoría estructural",
		Descripcion:    "Revisión de planos estructurales",
		Precio:         80.00,
		Disponibilidad: "fines de semana",
		IDasesor:       1,
	})

	contrat := a.CrearContratacion(models.Contratacion{
		Estudiante: "Carlos Mera",
		Fecha:      "2026-07-08",
		Estado:     "pendiente",
		IDservicio: serv.IDservicio,
	})
	if contrat.IDcontratacion == 0 {
		t.Fatal("esperaba ID asignado al crear contratación")
	}

	lista := a.ListarContrataciones()
	if len(lista) == 0 {
		t.Fatal("esperaba al menos 1 contratación en la lista")
	}

	encontrada, ok := a.BuscarContratacionPorID(contrat.IDcontratacion)
	if !ok {
		t.Fatal("esperaba encontrar la contratación recién creada")
	}
	if encontrada.Estudiante != "Carlos Mera" {
		t.Fatalf("esperaba 'Carlos Mera', obtuve '%s'", encontrada.Estudiante)
	}
	if encontrada.Estado != "pendiente" {
		t.Fatalf("esperaba estado 'pendiente', obtuve '%s'", encontrada.Estado)
	}

	_, ok = a.BuscarContratacionPorID(999)
	if ok {
		t.Fatal("no debía encontrar contratación inexistente")
	}

	actualizada, ok := a.ActualizarContratacion(contrat.IDcontratacion, models.Contratacion{
		Estudiante: "Carlos Mera",
		Fecha:      "2026-07-10",
		Estado:     "confirmada",
		IDservicio: serv.IDservicio,
	})
	if !ok {
		t.Fatal("esperaba actualizar contratación")
	}
	if actualizada.Estado != "confirmada" {
		t.Fatalf("esperaba estado 'confirmada', obtuve '%s'", actualizada.Estado)
	}
	if actualizada.Fecha != "2026-07-10" {
		t.Fatalf("esperaba fecha '2026-07-10', obtuve '%s'", actualizada.Fecha)
	}

	_, ok = a.ActualizarContratacion(999, models.Contratacion{Estudiante: "X"})
	if ok {
		t.Fatal("no debía actualizar contratación inexistente")
	}

	if !a.EliminarContratacion(contrat.IDcontratacion) {
		t.Fatal("esperaba eliminar contratación")
	}
	if a.EliminarContratacion(999) {
		t.Fatal("no debía eliminar contratación inexistente")
	}
}
