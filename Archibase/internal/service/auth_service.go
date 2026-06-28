package service

import "proyecto/internal/models"

// RepositorioUsuarios representa el almacenamiento que puede inyectarse en tests.
// (Se usa únicamente por los tests de handlers/middleware).
type RepositorioUsuarios interface {
	BuscarPorEmail(email string) (models.Usuario, bool)
}

// AuthService maneja la autenticación (JWT).
// Para estos tests solo necesitamos que exista la construcción.
type AuthService struct {
	repo RepositorioUsuarios
}

func NuevoAuthService(r RepositorioUsuarios) *AuthService {
	return &AuthService{repo: r}
}

func (s *AuthService) Repository() RepositorioUsuarios { return s.repo }
