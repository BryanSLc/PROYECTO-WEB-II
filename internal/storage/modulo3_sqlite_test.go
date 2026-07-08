package storage

import (
	"proyecto/internal/models"
	"testing"
)

// ==========================================
//         TESTS DE MATERIAL
// ==========================================

// TestSQLiteStorage_CrearYListarMaterial: prueba el repositorio real con GORM
// contra SQLite en memoria (:memory:) — no toca archibase.db.
func TestSQLiteStorage_CrearYListarMaterial(t *testing.T) {
	almacen := NuevoPostgresStorage("file::memory:?cache=shared&_pragma=foreign_keys(1)")

	material := models.MaterialProveedor{
		Nombre:            "MDF 15mm",
		Categoria:         "Madera",
		PrecioReferencial: 12.50,
		IDProveedor:       1,
	}

	creado := almacen.CrearMaterial(material)

	if creado.ID == 0 {
		t.Fatal("GORM debía asignar un ID al crear el material")
	}

	encontrado, ok := almacen.BuscarMaterialPorID(creado.ID)
	if !ok {
		t.Fatal("se esperaba encontrar el material recién creado")
	}
	if encontrado.Nombre != "MDF 15mm" {
		t.Fatalf("nombre esperado 'MDF 15mm', obtuve '%s'", encontrado.Nombre)
	}

	lista := almacen.ListarMateriales()
	if len(lista) != 1 {
		t.Fatalf("se esperaba 1 material en la lista, hay %d", len(lista))
	}
}

// ==========================================
//         TESTS DE PROVEEDOR
// ==========================================

// TestSQLiteStorage_CrearYListarProveedor: prueba el repositorio real con GORM
// contra SQLite en memoria (:memory:) — no toca archibase.db.
func TestSQLiteStorage_CrearYListarProveedor(t *testing.T) {
	almacen := NuevoPostgresStorage("file::memory:?cache=shared&_pragma=foreign_keys(1)")

	proveedor := models.Proveedor{
		Nombre:    "Ferretería Central",
		Ciudad:    "Manta",
		Provincia: "Manabí",
		Direccion: "Av. 4 de Noviembre",
		Telefono:  "052123456",
	}

	creado := almacen.CrearProveedor(proveedor)

	if creado.ID == 0 {
		t.Fatal("GORM debía asignar un ID al crear el proveedor")
	}

	encontrado, ok := almacen.BuscarProveedorPorID(creado.ID)
	if !ok {
		t.Fatal("se esperaba encontrar el proveedor recién creado")
	}
	if encontrado.Nombre != "Ferretería Central" {
		t.Fatalf("nombre esperado 'Ferretería Central', obtuve '%s'", encontrado.Nombre)
	}

	lista := almacen.ListarProveedores()
	if len(lista) != 1 {
		t.Fatalf("se esperaba 1 proveedor en la lista, hay %d", len(lista))
	}
}

// ==========================================
//         TESTS DE UBICACIÓN
// ==========================================

// TestSQLiteStorage_CrearYListarUbicacion: prueba el repositorio real con GORM
// contra SQLite en memoria (:memory:) — no toca archibase.db.
func TestSQLiteStorage_CrearYListarUbicacion(t *testing.T) {
	almacen := NuevoPostgresStorage("file::memory:?cache=shared&_pragma=foreign_keys(1)")

	ubicacion := models.Ubicacion{
		Provincia: "Manabí",
		Ciudad:    "Manta",
	}

	creada := almacen.CrearUbicacion(ubicacion)

	if creada.ID == 0 {
		t.Fatal("GORM debía asignar un ID al crear la ubicación")
	}

	encontrada, ok := almacen.BuscarUbicacionPorID(creada.ID)
	if !ok {
		t.Fatal("se esperaba encontrar la ubicación recién creada")
	}
	if encontrada.Provincia != "Manabí" {
		t.Fatalf("provincia esperada 'Manabí', obtuve '%s'", encontrada.Provincia)
	}

	lista := almacen.ListarUbicaciones()
	if len(lista) != 1 {
		t.Fatalf("se esperaba 1 ubicación en la lista, hay %d", len(lista))
	}
}
