package storage

import (
	"proyecto/internal/models"
)

type ProveedorStorage struct {
	listaProveedores  []models.Proveedor
	conteoProveedores int
}

func NuevoProveedorStorage() *ProveedorStorage {
	return &ProveedorStorage{
		listaProveedores:  []models.Proveedor{},
		conteoProveedores: 1,
	}
}

func (s *ProveedorStorage) CrearProveedor(proveedor models.Proveedor) models.Proveedor {
	proveedor.ID = s.conteoProveedores
	s.conteoProveedores++
	s.listaProveedores = append(s.listaProveedores, proveedor)
	return proveedor
}

func (s *ProveedorStorage) ListarProveedores() []models.Proveedor {
	return s.listaProveedores
}

func (s *ProveedorStorage) BuscarProveedorPorID(id int) (models.Proveedor, bool) {
	for _, proveedor := range s.listaProveedores {
		if proveedor.ID == id {
			return proveedor, true
		}
	}
	return models.Proveedor{}, false
}

func (s *ProveedorStorage) ActualizarProveedor(id int, datos models.Proveedor) (models.Proveedor, bool) {
	for i, proveedor := range s.listaProveedores {
		if proveedor.ID == id {
			datos.ID = id
			s.listaProveedores[i] = datos
			return datos, true
		}
	}
	return models.Proveedor{}, false
}

func (s *ProveedorStorage) EliminarProveedor(id int) bool {
	for i, proveedor := range s.listaProveedores {
		if proveedor.ID == id {
			s.listaProveedores = append(s.listaProveedores[:i], s.listaProveedores[i+1:]...)
			return true
		}
	}
	return false
}
