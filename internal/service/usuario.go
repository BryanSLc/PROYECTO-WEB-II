package service

import "proyecto/internal/models"

type RepositorioUsuariosService interface {
	CrearUsuario(models.Usuario) models.Usuario
	ListarUsuarios() []models.Usuario
	BuscarUsuarioPorID(int) (models.Usuario, bool)
	ActualizarUsuario(int, models.Usuario) (models.Usuario, bool)
	EliminarUsuario(int) bool
}

type UsuarioService struct {
	almacen RepositorioUsuariosService
}

func NuevoUsuarioService(a RepositorioUsuariosService) *UsuarioService {
	return &UsuarioService{almacen: a}
}

func (s *UsuarioService) Crear(u models.Usuario) (models.Usuario, error) {
	if u.Nombre == "" {
		return models.Usuario{}, ErrNombreObligatorio
	}
	if u.Email == "" {
		return models.Usuario{}, ErrEmailObligatorio
	}
	return s.almacen.CrearUsuario(u), nil
}

func (s *UsuarioService) Listar() []models.Usuario {
	return s.almacen.ListarUsuarios()
}

func (s *UsuarioService) BuscarPorID(id int) (models.Usuario, error) {
	u, encontrado := s.almacen.BuscarUsuarioPorID(id)
	if !encontrado {
		return models.Usuario{}, ErrUsuarioNoEncontrado
	}
	return u, nil
}

func (s *UsuarioService) Actualizar(id int, u models.Usuario) (models.Usuario, error) {
	if u.Nombre == "" {
		return models.Usuario{}, ErrNombreObligatorio
	}
	uActualizado, encontrado := s.almacen.ActualizarUsuario(id, u)
	if !encontrado {
		return models.Usuario{}, ErrUsuarioNoEncontrado
	}
	return uActualizado, nil
}

func (s *UsuarioService) Eliminar(id int) error {
	if !s.almacen.EliminarUsuario(id) {
		return ErrUsuarioNoEncontrado
	}
	return nil
}
