package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"proyecto/internal/models"
	"proyecto/internal/storage"
)

func GetAllAsesores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("--> Obteniendo todos los asesores")
	json.NewEncoder(w).Encode(storage.ListaAsesores)
}

func CreateAsesor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var nuevoAsesor models.Asesor
	if err := json.NewDecoder(r.Body).Decode(&nuevoAsesor); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}

	if nuevoAsesor.Nombre == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "El nombre es obligatorio"})
		return
	}

	nuevoAsesor.IDasesor = storage.ConteoAsesores
	storage.ConteoAsesores++

	storage.ListaAsesores = append(storage.ListaAsesores, nuevoAsesor)
	fmt.Println("--> Asesor creado con ID:", nuevoAsesor.IDasesor)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevoAsesor)
}

func GetAsesorByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID invalido"})
		return
	}

	for _, asesor := range storage.ListaAsesores {
		if asesor.IDasesor == id {
			fmt.Println("--> Asesor encontrado con ID:", id)
			json.NewEncoder(w).Encode(asesor)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Asesor no encontrado"})
}

func UpdateAsesor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID invalido"})
		return
	}

	var asesorActualizado models.Asesor
	if err := json.NewDecoder(r.Body).Decode(&asesorActualizado); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}

	if asesorActualizado.Nombre == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "El nombre es obligatorio"})
		return
	}

	for i, asesor := range storage.ListaAsesores {
		if asesor.IDasesor == id {
			asesorActualizado.IDasesor = id
			storage.ListaAsesores[i] = asesorActualizado
			fmt.Println("--> Asesor actualizado con ID:", id)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"mensaje": "Asesor actualizado",
				"asesor":  asesorActualizado,
			})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Asesor no encontrado"})
}

func DeleteAsesor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID invalido"})
		return
	}

	for i, asesor := range storage.ListaAsesores {
		if asesor.IDasesor == id {
			storage.ListaAsesores = append(storage.ListaAsesores[:i], storage.ListaAsesores[i+1:]...)
			fmt.Println("--> Asesor eliminado con ID:", id)
			json.NewEncoder(w).Encode(map[string]string{"mensaje": "Asesor eliminado correctamente"})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Asesor no encontrado"})
}
