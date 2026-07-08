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

// CrearAsesor crea un asesor usando el servicio inyectado.
func (s *Servidor) CrearAsesor(w http.ResponseWriter, r *http.Request) {
	var a models.Asesor
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		RespondError(w, http.StatusBadRequest, "body invalido")
		return
	}

	asesor, err := s.AsesorService.Crear(a)
	if err != nil {
		if errors.Is(err, service.ErrNombreAsesorObligatorio) {
			RespondError(w, http.StatusBadRequest, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, "Error interno del servidor")
		return
	}

	RespondJSON(w, http.StatusCreated, asesor)
}

func (s *Servidor) ObtenerAsesores(w http.ResponseWriter, _ *http.Request) {
	RespondJSON(w, http.StatusOK, s.AsesorService.Listar())
}

func (s *Servidor) ObtenerAsesorPorID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID invalido")
		return
	}

	asesor, err := s.AsesorService.BuscarPorID(id)
	if err != nil {
		if errors.Is(err, service.ErrAsesorNoEncontrado) {
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, "Error interno del servidor")
		return
	}

	RespondJSON(w, http.StatusOK, asesor)
}
