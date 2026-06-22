package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"proyecto/internal/models"
	"proyecto/internal/storage"

	"github.com/go-chi/chi/v5"
)

func GetAllAsesores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("--> Obteniendo todos los asesores")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(storage.GetAllAsesores())
}

func GetAsesorByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("--> ID invalido")
		http.Error(w, `{"error":"ID invalido"}`, http.StatusBadRequest)
		return
	}
	asesor, err := storage.GetAsesorByID(id)
	if err != nil {
		fmt.Println("--> Asesor no encontrado con ID:", id)
		http.Error(w, `{"error":"asesor no encontrado"}`, http.StatusNotFound)
		return
	}
	fmt.Println("--> Asesor encontrado con ID:", id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(asesor)
}

func CreateAsesor(w http.ResponseWriter, r *http.Request) {
	var a models.Asesor
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		fmt.Println("--> Body invalido")
		http.Error(w, `{"error":"body invalido"}`, http.StatusBadRequest)
		return
	}
	if a.Nombre == "" || a.Especialidad == "" || a.Contacto == "" || a.Modalidad == "" {
		fmt.Println("--> Campos requeridos faltantes")
		http.Error(w, `{"error":"nombre, especialidad, contacto y modalidad son requeridos"}`, http.StatusBadRequest)
		return
	}
	fmt.Println("--> Asesor creado")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(storage.CreateAsesor(a))
}

func UpdateAsesor(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("--> ID invalido")
		http.Error(w, `{"error":"ID invalido"}`, http.StatusBadRequest)
		return
	}
	var a models.Asesor
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		fmt.Println("--> Body invalido")
		http.Error(w, `{"error":"body invalido"}`, http.StatusBadRequest)
		return
	}
	if a.Nombre == "" || a.Especialidad == "" || a.Contacto == "" || a.Modalidad == "" {
		fmt.Println("--> Campos requeridos faltantes")
		http.Error(w, `{"error":"nombre, especialidad, contacto y modalidad son requeridos"}`, http.StatusBadRequest)
		return
	}
	actualizado, err := storage.UpdateAsesor(id, a)
	if err != nil {
		fmt.Println("--> Asesor no encontrado con ID:", id)
		http.Error(w, `{"error":"asesor no encontrado"}`, http.StatusNotFound)
		return
	}
	fmt.Println("--> Asesor actualizado con ID:", id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(actualizado)
}

func DeleteAsesor(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("--> ID invalido")
		http.Error(w, `{"error":"ID invalido"}`, http.StatusBadRequest)
		return
	}
	if err := storage.DeleteAsesor(id); err != nil {
		fmt.Println("--> Asesor no encontrado con ID:", id)
		http.Error(w, `{"error":"asesor no encontrado"}`, http.StatusNotFound)
		return
	}
	fmt.Println("--> Asesor eliminado con ID:", id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "asesor eliminado"})
}