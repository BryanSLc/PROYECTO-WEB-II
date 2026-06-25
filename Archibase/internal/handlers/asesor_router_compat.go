package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Compatibilidad: el router en cmd/api/main.go usa varios nombres incorrectos.
// Estas funciones puente permiten compilar sin rehacer el router.

func (s *Servidor) GetAllAsesores(w http.ResponseWriter, r *http.Request) {
	s.ObtenerAsesores(w, r)
}

func (s *Servidor) CreateAsesor(w http.ResponseWriter, r *http.Request) {
	s.CrearAsesor(w, r)
}

func (s *Servidor) GetAsesorByID(w http.ResponseWriter, r *http.Request) {
	s.ObtenerAsesorPorID(w, r)
}

func (s *Servidor) UpdateAsesor(w http.ResponseWriter, r *http.Request) {
	s.ActualizarAsesor(w, r)
}

func (s *Servidor) DeleteAsesor(w http.ResponseWriter, r *http.Request) {
	s.EliminarAsesor(w, r)
}

// Versiones sin receiver (para el router con handlers.UpdateAsesor / DeleteAsesor)
func UpdateAsesor(w http.ResponseWriter, r *http.Request) {
	_ = chi.URLParam(r, "id")
	w.WriteHeader(http.StatusNotImplemented)
}
func DeleteAsesor(w http.ResponseWriter, r *http.Request) {
	_ = chi.URLParam(r, "id")
	w.WriteHeader(http.StatusNotImplemented)
}
