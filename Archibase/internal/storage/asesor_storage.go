package storage

import "proyecto/internal/models"

func (s *SQLiteStorage) CrearAsesor(a models.Asesor) models.Asesor {
	s.db.Create(&a)
	return a
}

func (s *SQLiteStorage) ListarAsesores() []models.Asesor {
	var items []models.Asesor
	s.db.Find(&items)
	return items
}

func (s *SQLiteStorage) BuscarAsesorPorID(id int) (models.Asesor, bool) {
	var item models.Asesor
	if err := s.db.First(&item, id).Error; err != nil {
		return models.Asesor{}, false
	}
	return item, true
}

func (s *SQLiteStorage) ActualizarAsesor(id int, datos models.Asesor) (models.Asesor, bool) {
	var item models.Asesor
	if err := s.db.First(&item, id).Error; err != nil {
		return models.Asesor{}, false
	}

	item.IDasesor = datos.IDasesor
	item.Nombre = datos.Nombre
	item.Especialidad = datos.Especialidad
	item.Experiencia = datos.Experiencia
	item.Contacto = datos.Contacto
	item.Modalidad = datos.Modalidad

	s.db.Save(&item)
	return item, true
}

func (s *SQLiteStorage) EliminarAsesor(id int) bool {
	resultado := s.db.Delete(&models.Asesor{}, id)
	return resultado.RowsAffected > 0
}
