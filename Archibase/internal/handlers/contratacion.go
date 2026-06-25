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

func (s *Servidor) CrearContratacion(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")

	var nuevaContratacion models.Contratacion
	lector := json.NewDecoder(peticion.Body)
	err := lector.Decode(&nuevaContratacion)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}

	if nuevaContratacion.Estudiante == "" {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "El estudiante es obligatorio"})
		return
	}

	nuevaContratacion.IDcontratacion = storage.ConteoContrataciones
	storage.ConteoContrataciones++
	storage.ListaContrataciones = append(storage.ListaContrataciones, nuevaContratacion)

	fmt.Println("--> Contratacion creada con ID:", nuevaContratacion.IDcontratacion)
	respuesta.WriteHeader(http.StatusCreated)
	json.NewEncoder(respuesta).Encode(nuevaContratacion)
}

func (s *Servidor) ObtenerContrataciones(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	fmt.Println("--> Obteniendo todas las contrataciones")
	json.NewEncoder(respuesta).Encode(storage.ListaContrataciones)
}

func (s *Servidor) ObtenerContratacionPorID(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")

	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}

	for _, contratacion := range storage.ListaContrataciones {
		if contratacion.IDcontratacion == id {
			fmt.Println("--> Contratacion encontrada con ID:", id)
			json.NewEncoder(respuesta).Encode(contratacion)
			return
		}
	}

	respuesta.WriteHeader(http.StatusNotFound)
	json.NewEncoder(respuesta).Encode(map[string]string{"error": "Contratacion no encontrada"})
}

func (s *Servidor) ActualizarContratacion(respuesta http.ResponseWriter, peticion *http.Request) {
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

	if contratacionActualizada.Estudiante == "" {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "El estudiante es obligatorio"})
		return
	}

	for i, c := range storage.ListaContrataciones {
		if c.IDcontratacion == id {
			contratacionActualizada.IDcontratacion = id
			storage.ListaContrataciones[i] = contratacionActualizada
			fmt.Println("--> Contratacion actualizada con ID:", id)
			json.NewEncoder(respuesta).Encode(map[string]interface{}{
				"mensaje":      "Contratacion actualizada",
				"contratacion": contratacionActualizada,
			})
			return
		}
	}

	respuesta.WriteHeader(http.StatusNotFound)
	json.NewEncoder(respuesta).Encode(map[string]string{"error": "Contratacion no encontrada"})
}

func (s *Servidor) EliminarContratacion(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")

	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}

	for i, c := range storage.ListaContrataciones {
		if c.IDcontratacion == id {
			storage.ListaContrataciones = append(storage.ListaContrataciones[:i], storage.ListaContrataciones[i+1:]...)
			fmt.Println("--> Contratacion eliminada con ID:", id)
			json.NewEncoder(respuesta).Encode(map[string]string{"mensaje": "Contratacion eliminada correctamente"})
			return
		}
	}

	respuesta.WriteHeader(http.StatusNotFound)
	json.NewEncoder(respuesta).Encode(map[string]string{"error": "Contratacion no encontrada"})
}
