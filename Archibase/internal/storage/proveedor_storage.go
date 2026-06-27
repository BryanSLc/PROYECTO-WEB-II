package storage

import "proyecto/internal/models"

type MaterialMemoryStorage struct {
	listaMateriales  []models.MaterialProveedor
	conteoMateriales int
}

func NuevoMaterialMemoryStorage() *MaterialMemoryStorage {
	return &MaterialMemoryStorage{
		listaMateriales:  []models.MaterialProveedor{},
		conteoMateriales: 1,
	}
}

func (s *MaterialMemoryStorage) CrearMaterial(m models.MaterialProveedor) models.MaterialProveedor {
	m.ID = s.conteoMateriales
	s.conteoMateriales++
	s.listaMateriales = append(s.listaMateriales, m)
	return m
}

func (s *MaterialMemoryStorage) ListarMateriales() []models.MaterialProveedor {
	return s.listaMateriales
}

func (s *MaterialMemoryStorage) BuscarMaterialPorID(id int) (models.MaterialProveedor, bool) {
	for _, m := range s.listaMateriales {
		if m.ID == id {
			return m, true
		}
	}
	return models.MaterialProveedor{}, false
}

func (s *MaterialMemoryStorage) ActualizarMaterial(id int, datos models.MaterialProveedor) (models.MaterialProveedor, bool) {
	for i, m := range s.listaMateriales {
		if m.ID == id {
			datos.ID = id
			s.listaMateriales[i] = datos
			return datos, true
		}
	}
	return models.MaterialProveedor{}, false
}

func (s *MaterialMemoryStorage) EliminarMaterial(id int) bool {
	for i, m := range s.listaMateriales {
		if m.ID == id {
			s.listaMateriales = append(s.listaMateriales[:i], s.listaMateriales[i+1:]...)
			return true
		}
	}
	return false
}
