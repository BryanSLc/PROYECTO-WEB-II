package storage

import "proyecto/internal/models"

// CRUD de Contrataciones usando SQLiteStorage.

func (s *SQLiteStorage) CrearContratacion(c models.Contratacion) models.Contratacion {
	s.db.Create(&c)
	return c
}

func (s *SQLiteStorage) ListarContrataciones() []models.Contratacion {
	var items []models.Contratacion
	s.db.Find(&items)
	return items
}

func (s *SQLiteStorage) BuscarContratacionPorID(id int) (models.Contratacion, bool) {
	var item models.Contratacion
	if err := s.db.First(&item, id).Error; err != nil {
		return models.Contratacion{}, false
	}
	return item, true
}

func (s *SQLiteStorage) ActualizarContratacion(id int, datos models.Contratacion) (models.Contratacion, bool) {
	var item models.Contratacion
	if err := s.db.First(&item, id).Error; err != nil {
		return models.Contratacion{}, false
	}

	// Actualizamos campos
	item.Estudiante = datos.Estudiante
	item.Fecha = datos.Fecha
	item.Estado = datos.Estado
	item.IDservicio = datos.IDservicio

	s.db.Save(&item)
	return item, true
}

func (s *SQLiteStorage) EliminarContratacion(id int) bool {
	resultado := s.db.Delete(&models.Contratacion{}, id)
	return resultado.RowsAffected > 0
}
