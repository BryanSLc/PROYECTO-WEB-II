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

func CrearAsesor(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	var nuevoAsesor models.Asesor
	lector := json.NewDecoder(peticion.Body)
	err := lector.Decode(&nuevoAsesor)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}
	if nuevoAsesor.Nombre == "" { // Asumiendo que usas 'Nombre' en vez de 'Titulo' para un asesor
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "El nombre es obligatorio"})
		return
	}
	nuevoAsesor.ID = storage.ConteoAsesores
	storage.ConteoAsesores++
	storage.ListaAsesores = append(storage.ListaAsesores, nuevoAsesor)
	fmt.Println("--> Asesor creado con ID:", nuevoAsesor.ID)
	respuesta.WriteHeader(http.StatusCreated)
	json.NewEncoder(respuesta).Encode(nuevoAsesor)
}

func ObtenerAsesores(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	fmt.Println("--> Obteniendo todos los asesores")
	json.NewEncoder(respuesta).Encode(storage.ListaAsesores)
}

func ObtenerAsesorPorID(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}
	for _, asesor := range storage.ListaAsesores {
		if asesor.ID == id {
			fmt.Println("--> Asesor encontrado con ID:", id)
			json.NewEncoder(respuesta).Encode(asesor)
			return
		}
	}
	respuesta.WriteHeader(http.StatusNotFound) //error 404
	json.NewEncoder(respuesta).Encode(map[string]string{"error": "Asesor no encontrado"})
}

func ActualizarAsesor(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}
	var asesorActualizado models.Asesor
	lector := json.NewDecoder(peticion.Body)
	err = lector.Decode(&asesorActualizado)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}
	if asesorActualizado.Nombre == "" {
		respuesta.WriteHeader(http.StatusBadRequest) //400 error
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "El nombre es obligatorio"})
		return
	}
	for i, asesor := range storage.ListaAsesores {
		if asesor.ID == id {
			asesorActualizado.ID = id
			storage.ListaAsesores[i] = asesorActualizado
			fmt.Println("--> Asesor actualizado con ID:", id)
			json.NewEncoder(respuesta).Encode(map[string]interface{}{
				"mensaje": "Asesor actualizado",
				"asesor":  asesorActualizado,
			})
			return
		}
	}
	respuesta.WriteHeader(http.StatusNotFound)
	json.NewEncoder(respuesta).Encode(map[string]string{"error": "Asesor no encontrado"})
}

func EliminarAsesor(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}
	for i, asesor := range storage.ListaAsesores {
		if asesor.ID == id {
			storage.ListaAsesores = append(storage.ListaAsesores[:i], storage.ListaAsesores[i+1:]...)

			fmt.Println("--> Asesor eliminado con ID:", id)
			json.NewEncoder(respuesta).Encode(map[string]string{"mensaje": "Asesor eliminado correctamente"})
			return
		}
	}
	respuesta.WriteHeader(http.StatusNotFound)
	json.NewEncoder(respuesta).Encode(map[string]string{"error": "Asesor no encontrado"})
}
