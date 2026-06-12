package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"proyecto/internal/models"
	"proyecto/internal/storage"
)

func CrearUbicacion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var nuevaUbicacion models.Ubicacion
	err := json.NewDecoder(r.Body).Decode(&nuevaUbicacion)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if nuevaUbicacion.Provincia == "" || nuevaUbicacion.Ciudad == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	nuevaUbicacion.ID = storage.ConteoUbicaciones
	storage.ConteoUbicaciones++
	storage.ListaUbicaciones =
		append(storage.ListaUbicaciones, nuevaUbicacion)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevaUbicacion)
}

func ObtenerUbicaciones(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storage.ListaUbicaciones)
}

func ObtenerUbicacionPorID(w http.ResponseWriter, r *http.Request) {
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for _, ubicacion := range storage.ListaUbicaciones {
		if ubicacion.ID == id {
			json.NewEncoder(w).Encode(ubicacion)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func ActualizarUbicacion(w http.ResponseWriter, r *http.Request) {
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var ubicacionActualizada models.Ubicacion
	err = json.NewDecoder(r.Body).Decode(&ubicacionActualizada)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for i, ubicacion := range storage.ListaUbicaciones {
		if ubicacion.ID == id {
			ubicacionActualizada.ID = id
			storage.ListaUbicaciones[i] = ubicacionActualizada
			json.NewEncoder(w).Encode(ubicacionActualizada)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func EliminarUbicacion(w http.ResponseWriter, r *http.Request) {
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for i, ubicacion := range storage.ListaUbicaciones {
		if ubicacion.ID == id {
			storage.ListaUbicaciones =
				append(
					storage.ListaUbicaciones[:i],
					storage.ListaUbicaciones[i+1:]...,
				)
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
