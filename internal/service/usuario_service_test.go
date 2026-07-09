package service

import (
	"proyecto/internal/models"
	"testing"
)

type mockAlmacenUsuario struct {
	usuarios          []models.Usuario
	vecesCrearLlamado int
}

func (m *mockAlmacenUsuario) CrearUsuario(u models.Usuario) models.Usuario {
	m.vecesCrearLlamado++
	u.ID = len(m.usuarios) + 1
	m.usuarios = append(m.usuarios, u)
	return u
}
func (m *mockAlmacenUsuario) ListarUsuarios() []models.Usuario { return m.usuarios }
func (m *mockAlmacenUsuario) BuscarUsuarioPorID(id int) (models.Usuario, bool) {
	for _, u := range m.usuarios {
		if u.ID == id {
			return u, true
		}
	}
	return models.Usuario{}, false
}
func (m *mockAlmacenUsuario) ActualizarUsuario(id int, datos models.Usuario) (models.Usuario, bool) {
	for i, u := range m.usuarios {
		if u.ID == id {
			datos.ID = id
			m.usuarios[i] = datos
			return datos, true
		}
	}
	return models.Usuario{}, false
}
func (m *mockAlmacenUsuario) EliminarUsuario(id int) bool {
	for i, u := range m.usuarios {
		if u.ID == id {
			m.usuarios = append(m.usuarios[:i], m.usuarios[i+1:]...)
			return true
		}
	}
	return false
}

func TestUsuarioService_Crear_Exitoso(t *testing.T) {
	mock := &mockAlmacenUsuario{}
	svc := NuevoUsuarioService(mock)

	u, err := svc.Crear(models.Usuario{Nombre: "Ana", Email: "ana@test.com"})
	if err != nil {
		t.Fatalf("no se esperaba error, obtuve: %v", err)
	}
	if u.ID == 0 {
		t.Fatal("se esperaba un ID asignado")
	}
	if mock.vecesCrearLlamado != 1 {
		t.Fatalf("esperaba 1 llamada a CrearUsuario, hubo %d", mock.vecesCrearLlamado)
	}
}

func TestUsuarioService_Crear_RechazaNombreVacio(t *testing.T) {
	mock := &mockAlmacenUsuario{}
	svc := NuevoUsuarioService(mock)

	_, err := svc.Crear(models.Usuario{Email: "test@test.com"})
	if err != ErrNombreObligatorio {
		t.Fatalf("esperaba ErrNombreObligatorio, obtuve: %v", err)
	}
	if mock.vecesCrearLlamado != 0 {
		t.Fatal("CrearUsuario no debía llamarse")
	}
}

func TestUsuarioService_Crear_RechazaEmailVacio(t *testing.T) {
	mock := &mockAlmacenUsuario{}
	svc := NuevoUsuarioService(mock)

	_, err := svc.Crear(models.Usuario{Nombre: "Ana"})
	if err != ErrEmailObligatorio {
		t.Fatalf("esperaba ErrEmailObligatorio, obtuve: %v", err)
	}
}

func TestUsuarioService_Listar(t *testing.T) {
	mock := &mockAlmacenUsuario{}
	svc := NuevoUsuarioService(mock)

	svc.Crear(models.Usuario{Nombre: "Ana", Email: "ana@test.com"})
	svc.Crear(models.Usuario{Nombre: "Luis", Email: "luis@test.com"})

	lista := svc.Listar()
	if len(lista) != 2 {
		t.Fatalf("esperaba 2 usuarios, obtuve %d", len(lista))
	}
}

func TestUsuarioService_BuscarPorID_Existe(t *testing.T) {
	mock := &mockAlmacenUsuario{}
	svc := NuevoUsuarioService(mock)

	creado, _ := svc.Crear(models.Usuario{Nombre: "Ana", Email: "ana@test.com"})
	encontrado, err := svc.BuscarPorID(creado.ID)
	if err != nil {
		t.Fatalf("no se esperaba error, obtuve: %v", err)
	}
	if encontrado.Nombre != "Ana" {
		t.Fatalf("esperaba 'Ana', obtuve '%s'", encontrado.Nombre)
	}
}

func TestUsuarioService_BuscarPorID_NoExiste(t *testing.T) {
	mock := &mockAlmacenUsuario{}
	svc := NuevoUsuarioService(mock)

	_, err := svc.BuscarPorID(999)
	if err != ErrUsuarioNoEncontrado {
		t.Fatalf("esperaba ErrUsuarioNoEncontrado, obtuve: %v", err)
	}
}

func TestUsuarioService_Actualizar_Exitoso(t *testing.T) {
	mock := &mockAlmacenUsuario{}
	svc := NuevoUsuarioService(mock)

	creado, _ := svc.Crear(models.Usuario{Nombre: "Ana", Email: "ana@test.com"})
	actualizado, err := svc.Actualizar(creado.ID, models.Usuario{Nombre: "Ana Lopez", Email: "ana@test.com"})
	if err != nil {
		t.Fatalf("no se esperaba error, obtuve: %v", err)
	}
	if actualizado.Nombre != "Ana Lopez" {
		t.Fatalf("esperaba 'Ana Lopez', obtuve '%s'", actualizado.Nombre)
	}
}

func TestUsuarioService_Actualizar_RechazaNombreVacio(t *testing.T) {
	mock := &mockAlmacenUsuario{}
	svc := NuevoUsuarioService(mock)

	creado, _ := svc.Crear(models.Usuario{Nombre: "Ana", Email: "ana@test.com"})
	_, err := svc.Actualizar(creado.ID, models.Usuario{Email: "ana@test.com"})
	if err != ErrNombreObligatorio {
		t.Fatalf("esperaba ErrNombreObligatorio, obtuve: %v", err)
	}
}

func TestUsuarioService_Actualizar_NoExiste(t *testing.T) {
	mock := &mockAlmacenUsuario{}
	svc := NuevoUsuarioService(mock)

	_, err := svc.Actualizar(999, models.Usuario{Nombre: "Ana", Email: "ana@test.com"})
	if err != ErrUsuarioNoEncontrado {
		t.Fatalf("esperaba ErrUsuarioNoEncontrado, obtuve: %v", err)
	}
}

func TestUsuarioService_Eliminar_Exitoso(t *testing.T) {
	mock := &mockAlmacenUsuario{}
	svc := NuevoUsuarioService(mock)

	creado, _ := svc.Crear(models.Usuario{Nombre: "Ana", Email: "ana@test.com"})
	err := svc.Eliminar(creado.ID)
	if err != nil {
		t.Fatalf("no se esperaba error, obtuve: %v", err)
	}
}

func TestUsuarioService_Eliminar_NoExiste(t *testing.T) {
	mock := &mockAlmacenUsuario{}
	svc := NuevoUsuarioService(mock)

	err := svc.Eliminar(999)
	if err != ErrUsuarioNoEncontrado {
		t.Fatalf("esperaba ErrUsuarioNoEncontrado, obtuve: %v", err)
	}
}
