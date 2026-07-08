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

func (s *Servidor) CrearUsuario(w http.ResponseWriter, r *http.Request) {
	var nuevoUsuario models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&nuevoUsuario); err != nil {
		RespondError(w, http.StatusBadRequest, "Datos invalidos")
		return
	}

	// Delegamos la validación y creación completa al servicio de negocio
	usuarioCreado, err := s.UsuarioService.Crear(nuevoUsuario)
	if err != nil {
		if errors.Is(err, service.ErrNombreObligatorio) || errors.Is(err, service.ErrEmailObligatorio) {
			RespondError(w, http.StatusBadRequest, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, "Error interno del servidor")
		return
	}

	RespondJSON(w, http.StatusCreated, usuarioCreado)
}

func (s *Servidor) ObtenerUsuarios(w http.ResponseWriter, _ *http.Request) {
	RespondJSON(w, http.StatusOK, s.UsuarioService.Listar())
}

func (s *Servidor) ObtenerUsuarioPorID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID invalido")
		return
	}

	usuario, err := s.UsuarioService.BuscarPorID(id)
	if err != nil {
		if errors.Is(err, service.ErrUsuarioNoEncontrado) {
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, "Error interno del servidor")
		return
	}

	RespondJSON(w, http.StatusOK, usuario)
}

func (s *Servidor) ActualizarUsuario(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID invalido")
		return
	}

	var usuarioActualizado models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&usuarioActualizado); err != nil {
		RespondError(w, http.StatusBadRequest, "Datos invalidos")
		return
	}

	usuario, err := s.UsuarioService.Actualizar(id, usuarioActualizado)
	if err != nil {
		if errors.Is(err, service.ErrUsuarioNoEncontrado) {
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		if errors.Is(err, service.ErrNombreObligatorio) {
			RespondError(w, http.StatusBadRequest, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, "Error interno del servidor")
		return
	}

	RespondJSON(w, http.StatusOK, map[string]any{
		"mensaje": "Usuario actualizado correctamente",
		"usuario": usuario,
	})
}

func (s *Servidor) EliminarUsuario(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID invalido")
		return
	}

	if err := s.UsuarioService.Eliminar(id); err != nil {
		if errors.Is(err, service.ErrUsuarioNoEncontrado) {
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, "Error interno del servidor")
		return
	}

	RespondJSON(w, http.StatusOK, map[string]string{"mensaje": "Usuario eliminado correctamente"})
}
