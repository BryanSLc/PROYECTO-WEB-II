package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"proyecto/internal/models"
)

func (s *Servidor) CrearUbicacion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var nuevaUbicacion models.Ubicacion
	err := json.NewDecoder(r.Body).Decode(&nuevaUbicacion)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}
	if nuevaUbicacion.Provincia == "" || nuevaUbicacion.Ciudad == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Provincia y Ciudad son obligatorias"})
		return
	}

	ubicacionCreada := s.Almacen.CrearUbicacion(nuevaUbicacion)
	fmt.Println("--> Ubicacion creada con ID:", ubicacionCreada.ID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ubicacionCreada)
}

func (s *Servidor) ObtenerUbicaciones(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("--> Obteniendo todas las ubicaciones")
	ubicaciones := s.Almacen.ListarUbicaciones()
	json.NewEncoder(w).Encode(ubicaciones)
}

func (s *Servidor) ObtenerUbicacionPorID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID invalido"})
		return
	}

	ubicacion, encontrado := s.Almacen.BuscarUbicacionPorID(id)
	if !encontrado {
		fmt.Println("--> Ubicacion no encontrada con ID:", id)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Ubicacion no encontrada"})
		return
	}

	fmt.Println("--> Ubicacion encontrada con ID:", id)
	json.NewEncoder(w).Encode(ubicacion)
}

func (s *Servidor) ActualizarUbicacion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID invalido"})
		return
	}

	var ubicacionActualizada models.Ubicacion
	err = json.NewDecoder(r.Body).Decode(&ubicacionActualizada)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}

	ubicacion, encontrado := s.Almacen.ActualizarUbicacion(id, ubicacionActualizada)
	if !encontrado {
		fmt.Println("--> Ubicacion no encontrada con ID:", id)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Ubicacion no encontrada"})
		return
	}

	fmt.Println("--> Ubicacion actualizada con ID:", id)
	json.NewEncoder(w).Encode(ubicacion)
}

func (s *Servidor) EliminarUbicacion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID invalido"})
		return
	}

	exito := s.Almacen.EliminarUbicacion(id)
	if !exito {
		fmt.Println("--> Ubicacion no encontrada con ID:", id)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Ubicacion no encontrada"})
		return
	}

	fmt.Println("--> Ubicacion eliminada con ID:", id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Ubicacion eliminada correctamente"})
}
