package service

import (
	"proyecto/internal/models"
	"proyecto/internal/storage"
)

type ContratacionService struct {
	almacen *storage.SQLiteStorage
}

func NuevoContratacionService(a *storage.SQLiteStorage) *ContratacionService {
	return &ContratacionService{almacen: a}
}

func (s *ContratacionService) Crear(a models.Contratacion) (models.Contratacion, error) {
	if a.Estudiante == "" {
		return models.Contratacion{}, ErrDatosInvalidos
	}
	if a.IDservicio == 0 {
		return models.Contratacion{}, ErrIDMaquetaObligatorio
	}
	return s.almacen.CrearContratacion(a), nil
}

func (s *ContratacionService) Listar() []models.Contratacion {
	return s.almacen.ListarContrataciones()
}

func (s *ContratacionService) BuscarPorID(id int) (models.Contratacion, error) {
	a, encontrada := s.almacen.BuscarContratacionPorID(id)
	if !encontrada {
		return models.Contratacion{}, ErrContratacionNoEncontrada
	}
	return a, nil
}

func (s *ContratacionService) Actualizar(id int, a models.Contratacion) (models.Contratacion, error) {
	if a.Estudiante == "" {
		return models.Contratacion{}, ErrDatosInvalidos
	}
	if a.IDservicio == 0 {
		return models.Contratacion{}, ErrIDMaquetaObligatorio
	}
	aActualizado, encontrada := s.almacen.ActualizarContratacion(id, a)

	if !encontrada {
		return models.Contratacion{}, ErrContratacionNoEncontrada
	}
	return aActualizado, nil
}

func (s *ContratacionService) Eliminar(id int) error {
	if !s.almacen.EliminarContratacion(id) {
		return ErrContratacionNoEncontrada
	}
	return nil
}
