package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"proyecto/internal/models"
	"proyecto/internal/storage"

	"github.com/go-chi/chi/v5"
)

// Crear asesor
func CrearAsesor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var nuevoAsesor models.Asesor

	if err := json.NewDecoder(r.Body).Decode(&nuevoAsesor); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Datos inválidos",
		})
		return
	}

	if nuevoAsesor.Nombre == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "El nombre es obligatorio",
		})
		return
	}

	asesorCreado := storage.CreateAsesor(nuevoAsesor)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(asesorCreado)
}

// Obtener todos los asesores
func ObtenerAsesores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	asesores := storage.GetAllAsesores()

	json.NewEncoder(w).Encode(asesores)
}

// Obtener asesor por ID
func ObtenerAsesorPorID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "ID inválido",
		})
		return
	}

	asesor, err := storage.GetAsesorByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Asesor no encontrado",
		})
		return
	}

	json.NewEncoder(w).Encode(asesor)
}

// Actualizar asesor
func ActualizarAsesor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "ID inválido",
		})
		return
	}

	var asesorActualizado models.Asesor

	if err := json.NewDecoder(r.Body).Decode(&asesorActualizado); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Datos inválidos",
		})
		return
	}

	if asesorActualizado.Nombre == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "El nombre es obligatorio",
		})
		return
	}

	asesor, err := storage.UpdateAsesor(id, asesorActualizado)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Asesor no encontrado",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"mensaje": "Asesor actualizado correctamente",
		"asesor":  asesor,
	})
}

// Eliminar asesor
func EliminarAsesor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "ID inválido",
		})
		return
	}

	if err := storage.DeleteAsesor(id); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Asesor no encontrado",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"mensaje": "Asesor eliminado correctamente",
	})
}