package storage

import (
	"proyecto/internal/models"
)

type UbicacionStorage struct {
	storage *SQLiteStorage
}

func NuevoUbicacionStorage(s *SQLiteStorage) *UbicacionStorage {
	return &UbicacionStorage{
		storage: s,
	}
}

func (s *UbicacionStorage) CrearUbicacion(ubicacion models.Ubicacion) models.Ubicacion {
	return s.storage.CrearUbicacion(ubicacion)
}

func (s *UbicacionStorage) ListarUbicaciones() []models.Ubicacion {
	return s.storage.ListarUbicaciones()
}

func (s *UbicacionStorage) BuscarUbicacionPorID(id int) (models.Ubicacion, bool) {
	return s.storage.BuscarUbicacionPorID(id)
}

func (s *UbicacionStorage) ActualizarUbicacion(id int, datos models.Ubicacion) (models.Ubicacion, bool) {
	return s.storage.ActualizarUbicacion(id, datos)
}

func (s *UbicacionStorage) EliminarUbicacion(id int) bool {
	return s.storage.EliminarUbicacion(id)
}
