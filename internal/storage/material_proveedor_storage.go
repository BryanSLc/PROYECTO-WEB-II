package storage

import (
	"proyecto/internal/models"
)

type MaterialStorage struct {
	storage *SQLiteStorage
}

func NuevoMaterialStorage(s *SQLiteStorage) *MaterialStorage {
	return &MaterialStorage{
		storage: s,
	}
}

func (s *MaterialStorage) CrearMaterial(material models.MaterialProveedor) models.MaterialProveedor {
	return s.storage.CrearMaterial(material)
}

func (s *MaterialStorage) ListarMateriales() []models.MaterialProveedor {
	return s.storage.ListarMateriales()
}

func (s *MaterialStorage) BuscarMaterialPorID(id int) (models.MaterialProveedor, bool) {
	return s.storage.BuscarMaterialPorID(id)
}

func (s *MaterialStorage) ActualizarMaterial(id int, datos models.MaterialProveedor) (models.MaterialProveedor, bool) {
	return s.storage.ActualizarMaterial(id, datos)
}

func (s *MaterialStorage) EliminarMaterial(id int) bool {
	return s.storage.EliminarMaterial(id)
}
