package storage

import "proyecto/internal/models"

type UbicacionStorage struct {
	listaUbicaciones  []models.Ubicacion
	conteoUbicaciones int
}

func NuevoUbicacionStorage() *UbicacionStorage {
	return &UbicacionStorage{
		listaUbicaciones:  []models.Ubicacion{},
		conteoUbicaciones: 1,
	}
}

func (s *UbicacionStorage) CrearUbicacion(u models.Ubicacion) models.Ubicacion {
	u.ID = s.conteoUbicaciones
	s.conteoUbicaciones++
	s.listaUbicaciones = append(s.listaUbicaciones, u)
	return u
}

func (s *UbicacionStorage) ListarUbicaciones() []models.Ubicacion {
	return s.listaUbicaciones
}

func (s *UbicacionStorage) BuscarUbicacionPorID(id int) (models.Ubicacion, bool) {
	for _, u := range s.listaUbicaciones {
		if u.ID == id {
			return u, true
		}
	}
	return models.Ubicacion{}, false
}

func (s *UbicacionStorage) ActualizarUbicacion(id int, datos models.Ubicacion) (models.Ubicacion, bool) {
	for i, u := range s.listaUbicaciones {
		if u.ID == id {
			datos.ID = id
			s.listaUbicaciones[i] = datos
			return datos, true
		}
	}
	return models.Ubicacion{}, false
}

func (s *UbicacionStorage) EliminarUbicacion(id int) bool {
	for i, u := range s.listaUbicaciones {
		if u.ID == id {
			s.listaUbicaciones = append(s.listaUbicaciones[:i], s.listaUbicaciones[i+1:]...)
			return true
		}
	}
	return false
}
