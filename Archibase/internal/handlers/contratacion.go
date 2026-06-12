package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"proyecto/internal/models"
	"proyecto/internal/storage"

	"github.com/go-chi/chi/v5"
)

// Obtener todas las contrataciones
func ObtenerContrataciones(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	contrataciones := storage.GetAllContrataciones()

	json.NewEncoder(w).Encode(contrataciones)
}

// Obtener contratación por ID
func ObtenerContratacionPorID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "ID inválido",
		})
		return
	}

	contratacion, err := storage.GetContratacionByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Contratación no encontrada",
		})
		return
	}

	json.NewEncoder(w).Encode(contratacion)
}

// Crear contratación
func CrearContratacion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var nuevaContratacion models.Contratacion

	if err := json.NewDecoder(r.Body).Decode(&nuevaContratacion); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Datos inválidos",
		})
		return
	}

	if nuevaContratacion.Estudiante == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "El estudiante es obligatorio",
		})
		return
	}

	if nuevaContratacion.Fecha == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "La fecha es obligatoria",
		})
		return
	}

	if nuevaContratacion.Estado == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "El estado es obligatorio",
		})
		return
	}

	if nuevaContratacion.IDservicio == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "El ID del servicio es obligatorio",
		})
		return
	}

	contratacionCreada := storage.CreateContratacion(nuevaContratacion)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(contratacionCreada)
}

// Actualizar contratación
func ActualizarContratacion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "ID inválido",
		})
		return
	}

	var contratacionActualizada models.Contratacion

	if err := json.NewDecoder(r.Body).Decode(&contratacionActualizada); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Datos inválidos",
		})
		return
	}

	if contratacionActualizada.Estudiante == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "El estudiante es obligatorio",
		})
		return
	}

	if contratacionActualizada.Fecha == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "La fecha es obligatoria",
		})
		return
	}

	if contratacionActualizada.Estado == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "El estado es obligatorio",
		})
		return
	}

	if contratacionActualizada.IDservicio == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "El ID del servicio es obligatorio",
		})
		return
	}

	contratacion, err := storage.UpdateContratacion(id, contratacionActualizada)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Contratación no encontrada",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"mensaje":      "Contratación actualizada correctamente",
		"contratacion": contratacion,
	})
}

// Eliminar contratación
func EliminarContratacion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "ID inválido",
		})
		return
	}

	if err := storage.DeleteContratacion(id); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Contratación no encontrada",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"mensaje": "Contratación eliminada correctamente",
	})
}
