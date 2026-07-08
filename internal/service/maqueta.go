package service

import (
	"proyecto/internal/models"
	"proyecto/internal/storage"
)

type MaquetaService struct {
	almacen *storage.SQLiteStorage
}

func NuevoMaquetaService(a *storage.SQLiteStorage) *MaquetaService {
	return &MaquetaService{almacen: a}
}

func (s *MaquetaService) Crear(m models.Maqueta) (models.Maqueta, error) {
	if m.Titulo == "" {
		return models.Maqueta{}, ErrTituloObligatorio
	}
	return s.almacen.CrearMaqueta(m), nil
}

func (s *MaquetaService) Listar() []models.Maqueta {
	return s.almacen.ListarMaquetas()
}

func (s *MaquetaService) BuscarPorID(id int) (models.Maqueta, error) {
	m, encontrada := s.almacen.BuscarMaquetaPorID(id)
	if !encontrada {
		return models.Maqueta{}, ErrMaquetaNoEncontrada
	}
	return m, nil
}

func (s *MaquetaService) Actualizar(id int, m models.Maqueta) (models.Maqueta, error) {
	if m.Titulo == "" {
		return models.Maqueta{}, ErrTituloObligatorio
	}
	mActualizada, encontrada := s.almacen.ActualizarMaqueta(id, m)
	if !encontrada {
		return models.Maqueta{}, ErrMaquetaNoEncontrada
	}
	return mActualizada, nil
}

func (s *MaquetaService) Eliminar(id int) error {
	if !s.almacen.EliminarMaqueta(id) {
		return ErrMaquetaNoEncontrada
	}
	return nil
}

// Métodos de Evolución de la Maqueta
func (s *MaquetaService) AgregarEvolucion(e models.EvolucionMaqueta) (models.EvolucionMaqueta, error) {
	if e.MaquetaID == 0 {
		return models.EvolucionMaqueta{}, ErrIDMaquetaObligatorio
	}
	if e.Titulo == "" {
		return models.EvolucionMaqueta{}, ErrTituloAvanceObligatorio
	}
	if e.Paso <= 0 {
		return models.EvolucionMaqueta{}, ErrPasoInvalido
	}
	_, existe := s.almacen.BuscarMaquetaPorID(e.MaquetaID)
	if !existe {
		return models.EvolucionMaqueta{}, ErrMaquetaNoEncontrada
	}
	return s.almacen.AgregarEvolucion(e), nil
}

func (s *MaquetaService) ListarEvolucion(maquetaID int) ([]models.EvolucionMaqueta, error) {
	_, existe := s.almacen.BuscarMaquetaPorID(maquetaID)
	if !existe {
		return nil, ErrMaquetaNoEncontrada
	}
	return s.almacen.ListarEvolucionPorMaqueta(maquetaID), nil
}

func (s *MaquetaService) EliminarEvolucion(id int) error {
	if !s.almacen.EliminarEvolucion(id) {
		return ErrEvolucionNoEncontrada
	}
	return nil
}
