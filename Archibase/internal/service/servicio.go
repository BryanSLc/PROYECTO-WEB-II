package service

import (
	"proyecto/internal/models"
	"proyecto/internal/storage"
)

type ServicioService struct {
	almacen *storage.SQLiteStorage
}

func NuevoServicioService(a *storage.SQLiteStorage) *ServicioService {
	return &ServicioService{almacen: a}
}

func (s *ServicioService) Crear(a models.Servicio) (models.Servicio, error) {
	if a.Titulo == "" {
		return models.Servicio{}, ErrTituloObligatorio
	}
	if a.IDasesor == 0 {
		return models.Servicio{}, ErrIDMaquetaObligatorio
	}
	return s.almacen.CrearServicio(a), nil
}

func (s *ServicioService) Listar() []models.Servicio {
	return s.almacen.ListarServicios()
}

func (s *ServicioService) BuscarPorID(id int) (models.Servicio, error) {
	a, encontrada := s.almacen.BuscarServicioPorID(id)
	if !encontrada {
		return models.Servicio{}, ErrServicioNoEncontrado
	}
	return a, nil
}

func (s *ServicioService) Actualizar(id int, a models.Servicio) (models.Servicio, error) {
	if a.Titulo == "" {
		return models.Servicio{}, ErrTituloObligatorio
	}
	if a.IDasesor == 0 {
		return models.Servicio{}, ErrIDMaquetaObligatorio
	}

	aActualizado, encontrada := s.almacen.ActualizarServicio(id, a)
	if !encontrada {
		return models.Servicio{}, ErrServicioNoEncontrado
	}
	return aActualizado, nil
}

func (s *ServicioService) Eliminar(id int) error {
	if !s.almacen.EliminarServicio(id) {
		return ErrServicioNoEncontrado
	}
	return nil
}
