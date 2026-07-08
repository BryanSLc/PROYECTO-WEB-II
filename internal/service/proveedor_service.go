package service

import "proyecto/internal/models"

// RepositorioProveedores es la interfaz que permite inyectar un mock en los tests
// sin tocar la base de datos real. SQLiteStorage la cumple automáticamente.
type RepositorioProveedores interface {
	CrearProveedor(p models.Proveedor) models.Proveedor
	ListarProveedores() []models.Proveedor
	BuscarProveedorPorID(id int) (models.Proveedor, bool)
	ActualizarProveedor(id int, datos models.Proveedor) (models.Proveedor, bool)
	EliminarProveedor(id int) bool
}

type ProveedorService struct {
	almacen RepositorioProveedores
}

func NuevoProveedorService(a RepositorioProveedores) *ProveedorService {
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
