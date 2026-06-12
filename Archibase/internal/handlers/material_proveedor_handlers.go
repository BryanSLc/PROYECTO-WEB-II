package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"proyecto/internal/models"
	"proyecto/internal/storage"
)

func CrearMaterial(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var nuevoMaterial models.MaterialProveedor
	err := json.NewDecoder(r.Body).Decode(&nuevoMaterial)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if nuevoMaterial.Nombre == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	nuevoMaterial.ID = storage.ConteoMateriales
	storage.ConteoMateriales++
	storage.ListaMateriales =
		append(storage.ListaMateriales, nuevoMaterial)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevoMaterial)
}

func ObtenerMateriales(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storage.ListaMateriales)
}

func ObtenerMaterialPorID(w http.ResponseWriter, r *http.Request) {
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for _, material := range storage.ListaMateriales {
		if material.ID == id {
			json.NewEncoder(w).Encode(material)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func ActualizarMaterial(w http.ResponseWriter, r *http.Request) {
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var materialActualizado models.MaterialProveedor
	err = json.NewDecoder(r.Body).Decode(&materialActualizado)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for i, material := range storage.ListaMateriales {
		if material.ID == id {
			materialActualizado.ID = id
			storage.ListaMateriales[i] = materialActualizado
			json.NewEncoder(w).Encode(materialActualizado)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func EliminarMaterial(w http.ResponseWriter, r *http.Request) {
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for i, material := range storage.ListaMateriales {
		if material.ID == id {
			storage.ListaMateriales =
				append(
					storage.ListaMateriales[:i],
					storage.ListaMateriales[i+1:]...,
				)
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
