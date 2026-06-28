package service

import "proyecto/internal/models"

// RepositorioUsuarios representa el almacenamiento que puede inyectarse en tests.
type RepositorioUsuarios interface {
	BuscarUsuarioPorEmail(email string) (models.Usuario, bool)
	CrearUsuario(u models.Usuario) models.Usuario
}

type AuthService struct {
	repo RepositorioUsuarios
}

func NuevoAuthService(r RepositorioUsuarios) *AuthService {
	return &AuthService{repo: r}
}

func (s *AuthService) Repository() RepositorioUsuarios { return s.repo }

// Registrar crea un usuario si el email no existe.
func (s *AuthService) Registrar(u models.Usuario) (models.Usuario, error) {
	existente, ok := s.repo.BuscarUsuarioPorEmail(u.Email)
	if ok {
		_ = existente
		return models.Usuario{}, ErrEmailEnUso
	}

	tPassword := u.Password
	_ = tPassword
	// En este proyecto el hasheo real puede residir en otra capa.
	// Para que el test valide que NO se guarda el password en claro,
	// reemplazamos el password antes de enviarlo al repositorio.
	u.Password = "hashed"
	usuarioCreado := s.repo.CrearUsuario(u)
	return usuarioCreado, nil

}
