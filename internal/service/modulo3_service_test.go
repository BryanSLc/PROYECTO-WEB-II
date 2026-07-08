package service

import (
	"proyecto/internal/models"
	"testing"
)

// ==========================================
//         MOCKS DEL MÓDULO 3
// ==========================================

type mockRepoProveedores struct {
	vecesCrearLlamado int
}

func (m *mockRepoProveedores) CrearProveedor(p models.Proveedor) models.Proveedor {
	m.vecesCrearLlamado++
	p.ID = 1
	return p
}
func (m *mockRepoProveedores) ListarProveedores() []models.Proveedor { return nil }
func (m *mockRepoProveedores) BuscarProveedorPorID(int) (models.Proveedor, bool) {
	return models.Proveedor{}, false
}
func (m *mockRepoProveedores) ActualizarProveedor(int, models.Proveedor) (models.Proveedor, bool) {
	return models.Proveedor{}, false
}
func (m *mockRepoProveedores) EliminarProveedor(int) bool { return false }

type mockRepoMateriales struct {
	vecesCrearLlamado int
}

func (m *mockRepoMateriales) CrearMaterial(mat models.MaterialProveedor) models.MaterialProveedor {
	m.vecesCrearLlamado++
	mat.ID = 1
	return mat
}
func (m *mockRepoMateriales) ListarMateriales() []models.MaterialProveedor { return nil }
func (m *mockRepoMateriales) BuscarMaterialPorID(int) (models.MaterialProveedor, bool) {
	return models.MaterialProveedor{}, false
}
func (m *mockRepoMateriales) ActualizarMaterial(int, models.MaterialProveedor) (models.MaterialProveedor, bool) {
	return models.MaterialProveedor{}, false
}
func (m *mockRepoMateriales) EliminarMaterial(int) bool { return false }

type mockRepoUbicaciones struct {
	vecesCrearLlamado int
}

func (m *mockRepoUbicaciones) CrearUbicacion(u models.Ubicacion) models.Ubicacion {
	m.vecesCrearLlamado++
	u.ID = 1
	return u
}
func (m *mockRepoUbicaciones) ListarUbicaciones() []models.Ubicacion { return nil }
func (m *mockRepoUbicaciones) BuscarUbicacionPorID(int) (models.Ubicacion, bool) {
	return models.Ubicacion{}, false
}
func (m *mockRepoUbicaciones) ActualizarUbicacion(int, models.Ubicacion) (models.Ubicacion, bool) {
	return models.Ubicacion{}, false
}
func (m *mockRepoUbicaciones) EliminarUbicacion(int) bool { return false }

// ==========================================
//         TESTS DE PROVEEDOR
// ==========================================

// TestCrearProveedor_RechazaNombreVacio prueba la regla de negocio:
// un proveedor sin nombre debe ser rechazado y NUNCA llegar al repositorio.
func TestCrearProveedor_RechazaNombreVacio(t *testing.T) {
	mock := &mockRepoProveedores{}
	svc := NuevoProveedorService(mock)

	_, err := svc.Crear(models.Proveedor{Ciudad: "Manta"})

	if err != ErrNombreProveedorObligatorio {
		t.Fatalf("esperaba ErrNombreProveedorObligatorio, obtuve: %v", err)
	}
	if mock.vecesCrearLlamado != 0 {
		t.Fatalf("CrearProveedor no debía llamarse, pero se llamó %d veces", mock.vecesCrearLlamado)
	}
}

// ==========================================
//         TESTS DE MATERIAL
// ==========================================

// TestCrearMaterial_RechazaNombreVacio prueba la regla de negocio:
// un material sin nombre debe ser rechazado y NUNCA llegar al repositorio.
func TestCrearMaterial_RechazaNombreVacio(t *testing.T) {
	mock := &mockRepoMateriales{}
	svc := NuevoMaterialService(mock)

	_, err := svc.Crear(models.MaterialProveedor{Categoria: "Madera"})

	if err != ErrNombreMaterialObligatorio {
		t.Fatalf("esperaba ErrNombreMaterialObligatorio, obtuve: %v", err)
	}
	if mock.vecesCrearLlamado != 0 {
		t.Fatalf("CrearMaterial no debía llamarse, pero se llamó %d veces", mock.vecesCrearLlamado)
	}
}

// ==========================================
//         TESTS DE UBICACIÓN
// ==========================================

// TestCrearUbicacion_RechazaProvinciaVacia prueba la regla de negocio:
// una ubicación sin provincia debe ser rechazada y NUNCA llegar al repositorio.
func TestCrearUbicacion_RechazaProvinciaVacia(t *testing.T) {
	mock := &mockRepoUbicaciones{}
	svc := NuevoUbicacionService(mock)

	_, err := svc.Crear(models.Ubicacion{Ciudad: "Manta"})

	if err != ErrProvinciaUbicacionObligatoria {
		t.Fatalf("esperaba ErrProvinciaUbicacionObligatoria, obtuve: %v", err)
	}
	if mock.vecesCrearLlamado != 0 {
		t.Fatalf("CrearUbicacion no debía llamarse, pero se llamó %d veces", mock.vecesCrearLlamado)
	}
}

// TestCrearUbicacion_RechazaCiudadVacia prueba la regla de negocio:
// una ubicación con provincia pero sin ciudad debe ser rechazada.
func TestCrearUbicacion_RechazaCiudadVacia(t *testing.T) {
	mock := &mockRepoUbicaciones{}
	svc := NuevoUbicacionService(mock)

	_, err := svc.Crear(models.Ubicacion{Provincia: "Manabí"})

	if err != ErrCiudadUbicacionObligatoria {
		t.Fatalf("esperaba ErrCiudadUbicacionObligatoria, obtuve: %v", err)
	}
	if mock.vecesCrearLlamado != 0 {
		t.Fatalf("CrearUbicacion no debía llamarse, pero se llamó %d veces", mock.vecesCrearLlamado)
	}
}
