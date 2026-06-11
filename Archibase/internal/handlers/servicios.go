package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"proyecto/internal/models"
	"proyecto/internal/storage"

	"github.com/go-chi/chi/v5"
)

func GetAllAsesores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	json.NewEncoder(w).Encode(storage.GetAllAsesores())
}

func GetAsesorByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, `{"error":"ID invalido"}`, http.StatusBadRequest) // 400
		return
	}
	asesor, err := storage.GetAsesorByID(id)
	if err != nil {
		http.Error(w, `{"error":"asesor no encontrado"}`, http.StatusNotFound) // 404
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	json.NewEncoder(w).Encode(asesor)
}

func CreateAsesor(w http.ResponseWriter, r *http.Request) {
	var a models.Asesor
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, `{"error":"body invalido"}`, http.StatusBadRequest) // 400
		return
	}
	if a.Nombre == "" || a.Especialidad == "" || a.Contacto == "" || a.Modalidad == "" {
		http.Error(w, `{"error":"nombre, especialidad, contacto y modalidad son requeridos"}`, http.StatusBadRequest) // 400
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	json.NewEncoder(w).Encode(storage.CreateAsesor(a))
}

func UpdateAsesor(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, `{"error":"ID invalido"}`, http.StatusBadRequest) // 400
		return
	}
	var a models.Asesor
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, `{"error":"body invalido"}`, http.StatusBadRequest) // 400
		return
	}
	if a.Nombre == "" || a.Especialidad == "" || a.Contacto == "" || a.Modalidad == "" {
		http.Error(w, `{"error":"nombre, especialidad, contacto y modalidad son requeridos"}`, http.StatusBadRequest) // 400
		return
	}
	actualizado, err := storage.UpdateAsesor(id, a)
	if err != nil {
		http.Error(w, `{"error":"asesor no encontrado"}`, http.StatusNotFound) // 404
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	json.NewEncoder(w).Encode(actualizado)
}

func DeleteAsesor(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, `{"error":"ID invalido"}`, http.StatusBadRequest) // 400
		return
	}
	if err := storage.DeleteAsesor(id); err != nil {
		http.Error(w, `{"error":"asesor no encontrado"}`, http.StatusNotFound) // 404
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "asesor eliminado"})
}
