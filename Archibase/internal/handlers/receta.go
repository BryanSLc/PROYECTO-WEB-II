package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"proyecto/internal/models"
)

func (s *Servidor) CrearReceta(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")

	var nueva models.Receta
	lector := json.NewDecoder(peticion.Body)
	if err := lector.Decode(&nueva); err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}

	recetaCreada, svcErr := s.RecetaSvc.Crear(nueva)
	if svcErr != nil {
		responderError(respuesta, svcErr)
		return
	}

	fmt.Println("--> Receta creada con ID:", recetaCreada.ID)
	respuesta.WriteHeader(http.StatusCreated)
	json.NewEncoder(respuesta).Encode(recetaCreada)
}

func (s *Servidor) ObtenerRecetas(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	fmt.Println("--> Obteniendo todas las recetas")
	json.NewEncoder(respuesta).Encode(s.RecetaSvc.Listar())
}

func (s *Servidor) ObtenerRecetaPorID(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")

	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}

	receta, svcErr := s.RecetaSvc.BuscarPorID(id)
	if svcErr != nil {
		responderError(respuesta, svcErr)
		return
	}

	fmt.Println("--> Receta encontrada con ID:", id)
	json.NewEncoder(respuesta).Encode(receta)
}

func (s *Servidor) ActualizarReceta(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")

	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}

	var recetaActualizada models.Receta
	lector := json.NewDecoder(peticion.Body)
	if err := lector.Decode(&recetaActualizada); err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}

	receta, svcErr := s.RecetaSvc.Actualizar(id, recetaActualizada)
	if svcErr != nil {
		responderError(respuesta, svcErr)
		return
	}

	fmt.Println("--> Receta actualizada con ID:", id)
	json.NewEncoder(respuesta).Encode(map[string]interface{}{
		"mensaje": "Receta actualizada",
		"receta":  receta,
	})
}

func (s *Servidor) EliminarReceta(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")

	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}

	if svcErr := s.RecetaSvc.Eliminar(id); svcErr != nil {
		responderError(respuesta, svcErr)
		return
	}

	fmt.Println("--> Receta eliminada con ID:", id)
	json.NewEncoder(respuesta).Encode(map[string]string{"mensaje": "Receta eliminada correctamente"})
}
