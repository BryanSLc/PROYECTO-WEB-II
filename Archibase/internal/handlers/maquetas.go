package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"proyecto/internal/models"
	"proyecto/internal/service"

	"github.com/go-chi/chi/v5"
)

func (s *Servidor) CrearMaqueta(w http.ResponseWriter, r *http.Request) {
	var nuevaMaqueta models.Maqueta
	if err := json.NewDecoder(r.Body).Decode(&nuevaMaqueta); err != nil {
		RespondError(w, http.StatusBadRequest, "Datos invalidos")
		return
	}

	maquetaCreada, err := s.MaquetaService.Crear(nuevaMaqueta)
	if err != nil {
		if errors.Is(err, service.ErrTituloObligatorio) {
			RespondError(w, http.StatusBadRequest, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, "Error interno del servidor")
		return
	}

	RespondJSON(w, http.StatusCreated, maquetaCreada)
}

func (s *Servidor) ObtenerMaquetas(w http.ResponseWriter, _ *http.Request) {
	RespondJSON(w, http.StatusOK, s.MaquetaService.Listar())
}

func (s *Servidor) ObtenerMaquetaPorID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID invalido")
		return
	}

	maqueta, err := s.MaquetaService.BuscarPorID(id)
	if err != nil {
		if errors.Is(err, service.ErrMaquetaNoEncontrada) {
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, "Error interno del servidor")
		return
	}

	RespondJSON(w, http.StatusOK, maqueta)
}

func (s *Servidor) ActualizarMaqueta(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID invalido")
		return
	}

	var maquetaActualizada models.Maqueta
	if err := json.NewDecoder(r.Body).Decode(&maquetaActualizada); err != nil {
		RespondError(w, http.StatusBadRequest, "Datos invalidos")
		return
	}

	maqueta, err := s.MaquetaService.Actualizar(id, maquetaActualizada)
	if err != nil {
		if errors.Is(err, service.ErrMaquetaNoEncontrada) {
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		if errors.Is(err, service.ErrTituloObligatorio) {
			RespondError(w, http.StatusBadRequest, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, "Error interno del servidor")
		return
	}

	RespondJSON(w, http.StatusOK, map[string]any{
		"mensaje": "Maqueta actualizada correctamente",
		"maqueta": maqueta,
	})
}

func (s *Servidor) EliminarMaqueta(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID invalido")
		return
	}

	if err := s.MaquetaService.Eliminar(id); err != nil {
		if errors.Is(err, service.ErrMaquetaNoEncontrada) {
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, "Error interno del servidor")
		return
	}

	RespondJSON(w, http.StatusOK, map[string]string{"mensaje": "Maqueta eliminada correctamente"})
}

func (s *Servidor) AgregarEvolucionMaqueta(w http.ResponseWriter, r *http.Request) {
	var nuevaEvolucion models.EvolucionMaqueta
	if err := json.NewDecoder(r.Body).Decode(&nuevaEvolucion); err != nil {
		RespondError(w, http.StatusBadRequest, "Datos invalidos")
		return
	}

	evolucion, err := s.MaquetaService.AgregarEvolucion(nuevaEvolucion)
	if err != nil {
		if errors.Is(err, service.ErrMaquetaNoEncontrada) {
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		if errors.Is(err, service.ErrIDMaquetaObligatorio) || errors.Is(err, service.ErrTituloAvanceObligatorio) || errors.Is(err, service.ErrPasoInvalido) {
			RespondError(w, http.StatusBadRequest, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, "Error interno del servidor")
		return
	}

	RespondJSON(w, http.StatusCreated, evolucion)
}

func (s *Servidor) ObtenerEvolucionPorMaqueta(w http.ResponseWriter, r *http.Request) {
	maquetaID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID de maqueta invalido")
		return
	}

	historial, err := s.MaquetaService.ListarEvolucion(maquetaID)
	if err != nil {
		// Asumiendo que tu método del service lanza este error si la maqueta padre no existe
		if errors.Is(err, service.ErrMaquetaNoEncontrada) {
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, "Error interno del servidor")
		return
	}

	RespondJSON(w, http.StatusOK, historial)
}

func (s *Servidor) EliminarEvolucion(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID invalido")
		return
	}

	if err := s.MaquetaService.EliminarEvolucion(id); err != nil {
		if errors.Is(err, service.ErrEvolucionNoEncontrada) {
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, "Error interno del servidor")
		return
	}

	RespondJSON(w, http.StatusOK, map[string]string{"mensaje": "Evolucion eliminada correctamente"})
}
