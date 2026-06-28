package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"proyecto/internal/models"
	"proyecto/internal/service"
)

// registroEntrada se usa solo para decodificar el body del registro,
// porque models.Usuario tiene Password con `json:"-"` (para nunca
// exponerlo en las respuestas), lo cual también haría que se IGNORE
// al decodificar el body de entrada si decodificáramos directo ahí.
type registroEntrada struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Semestre int    `json:"semestre"`
	Telefono string `json:"telefono"`
	Rol      string `json:"rol"`
}

// Registrar atiende POST /api/v1/auth/registro
func (s *Servidor) Registrar(w http.ResponseWriter, r *http.Request) {
	var entrada registroEntrada
	if err := json.NewDecoder(r.Body).Decode(&entrada); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	nuevoUsuario := models.Usuario{
		Nombre:   entrada.Nombre,
		Apellido: entrada.Apellido,
		Email:    entrada.Email,
		Password: entrada.Password,
		Semestre: entrada.Semestre,
		Telefono: entrada.Telefono,
		Rol:      entrada.Rol,
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
