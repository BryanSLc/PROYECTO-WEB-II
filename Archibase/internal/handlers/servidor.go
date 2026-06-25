package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"proyecto/internal/service"
	"proyecto/internal/storage"
)

type Servidor struct {
	UsuarioSvc      *service.UsuarioService
	MaquetaSvc      *service.MaquetaService
	RecetaSvc       *service.RecetaService
	ContratacionSvc *service.ContratacionService
	AsesorSvc       *service.AsesorService
}

func NuevoServidor(a *storage.SQLiteStorage) *Servidor {
	return &Servidor{
		UsuarioSvc:      service.NuevoUsuarioService(a),
		MaquetaSvc:      service.NuevoMaquetaService(a),
		RecetaSvc:       service.NuevoRecetaService(a),
		ContratacionSvc: service.NuevoContratacionService(a),
		AsesorSvc:       service.NuevoAsesorService(a),
	}
}

// responderError es una función de ayuda para limpiar aún más los handlers
func responderError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	switch {
	case errors.Is(err, service.ErrDatosInvalidos),
		errors.Is(err, service.ErrNombreObligatorio),
		errors.Is(err, service.ErrEmailObligatorio),
		errors.Is(err, service.ErrTituloObligatorio),
		errors.Is(err, service.ErrIDMaquetaObligatorio),
		errors.Is(err, service.ErrTituloAvanceObligatorio),
		errors.Is(err, service.ErrPasoInvalido):
		w.WriteHeader(http.StatusBadRequest)

	case errors.Is(err, service.ErrUsuarioNoEncontrado),
		errors.Is(err, service.ErrMaquetaNoEncontrada),
		errors.Is(err, service.ErrEvolucionNoEncontrada),
		errors.Is(err, service.ErrRecetaNoEncontrada):
		w.WriteHeader(http.StatusNotFound)

	default:
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "error interno del servidor"})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}
