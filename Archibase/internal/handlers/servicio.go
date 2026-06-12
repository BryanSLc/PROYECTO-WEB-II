package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"proyecto/internal/models"
	"proyecto/internal/storage"

	"github.com/go-chi/chi/v5"
)

// Obtener todos los servicios
func ObtenerServicios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	servicios := storage.GetAllServicios()

	json.NewEncoder(w).Encode(servicios)
}

// Obtener servicio por ID
func ObtenerServicioPorID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "ID inválido",
		})
		return
	}

	servicio, err := storage.GetServicioByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Servicio no encontrado",
		})
		return
	}

	json.NewEncoder(w).Encode(servicio)
}

// Crear servicio
func CrearServicio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var nuevoServicio models.Servicio

	if err := json.NewDecoder(r.Body).Decode(&nuevoServicio); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Datos inválidos",
		})
		return
	}

	if nuevoServicio.Titulo == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "El título es obligatorio",
		})
		return
	}

	if nuevoServicio.Descripcion == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "La descripción es obligatoria",
		})
		return
	}

	if nuevoServicio.Disponibilidad == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "La disponibilidad es obligatoria",
		})
		return
	}

	if nuevoServicio.IDasesor == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "El ID del asesor es obligatorio",
		})
		return
	}

	if nuevoServicio.Precio < 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "El precio no puede ser negativo",
		})
		return
	}

	servicioCreado := storage.CreateServicio(nuevoServicio)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(servicioCreado)
}

// Actualizar servicio
func ActualizarServicio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "ID inválido",
		})
		return
	}

	var servicioActualizado models.Servicio

	if err := json.NewDecoder(r.Body).Decode(&servicioActualizado); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Datos inválidos",
		})
		return
	}

	if servicioActualizado.Titulo == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "El título es obligatorio",
		})
		return
	}

	if servicioActualizado.Descripcion == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "La descripción es obligatoria",
		})
		return
	}

	if servicioActualizado.Disponibilidad == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "La disponibilidad es obligatoria",
		})
		return
	}

	if servicioActualizado.IDasesor == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "El ID del asesor es obligatorio",
		})
		return
	}

	if servicioActualizado.Precio < 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "El precio no puede ser negativo",
		})
		return
	}

	servicio, err := storage.UpdateServicio(id, servicioActualizado)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Servicio no encontrado",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"mensaje":  "Servicio actualizado correctamente",
		"servicio": servicio,
	})
}

// Eliminar servicio
func EliminarServicio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "ID inválido",
		})
		return
	}

	if err := storage.DeleteServicio(id); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Servicio no encontrado",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"mensaje": "Servicio eliminado correctamente",
	})
}