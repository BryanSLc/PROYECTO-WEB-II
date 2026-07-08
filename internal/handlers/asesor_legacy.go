package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"proyecto/internal/models"
	"proyecto/internal/service"

	"github.com/go-chi/chi/v5"
)

// En caso de que el servidor principal aún use handlers como funciones globales,
// proveemos un set legacy que permite compilar.
// Estos wrappers solo funcionan si existe una instancia *Servidor global, por lo que
// se recomienda enrutar usando los métodos de *Servidor.

// NOTE: No se usa en tests actuales.

var legacyServidor *Servidor

func SetLegacyServidor(s *Servidor) { legacyServidor = s }

func GetAllAsesores(w http.ResponseWriter, r *http.Request) {
	if legacyServidor == nil {
		RespondError(w, http.StatusInternalServerError, "servidor no inicializado")
		return
	}
	legacyServidor.ObtenerAsesores(w, r)
}

func GetAsesorByID(w http.ResponseWriter, r *http.Request) {
	if legacyServidor == nil {
		RespondError(w, http.StatusInternalServerError, "servidor no inicializado")
		return
	}
	legacyServidor.ObtenerAsesorPorID(w, r)
}

func CreateAsesor(w http.ResponseWriter, r *http.Request) {
	if legacyServidor == nil {
		RespondError(w, http.StatusInternalServerError, "servidor no inicializado")
		return
	}
	legacyServidor.CrearAsesor(w, r)
}

func UpdateAsesor(w http.ResponseWriter, r *http.Request) {
	// Handler legacy no implementado
	_ = json.NewDecoder(r.Body)
	RespondError(w, http.StatusNotImplemented, "no implementado")
}

func DeleteAsesor(w http.ResponseWriter, r *http.Request) {
	// Handler legacy no implementado
	_, _ = strconv.Atoi(chi.URLParam(r, "id"))
	_ = service.ErrAsesorNoEncontrado

	RespondError(w, http.StatusNotImplemented, "no implementado")
}

// compile-time check
var _ = models.Asesor{}
