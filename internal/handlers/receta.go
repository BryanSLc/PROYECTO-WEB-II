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
	var nuevaReceta models.Receta
	lector := json.NewDecoder(peticion.Body)
	err := lector.Decode(&nuevaReceta)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}
	if nuevaReceta.Titulo == "" {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "El titulo es obligatorio"})
		return
	}
	recetaCreada := s.Almacen.CrearReceta(nuevaReceta)
	fmt.Println("--> Receta creada con ID:", recetaCreada.ID)
	respuesta.WriteHeader(http.StatusCreated)
	json.NewEncoder(respuesta).Encode(recetaCreada)
}

func (s *Servidor) ObtenerRecetas(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	maquetaIDTexto := peticion.URL.Query().Get("maqueta_id")
	if maquetaIDTexto == "" {
		fmt.Println("--> Obteniendo todas las recetas")
		json.NewEncoder(respuesta).Encode(s.Almacen.ListarRecetas())
		return
	}
	maquetaID, err := strconv.Atoi(maquetaIDTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "maqueta_id invalido"})
		return
	}
	fmt.Println("--> Obteniendo recetas de la maqueta con ID:", maquetaID)
	json.NewEncoder(respuesta).Encode(s.Almacen.ListarRecetasPorMaqueta(maquetaID))
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
	receta, encontrada := s.Almacen.BuscarRecetaPorID(id)
	if !encontrada {
		respuesta.WriteHeader(http.StatusNotFound)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Receta no encontrada"})
		return
	}
	fmt.Println("--> Receta encontrado con ID:", id)
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
	err = lector.Decode(&recetaActualizada)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}
	if recetaActualizada.Titulo == "" {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "El titulo es obligatorio"})
		return
	}
	receta, encontrada := s.Almacen.ActualizarReceta(id, recetaActualizada)
	if !encontrada {
		respuesta.WriteHeader(http.StatusNotFound)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Receta no encontrada"})
		return
	}
	fmt.Println("--> Receta actualizada con ID:", id)
	json.NewEncoder(respuesta).Encode(map[string]interface{}{
		"mensaje": "Receta actualizada correctamente",
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
	if !s.Almacen.EliminarReceta(id) {
		respuesta.WriteHeader(http.StatusNotFound)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Receta no encontrada"})
		return
	}
	fmt.Println("--> Receta eliminada con ID:", id)
	json.NewEncoder(respuesta).Encode(map[string]string{"mensaje": "Receta eliminada correctamente"})
}
