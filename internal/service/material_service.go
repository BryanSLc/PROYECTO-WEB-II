package service

import "proyecto/internal/models"

// RepositorioMateriales es la interfaz que permite inyectar un mock en los tests
// sin tocar la base de datos real. SQLiteStorage la cumple automáticamente.
type RepositorioMateriales interface {
	CrearMaterial(m models.MaterialProveedor) models.MaterialProveedor
	ListarMateriales() []models.MaterialProveedor
	BuscarMaterialPorID(id int) (models.MaterialProveedor, bool)
	ActualizarMaterial(id int, datos models.MaterialProveedor) (models.MaterialProveedor, bool)
	EliminarMaterial(id int) bool
}

type MaterialService struct {
	almacen RepositorioMateriales
}

func NuevoMaterialService(a RepositorioMateriales) *MaterialService {
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
