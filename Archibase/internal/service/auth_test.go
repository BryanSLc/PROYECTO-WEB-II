// Archivo: internal/service/auth_test.go
package service

import (
	"testing"

	"proyecto/internal/models"
)

// mockRepositorioUsuarios es un doble de prueba manual (no toca base de
// datos real). Implementa la interfaz RepositorioUsuarios y nos permite
// controlar exactamente qué responde BuscarUsuarioPorEmail, y contar
// cuántas veces se llamó CrearUsuario.
type mockRepositorioUsuarios struct {
	usuarioExistente  models.Usuario
	existe            bool
	vecesCrearLlamado int
}

func (m *mockRepositorioUsuarios) BuscarUsuarioPorEmail(email string) (models.Usuario, bool) {
	return m.usuarioExistente, m.existe
}

func (m *mockRepositorioUsuarios) CrearUsuario(u models.Usuario) models.Usuario {
	m.vecesCrearLlamado++
	u.ID = 1
	return u
}

// TestRegistrar_RechazaEmailEnUso prueba una regla de negocio real:
// si el email ya existe, Registrar debe devolver ErrEmailEnUso y
// CrearUsuario NUNCA debe ejecutarse (el dato inválido no debe llegar
// al repositorio).
func TestRegistrar_RechazaEmailEnUso(t *testing.T) {
	mock := &mockRepositorioUsuarios{
		usuarioExistente: models.Usuario{ID: 5, Email: "ya.existe@uleam.edu.ec"},
		existe:           true, // simulamos que el email ya está registrado
	}
	authService := NuevoAuthService(mock)

	_, err := authService.Registrar(models.Usuario{
		Nombre:   "Pedro",
		Email:    "ya.existe@uleam.edu.ec",
		Password: "claveValida123",
	})

	if err != ErrEmailEnUso {
		t.Fatalf("se esperaba ErrEmailEnUso, se obtuvo: %v", err)
	}

	if mock.vecesCrearLlamado != 0 {
		t.Fatalf("CrearUsuario no debía llamarse, pero se llamó %d veces", mock.vecesCrearLlamado)
	}
}

// TestRegistrar_CreaUsuarioCuandoEmailEsNuevo confirma el caso contrario:
// si el email NO existe, sí debe llegar a CrearUsuario.
func TestRegistrar_CreaUsuarioCuandoEmailEsNuevo(t *testing.T) {
	mock := &mockRepositorioUsuarios{
		existe: false, // el email no existe todavía
	}
	authService := NuevoAuthService(mock)

	usuario, err := authService.Registrar(models.Usuario{
		Nombre:   "Maria",
		Email:    "nueva@uleam.edu.ec",
		Password: "claveValida123",
	})

	if err != nil {
		t.Fatalf("no se esperaba error, se obtuvo: %v", err)
	}

	if mock.vecesCrearLlamado != 1 {
		t.Fatalf("se esperaba 1 llamada a CrearUsuario, hubo %d", mock.vecesCrearLlamado)
	}

	if usuario.Password == "claveValida123" {
		t.Fatal("el password no fue hasheado antes de guardarlo")
	}
}
