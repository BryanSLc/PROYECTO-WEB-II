package storage

import (
	"proyecto/internal/models"
	"testing"
)

// TestSQLiteStorage_CrearYListarAsesor: prueba el repositorio real con GORM
// contra SQLite en memoria (:memory:) — no toca archibase.db.
func TestSQLiteStorage_CrearYListarAsesor(t *testing.T) {
	almacen := NuevoSQLiteStorage("file::memory:?cache=shared&_pragma=foreign_keys(1)")

	asesor := models.Asesor{
		Nombre:       "Arq. Pérez",
		Especialidad: "Diseño Urbano",
		Experiencia:  "10 años",
		Contacto:     "0991234567",
		Modalidad:    "presencial",
	}

	creado := almacen.CrearAsesor(asesor)

	if creado.IDasesor == 0 {
		t.Fatal("GORM debía asignar un ID al crear el asesor")
	}

	encontrado, ok := almacen.BuscarAsesorPorID(creado.IDasesor)
	if !ok {
		t.Fatal("se esperaba encontrar el asesor recién creado")
	}
	if encontrado.Nombre != "Arq. Pérez" {
		t.Fatalf("nombre esperado 'Arq. Pérez', obtuve '%s'", encontrado.Nombre)
	}

	lista := almacen.ListarAsesores()
	if len(lista) != 1 {
		t.Fatalf("se esperaba 1 asesor en la lista, hay %d", len(lista))
	}
}
