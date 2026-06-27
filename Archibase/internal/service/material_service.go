package service

import (
	"proyecto/internal/models"
	"proyecto/internal/storage"
)

type MaterialService struct {
	almacen *storage.SQLiteStorage
}

func NuevoMaterialService(a *storage.SQLiteStorage) *MaterialService {
	return &MaterialService{almacen: a}
}

func (s *MaterialService) Crear(m models.MaterialProveedor) (models.MaterialProveedor, error) {
	if m.Nombre == "" {
		return models.MaterialProveedor{}, ErrNombreMaterialObligatorio
	}
	return s.almacen.CrearMaterial(m), nil
}

func (s *MaterialService) Listar() []models.MaterialProveedor {
	return s.almacen.ListarMateriales()
}

func (s *MaterialService) BuscarPorID(id int) (models.MaterialProveedor, error) {
	material, encontrado := s.almacen.BuscarMaterialPorID(id)
	if !encontrado {
		return models.MaterialProveedor{}, ErrMaterialNoEncontrado
	}
	return material, nil
}

func (s *MaterialService) Actualizar(id int, m models.MaterialProveedor) (models.MaterialProveedor, error) {
	if m.Nombre == "" {
		return models.MaterialProveedor{}, ErrNombreMaterialObligatorio
	}
	materialActualizado, encontrado := s.almacen.ActualizarMaterial(id, m)
	if !encontrado {
		return models.MaterialProveedor{}, ErrMaterialNoEncontrado
	}
	return materialActualizado, nil
}

func (s *MaterialService) Eliminar(id int) error {
	if !s.almacen.EliminarMaterial(id) {
		return ErrMaterialNoEncontrado
	}
	return nil
}
