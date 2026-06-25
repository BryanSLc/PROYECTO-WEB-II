package service

import (
	"proyecto/internal/models"
	"proyecto/internal/storage"
)

type AsesorService struct {
	almacen *storage.SQLiteStorage
}

func NuevoAsesorService(a *storage.SQLiteStorage) *AsesorService {
	return &AsesorService{almacen: a}
}

func (s *AsesorService) Crear(a models.Asesor) (models.Asesor, error) {
	if a.Nombre == "" {
		return models.Asesor{}, ErrNombreObligatorio
	}
	return s.almacen.CrearAsesor(a), nil
}

func (s *AsesorService) Listar() []models.Asesor {
	return s.almacen.ListarAsesores()
}

func (s *AsesorService) BuscarPorID(id int) (models.Asesor, error) {
	a, encontrada := s.almacen.BuscarAsesorPorID(id)
	if !encontrada {
		return models.Asesor{}, ErrAsesorNoEncontrado
	}
	return a, nil
}

func (s *AsesorService) Actualizar(id int, a models.Asesor) (models.Asesor, error) {
	if a.Nombre == "" {
		return models.Asesor{}, ErrNombreObligatorio
	}
	aActualizado, encontrada := s.almacen.ActualizarAsesor(id, a)
	if !encontrada {
		return models.Asesor{}, ErrAsesorNoEncontrado
	}
	return aActualizado, nil
}

func (s *AsesorService) Eliminar(id int) error {
	if !s.almacen.EliminarAsesor(id) {
		return ErrAsesorNoEncontrado
	}
	return nil
}
