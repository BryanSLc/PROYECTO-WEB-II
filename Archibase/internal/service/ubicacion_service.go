package service

import (
	"proyecto/internal/models"
	"proyecto/internal/storage"
)

type UbicacionService struct {
	almacen *storage.SQLiteStorage
}

func NuevoUbicacionService(a *storage.SQLiteStorage) *UbicacionService {
	return &UbicacionService{almacen: a}
}

func (s *UbicacionService) Crear(u models.Ubicacion) (models.Ubicacion, error) {
	if u.Provincia == "" {
		return models.Ubicacion{}, ErrProvinciaUbicacionObligatoria
	}
	if u.Ciudad == "" {
		return models.Ubicacion{}, ErrCiudadUbicacionObligatoria
	}
	return s.almacen.CrearUbicacion(u), nil
}

func (s *UbicacionService) Listar() []models.Ubicacion {
	return s.almacen.ListarUbicaciones()
}

func (s *UbicacionService) BuscarPorID(id int) (models.Ubicacion, error) {
	ubicacion, encontrado := s.almacen.BuscarUbicacionPorID(id)
	if !encontrado {
		return models.Ubicacion{}, ErrUbicacionNoEncontrada
	}
	return ubicacion, nil
}

func (s *UbicacionService) Actualizar(id int, u models.Ubicacion) (models.Ubicacion, error) {
	if u.Provincia == "" {
		return models.Ubicacion{}, ErrProvinciaUbicacionObligatoria
	}
	if u.Ciudad == "" {
		return models.Ubicacion{}, ErrCiudadUbicacionObligatoria
	}
	ubicacionActualizada, encontrado := s.almacen.ActualizarUbicacion(id, u)
	if !encontrado {
		return models.Ubicacion{}, ErrUbicacionNoEncontrada
	}
	return ubicacionActualizada, nil
}

func (s *UbicacionService) Eliminar(id int) error {
	if !s.almacen.EliminarUbicacion(id) {
		return ErrUbicacionNoEncontrada
	}
	return nil
}
