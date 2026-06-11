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

func CrearMaqueta(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	var nuevaMaqueta models.Maqueta
	lector := json.NewDecoder(peticion.Body)
	err := lector.Decode(&nuevaMaqueta)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}
	if nuevaMaqueta.Titulo == "" {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "El titulo es obligatorio"})
		return
	}
	nuevaMaqueta.ID = storage.ConteoMaquetas
	storage.ConteoMaquetas++
	storage.ListaMaquetas = append(storage.ListaMaquetas, nuevaMaqueta)
	fmt.Println("--> Maqueta creada con ID:", nuevaMaqueta.ID)
	respuesta.WriteHeader(http.StatusCreated)
	json.NewEncoder(respuesta).Encode(nuevaMaqueta)
}

func ObtenerMaquetas(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	fmt.Println("--> Obteniendo todas las maquetas")
	json.NewEncoder(respuesta).Encode(storage.ListaMaquetas)
}

func ObtenerMaquetaPorID(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}
	for _, maqueta := range storage.ListaMaquetas {
		if maqueta.ID == id {
			fmt.Println("--> Maqueta encontrada con ID:", id)
			json.NewEncoder(respuesta).Encode(maqueta)
			return
		}
	}
	respuesta.WriteHeader(http.StatusNotFound)
	json.NewEncoder(respuesta).Encode(map[string]string{"error": "Maqueta no encontrada"})
}

func ActualizarMaqueta(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}
	var maquetaActualizada models.Maqueta
	lector := json.NewDecoder(peticion.Body)
	err = lector.Decode(&maquetaActualizada)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}
	if maquetaActualizada.Titulo == "" {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "El titulo es obligatorio"})
		return
	}
	for i, maqueta := range storage.ListaMaquetas {
		if maqueta.ID == id {
			maquetaActualizada.ID = id
			storage.ListaMaquetas[i] = maquetaActualizada
			fmt.Println("--> Maqueta actualizada con ID:", id)
			json.NewEncoder(respuesta).Encode(map[string]interface{}{
				"mensaje": "Maqueta actualizada",
				"maqueta": maquetaActualizada,
			})
			return
		}
	}
	respuesta.WriteHeader(http.StatusNotFound)
	json.NewEncoder(respuesta).Encode(map[string]string{"error": "Maqueta no encontrada"})
}

func EliminarMaqueta(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}
	for i, maqueta := range storage.ListaMaquetas {
		if maqueta.ID == id {
			storage.ListaMaquetas = append(storage.ListaMaquetas[:i], storage.ListaMaquetas[i+1:]...)

			fmt.Println("--> Maqueta eliminada con ID:", id)
			json.NewEncoder(respuesta).Encode(map[string]string{"mensaje": "Maqueta eliminada correctamente"})
			return
		}
	}
	respuesta.WriteHeader(http.StatusNotFound)
	json.NewEncoder(respuesta).Encode(map[string]string{"error": "Maqueta no encontrada"})
}
