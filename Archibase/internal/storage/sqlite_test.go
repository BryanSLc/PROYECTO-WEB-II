// Archivo: internal/storage/sqlite_test.go
package storage

import (
	"testing"

	"proyecto/internal/models"
)

// TestSQLiteStorage_CrearYBuscarUsuario prueba el repositorio real (GORM +
// SQLite), pero contra una base en memoria (no el archibase.db real), para
// no ensuciar la base de datos del proyecto con datos de prueba.
func TestSQLiteStorage_CrearYBuscarUsuario(t *testing.T) {
	// "file::memory:?cache=shared" crea una base SQLite que vive solo en RAM
	// durante la ejecución del test, y se descarta al terminar.
	almacen := NuevoSQLiteStorage("file::memory:?cache=shared")

	usuario := models.Usuario{
		Nombre:   "Ana",
		Apellido: "Torres",
		Email:    "ana.torres@uleam.edu.ec",
		Password: "hashficticio",
		Semestre: 3,
		Telefono: "0990001122",
		Rol:      "estudiante",
	}

	creado := almacen.CrearUsuario(usuario)

	if creado.ID == 0 {
		t.Fatal("se esperaba que GORM asignara un ID al crear el usuario")
	}

	encontrado, existe := almacen.BuscarUsuarioPorEmail("ana.torres@uleam.edu.ec")
	if !existe {
		t.Fatal("se esperaba encontrar el usuario recién creado por email")
	}

	if encontrado.Nombre != "Ana" {
		t.Fatalf("se esperaba nombre 'Ana', se obtuvo '%s'", encontrado.Nombre)
	}

	lista := almacen.ListarUsuarios()
	if len(lista) != 1 {
		t.Fatalf("se esperaba 1 usuario en la lista, hay %d", len(lista))
	}
}
