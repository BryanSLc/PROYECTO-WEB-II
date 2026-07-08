package storage

import (
	"proyecto/internal/models"
	"time"
)

type Memoria struct {
	maquetas        []models.Maqueta
	nextMaquetaID   int
	evoluciones     []models.EvolucionMaqueta
	nextEvolucionID int
	usuarios        []models.Usuario
	nextUsuarioID   int
	recetas         []models.Receta
	nextRecetaID    int
}

func NuevaMemoria() *Memoria {
	return &Memoria{
		maquetas:        []models.Maqueta{},
		nextMaquetaID:   1,
		evoluciones:     []models.EvolucionMaqueta{},
		nextEvolucionID: 1,
		usuarios:        []models.Usuario{},
		nextUsuarioID:   1,
		recetas:         []models.Receta{},
		nextRecetaID:    1,
	}
}

func (m *Memoria) CrearMaqueta(maqueta models.Maqueta) models.Maqueta {
	maqueta.ID = m.nextMaquetaID
	m.nextMaquetaID++
	m.maquetas = append(m.maquetas, maqueta)
	return maqueta
}

func (m *Memoria) ListarMaquetas() []models.Maqueta {
	return m.maquetas
}

func (m *Memoria) BuscarMaquetaPorID(id int) (models.Maqueta, bool) {
	for _, maqueta := range m.maquetas {
		if maqueta.ID == id {
			return maqueta, true
		}
	}
	return models.Maqueta{}, false
}

func (m *Memoria) ActualizarMaqueta(id int, datos models.Maqueta) (models.Maqueta, bool) {
	for i, maqueta := range m.maquetas {
		if maqueta.ID == id {
			datos.ID = id
			m.maquetas[i] = datos
			return datos, true
		}
	}
	return models.Maqueta{}, false
}

func (m *Memoria) EliminarMaqueta(id int) bool {
	for i, maqueta := range m.maquetas {
		if maqueta.ID == id {
			m.maquetas = append(m.maquetas[:i], m.maquetas[i+1:]...)
			return true
		}
	}
	return false
}

func (m *Memoria) AgregarEvolucion(evolucion models.EvolucionMaqueta) models.EvolucionMaqueta {
	evolucion.ID = m.nextEvolucionID
	m.nextEvolucionID++
	evolucion.Fecha = time.Now()
	m.evoluciones = append(m.evoluciones, evolucion)
	return evolucion
}

func (m *Memoria) ListarEvolucionPorMaqueta(maquetaID int) []models.EvolucionMaqueta {
	var historial []models.EvolucionMaqueta
	for _, ev := range m.evoluciones {
		if ev.MaquetaID == maquetaID {
			historial = append(historial, ev)
		}
	}
	return historial
}

func (m *Memoria) EliminarEvolucion(id int) bool {
	for i, ev := range m.evoluciones {
		if ev.ID == id {
			m.evoluciones = append(m.evoluciones[:i], m.evoluciones[i+1:]...)
			return true
		}
	}
	return false
}

func (m *Memoria) CrearUsuario(usuario models.Usuario) models.Usuario {
	usuario.ID = m.nextUsuarioID
	m.nextUsuarioID++
	usuario.FechaCreacion = time.Now()
	m.usuarios = append(m.usuarios, usuario)
	return usuario
}

func (m *Memoria) ListarUsuarios() []models.Usuario {
	return m.usuarios
}

func (m *Memoria) BuscarUsuarioPorID(id int) (models.Usuario, bool) {
	for _, usuario := range m.usuarios {
		if usuario.ID == id {
			return usuario, true
		}
	}
	return models.Usuario{}, false
}

func (m *Memoria) ActualizarUsuario(id int, datos models.Usuario) (models.Usuario, bool) {
	for i, usuario := range m.usuarios {
		if usuario.ID == id {
			datos.ID = id
			datos.FechaCreacion = usuario.FechaCreacion
			m.usuarios[i] = datos
			return datos, true
		}
	}
	return models.Usuario{}, false
}

func (m *Memoria) EliminarUsuario(id int) bool {
	for i, usuario := range m.usuarios {
		if usuario.ID == id {
			m.usuarios = append(m.usuarios[:i], m.usuarios[i+1:]...)
			return true
		}
	}
	return false
}

func (m *Memoria) CrearReceta(receta models.Receta) models.Receta {
	receta.ID = m.nextRecetaID
	m.nextRecetaID++
	m.recetas = append(m.recetas, receta)
	return receta
}

func (m *Memoria) ListarRecetas() []models.Receta {
	return m.recetas
}

func (m *Memoria) ListarRecetasPorMaqueta(maquetaID int) []models.Receta {
	var recetasFiltradas []models.Receta
	for _, receta := range m.recetas {
		if receta.MaquetaID == maquetaID {
			recetasFiltradas = append(recetasFiltradas, receta)
		}
	}
	return recetasFiltradas
}

func (m *Memoria) BuscarRecetaPorID(id int) (models.Receta, bool) {
	for _, receta := range m.recetas {
		if receta.ID == id {
			return receta, true
		}
	}
	return models.Receta{}, false
}

func (m *Memoria) ActualizarReceta(id int, datos models.Receta) (models.Receta, bool) {
	for i, receta := range m.recetas {
		if receta.ID == id {
			datos.ID = id
			m.recetas[i] = datos
			return datos, true
		}
	}
	return models.Receta{}, false
}

func (m *Memoria) EliminarReceta(id int) bool {
	for i, receta := range m.recetas {
		if receta.ID == id {
			m.recetas = append(m.recetas[:i], m.recetas[i+1:]...)
			return true
		}
	}
	return false
}
