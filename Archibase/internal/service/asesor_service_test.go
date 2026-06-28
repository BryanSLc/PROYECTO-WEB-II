package service

import (
	"proyecto/internal/models"
	"testing"
)

type mockRepoAsesores struct {
	vecesCrearLlamado int
}

func (m *mockRepoAsesores) CrearAsesor(a models.Asesor) models.Asesor {
	m.vecesCrearLlamado++
	a.IDasesor = 1
	return a
}
func (m *mockRepoAsesores) ListarAsesores() []models.Asesor { return nil }
func (m *mockRepoAsesores) BuscarAsesorPorID(int) (models.Asesor, bool) {
	return models.Asesor{}, false
}
func (m *mockRepoAsesores) ActualizarAsesor(int, models.Asesor) (models.Asesor, bool) {
	return models.Asesor{}, false
}
func (m *mockRepoAsesores) EliminarAsesor(int) bool { return false }

// TestCrear_RechazaNombreVacio prueba la regla de negocio real:
// un asesor sin nombre debe ser rechazado y NUNCA llegar al repositorio.
func TestCrear_RechazaNombreVacio(t *testing.T) {
	mock := &mockRepoAsesores{}
	svc := NuevoAsesorService(mock)

	_, err := svc.Crear(models.Asesor{Especialidad: "Estructuras"})

	if err != ErrNombreAsesorObligatorio {
		t.Fatalf("esperaba ErrNombreAsesorObligatorio, obtuve: %v", err)
	}
	if mock.vecesCrearLlamado != 0 {
		t.Fatalf("CrearAsesor no debía llamarse, pero se llamó %d veces", mock.vecesCrearLlamado)
	}
}
