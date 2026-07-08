package service

import "proyecto/internal/models"

// RepositorioAsesores es la interfaz que permite inyectar un mock en los tests
// sin tocar la base de datos real. SQLiteStorage la cumple automáticamente.
type RepositorioAsesores interface {
	CrearAsesor(a models.Asesor) models.Asesor
	ListarAsesores() []models.Asesor
	BuscarAsesorPorID(id int) (models.Asesor, bool)
	ActualizarAsesor(id int, datos models.Asesor) (models.Asesor, bool)
	EliminarAsesor(id int) bool
}

type AsesorService struct {
	almacen RepositorioAsesores
}

func NuevoAsesorService(a RepositorioAsesores) *AsesorService {
	return &AsesorService{almacen: a}
}

func (s *AsesorService) Crear(a models.Asesor) (models.Asesor, error) {
	if a.Nombre == "" {
		return models.Asesor{}, ErrNombreAsesorObligatorio
	}
	return s.almacen.CrearAsesor(a), nil
}

func (s *AsesorService) Listar() []models.Asesor {
	return s.almacen.ListarAsesores()
}

func (s *AsesorService) BuscarPorID(id int) (models.Asesor, error) {
	a, ok := s.almacen.BuscarAsesorPorID(id)
	if !ok {
		return models.Asesor{}, ErrAsesorNoEncontrado
	}
	return a, nil
}

func (s *AsesorService) Actualizar(id int, a models.Asesor) (models.Asesor, error) {
	if a.Nombre == "" {
		return models.Asesor{}, ErrNombreAsesorObligatorio
	}
	actualizado, ok := s.almacen.ActualizarAsesor(id, a)
	if !ok {
		return models.Asesor{}, ErrAsesorNoEncontrado
	}
	return actualizado, nil
}

func (s *AsesorService) Eliminar(id int) error {
	if !s.almacen.EliminarAsesor(id) {
		return ErrAsesorNoEncontrado
	}
	return nil
}
