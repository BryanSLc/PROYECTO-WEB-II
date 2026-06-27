package service

import (
	"proyecto/internal/models"
	"proyecto/internal/storage"
)

type ProveedorService struct {
	almacen *storage.SQLiteStorage
}

func NuevoProveedorService(a *storage.SQLiteStorage) *ProveedorService {
	return &ProveedorService{almacen: a}
}

func (s *ProveedorService) Crear(p models.Proveedor) (models.Proveedor, error) {
	if p.Nombre == "" {
		return models.Proveedor{}, ErrNombreProveedorObligatorio
	}
	return s.almacen.CrearProveedor(p), nil
}

func (s *ProveedorService) Listar() []models.Proveedor {
	return s.almacen.ListarProveedores()
}

func (s *ProveedorService) BuscarPorID(id int) (models.Proveedor, error) {
	proveedor, encontrado := s.almacen.BuscarProveedorPorID(id)
	if !encontrado {
		return models.Proveedor{}, ErrProveedorNoEncontrado
	}
	return proveedor, nil
}

func (s *ProveedorService) Actualizar(id int, p models.Proveedor) (models.Proveedor, error) {
	if p.Nombre == "" {
		return models.Proveedor{}, ErrNombreProveedorObligatorio
	}
	proveedorActualizado, encontrado := s.almacen.ActualizarProveedor(id, p)
	if !encontrado {
		return models.Proveedor{}, ErrProveedorNoEncontrado
	}
	return proveedorActualizado, nil
}

func (s *ProveedorService) Eliminar(id int) error {
	if !s.almacen.EliminarProveedor(id) {
		return ErrProveedorNoEncontrado
	}
	return nil
}
