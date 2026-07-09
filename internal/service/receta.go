package service

import "proyecto/internal/models"

type RepositorioRecetas interface {
	CrearReceta(models.Receta) models.Receta
	ListarRecetas() []models.Receta
	ListarRecetasPorMaqueta(int) []models.Receta
	BuscarRecetaPorID(int) (models.Receta, bool)
	ActualizarReceta(int, models.Receta) (models.Receta, bool)
	EliminarReceta(int) bool
}

type RecetaService struct {
	almacen RepositorioRecetas
}

func NuevoRecetaService(a RepositorioRecetas) *RecetaService {
	return &RecetaService{almacen: a}
}

func (s *RecetaService) Crear(r models.Receta) (models.Receta, error) {
	if r.Titulo == "" {
		return models.Receta{}, ErrTituloObligatorio
	}
	return s.almacen.CrearReceta(r), nil
}

func (s *RecetaService) Listar() []models.Receta {
	return s.almacen.ListarRecetas()
}

func (s *RecetaService) ListarPorMaqueta(maquetaID int) []models.Receta {
	return s.almacen.ListarRecetasPorMaqueta(maquetaID)
}

func (s *RecetaService) BuscarPorID(id int) (models.Receta, error) {
	r, encontrada := s.almacen.BuscarRecetaPorID(id)
	if !encontrada {
		return models.Receta{}, ErrRecetaNoEncontrada
	}
	return r, nil
}

func (s *RecetaService) Actualizar(id int, r models.Receta) (models.Receta, error) {
	if r.Titulo == "" {
		return models.Receta{}, ErrTituloObligatorio
	}
	rActualizada, encontrada := s.almacen.ActualizarReceta(id, r)
	if !encontrada {
		return models.Receta{}, ErrRecetaNoEncontrada
	}
	return rActualizada, nil
}

func (s *RecetaService) Eliminar(id int) error {
	if !s.almacen.EliminarReceta(id) {
		return ErrRecetaNoEncontrada
	}
	return nil
}
