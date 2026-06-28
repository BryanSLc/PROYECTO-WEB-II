package handlers

import "proyecto/internal/models"

// fakeRepositorioUsuarios es un doble para auth.
// Está aquí para que puedan compilarse los tests de handlers.
// No implementa lógica real.
type fakeRepositorioUsuarios struct{}

func (f *fakeRepositorioUsuarios) BuscarPorEmail(email string) (models.Usuario, bool) {
	return models.Usuario{}, false
}
