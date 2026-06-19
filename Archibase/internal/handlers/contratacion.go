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

func CrearContratacion(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	var nuevaContratacion models.Contratacion
	lector := json.NewDecoder(peticion.Body)
	err := lector.Decode(&nuevaContratacion)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}
	if nuevaContratacion.Detalle == "" { // Nota: Cambié 'Titulo' por 'Detalle'. Ajusta según tu struct.
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "El detalle es obligatorio"})
		return
	}
	nuevaContratacion.ID = storage.ConteoContrataciones
	storage.ConteoContrataciones++
	storage.ListaContrataciones = append(storage.ListaContrataciones, nuevaContratacion)
	fmt.Println("--> Contratación creada con ID:", nuevaContratacion.ID)
	respuesta.WriteHeader(http.StatusCreated)
	json.NewEncoder(respuesta).Encode(nuevaContratacion)
}

func ObtenerContrataciones(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	fmt.Println("--> Obteniendo todas las contrataciones")
	json.NewEncoder(respuesta).Encode(storage.ListaContrataciones)
}

func ObtenerContratacionPorID(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}
	for _, contratacion := range storage.ListaContrataciones {
		if contratacion.ID == id {
			fmt.Println("--> Contratación encontrada con ID:", id)
			json.NewEncoder(respuesta).Encode(contratacion)
			return
		}
	}
	respuesta.WriteHeader(http.StatusNotFound) // error 404
	json.NewEncoder(respuesta).Encode(map[string]string{"error": "Contratación no encontrada"})
}

func ActualizarContratacion(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}
	var contratacionActualizada models.Contratacion
	lector := json.NewDecoder(peticion.Body)
	err = lector.Decode(&contratacionActualizada)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}
	if contratacionActualizada.Detalle == "" {
		respuesta.WriteHeader(http.StatusBadRequest) // 400 error
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "El detalle es obligatorio"})
		return
	}
	for i, contratacion := range storage.ListaContrataciones {
		if contratacion.ID == id {
			contratacionActualizada.ID = id
			storage.ListaContrataciones[i] = contratacionActualizada
			fmt.Println("--> Contratación actualizada con ID:", id)
			json.NewEncoder(respuesta).Encode(map[string]interface{}{
				"mensaje":      "Contratación actualizada",
				"contratacion": contratacionActualizada,
			})
			return
		}
	}
	respuesta.WriteHeader(http.StatusNotFound)
	json.NewEncoder(respuesta).Encode(map[string]string{"error": "Contratación no encontrada"})
}

func EliminarContratacion(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}
	for i, contratacion := range storage.ListaContrataciones {
		if contratacion.ID == id {
			storage.ListaContrataciones = append(storage.ListaContrataciones[:i], storage.ListaContrataciones[i+1:]...)

			fmt.Println("--> Contratación eliminada con ID:", id)
			json.NewEncoder(respuesta).Encode(map[string]string{"mensaje": "Contratación eliminada correctamente"})
			return
		}
	}
	respuesta.WriteHeader(http.StatusNotFound)
	json.NewEncoder(respuesta).Encode(map[string]string{"error": "Contratación no encontrada"})
}
