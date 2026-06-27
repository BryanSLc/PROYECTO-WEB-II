package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"proyecto/internal/models"
	"proyecto/internal/service"
)

// Registrar atiende POST /api/v1/auth/registro
func (s *Servidor) Registrar(w http.ResponseWriter, r *http.Request) {
	var nuevoUsuario models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&nuevoUsuario); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	usuarioCreado, err := s.AuthService.Registrar(nuevoUsuario)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNombreObligatorio),
			errors.Is(err, service.ErrEmailObligatorio),
			errors.Is(err, service.ErrEmailEnUso),
			errors.Is(err, service.ErrCredencialesInvalidas):
			RespondError(w, http.StatusBadRequest, err.Error())
			return
		default:
			RespondError(w, http.StatusInternalServerError, "Error interno del servidor")
			return
		}
	}

	RespondJSON(w, http.StatusCreated, usuarioCreado)
}

// Login atiende POST /api/v1/auth/login
func (s *Servidor) Login(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	token, err := s.AuthService.Login(creds.Email, creds.Password)
	if err != nil {
		RespondError(w, http.StatusUnauthorized, err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, map[string]string{"token": token})
}
